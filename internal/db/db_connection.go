package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	dbConfig "github.com/KrizzMU/coolback-alkol/internal/config/dbConf"
	"github.com/KrizzMU/coolback-alkol/internal/core"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("postgres", dbConfig.GetConnectionString())
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&core.User{})

	db.AutoMigrate(&core.Contact{})
	db.Model(&core.Contact{}).AddForeignKey("owner_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("contact_id", "users(id)", "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Session{})
	db.Model(&core.Session{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Chat{})

	db.AutoMigrate(&core.ChatMembers{})
	db.Model(&core.ChatMembers{}).AddForeignKey("chat_id", "chats(id)", "CASCADE", "CASCADE").AddForeignKey("member_id", "users(id)", "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Message{})
	db.Model(&core.Message{}).AddForeignKey("sender_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("chat_id", "chats(id)", "CASCADE", "CASCADE")

	return db
}
