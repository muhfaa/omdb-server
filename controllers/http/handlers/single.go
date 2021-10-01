package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhfaa/omdb-server/controllers/http/presenters"
	"github.com/muhfaa/omdb-server/service"
)

type SingleHandler struct {
	singleService service.SingleService
}

func NewSingleHandler(singleService service.SingleService) SingleHandler {
	return SingleHandler{
		singleService: singleService,
	}
}

func (han SingleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	single, err := han.singleService.Single(r.Context(), id)
	if err != nil || single == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	presenters.WriteHTTPJSON(w, single)
}
