package repository

import (
	"rates/internal/app/service"
	"time"
)

type ImportLogData struct {
	ID          int64     `gorm:"primaryKey"`
	Date        time.Time `gorm:"type:timestamp"`
	DataDate    time.Time `gorm:"type:date"`
	Auto        bool
	Success     bool
	Description string
}

func (s *Store) LogImport(dataDate time.Time, auto bool, success bool, description string) error {

	importData := ImportLogData{
		Date:        time.Now(),
		DataDate:    dataDate,
		Auto:        auto,
		Success:     success,
		Description: description,
	}

	result := s.Client.Save(&importData)
	return result.Error

}

func (s *Store) GetImportLogs(offset int, limit int) (service.ImportLogPage, error) {

	var total int64
	result := s.Client.Model(&ImportLogData{}).Count(&total)
	if result.Error != nil {
		return service.ImportLogPage{}, result.Error
	}

	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	var tmpImportLogData []ImportLogData
	result = s.Client.Offset(offset).Limit(limit).Find(&tmpImportLogData)
	if result.Error != nil {
		return service.ImportLogPage{}, result.Error
	}

	var importData []service.ImportLogData
	for _, tmpData := range tmpImportLogData {

		data := service.ImportLogData{
			Date:        tmpData.Date,
			DataDate:    tmpData.DataDate,
			Auto:        tmpData.Auto,
			Success:     tmpData.Success,
			Description: tmpData.Description,
		}
		importData = append(importData, data)
	}

	page := service.ImportLogPage{
		Total:  int(total),
		Offset: offset,
		Limit:  limit,
		Data:   importData,
	}

	return page, result.Error
}
