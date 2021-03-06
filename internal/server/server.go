package server

import (
	"github.com/gorilla/mux"
	"github.com/po1yb1ank/FSMOrchestrator/config"
	"github.com/po1yb1ank/FSMOrchestrator/internal"
	"github.com/po1yb1ank/FSMOrchestrator/internal/fsm"
	"github.com/po1yb1ank/FSMOrchestrator/internal/rest"
	"log"
	"net/http"
)

type Server struct {
	Cfg config.Cfg
}

func (s *Server) Start() error {

	fsm.InitMachine()
	log.Println("FSM init: done")
	internal.RemotePath = s.Cfg.Remote
	log.Println("Set remote to:", s.Cfg.Remote)
	r := mux.NewRouter()

	r.HandleFunc("/request", rest.RequestHandler).Methods("POST")
	log.Println("Listening on:", s.Cfg.Port)
	if err := http.ListenAndServe(":"+s.Cfg.Port, r); err != nil {
		return err
	}

	return nil
}
func (s *Server) GetRemote() string {
	return s.Cfg.Remote
}
