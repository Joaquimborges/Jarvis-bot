package wake_up_server

import (
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/rest"
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
)

type WakeServers struct {
	client rest.Waitress
	urls   []string
}

func NewWakeServersUsecase(client rest.Waitress, urls ...string) *WakeServers {
	return &WakeServers{
		client: client,
		urls:   urls,
	}
}

func (*WakeServers) IsValid(message string) bool {
	return util.ContainsValue(
		message,
		[]string{"wake", "ping server", "server"},
	)
}

func (w *WakeServers) BuildResponse(_ string) string {
	if w.urls == nil || len(w.urls) < 1 {
		return "You don't have any server URL"
	}

	respChan := make(chan interface{})
	for _, url := range w.urls {
		go func(url string, ch chan<- interface{}) {
			bytes, err := w.client.Get(url)
			if err != nil {
				ch <- fmt.Errorf("%v, URL: %s", err, url)
				return
			}
			ch <- bytes
		}(url, respChan)
	}

	for {
		response := <-respChan
		switch resp := response.(type) {
		case error:
			return resp.Error()
		case []byte:
			if response != nil {
				return "The servers are ready"
			}
			return "the request may not have been successful"
		}
	}
}
