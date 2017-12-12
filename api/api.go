package api

import (
	"encoding/json"
	"fmt"
	"github.com/msergeant/game-of-life-go/model"
	"net"
	"net/http"
	"time"
)

type Config struct {
}

func Start(cfg Config, listener net.Listener) {
	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	http.Handle("/worlds/next", worldHandler())

	go server.Serve(listener)
}

func worldHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-Prototype-Version, Token, Content-Type")
		w.Header().Set("Access-Control-Max-Age", "1728000")

		if r.Method == "POST" {
			fmt.Println("Started POST /worlds/next")
			decoder := json.NewDecoder(r.Body)
			var world model.World
			err := decoder.Decode(&world)
			if err != nil {
				fmt.Println("Decoding Error")
			}
			defer r.Body.Close()

			next := world.GenerateNextIteration()
			js, err := json.Marshal(next)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Println(next)

			elapsed := time.Since(start)
			fmt.Printf("Response took", elapsed, "\n")

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		} else if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println("The request was not a post")
		}
	})
}
