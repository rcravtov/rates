package service

import "time"

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
