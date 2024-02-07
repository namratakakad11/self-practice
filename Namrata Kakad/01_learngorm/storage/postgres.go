package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
	Host     string
}

func NewConnection(config *Config) (*gorm.DB, error) {

	dsn := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"

		// fmt.Sprintf(
		// 	"user=%s password=%s  dbname=%s host=%s port=%s sslmode=%s ",

		// 	config.User, config.Password, config.DBName, config.Host, config.Port, config.SSLMode,
		// )
	// fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.User, config.Password, config.Host, config.Port, config.DBName)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("Error : ", err)
		return db, err
	}
	return db, nil

}
