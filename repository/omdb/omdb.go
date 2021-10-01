package omdb

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/muhfaa/omdb-server/core"
	"github.com/muhfaa/omdb-server/repository"
)

var (
	ErrNoData = errors.New("no data")
)

type OMDBRepo struct {
	client  repository.HTTPClient
	baseURL string
	apiKey  string
}

func NewOMDBRepo(client repository.HTTPClient, baseURL string, apiKey string) OMDBRepo {
	return OMDBRepo{
		client:  client,
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (s OMDBRepo) httpGetByte(ctx context.Context, endpoint string) ([]byte, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(r)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (s OMDBRepo) Search(ctx context.Context, text string, page uint) (*core.OMDBSearchResult, error) {
	qs := url.Values{
		"apikey": []string{s.apiKey},
		"s":      []string{text},
		"page":   []string{strconv.Itoa(int(page))},
	}

	endpoint := s.baseURL + "?" + qs.Encode()

	body, err := s.httpGetByte(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var out core.OMDBSearchResult

	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}

	if out.Response == "False" {
		return nil, ErrNoData
	}

	return &out, nil
}

func (s OMDBRepo) GetByID(ctx context.Context, id string) (*core.OMDBResultSingle, error) {
	qs := url.Values{
		"apikey": []string{s.apiKey},
		"i":      []string{id},
	}

	endpoint := s.baseURL + "?" + qs.Encode()

	body, err := s.httpGetByte(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var out core.OMDBResultSingle

	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}

	if out.Response == "False" {
		return nil, ErrNoData
	}

	return &out, nil
}
