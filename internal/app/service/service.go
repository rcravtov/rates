package service

import "time"

type IStore interface {
	ImportRates([]Rate) error
	GetCurrencies() ([]Currency, error)
	GetRates(time.Time) ([]Rate, error)
}

type Service struct {
	Store IStore
}

type Currency struct {
	Code     int
	CharCode string
	Name     string
}

type Rate struct {
	Date       time.Time
	Currency   Currency
	Multiplier int
	Value      float32
	Change     float32
}

func New(store IStore) *Service {
	return &Service{Store: store}
}

func (s *Service) ImportRates(date time.Time) ([]Rate, error) {

	rates, err := s.GetNBMRates(date)
	if err != nil {
		return nil, err
	}

	err = s.Store.ImportRates(rates)
	if err != nil {
		return nil, err
	}

	return rates, nil

}

func (s *Service) GetCurrencies() ([]Currency, error) {

	currencies, err := s.Store.GetCurrencies()
	if err != nil {
		return nil, err
	}

	return currencies, nil

}

func (s *Service) GetRates(date time.Time) ([]Rate, error) {

	rates, err := s.Store.GetRates(date)
	if err != nil {
		return nil, err
	}

	return rates, nil

}

func GetLocalCurrency() Currency {

	return Currency{
		Code:     498,
		CharCode: "MDL",
		Name:     "Leu moldovenesc",
	}

}
