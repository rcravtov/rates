package repository

import "rates/internal/app/service"

type Settings struct {
	ID            int `gorm:"primaryKey;autoIncrement:false"`
	AutoImport    bool
	ImportHours   int
	ImportMinutes int
}

type AuthSettings struct {
	ID           int `gorm:"primaryKey;autoIncrement:false"`
	Login        string
	PasswordHash string
}

func (s *Store) GetSettingsCount() (int64, error) {

	var count int64

	result := s.Client.Model(&Settings{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (s *Store) SetSettings(serviceSettings service.Settings) error {

	settings := Settings{
		ID:            1,
		AutoImport:    serviceSettings.AutoImport,
		ImportHours:   serviceSettings.ImportHours,
		ImportMinutes: serviceSettings.ImportMinutes,
	}

	result := s.Client.Save(&settings)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetSettings() (service.Settings, error) {

	var settings Settings
	result := s.Client.Take(&settings)
	if result.Error != nil {
		return service.Settings{}, result.Error
	}

	serviceSettings := service.Settings{
		AutoImport:    settings.AutoImport,
		ImportHours:   settings.ImportHours,
		ImportMinutes: settings.ImportMinutes,
	}

	return serviceSettings, nil

}

func (s *Store) GetAuthSettingsCount() (int64, error) {

	var count int64

	result := s.Client.Model(&AuthSettings{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (s *Store) SetAuthSettings(serviceAuthSettings service.AuthSettings) error {

	authSettings := AuthSettings{
		ID:           1,
		Login:        serviceAuthSettings.Login,
		PasswordHash: serviceAuthSettings.PasswordHash,
	}

	result := s.Client.Save(&authSettings)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetAuthSettings() (service.AuthSettings, error) {

	var authSettings AuthSettings
	result := s.Client.Take(&authSettings)
	if result.Error != nil {
		return service.AuthSettings{}, result.Error
	}

	serviceAuthSettings := service.AuthSettings{
		Login:        authSettings.Login,
		PasswordHash: authSettings.PasswordHash,
	}

	return serviceAuthSettings, nil

}
