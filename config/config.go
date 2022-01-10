package config

import (
	"erajaya/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// inisialisasi database
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	config := os.Getenv("CONNECTION_DB")

	var e error

	DB, e = gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

// automigrate -> for create automation table if table not exist
func InitMigrate() {
	DB.AutoMigrate(&models.Product{})
}

func InitDBTest() {
	config := "root:root@tcp(localhost:3306)/erajaya_test?charset=utf8&parseTime=True&loc=Local"

	var e error

	DB, e = gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&models.Product{})
	DB.AutoMigrate(&models.Product{})
}
