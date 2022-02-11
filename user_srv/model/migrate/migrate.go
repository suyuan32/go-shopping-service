package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"user_srv/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	db := Connect()
	err := db.AutoMigrate(&model.User{}, &model.UserAddress{})
	if err != nil {
		fmt.Println("error:", err.Error())
	}

}

func Connect() *gorm.DB {
	dsn := "ryan:123@tcp(127.0.0.1:3306)/go_shopping_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "go_shopping_",
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		panic("mysql connect fail")
	}
	return db
}
