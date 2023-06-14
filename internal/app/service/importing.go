package service

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type XMLRate struct {
	XMLName    xml.Name `xml:"Valute"`
	Code       int      `xml:"NumCode"`
	CharCode   string   `xml:"CharCode"`
	Name       string   `xml:"Name"`
	Multiplier int      `xml:"Nominal"`
	Value      float32  `xml:"Value"`
}

type XMLRates struct {
	XMLName xml.Name  `xml:"ValCurs"`
	DateStr string    `xml:"Date,attr"`
	Date    time.Time `xml:"-"`
	Rates   []XMLRate `xml:"Valute"`
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

func (s *Service) GetNBMRatesXML(date time.Time) (XMLRates, error) {

	var xmlRates XMLRates

	dateFormated := date.Format("02.01.2006")

	url := fmt.Sprintf("http://www.bnm.md/ro/official_exchange_rates?get_xml=1&date=%s", dateFormated)
	resp, err := http.Get(url)
	if err != nil {
		return xmlRates, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return xmlRates, err
	}

	xml.Unmarshal(body, &xmlRates)

	xmlRates.Date, err = time.Parse("02.01.2006", xmlRates.DateStr)
	if err != nil {
		return xmlRates, err
	}

	return xmlRates, nil

}

func (s *Service) GetNBMRates(date time.Time) ([]Rate, error) {

	var rates []Rate

	xmlRates, err := s.GetNBMRatesXML(date)
	if err != nil {
		return nil, err
	}

	for _, xmlRate := range xmlRates.Rates {

		currency := Currency{
			Code:     xmlRate.Code,
			CharCode: xmlRate.CharCode,
			Name:     xmlRate.Name,
		}

		rate := Rate{
			Date:       xmlRates.Date,
			Currency:   currency,
			Multiplier: xmlRate.Multiplier,
			Value:      xmlRate.Value,
		}

		rates = append(rates, rate)
	}

	return rates, nil

}
