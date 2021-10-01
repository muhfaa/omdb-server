package presenters

import (
	"encoding/json"
	"net/http"
)

func WriteHTTPJSON(w http.ResponseWriter, response interface{}) {
	resByte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resByte)
}
