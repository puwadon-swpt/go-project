package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	// อ่านค่าจาก environment variables
	MYSQL_HOST := os.Getenv("MYSQL_HOST") // ใช้ MYSQL_HOST จาก environment variables
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DB := os.Getenv("MYSQL_DB")

	DSN := MYSQL_USER + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_HOST + ":3306)/" + MYSQL_DB + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	log.Println("Successfully connected to the database")
}
