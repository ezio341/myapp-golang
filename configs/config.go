package configs

import (
	"fmt"
	roleDB "myproject/models/role/database"
	userDB "myproject/models/user/database"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TARGET"),
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("database initialization failed")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(userDB.User{})
	DB.AutoMigrate(roleDB.Role{})
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Failed load env file")
	}
}

// func CustomErrorHandler(next echo.HandlerFunc) echo.HandlerFunc{
// 	return
// }
