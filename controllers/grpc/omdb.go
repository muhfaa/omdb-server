package grpccontroller

import (
	"context"

	"github.com/muhfaa/omdb-server/controllers/grpc/presenters"
	"github.com/muhfaa/omdb-server/core"
	"github.com/muhfaa/omdb-server/repository/grpcstub"
	"github.com/muhfaa/omdb-server/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound = status.Errorf(codes.NotFound, "not found")
)

type GRPCServiceImpl struct {
	grpcstub.UnimplementedOmdbServer

	searchService service.SearchService
	singleService service.SingleService
}

func NewGRPCService(searchService service.SearchService, singleService service.SingleService) GRPCServiceImpl {
	return GRPCServiceImpl{
		searchService: searchService,
		singleService: singleService,
	}
}

func (impl GRPCServiceImpl) Search(ctx context.Context, r *core.SearchRequest) (*core.SearchReply, error) {
	response, err := impl.searchService.Search(ctx, r.GetSearchword(), uint(r.GetPage()))
	if err != nil || response == nil {
		return nil, ErrNotFound
	}

	return presenters.SearchResultToProto(response), nil
}

func (impl GRPCServiceImpl) Single(ctx context.Context, r *core.SingleRequest) (*core.SingleReply, error) {
	single, err := impl.singleService.Single(ctx, r.GetId())
	if err != nil || single == nil {
		return nil, ErrNotFound
	}

	return presenters.SingleToProto(single), nil
}
