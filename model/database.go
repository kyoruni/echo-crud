package model

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
}

func GetDB() (*gorm.DB, error) {
	userName := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(host.docker.internal:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, pass, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
