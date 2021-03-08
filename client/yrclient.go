package client

import (
	"net/http"
)

type YrClient struct {
	HttpClient *http.Client
	UserAgent  string
}

func NewYrClient(httpClient *http.Client, userAgent string) *YrClient {
	return &YrClient{HttpClient: httpClient, UserAgent: userAgent}
}

func (c *YrClient) GetRequest(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", c.UserAgent)

	return c.HttpClient.Do(request)
}
