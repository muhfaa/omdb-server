package grpccontroller_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	grpccontroller "github.com/muhfaa/omdb-server/controllers/grpc"
	"github.com/muhfaa/omdb-server/core"
	"github.com/muhfaa/omdb-server/service"
)

var _ = Describe("Impl", func() {

	mockSearchService := service.NewMockSearchService()
	mockSingleService := service.NewMockSingleService()

	Describe("Search", func() {
		Context("there's data", func() {
			It("ok", func() {
				mockSearchService.MockSearch = func(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error) {
					return &core.OMDBSearchResult{
						Search: []core.OMDBResultCompact{
							{
								IMDBID: "DavidBowie",
							},
						},
					}, nil
				}

				service := grpccontroller.NewGRPCService(mockSearchService, mockSingleService)

				reply, err := service.Search(context.Background(), &core.SearchRequest{
					Page:       1,
					Searchword: "Batman",
				})

				Expect(err).NotTo(HaveOccurred())

				Expect(reply).NotTo(BeNil())
				Expect(len(reply.GetSearch())).To(Equal(1))
			})
		})

		Context("no data", func() {
			It("error", func() {
				mockSearchService.MockSearch = func(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error) {
					return nil, nil
				}

				service := grpccontroller.NewGRPCService(mockSearchService, mockSingleService)

				_, err := service.Search(context.Background(), &core.SearchRequest{
					Page:       1,
					Searchword: "Batman",
				})

				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Single", func() {
		Context("there's data", func() {
			It("ok", func() {
				mockSingleService.MockSingle = func(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
					return &core.OMDBResultSingle{
						Response: "True",
					}, nil
				}

				service := grpccontroller.NewGRPCService(mockSearchService, mockSingleService)

				reply, err := service.Single(context.Background(), &core.SingleRequest{
					Id: "DavidBowie",
				})

				Expect(err).NotTo(HaveOccurred())

				Expect(reply).NotTo(BeNil())
				Expect(reply.Response).To(Equal("True"))
			})
		})

		Context("no data", func() {
			It("error", func() {
				mockSingleService.MockSingle = func(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
					return nil, nil
				}

				service := grpccontroller.NewGRPCService(mockSearchService, mockSingleService)

				_, err := service.Single(context.Background(), &core.SingleRequest{
					Id: "DavidBowie",
				})

				Expect(err).To(HaveOccurred())
			})
		})
	})

})
