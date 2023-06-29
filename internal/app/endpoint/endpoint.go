package endpoint

import (
	"net/http"
	"strconv"
	"time"

	"rates/internal/app/service"

	"github.com/gin-gonic/gin"
)

type IService interface {
	GetNBMRates(time.Time) ([]service.Rate, error)
	ImportRates(time.Time, bool) ([]service.Rate, error)
	GetCurrencies() ([]service.Currency, error)
	GetRates(time.Time) ([]service.Rate, error)
	GetSettings() (service.Settings, error)
	SetSettings(service.RawSettings) (service.Settings, error)
	GetAuthSettings() (service.AuthSettings, error)
	SetAuthSettings(service.RawAuthSettings) (service.AuthSettings, error)
	AuthorizeUser(string, string) (string, error)
	GetImportLogs(int, int) (service.ImportLogPage, error)
}

type Endpoint struct {
	service IService
}

func New(s IService) *Endpoint {
	return &Endpoint{
		service: s,
	}
}

func (e *Endpoint) Status(c *gin.Context) {

	c.String(http.StatusOK, "OK")

}

func (e *Endpoint) ImportRates(c *gin.Context) {

	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	rates, err := e.service.ImportRates(date, false)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, rates)

}

func (e *Endpoint) GetCurrencies(c *gin.Context) {

	currencies, err := e.service.GetCurrencies()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, currencies)

}

func (e *Endpoint) GetRates(c *gin.Context) {

	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	rates, err := e.service.GetRates(date)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, rates)

}

func (e *Endpoint) GetSettings(c *gin.Context) {

	settings, err := e.service.GetSettings()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, settings)

}

func (e *Endpoint) SetSettings(c *gin.Context) {

	var input service.RawSettings
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	settings, err := e.service.SetSettings(input)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, settings)

}

func (e *Endpoint) GetAuthSettings(c *gin.Context) {

	settings, err := e.service.GetAuthSettings()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, settings)

}

func (e *Endpoint) SetAuthSettings(c *gin.Context) {

	var input service.RawAuthSettings
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	settings, err := e.service.SetAuthSettings(input)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, settings)

}

func (e *Endpoint) Auth(c *gin.Context) {

	type AuthInput struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	type AuthResponse struct {
		Token string `json:"token"`
	}

	var input AuthInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := e.service.AuthorizeUser(input.Login, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	response := AuthResponse{
		Token: token,
	}

	c.JSON(http.StatusOK, response)

}

func (e *Endpoint) GetImportLogs(c *gin.Context) {

	offsetStr := c.Query("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 0
	}

	logs, err := e.service.GetImportLogs(offset, limit)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, logs)

}
