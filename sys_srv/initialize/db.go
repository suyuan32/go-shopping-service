//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"sys_srv/global"
)

func InitGorm() {
	cfg := global.ServerConfig.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/go_shopping_sys_srv?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port)
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
		zap.S().Fatal(global.GetMessage("FailStartGorm"), " err: ", err.Error())
	}

	global.DB = db
}
