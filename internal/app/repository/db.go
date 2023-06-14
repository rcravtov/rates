package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	Client *gorm.DB
}

func New() (*Store, error) {

	var err error
	store := &Store{}

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_DB"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	store.Client, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return store, err
	}

	err = store.Client.AutoMigrate(&Currency{}, &Rate{}, &Settings{})
	if err != nil {
		return store, err
	}

	return store, nil

}
