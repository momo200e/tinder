package service

import (
	"tinder/domain"
)

type service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) domain.Service {
	return &service{
		repo: repo,
	}
}
