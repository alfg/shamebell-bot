package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/antage/eventsource"
)

var (
	apiBaseURL = "https://discordapp.com/api"

	es eventsource.EventSource
)

func broadcastLoop() {
	var id int
	for {
		time.Sleep(time.Second * 5)
		es.SendEventMessage(string("testing"), "message", strconv.Itoa(id))
		id++
	}
}

func server() {
	s := http.NewServeMux()
	port := "4000"
	// s.Handle("/", http.FileServer(http.Dir("static/build")))
	s.Handle("/events", es)
	fmt.Println("server running on ", port)
	http.ListenAndServe(":"+port, s)
}

func main() {
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
