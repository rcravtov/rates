package repository

import (
	"rates/internal/app/service"
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	Code     int `gorm:"primaryKey;autoIncrement:false"`
	CharCode string
	Name     string
}

type Rate struct {
	Date         time.Time `gorm:"primaryKey;type:date"`
	CurrencyCode int       `gorm:"primaryKey"`
	Currency     Currency  `gorm:"foreignKey:CurrencyCode"`
	Multiplier   int       `gorm:"type:decimal(10,0)"`
	Value        float32   `gorm:"type:decimal(10,4)"`
	Change       float32   `gorm:"-"`
}

func (store *Store) ImportRates(rates []service.Rate) error {

	for _, serviceRate := range rates {

		rate := ConvertServiceRateToRate(serviceRate)

		filter := Rate{Date: rate.Date, Currency: rate.Currency}
		queryDB := store.Client.Session(&gorm.Session{FullSaveAssociations: true})
		queryDB.Where(filter)
		queryDB.Assign(rate).FirstOrCreate(&rate)

	}

	return nil

}

func (store *Store) GetCurrencies() ([]service.Currency, error) {

	var currencies []Currency
	store.Client.Order("char_code").Find(&currencies)

	var result []service.Currency
	localCurrency := service.GetLocalCurrency()
	result = append(result, localCurrency)

	for _, currency := range currencies {
		serviceCurrency := ConvertCurrencyToServiceCurrency(currency)
		result = append(result, serviceCurrency)
	}

	return result, nil

}

func (store *Store) GetRates(date time.Time) ([]service.Rate, error) {

	// filter := Rate{Date: date}
	// var rates []Rate
	// result := store.Client.Joins("Currency").Order("char_code").Find(&rates, filter)

	var rates []service.Rate
	localRate := service.Rate{
		Date:       date,
		Multiplier: 1,
		Value:      1,
		Change:     0,
		Currency:   service.GetLocalCurrency(),
	}
	rates = append(rates, localRate)

	type Row struct {
		Date         time.Time
		CurrencyCode int
		Multiplier   int
		Value        float32
		Change       float32
		CharCode     string
		Name         string
	}

	var rows []Row
	previousDate := date.AddDate(0, 0, -1)

	query := `SELECT 
				rates.date, 
				rates.currency_code, 
				rates.multiplier,
				rates.value,
				currencies.char_code,
				currencies.name,
				COALESCE(rates.value - previous_rates.value * rates.multiplier / previous_rates.multiplier, 0.0000)::numeric(10,4) AS change
			FROM rates
				LEFT JOIN currencies ON rates.currency_code = currencies.code
				LEFT JOIN rates AS previous_rates ON 
					rates.currency_code = previous_rates.currency_code AND
					previous_rates.date = ?
					
			WHERE
				rates.date = ?
			ORDER BY
				currencies.char_code`

	store.Client.Raw(query, previousDate, date).Scan(&rows)

	for _, row := range rows {
		rate := service.Rate{
			Date:       row.Date,
			Multiplier: row.Multiplier,
			Value:      row.Value,
			Change:     row.Change,
			Currency: service.Currency{
				Code:     row.CurrencyCode,
				CharCode: row.CharCode,
				Name:     row.Name,
			},
		}
		rates = append(rates, rate)
	}

	return rates, nil

}

func ConvertServiceCurrencyToCurrency(currency service.Currency) Currency {

	result := Currency{
		Code:     currency.Code,
		CharCode: currency.CharCode,
		Name:     currency.Name,
	}

	return result

}

func ConvertServiceRateToRate(rate service.Rate) Rate {

	currency := ConvertServiceCurrencyToCurrency(rate.Currency)

	result := Rate{
		Date:         rate.Date,
		CurrencyCode: currency.Code,
		Currency:     currency,
		Multiplier:   rate.Multiplier,
		Value:        rate.Value,
	}

	return result

}

func ConvertCurrencyToServiceCurrency(currency Currency) service.Currency {

	result := service.Currency{
		Code:     currency.Code,
		CharCode: currency.CharCode,
		Name:     currency.Name,
	}

	return result

}

func ConvertRateToServiceRate(rate Rate) service.Rate {

	currency := ConvertCurrencyToServiceCurrency(rate.Currency)

	result := service.Rate{
		Date:       rate.Date,
		Currency:   currency,
		Multiplier: rate.Multiplier,
		Value:      rate.Value,
	}

	return result

}
