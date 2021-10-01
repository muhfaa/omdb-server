package service_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/muhfaa/omdb-server/core"
	omdbrepo "github.com/muhfaa/omdb-server/repository/omdb"
	"github.com/muhfaa/omdb-server/service"
)

var _ = Describe("Single", func() {
	mockOMDBRepo := omdbrepo.NewMock()

	Context("there's data", func() {
		It("ok", func() {
			mockOMDBRepo.MockGetByID = func(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
				return &core.OMDBResultSingle{}, nil
			}
			use := service.NewSingleService(mockOMDBRepo)

			resp, err := use.Single(context.Background(), "Bowie")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).NotTo(BeNil())
		})
	})

	Context("no data", func() {
		It("error", func() {
			mockOMDBRepo.MockGetByID = func(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
				return nil, nil
			}
			use := service.NewSingleService(mockOMDBRepo)

			resp, err := use.Single(context.Background(), "Bowie")
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})
	})
})
