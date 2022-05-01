package db

import (
	"github.com/Abdumalik92/bank_card/logs"
	"github.com/Abdumalik92/bank_card/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var (
	Dbs *gorm.DB
)

func Open(logger *logrus.Logger) *gorm.DB {
	db, err := gorm.Open("mysql", "root:1111@tcp(localhost:3306)/bank?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic("Couldn't opendatabase", err.Error())
	}

	db.LogMode(true)
	db.SetLogger(&logs.GormLogger{
		Name:   "db gorm logger",
		Logger: logger,
	})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Order{}, &models.Client{}, &models.Department{}, &models.Orderprocessing{}, &models.Personal{}, &models.Position{}, &models.Status{}, &models.Typecard{}, &models.Bankaccount{}, &models.Card{}, &models.Notification{})
	return db
}
