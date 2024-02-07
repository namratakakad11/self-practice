package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbtb struct {
	Id   string `gorm:"primaryKey"`
	Name string `json:"name"`
}

var db *gorm.DB
var err error

func init() {
	dsn := `host=localhost
			user=postgres
			password=postgres
			dbname=postgres
			port=5432
			sslmode=disable`

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	fmt.Println("connection ", db)
	if err != nil {
		return

	}
	db.AutoMigrate(&Dbtb{})
	fmt.Println("created")

}

func main() {

	createDBRecord(db)

}

func createDBRecord(db *gorm.DB) {

	u := Dbtb{Id: "111", Name: "abcd"}
	response := db.Create(&u)
	fmt.Println(response.Error, response.RowsAffected)

}
