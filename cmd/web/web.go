package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	redis "gopkg.in/redis.v3"

	"github.com/antage/eventsource"
)

var (
	r *redis.Client
	es eventsource.EventSource
)

type Stats struct {
	Total    string `json:"total"`
	Guilds   string `json:"guilds"`
	Channels string `json:"channels"`
}

func (c *Stats) ToJSON() []byte {
	data, _ := json.Marshal(c)
	return data
}

func NewStatsUpdate() *Stats {
	var (
		total    *redis.StringCmd
		guilds   *redis.IntCmd
		channels *redis.IntCmd
	)

	errors, err := r.Pipelined(func(pipe *redis.Pipeline) error {
		total = pipe.Get("shamebell:total")
		guilds = pipe.SCard("shamebell:guilds")
		channels = pipe.SCard("shamebell:channels")
		return nil
	})

	if err != nil {
		fmt.Println("Failed to get update from redis", errors, err)
	}

	return &Stats{
		Total:    total.Val(),
		Guilds:   strconv.FormatInt(guilds.Val(), 10),
		Channels: strconv.FormatInt(channels.Val(), 10),
	}
}

func broadcastLoop() {
	var id int
	for {
		time.Sleep(time.Second * 1)
		es.SendEventMessage(string(NewStatsUpdate().ToJSON()), "message", strconv.Itoa(id))
		id++
	}
}

func connectRedis(conn string) error {
	r = redis.NewClient(&redis.Options{Addr: conn, DB: 0})

	_, err := r.Ping().Result()
	if err != nil {
		fmt.Println("failed to connect to redis", err)
		return err
	}

	return nil
}

func server() {
	s := http.NewServeMux()

	s.Handle("/", http.FileServer(http.Dir("static/build")))
	s.Handle("/events", es)

	port := "4000"
	fmt.Println("server running on ", port)
	http.ListenAndServe(":"+port, s)
}

func main() {
	var (
		Redis = flag.String("r", "", "Redis Connection String")
	)
	flag.Parse()

	if *Redis != "" {
		if connectRedis(*Redis) != nil {
			return
		}
	}

	// Start event loop.
	es = eventsource.New(
		eventsource.DefaultSettings(),

		// We need to set custom headers for the SSE due to a webpack bug.
		// See: https://github.com/facebookincubator/create-react-app/issues/966#issuecomment-271311044
		func(req *http.Request) [][]byte {
			return [][]byte{
				[]byte("Cache-Control: no-transform"),
			}
		},
	)
	defer es.Close()
	go broadcastLoop()

	// Start server.
	server()
}
