package repository

import (
	"fmt"
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	var dbUser core.User

	if result := r.db.Where("login = ?", user.Login).First(&dbUser); result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, fmt.Errorf("login or password is incorrect")
		}

		return 0, result.Error
	}

	if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
		return 0, fmt.Errorf("login or password is incorrect")
	}

	return dbUser.ID, nil
}

func (r *UserRepo) SetUserName(userId uint64, userName string) error {
	var user core.User

	if result := r.db.First(&user, userId); result.Error != nil {
		return result.Error
	}

	user.UserName = userName

	if result := r.db.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepo) SignUp(user core.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}

	chatMember := core.ChatMembers{
		MemberId: user.ID,
		ChatId:   0,
	}

	return r.db.Create(&chatMember).Error
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

func (r *UserRepo) GetById(userId uint64) (core.User, error) {
	var user core.User

	if result := r.db.Where("id = ?", userId).First(&user); result.Error != nil {
		return core.User{}, result.Error
	}

	return user, nil
}
