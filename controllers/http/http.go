package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhfaa/omdb-server/controllers/http/handlers"
	"github.com/muhfaa/omdb-server/service"
)

func RunServer(port string, searchService service.SearchService, singleService service.SingleService) {
	routes := mux.NewRouter()

	routes.Handle("/search", handlers.NewSearchHandler(searchService))
	routes.Handle("/single/{id}", handlers.NewSingleHandler(singleService))

	http.ListenAndServe("0.0.0.0:"+port, routes)
}
