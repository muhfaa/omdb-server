package repository

import "net/http"

type HTTPClient interface {
	Do(r *http.Request) (*http.Response, error)
}

type MockHTTPClient struct {
	MockDo func(req *http.Request) (*http.Response, error)
}

func (m MockHTTPClient) Do(r *http.Request) (*http.Response, error) {
	return m.MockDo(r)
}

func NewMockHTTPClient(responseDo func(req *http.Request) (*http.Response, error)) MockHTTPClient {
	return MockHTTPClient{
		MockDo: responseDo,
	}
}
