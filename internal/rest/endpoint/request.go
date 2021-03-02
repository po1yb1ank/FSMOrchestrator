package endpoint

import (
	json "github.com/json-iterator/go"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request){

	m := Machine{}

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
