package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func (j *JarvisUsecase) WakeAllTestServers() error {
	socketServer := os.Getenv("MACHINE_SOCKET_SERVER_URL")
	machineAPI := os.Getenv("MACHINE_API_URL")
	bossAPI := os.Getenv("BOSS_YM_API_URL")

	respChan := make(chan interface{})

	go j.makePingRequest(socketServer, respChan)
	go j.makePingRequest(machineAPI, respChan)
	go j.makePingRequest(bossAPI, respChan)

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

func (j *JarvisUsecase) makePingRequest(
	url string,
	ch chan interface{},
) {
	bytes, err := j.client.Get(context.Background(), url)
	if err != nil {
		ch <- fmt.Errorf("%v, URL: %s", err, url)
		return
	}
	ch <- bytes
	return
}
