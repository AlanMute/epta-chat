package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserSevice(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Add(login, password string) error {
	return nil
}
