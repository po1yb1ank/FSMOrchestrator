package rest

import (
	"bytes"
	json "github.com/json-iterator/go"
	"github.com/po1yb1ank/FSMOrchestrator/internal"
	"github.com/po1yb1ank/FSMOrchestrator/internal/fsm"
	"log"
	"net/http"
	"net/http/httputil"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {

	m := internal.Machine{}
	ch := make(chan internal.Machine, 2)
	log.Println(r.Body)
	b, _ := httputil.DumpRequest(r, true)
	log.Println(string(b))
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
		jsonRes, err := json.Marshal(&m)
		if err != nil{
			log.Println("err on marshall", err)
		}
		jsonReader := bytes.NewReader(jsonRes)

		cli := http.DefaultClient
		req,_ := http.NewRequest("POST", internal.RemotePath, jsonReader)
		req.Header = r.Header
		req.Header.Add("DASHA-LENGTH", r.Header.Get("Content-Length"))


		_, err = cli.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			log.Println("Error while POST to: ", internal.RemotePath, err)
			return
		}
	case false:
	}
}