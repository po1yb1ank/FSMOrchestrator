package endpoint

import (
	json "github.com/json-iterator/go"
	"github.com/po1yb1ank/FSMOrchestrator/internal/fsm"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request){

	m := Machine{}
	ch := make(chan Machine, 2)
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ch <- m
	fsm.PushMachine(ch)

}
