package DB

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "ACCOUNT:PASSWORD@tcp(HOST:3306)/DB?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println("資料庫連接失敗")
	}
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Minute)
	return db
}
