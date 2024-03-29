package db

import (
	"log"
	"sut-notification-go/domain/notification/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	models := []interface{}{&model.Notification{}}
	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return Handler{db}
}
