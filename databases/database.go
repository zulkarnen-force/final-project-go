package databases

import (
	"final-project-go/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

var (
	HOSTNAME = "localhost"
	USER     = "postgres"
	PASSWORD = "root"
	DBPORT   = "5432"
	DBNAME   = "final_go"

	db *gorm.DB
	err error
)

func Migrate() {
	ConnectionDB()
	db.Debug().AutoMigrate(entity.User{}, entity.Photo{}, entity.SocialMedia{}, entity.Comment{})
}

func ConnectionDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s, dbname=%s sslmode=disable", 
                HOSTNAME, USER, PASSWORD, DBPORT, DBNAME)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db

}


func GetDB() *gorm.DB {
	return db
}