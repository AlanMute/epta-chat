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
	db.Model(&core.Contact{}).AddForeignKey("owner_id", "user(id)", "CASCADE", "CASCADE").AddForeignKey("contact_id", "user(id)", "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Token{})
	db.Model(&core.Token{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Chat{})

	db.AutoMigrate(&core.ChatMembers{})
	db.Model(&core.ChatMembers{}).AddForeignKey("chat_id", "chat(id)", "CASCADE", "CASCADE").AddForeignKey("member_id", "user(id)", "CASCADE", "CASCADE")

	db.AutoMigrate(&core.Message{})
	db.Model(&core.ChatMembers{}).AddForeignKey("sender_id", "user(id)", "CASCADE", "CASCADE").AddForeignKey("chat_id", "chat(id)", "CASCADE", "CASCADE")

	return db
}
