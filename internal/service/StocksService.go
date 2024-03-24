package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sample/internal/repository"
	"time"

	"github.com/jmoiron/sqlx"
)

type StocksService struct {
	R *repository.Repository
}

func InitStocksService(db *sqlx.DB) *StocksService {
	return &StocksService{
		R: repository.NewRepository(db),
	}
}

func (s *StocksService) GetStockMarketOpeningAndClosingDates(requestAllData bool) ([]string, error) {
	type HolidayScheduleResponse struct {
		Data [][]string `json:"data"`
	}
	apiURL := fmt.Sprintf("https://www.twse.com.tw/rwd/zh/holidaySchedule/holidaySchedule?response=json&_=%d", time.Now().Unix())

	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %d", response.StatusCode)
	}

	var responseBody HolidayScheduleResponse
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Extract the dates from the JSON response
	var dates []string
	for _, item := range responseBody.Data {
		dates = append(dates, item[0])
	}

	return dates, nil
}

// GetTheLatestOpeningDate 返回最新的开盘日期
func (s *StocksService) GetTheLatestOpeningDate() (string, error) {
	responseClosingDates, err := s.GetStockMarketOpeningAndClosingDates(false)
	if err != nil {
		return "", err
	}

	currentDate := time.Now()
	if currentDate.Hour() < 20 {
		currentDate = currentDate.AddDate(0, 0, -1)
	}

	for currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday || contains(responseClosingDates, currentDate.Format("2006-01-02")) {
		currentDate = currentDate.AddDate(0, 0, -1)
	}

	return currentDate.Format("20060102"), nil
}
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
