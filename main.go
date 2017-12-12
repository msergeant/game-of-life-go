package main

import (
	"flag"
	"github.com/msergeant/game-of-life-go/daemon"
	"log"
)

func processFlags() *daemon.Config {
	cfg := &daemon.Config{}

	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")

	flag.Parse()
	return cfg
}

func main() {
	cfg := processFlags()

	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}
