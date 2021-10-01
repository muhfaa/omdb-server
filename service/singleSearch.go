package service

import (
	"context"

	"github.com/muhfaa/omdb-server/core"
	omdbrepo "github.com/muhfaa/omdb-server/repository/omdb"
)

type SingleService interface {
	Single(ctx context.Context, id string) (*core.OMDBResultSingle, error)
}

type singleservice struct {
	omdbRepo omdbrepo.Repository
}

func NewSingleService(omdbRepo omdbrepo.Repository) singleservice {
	return singleservice{
		omdbRepo: omdbRepo,
	}
}

func (s singleservice) Single(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
	single, err := s.omdbRepo.GetByID(ctx, id)
	if err != nil || single == nil {
		return nil, ErrNoData
	}
	return single, nil
}

type MockSingleService struct {
	MockSingle func(ctx context.Context, id string) (*core.OMDBResultSingle, error)
}

func (m MockSingleService) Single(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
	return m.MockSingle(ctx, id)
}

func NewMockSingleService() MockSingleService {
	return MockSingleService{
		MockSingle: func(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
			return nil, nil
		},
	}
}
