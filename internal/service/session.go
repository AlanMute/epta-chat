package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

// TODO: Коля делает
type SessionService struct {
	repo repository.Session
}

func NewSessionService(repo repository.Session) *SessionService {
	return &SessionService{repo: repo}
}

func (s *SessionService) Add(session core.Session) error {
	return nil
}

func (s *SessionService) CheckRefresh(token string) error {
	return nil
}
