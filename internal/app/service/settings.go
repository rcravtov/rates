package service

import (
	"errors"
	"fmt"
	"log"
	"time"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthSettings struct {
type AuthSettings struct {
	Login        string `json:"login"`
	PasswordHash string `json:"-"`
	Token        string `json:"token"`
}

type RawAuthSettings struct {
type RawAuthSettings struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Settings struct {
	AutoImport    bool `json:"auto_import"`
	ImportHours   int  `json:"import_hours"`
	ImportMinutes int  `json:"import_minutes"`
}

type RawSettings struct {
	AutoImport    bool `json:"auto_import"`
	ImportHours   int  `json:"import_hours"`
	ImportMinutes int  `json:"import_minutes"`
}

func (s *Service) CheckCreateDefaultAuthSettings() error {

	authSettingsCount, err := s.Store.GetAuthSettingsCount()
	if err != nil {
		panic(fmt.Errorf("error getting auth settings count: %w", err))
	}

	if authSettingsCount == 0 {

		log.Println("No auth settings found. Creating default auth settings")

		rawAuthSettings := RawAuthSettings{
			Login:    "admin",
			Password: "admin",
		}

		_, err := s.SetAuthSettings(rawAuthSettings)
		if err != nil {
			panic(fmt.Errorf("error creating default auth settings: %w", err))
		}

	}

	return nil
}

func (s *Service) CheckCreateDefaultSettings() error {

	settingsCount, err := s.Store.GetSettingsCount()
	if err != nil {
		panic(fmt.Errorf("error getting settings count: %w", err))
		panic(fmt.Errorf("error getting settings count: %w", err))
	}

	if settingsCount == 0 {

		log.Println("No settings found. Creating default settings")

		rawSettings := RawSettings{
			AutoImport:    false,
			ImportHours:   0,
			ImportMinutes: 0,
			AutoImport:    false,
			ImportHours:   0,
			ImportMinutes: 0,
		}

		_, err := s.SetSettings(rawSettings)
		if err != nil {
			panic(fmt.Errorf("error creating default settings: %w", err))
		}

	} else {

		settings, err := s.GetSettings()
		if err != nil {
			return err
		}

		s.StartImportJobs(settings)


	} else {

		settings, err := s.GetSettings()
		if err != nil {
			return err
		}

		s.StartImportJobs(settings)

	}

	return nil
}

func (s *Service) GetAuthSettings() (AuthSettings, error) {

	return s.Store.GetAuthSettings()

}

func (s *Service) GetAuthSettings() (AuthSettings, error) {

	return s.Store.GetAuthSettings()

}

func (s *Service) GetSettings() (Settings, error) {

	return s.Store.GetSettings()

}

func (s *Service) SetAuthSettings(rawAuthSettings RawAuthSettings) (AuthSettings, error) {

	authSettings := AuthSettings{}

	if len(rawAuthSettings.Login) == 0 {
		return authSettings, errors.New("invalid login")
	}

	if len(rawAuthSettings.Password) == 0 {
		return authSettings, errors.New("invalid password")
	}

	authSettings.Login = rawAuthSettings.Login
	authSettings.PasswordHash = GetPasswordHash(rawAuthSettings.Password)

	err := s.Store.SetAuthSettings(authSettings)
	if err != nil {
		return authSettings, err
	}

	token, err := GenerateToken()
	if err != nil {
		return authSettings, err
	}

	authSettings.Token = token

	return authSettings, nil

}

func (s *Service) SetSettings(rawSettings RawSettings) (Settings, error) {

	settings := Settings{}

	if rawSettings.ImportHours < 0 || rawSettings.ImportHours > 23 {
		return settings, errors.New("invalid import hours value")
	}

	if rawSettings.ImportMinutes < 0 || rawSettings.ImportMinutes > 59 {
		return settings, errors.New("invalid import minutes value")
	}

	settings.AutoImport = rawSettings.AutoImport
	settings.ImportHours = rawSettings.ImportHours
	settings.ImportMinutes = rawSettings.ImportMinutes

	err := s.Store.SetSettings(settings)
	if err != nil {
		return settings, err
	}

	s.StartImportJobs(settings)

	s.StartImportJobs(settings)

	return settings, nil

}

func (s *Service) AuthorizeUser(login string, password string) (string, error) {

	authSettings, err := s.Store.GetAuthSettings()
	authSettings, err := s.Store.GetAuthSettings()
	if err != nil {
		return "", err
	}

	if authSettings.Login != login {
	if authSettings.Login != login {
		return "", errors.New("bad login or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(authSettings.PasswordHash), []byte(password))
	err = bcrypt.CompareHashAndPassword([]byte(authSettings.PasswordHash), []byte(password))
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

func (s *Service) StartImportJobs(settings Settings) error {

	s.Jobs.Stop()

	for _, entry := range s.Jobs.Entries() {
		s.Jobs.Remove(entry.ID)
	}

	if settings.AutoImport {

		spec := fmt.Sprintf("%d %d * * *", settings.ImportMinutes, settings.ImportHours)
		s.Jobs.AddFunc(spec, s.AutoImport)
		s.Jobs.Start()

		log.Printf("Auto import is set to %02d:%02d\n", settings.ImportHours, settings.ImportMinutes)

	}

	return nil

}

func (s *Service) AutoImport() {

	date := time.Now().AddDate(0, 0, 1)
	log.Println("Auto import started, data date:", date)
	s.ImportRates(date, true)

}

func (s *Service) StartImportJobs(settings Settings) error {

	s.Jobs.Stop()

	for _, entry := range s.Jobs.Entries() {
		s.Jobs.Remove(entry.ID)
	}

	if settings.AutoImport {

		spec := fmt.Sprintf("%d %d * * *", settings.ImportMinutes, settings.ImportHours)
		s.Jobs.AddFunc(spec, s.AutoImport)
		s.Jobs.Start()

		log.Printf("Auto import is set to %02d:%02d\n", settings.ImportHours, settings.ImportMinutes)

	}

	return nil

}

func (s *Service) AutoImport() {

	date := time.Now().AddDate(0, 0, 1)
	log.Println("Auto import started, data date:", date)
	s.ImportRates(date, true)

}
