package service_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/muhfaa/omdb-server/core"
	mysqlrepo "github.com/muhfaa/omdb-server/repository/mysql"
	omdbrepo "github.com/muhfaa/omdb-server/repository/omdb"
	"github.com/muhfaa/omdb-server/service"
)

var _ = Describe("Search", func() {

	mockOMDBRepo := omdbrepo.NewMock()
	mockMysqlRepo := mysqlrepo.NewMock()

	Context("there's data", func() {
		It("ok", func() {
			mockOMDBRepo.MockSearch = func(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error) {
				return &core.OMDBSearchResult{}, nil
			}
			use := service.NewSearchService(mockOMDBRepo, mockMysqlRepo)

			resp, err := use.Search(context.Background(), "Lebron James", 1)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).NotTo(BeNil())
		})
	})

	Context("no data", func() {
		It("error", func() {
			mockOMDBRepo.MockSearch = func(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error) {
				return nil, nil
			}
			use := service.NewSearchService(mockOMDBRepo, mockMysqlRepo)

			resp, err := use.Search(context.Background(), "Lebron James", 1)
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})
	})
})
