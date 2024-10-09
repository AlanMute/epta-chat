package service

import (
	"strconv"
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo         repository.User
	tokenManager auth.TokenManager
}

func NewUserService(repo repository.User, t auth.TokenManager) *UserService {
	return &UserService{
		repo:         repo,
		tokenManager: t,
	}
}

func (s *UserService) SignUp(login, password string) error {
	hashPsw, err := s.hashPassword(password)
	if err != nil {
		return err
	}

	return s.repo.SignUp(core.User{
		Login:    login,
		Password: hashPsw,
	})
}

func (s *UserService) SignIn(login, password string) (core.Tokens, error) {
	hashPsw, err := s.hashPassword(password)
	if err != nil {
		return core.Tokens{}, err
	}

	userId, err := s.repo.SignIn(core.User{
		Login:    login,
		Password: hashPsw,
	})
	if err != nil {
		return core.Tokens{}, err
	}

	accessToken, err := s.tokenManager.NewAccessToken(strconv.FormatUint(userId, 10), time.Hour)
	if err != nil {
		return core.Tokens{}, err
	}

	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return core.Tokens{}, err
	}

	t := time.Now().AddDate(0, 0, 30)
	session := core.Session{
		RefreshToken:   refreshToken,
		ExpirationTime: t,
	}

	if err := s.repo.AddSession(session); err != nil {

		return core.Tokens{}, nil
	}

	token := core.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}

func (s *UserService) Refresh(userId uint64, refreshToken string) (string, error) {

	if err := s.repo.CheckRefresh(refreshToken); err != nil {
		return "", err
	}

	accessToken, err := s.tokenManager.NewAccessToken(strconv.FormatUint(userId, 10), time.Hour)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *UserService) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
