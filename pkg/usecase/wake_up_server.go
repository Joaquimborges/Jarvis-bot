package usecase

import (
	"context"
	"errors"
	"fmt"
)

/*
WakeAllTestServers It's a really cool feature that
allows you to do a health check on your test or even prod servers.
You just need a public route that only returns a valid status code.
*/
func (j *JarvisUsecase) WakeAllTestServers(urls ...string) error {
	if urls == nil {
		return errors.New("request without url")
	}

	respChan := make(chan interface{})
	for _, url := range urls {
		go func(url string, ch chan<- interface{}) {
			bytes, err := j.client.Get(context.Background(), url)
			if err != nil {
				ch <- fmt.Errorf("%v, URL: %s", err, url)
				return
			}
			ch <- bytes
			return
		}(url, respChan)
	}

	for {
		select {
		case response := <-respChan:
			switch response.(type) {
			case error:
				return response.(error)
			case []byte:
				if response != nil {
					return nil
				}
				return errors.New("the request may not have been successful")
			}
		}
		return nil
	}
}
