package main

import (
	"crypto/md5"
	"easy-go-iot/user-srv/model"
	"encoding/hex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
)

func genMd5(code string) string {
	md5 := md5.New()
	io.WriteString(md5, code)
	md5.Sum([]byte(code))
	return hex.EncodeToString(md5.Sum(nil))
}

func AutoMigrate() {
	dsn := "root:root@tcp(127.0.0.1:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})
}

func main() {
	AutoMigrate()
}
