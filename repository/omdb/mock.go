package omdb

import (
	"context"

	"github.com/muhfaa/omdb-server/core"
)

type MockOMDB struct {
	MockSearch  func(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error)
	MockGetByID func(ctx context.Context, id string) (*core.OMDBResultSingle, error)
}

func NewMock() MockOMDB {
	return MockOMDB{
		MockSearch: func(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error) {
			return nil, nil
		},
		MockGetByID: func(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
			return nil, nil
		},
	}
}

func (m MockOMDB) Search(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error) {
	return m.MockSearch(ctx, text, page)
}
func (m MockOMDB) GetByID(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
	return m.MockGetByID(ctx, id)
}
