//go:build wireinject
// +build wireinject

package provider

import (
	"sync"

	"tinder/config"
	"tinder/domain"
	repo "tinder/internal/user/repo"
	svc "tinder/internal/user/service"

	"github.com/google/wire"
)

var cg *config.Config
var configOnce sync.Once

func NewConfig() *config.Config {
	configOnce.Do(func() {
		cg = config.NewConfig()
	})
	return cg
}

var repoTemp domain.Repository
var repoOnce sync.Once

func NewRepo() domain.Repository {
	repoOnce.Do(func() {
		repoTemp = repo.NewUserRepository()
	})
	return repoTemp
}

func NewService() (domain.Service, error) {
	panic(wire.Build(svc.NewService, NewRepo))
}
