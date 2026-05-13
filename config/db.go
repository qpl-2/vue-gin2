package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
	"vuegin/global"
)

func initDB() {
	db, err := gorm.Open(sqlite.Open("vue-gin"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	sqlDB.SetMaxIdleConns(Appconfig.Database.MaxTdleConns)
	sqlDB.SetMaxOpenConns(Appconfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour) //每个链接使用一小时后就会关闭

	global.Db = db
}
