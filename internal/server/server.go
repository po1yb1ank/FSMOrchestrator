package server

import (
	"github.com/gorilla/mux"
	"github.com/po1yb1ank/FSMOrchestrator/config"
	"github.com/po1yb1ank/FSMOrchestrator/internal/fsm"
	"github.com/po1yb1ank/FSMOrchestrator/internal/rest/endpoint"
	"log"
	"net/http"
)

type Server struct {
	Cfg config.Cfg
}
func(s *Server)Start() error{

	fsm.InitMachine()
	endpoint.RemotePath = s.Cfg.Remote

	r := mux.NewRouter()

	r.HandleFunc("/request",endpoint.RequestHandler).Methods("POST")
	log.Println("Listening on:", s.Cfg.Port)
	if err := http.ListenAndServe(":"+s.Cfg.Port, r); err != nil{
		return err
	}

	return nil
}
func(s *Server)GetRemote() string{
	return s.Cfg.Remote
}
