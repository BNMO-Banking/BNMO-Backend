package database

import (
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("Error opening database connection")
	} else {
		fmt.Println("Connected successfully")
	}
	db.AutoMigrate(
		&gormmodels.Base{}, &gormmodels.Account{}, &gormmodels.Admin{}, &gormmodels.Customer{}, &gormmodels.CustomerAddress{}, &gormmodels.Request{}, &gormmodels.Transfer{})

	DB = db
	// seed(DB)
}

func seed(db *gorm.DB) {
	var accounts []models.Account
	admin := models.Account{
		IsAdmin:       sql.NullBool{Bool: true, Valid: true},
		AccountStatus: sql.NullString{String: "accepted", Valid: true},
		FirstName:     "Admin",
		LastName:      "Admin",
		Email:         "admin@gmail.com",
		Username:      "admin",
		ImagePath:     "./images/Admin.png",
		AccountNumber: "1",
		Balance:       0,
	}
	admin.SetPassword("admin")

	user1 := models.Account{
		IsAdmin:       sql.NullBool{Bool: false, Valid: true},
		AccountStatus: sql.NullString{String: "accepted", Valid: true},
		FirstName:     "John",
		LastName:      "Doe",
		Email:         "johndoe@gmail.com",
		Username:      "johndoe123",
		ImagePath:     "./images/johndoe.png",
		AccountNumber: "100-100-1000",
		Balance:       0,
	}
	user1.SetPassword("user1")

	user2 := models.Account{
		IsAdmin:       sql.NullBool{Bool: false, Valid: true},
		AccountStatus: sql.NullString{String: "accepted", Valid: true},
		FirstName:     "Sarah",
		LastName:      "Baker",
		Email:         "sarahbaker@gmail.com",
		Username:      "sarahbaker123",
		ImagePath:     "./images/sarahbaker.png",
		AccountNumber: "100-100-1001",
		Balance:       0,
	}
	user2.SetPassword("user2")

	user3 := models.Account{
		IsAdmin:       sql.NullBool{Bool: false, Valid: true},
		AccountStatus: sql.NullString{String: "accepted", Valid: true},
		FirstName:     "Sam",
		LastName:      "Smith",
		Email:         "samsmith@gmail.com",
		Username:      "samsmith123",
		ImagePath:     "./images/samsmith.png",
		AccountNumber: "100-100-1002",
		Balance:       0,
	}
	user3.SetPassword("user3")

	accounts = append(accounts, admin, user1, user2, user3)

	db.Create(&accounts)
}
