package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

//go:generate mockgen -source ./rest.go -destination ../../internal/mocks/gateway/restclient_mock.go -package mocks_gateway
type Waitress interface {
	Get(url string) ([]byte, error)
	Post(url string, body interface{}, headers map[string]string) error
}

type Client struct {
	Rest *http.Client
}

func NewRestClient() Waitress {
	client := http.Client{Timeout: time.Second * 5}
	return &Client{
		Rest: &client,
	}
}

func (client *Client) Get(url string) ([]byte, error) {
	response, err := client.Rest.Get(url)
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

func (client *Client) Post(url string, body interface{}, headers map[string]string) error {
	byts, bErr := json.Marshal(body)
	if bErr != nil {
		return bErr
	}

	bodyReader := bytes.NewReader(byts)
	request, er := http.NewRequest(http.MethodPost, url, bodyReader)
	if er != nil {
		return er
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}
	response, err := client.Rest.Do(request)
	if err != nil {
		return err
	}

	_, err = client.handleResponseError(response)
	if err != nil {
		return err
	}
	return nil
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
