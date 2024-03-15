package usecase

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/rest"
)

type JarvisUsecase struct {
	client *rest.Client
}

func NewJarvisUsecase() *JarvisUsecase {
	return &JarvisUsecase{client: rest.NewRestClient()}
}
