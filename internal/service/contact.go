package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type ContactService struct {
	repo repository.Contact
}

func NewContactService(repo repository.Contact) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) Add(ownerId uint64, contactLogin string) error {
	return s.repo.Add(ownerId, contactLogin)
}

func (s *ContactService) Delete(ownerId, conactId uint64) error {
	return s.repo.Delete(ownerId, conactId)
}

func (s *ContactService) GetAll(ownerId uint64) ([]core.UserInfo, error) {
	return s.repo.GetAll(ownerId)
}

func (s *ContactService) GetById(id uint64) (core.UserInfo, error) {
	return s.repo.GetById(id)
}
