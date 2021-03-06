package main

import (
	"github.com/po1yb1ank/FSMOrchestrator/config"
	"github.com/po1yb1ank/FSMOrchestrator/internal/server"
	"log"
	"sync"
)

func main() {
	s := server.Server{
		Cfg: config.Cfg{
			Port:   "8080",
			Remote: "http://ce80dce81eb0.ngrok.io",
		},
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
