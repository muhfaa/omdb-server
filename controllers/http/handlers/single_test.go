package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/muhfaa/omdb-server/core"
	"github.com/muhfaa/omdb-server/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SingleHandler", func() {

	Describe("ServeHTTP", func() {
		mockSingleService := service.NewMockSingleService()

		Context("there's any data", func() {
			It("ok", func() {
				id := "davidBowie"
				mockSingleService.MockSingle = func(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
					return &core.OMDBResultSingle{
						OMDBResultCompact: core.OMDBResultCompact{
							IMDBID: id,
						},
					}, nil
				}

				han := NewSingleHandler(mockSingleService)

				w := httptest.NewRecorder()

				r, err := http.NewRequest(http.MethodGet, "/", nil)
				Expect(err).NotTo(HaveOccurred())

				r = mux.SetURLVars(r, map[string]string{
					"id": id,
				})

				han.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("no data", func() {
			It("error", func() {
				id := "davidBowie"
				mockSingleService.MockSingle = func(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
					return nil, nil
				}

				han := NewSingleHandler(mockSingleService)

				w := httptest.NewRecorder()

				r, err := http.NewRequest(http.MethodGet, "/", nil)
				Expect(err).NotTo(HaveOccurred())

				r = mux.SetURLVars(r, map[string]string{
					"id": id,
				})

				han.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

})
