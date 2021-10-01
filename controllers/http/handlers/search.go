package handlers

import (
	"net/http"
	"strconv"

	"github.com/muhfaa/omdb-server/controllers/http/presenters"
	"github.com/muhfaa/omdb-server/service"
)

type SearchHandler struct {
	searchService service.SearchService
}

func NewSearchHandler(searchService service.SearchService) SearchHandler {
	return SearchHandler{
		searchService: searchService,
	}
}

func (han SearchHandler) getPageAndSearchWord(r *http.Request) (uint, string) {
	qs := r.URL.Query()

	page, err := strconv.Atoi(qs.Get("pagination"))
	if err != nil || page < 0 {
		page = 1
	}

	return uint(page), qs.Get("searchword")
}

func (han SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page, searchword := han.getPageAndSearchWord(r)

	response, err := han.searchService.Search(r.Context(), searchword, page)
	if err != nil || response == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	presenters.WriteHTTPJSON(w, response)
}
