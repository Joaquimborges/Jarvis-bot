package rest

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURl string
	Rest    *http.Client
}

func NewRestClient(baseURL string) *Client {
	client := http.Client{Timeout: time.Second * 5}
	return &Client{
		BaseURl: baseURL,
		Rest:    &client,
	}
}

func (client *Client) Get(ctx context.Context, path string) ([]byte, error) {
	response, err := client.Rest.Get(fmt.Sprintf("%s/%s", client.BaseURl, path))
	if err != nil {
		return nil, fmt.Errorf("unable to make GET request, %v", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)
	return client.handleResponseError(response)
}

func (client *Client) handleResponseError(response *http.Response) ([]byte, error) {
	bytes, er := io.ReadAll(response.Body)
	if er != nil {
		return nil, er
	}

	if !client.IsAcceptedStatusCode(response.StatusCode) {
		return nil, fmt.Errorf("http request failed, %v", string(bytes))
	}
	return bytes, nil
}

func (*Client) IsAcceptedStatusCode(status int) bool {
	return status == http.StatusOK ||
		status == http.StatusCreated ||
		status == http.StatusNoContent
}
