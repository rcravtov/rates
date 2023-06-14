package service

import (
	"time"
)

type IStore interface {
	ImportRates([]Rate) error
	GetCurrencies() ([]Currency, error)
	GetRates(time.Time) ([]Rate, error)
	GetSettingsCount() (int64, error)
	SetSettings(Settings) error
	GetSettings() (Settings, error)
}

type Service struct {
	Store IStore
}

func New(store IStore) *Service {

	service := &Service{Store: store}
	service.CheckCreateDefaultSettings()

	return service
}
