package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type ContactService struct {
	repo repository.Contact
}

func NewContactSevice(repo repository.Contact) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) Add(ownerId, conactId uint64) error {
	return nil
}

func (s *ContactService) Delete(ownerId, conactId uint64) error {
	return nil
}

func (s *ContactService) GetAll(ownerId uint64) ([]core.UserInfo, error) {
	return nil, nil
}

func (s *ContactService) GetById(id uint64) (core.UserInfo, error) {
	return core.UserInfo{}, nil
}
