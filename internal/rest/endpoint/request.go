package endpoint

import (
	json "github.com/json-iterator/go"
	"github.com/po1yb1ank/FSMOrchestrator/internal/fsm"
	"log"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {

	m := Machine{}
	ch := make(chan Machine, 2)
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error on parsing json from /request:", err)
		return
	}
	ch <- m
	switch fsm.PushMachine(ch) {
	case true:
		log.Println("Passing to remote")
		_, err := http.Post(RemotePath, "application/json", r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			log.Println("Error while POST to: ", RemotePath, err)
			return
		}
	case false:
	}
}
