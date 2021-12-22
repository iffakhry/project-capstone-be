package config

import (
	"os"

	"final-project/models"

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

// auto migrate -> untuk membuat tabel otomatis jika tabel tidak terdapat pada database
func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.GroupProduct{})
	DB.AutoMigrate(&models.Products{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.Order{})
}

func InitDBTest() {
	//Set data source that will be used
	connection := "root:qwerty@tcp(127.0.0.1:3306)/todo_list_test?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	//Initialize DB session
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	//Migrate the database schema
	InitMigrationTest()
}

//Declare function to auto-migrate the schema
func InitMigrationTest() {
	DB.AutoMigrate(&models.Users{})
}
