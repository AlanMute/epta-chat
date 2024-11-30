package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"

	dbConfig "github.com/KrizzMU/coolback-alkol/internal/config/dbConf"
	"github.com/KrizzMU/coolback-alkol/internal/core"
)

const userId = "users(id)"

func GetConnection() *gorm.DB {
	db, err := gorm.Open("postgres", dbConfig.GetConnectionString())
	if err != nil {
		panic(err)
	}

	logrus.Info("Stating creating tables")
	db.AutoMigrate(&core.User{})

	db.AutoMigrate(&core.Contact{})
	db.Model(&core.Contact{}).AddForeignKey("owner_id", userId, "CASCADE", "CASCADE").AddForeignKey("contact_id", userId, "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Session{})
	db.Model(&core.Session{}).AddForeignKey("user_id", userId, "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Chat{})

	db.AutoMigrate(&core.ChatMembers{})
	db.Model(&core.ChatMembers{}).AddForeignKey("chat_id", "chats(id)", "CASCADE", "CASCADE").AddForeignKey("member_id", userId, "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Message{})
	db.Model(&core.Message{}).AddForeignKey("sender_id", userId, "CASCADE", "CASCADE").AddForeignKey("chat_id", "chats(id)", "CASCADE", "CASCADE")

	logrus.Info("Finish creating tables")

	return db
}
