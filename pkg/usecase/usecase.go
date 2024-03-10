package usecase

import "github.com/Joaquimborges/jarvis-bot/pkg/rest"

type JarvisUsecase struct {
	client *rest.Client
}

func NewJarvisUsecase() *JarvisUsecase {
	return &JarvisUsecase{client: rest.NewRestClient()}
}
