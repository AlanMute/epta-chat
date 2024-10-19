package repository

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type ContactRepo struct {
	db *gorm.DB
}

func NewContactPostgres(db *gorm.DB) *ContactRepo {
	return &ContactRepo{
		db: db,
	}
}

func (r *ContactRepo) Add(ownerId uint64, contactLogin string) error {
	var user core.User

	if err := r.db.Where("login = ?", contactLogin).First(&user).Error; err != nil {
		return err
	}

	newContact := core.Contact{
		OwnerId:   ownerId,
		ContactId: user.ID,
	}

	if result := r.db.Create(&newContact); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ContactRepo) Delete(ownerId, ContactId uint64) error {
	if result := r.db.
		Where("owner_id = ? AND contact_id = ?", ownerId, ContactId).
		Delete(&core.Contact{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ContactRepo) GetAll(ownerId uint64) ([]core.UserInfo, error) {
	var contacts []core.Contact

	if result := r.db.Where("owner_id = ?", ownerId).Find(&contacts); result.Error != nil {
		return nil, result.Error
	}

	var usersInfo []core.UserInfo

	for _, contact := range contacts {
		var user core.User
		if result := r.db.Where("id = ?", contact.ContactId).First(&user); result.Error != nil {
			return nil, result.Error
		}

		userInfo := core.UserInfo{
			ID:       user.ID,
			Login:    user.Login,
			UserName: user.UserName,
		}

		usersInfo = append(usersInfo, userInfo)
	}

	return usersInfo, nil
}

func (r *ContactRepo) GetById(id uint64) (core.UserInfo, error) {
	var user core.User

	result := r.db.First(&user, id)
	if result.Error != nil {
		return core.UserInfo{}, result.Error
	}

	return core.UserInfo{
		ID:       user.ID,
		Login:    user.Login,
		UserName: user.UserName,
	}, nil
}
