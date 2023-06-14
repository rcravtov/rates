package repository

import "rates/internal/app/service"

type Settings struct {
	ID           int `gorm:"primaryKey;autoIncrement:false"`
	Login        string
	PasswordHash string
}

func ConvertSettingsToServiceSettings(s Settings) service.Settings {

	return service.Settings{
		Login:        s.Login,
		PasswordHash: s.PasswordHash,
	}

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
		ID:           1,
		Login:        serviceSettings.Login,
		PasswordHash: serviceSettings.PasswordHash,
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

	serviceSettings := ConvertSettingsToServiceSettings(settings)
	return serviceSettings, nil

}
