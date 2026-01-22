// Package service is to foo
package service

import (
	"authServer/internal/model"
	"authServer/internal/repository"
)

type IKeyService interface {
	AuthStreamingKey(name, key string) (*model.Keys, error)
}

type keysService struct {
	keysRepository repository.IKeyRepository
}

func NewKeysService(repo repository.IKeyRepository) IKeyService {
	return &keysService{
		keysRepository: repo,
	}
}

func (s *keysService) AuthStreamingKey(name, key string) (*model.Keys, error) {
	return s.keysRepository.FindStreamKey(name, key)
}
