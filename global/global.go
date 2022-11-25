package global

import (
	"easy-go-iot/user-srv/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)

//func init() {
//	dsn := "root:root@tcp(127.0.0.1:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags),
//		logger.Config{
//			SlowThreshold: time.Second, // 慢 SQL 阈值
//			LogLevel:      logger.Info, // log level
//			Colorful:      true,        // 禁用彩色打印
//		},
//	)
//	var err error
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true,
//		},
//		Logger: newLogger,
//	})
//
//	if err != nil {
//		panic(err)
//	}
//}
