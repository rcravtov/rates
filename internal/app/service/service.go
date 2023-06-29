package service

import (
	"time"

	"github.com/robfig/cron/v3"
)

type IStore interface {
	ImportRates([]Rate) error
	GetCurrencies() ([]Currency, error)
	GetRates(time.Time) ([]Rate, error)
	GetSettingsCount() (int64, error)
	SetSettings(Settings) error
	GetSettings() (Settings, error)
	GetAuthSettingsCount() (int64, error)
	SetAuthSettings(AuthSettings) error
	GetAuthSettings() (AuthSettings, error)
	LogImport(time.Time, bool, bool, string) error
	GetImportLogs(int, int) (ImportLogPage, error)
}

type Service struct {
	Store IStore
	Jobs  *cron.Cron
}

func New(store IStore, cron *cron.Cron) *Service {

	service := &Service{
		Store: store,
		Jobs:  cron,
	}

	service.CheckCreateDefaultAuthSettings()
	service.CheckCreateDefaultSettings()

	return service
}
