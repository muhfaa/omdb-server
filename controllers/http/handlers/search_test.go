package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/muhfaa/omdb-server/core"
	omdbrepo "github.com/muhfaa/omdb-server/repository/omdb"
	"github.com/muhfaa/omdb-server/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchHandler", func() {

	Describe("getPageAndSearchWord", func() {

		mockSearchService := service.NewMockSearchService()

		Context("given pagination and searchword", func() {
			It("yield correct", func() {
				han := NewSearchHandler(mockSearchService)

				r, err := http.NewRequest(http.MethodGet, "/?pagination=2&searchword=Batman", nil)
				Expect(err).NotTo(HaveOccurred())

				page, searchWord := han.getPageAndSearchWord(r)
				Expect(page).To(Equal(uint(2)))
				Expect(searchWord).To(Equal("Batman"))
			})
		})
		Context("given pagination only", func() {
			It("yield correct", func() {
				han := NewSearchHandler(mockSearchService)

				r, err := http.NewRequest(http.MethodGet, "/?pagination=2", nil)
				Expect(err).NotTo(HaveOccurred())

				page, searchWord := han.getPageAndSearchWord(r)
				Expect(page).To(Equal(uint(2)))
				Expect(searchWord).To(Equal(""))
			})
		})
		Context("given searchword only", func() {
			It("yield correct", func() {
				han := NewSearchHandler(mockSearchService)

				r, err := http.NewRequest(http.MethodGet, "/?searchword=Batman", nil)
				Expect(err).NotTo(HaveOccurred())

				page, searchWord := han.getPageAndSearchWord(r)
				Expect(page).To(Equal(uint(1)))
				Expect(searchWord).To(Equal("Batman"))
			})
		})
	})

	Describe("ServeHTTP", func() {
		Context("there's data", func() {
			It("return http.StatusOK", func() {
				mockSearchService := service.NewMockSearchService()

				mockSearchService.MockSearch = func(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error) {
					return &core.OMDBSearchResult{}, nil
				}

				r, err := http.NewRequest(http.MethodGet, "/", nil)
				Expect(err).NotTo(HaveOccurred())

				w := httptest.NewRecorder()

				han := NewSearchHandler(mockSearchService)

				han.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
			})
		})
		Context("no data", func() {
			It("return http.StatusNotFound", func() {
				mockSearchService := omdbrepo.NewMock()

				mockSearchService.MockSearch = func(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error) {
					return nil, nil
				}

				r, err := http.NewRequest(http.MethodGet, "/", nil)
				Expect(err).NotTo(HaveOccurred())

				w := httptest.NewRecorder()

				han := NewSearchHandler(mockSearchService)

				han.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
			})
		})

	})

})
