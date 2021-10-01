package service

import (
	"context"
	"errors"

	"github.com/muhfaa/omdb-server/core"
	mysqlrepo "github.com/muhfaa/omdb-server/repository/mysql"
	omdbrepo "github.com/muhfaa/omdb-server/repository/omdb"
)

var (
	ErrNoData = errors.New("no data")
)

type SearchService interface {
	Search(ctx context.Context, searchWord string, page uint) (*core.OMDBSearchResult, error)
}

type searchservice struct {
	omdbRepo  omdbrepo.Repository
	mysqlRepo mysqlrepo.Repository
}

func NewSearchService(omdbRepo omdbrepo.Repository, mysqlRepo mysqlrepo.Repository) searchservice {
	return searchservice{
		omdbRepo:  omdbRepo,
		mysqlRepo: mysqlRepo,
	}
}

func (s searchservice) Search(ctx context.Context, searchWord string, page uint) (*core.OMDBSearchResult, error) {
	go s.mysqlRepo.SaveSearchActivity(searchWord)

	response, err := s.omdbRepo.Search(ctx, searchWord, page)
	if err != nil || response == nil {
		return nil, ErrNoData
	}

	return response, nil
}

type MockSearchService struct {
	MockSearch func(ctx context.Context, searchWord string, page uint) (*core.OMDBSearchResult, error)
}

func (m MockSearchService) Search(ctx context.Context, searchWord string, page uint) (*core.OMDBSearchResult, error) {
	return m.MockSearch(ctx, searchWord, page)
}

func NewMockSearchService() MockSearchService {
	return MockSearchService{
		MockSearch: func(ctx context.Context, searchWord string, page uint) (*core.OMDBSearchResult, error) {
			return nil, nil
		},
	}
}
