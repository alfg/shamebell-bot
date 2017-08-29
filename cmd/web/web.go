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
	apiBaseURL = "https://discordapp.com/api"

	r *redis.Client

	es eventsource.EventSource
)

type Stats struct {
	Total string `json:"total"`
}

func (c *Stats) ToJSON() []byte {
	data, _ := json.Marshal(c)
	return data
}

func NewStatsUpdate() *Stats {
	var total *redis.StringCmd

	errors, err := r.Pipelined(func(pipe *redis.Pipeline) error {
		total = pipe.Get("shamebell:total")
		return nil
	})
	fmt.Println(total.Val())

	if err != nil {
		fmt.Println("Failed to get update from redis", errors, err)
	}

	return &Stats{
		Total: total.Val(),
	}
}

func broadcastLoop() {
	var id int
	for {
		time.Sleep(time.Second * 5)
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
