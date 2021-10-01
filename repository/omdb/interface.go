package omdb

import (
	"context"

	"github.com/muhfaa/omdb-server/core"
)

type Repository interface {
	Search(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error)
	GetByID(ctx context.Context, id string) (*core.OMDBResultSingle, error)
}
