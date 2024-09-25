package database

import (
	"ConfessionWall/config/config"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	database := config.Config.GetString("database.database")
	host := config.Config.GetString("database.host")
	port := config.Config.GetString("database.port")
	user := config.Config.GetString("database.user")
	password := config.Config.GetString("database.password")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		zap.L().Fatal("Database connect failed", zap.Error(err))
	}

	err = autoMigrate(db)
	if err != nil {
		zap.L().Fatal("Database migrate failed", zap.Error(err))
	}

	DB = db
}
