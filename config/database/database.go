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

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束 提升数据库速度
	})

	if err != nil {
		zap.L().Fatal("Database connect failed", zap.Error(err))
	} else {
		zap.L().Info("Database connect success")
	}

	err = autoMigrate(db)
	if err != nil {
		zap.L().Fatal("Database migrate failed", zap.Error(err))
	} else {
		zap.L().Info("Database migrate success")
	}

	DB = db
}
