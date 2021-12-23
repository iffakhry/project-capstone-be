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
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Error loading .env file")
	// }
	config := "root:nuril123@tcp(127.0.0.1:3306)/db_final_project_testing?charset=utf8&parseTime=True&loc=Local"

	var e error

	DB, e = gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&models.Users{})
	DB.Migrator().DropTable(&models.GroupProduct{})
	DB.Migrator().DropTable(&models.Products{})
	DB.Migrator().DropTable(&models.Payment{})
	DB.Migrator().DropTable(&models.Order{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.GroupProduct{})
	DB.AutoMigrate(&models.Products{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.Order{})
}
