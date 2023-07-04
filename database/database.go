package database

import (
	"BNMO/enum"
	gormmodels "BNMO/gorm_models"
	"BNMO/utils"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Initialize() {
	host := os.Getenv("MARIADB_HOST")
	port := os.Getenv("MARIADB_PORT")
	user := os.Getenv("MARIADB_USER")
	pass := os.Getenv("MARIADB_PASS")
	db_name := os.Getenv("MARIADB_DB")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db_name)

	// Connect to database using gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error opening database connection")
	} else {
		fmt.Println("Connected successfully")
	}
	db.AutoMigrate(
		&gormmodels.Base{}, &gormmodels.Account{}, &gormmodels.Admin{}, &gormmodels.Customer{}, &gormmodels.CustomerAddress{}, &gormmodels.Request{}, &gormmodels.Transfer{})

	DB = db
	seed(DB)
}

func seed(db *gorm.DB) {
	pass, _ := utils.HashPassword("password")
	db.Create(&gormmodels.Admin{
		Role: enum.CHIEF,
		Account: gormmodels.Account{
			Email:       "admin@gmail.com",
			Username:    "admin",
			FirstName:   "Super",
			LastName:    "Admin",
			Password:    pass,
			AccountType: enum.ADMIN,
		},
	})
}
