package service

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Settings struct {
	Login        string `json:"login"`
	PasswordHash string `json:"-"`
}

type RawSettings struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Service) CheckCreateDefaultSettings() error {

	settingsCount, err := s.Store.GetSettingsCount()
	if err != nil {
		panic(fmt.Errorf("error getting user count: %w", err))
	}

	if settingsCount == 0 {

		log.Println("No settings found. Creating default settings")

		rawSettings := RawSettings{
			Login:    "admin",
			Password: "admin",
		}

		_, err := s.SetSettings(rawSettings)
		if err != nil {
			panic(fmt.Errorf("error creating default settings: %w", err))
		}
	}

	return nil
}

func (s *Service) GetSettings() (Settings, error) {

	return s.Store.GetSettings()

}

func (s *Service) SetSettings(rawSettings RawSettings) (Settings, error) {

	settings := Settings{
		Login:        rawSettings.Login,
		PasswordHash: GetPasswordHash(rawSettings.Password),
	}

	err := s.Store.SetSettings(settings)
	if err != nil {
		return settings, err
	}

	return settings, nil

}

func (s *Service) AuthorizeUser(login string, password string) (string, error) {

	settings, err := s.Store.GetSettings()
	if err != nil {
		return "", err
	}

	if settings.Login != login {
		return "", errors.New("bad login or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(settings.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := GenerateToken()
	if err != nil {
		return "", err
	}

	return token, nil

}

func GetPasswordHash(password string) string {

	bhash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bhash)

}
