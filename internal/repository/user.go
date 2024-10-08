package repository

import (
	"fmt"
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db                   *gorm.DB
	deleteRoutineStarted bool
}

func NewUserPostgres(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) SignIn(user core.User) (uint64, error) {
	if result := r.db.Where("login = ? AND password = ?", user.Login, user.Password).First(&user); result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, fmt.Errorf("login or password is incorrect")
		}

		return 0, result.Error
	}

	return 0, nil
}

func (r *UserRepo) SignUp(user core.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepo) AddSession(session core.Session) error {
	if result := r.db.Create(&session); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepo) CheckRefresh(token string) error {
	var session core.Session

	if !r.deleteRoutineStarted {
		go r.deleteExpiredTokens()
		r.deleteRoutineStarted = true
	}

	if result := r.db.Where("refresh_token = ? AND expiration_time > ?", token, time.Now()).First(&session); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepo) deleteExpiredTokens() {
	for {
		r.db.Where("expiration_time < ?", time.Now()).Delete(&core.Session{})

		time.Sleep(24 * time.Hour)
	}
}
