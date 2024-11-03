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
		UserName: login,
		Password: hashPsw,
	})
}

func (s *UserService) SignIn(login, password string) (uint64, core.Tokens, error) {
	userId, err := s.repo.SignIn(core.User{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return userId, core.Tokens{}, err
	}

	accessToken, err := s.tokenManager.NewAccessToken(strconv.FormatUint(userId, 10), time.Hour)
	if err != nil {
		return userId, core.Tokens{}, err
	}

	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return userId, core.Tokens{}, err
	}

	t := time.Now().AddDate(0, 0, 30)
	session := core.Session{
		UserId:         userId,
		RefreshToken:   refreshToken,
		ExpirationTime: t,
	}

	if err := s.repo.AddSession(session); err != nil {

		return userId, core.Tokens{}, nil
	}

	token := core.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return userId, token, nil
}

func (s *UserService) SetUserName(userId uint64, userName string) error {
	return s.repo.SetUserName(userId, userName)
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
