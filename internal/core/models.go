package core

import (
	"time"
)

type GormModel struct {
	ID uint64 `gorm:"primary_key" json:"id"`
}

type User struct {
	GormModel

	Login    string `gorm:"not null; unique"`
	UserName string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type Contact struct {
	GormModel

	OwnerId uint64
	Owner   User `gorm:"foreignKey:OwnerId"`

	ConactId uint64
	Contact  User `gorm:"foreignKey:ConactId"`
}

type Token struct {
	GormModel

	UserId         uint64
	RefreshToken   string `gorm:"not null"`
	ExpirationTime time.Time

	User User `gorm:"foreignKey:UserId"`
}

type Chat struct {
	GormModel

	Name     string `gorm:"not null"`
	IsDirect bool   `gorm:"not null"`
	OwnerId  uint64

	Owner User `gorm:"foreignKey:OwnerId"`
}

type ChatMembers struct {
	GormModel

	MemberId uint64 `gorm:"not null"`
	ChatId   uint64 `gorm:"not null"`

	Member User `gorm:"foreignKey:MemberId"`
	Chat   Chat `gorm:"foreignKey:ChatId"`
}

type Message struct {
	GormModel

	Text        string `gorm:"not null"`
	SendingTime time.Time
	SenderId    uint64 `gorm:"not null"`
	ChatId      uint64 `gorm:"not null"`

	Sender User `gorm:"foreignKey:SenderId"`
	Chat   Chat `gorm:"foreignKey:ChatId"`
}
