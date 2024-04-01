package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sample/internal/repository/model"
	"sample/internal/util"
	"time"
)

type UserService struct {
	structInterface i_UserRepository
}

type i_UserRepository interface {
	Signup(data *model.User) error
}

func NewUserService(varName i_UserRepository) *UserService {
	util.PrintLogWithColor("Enter stockService", "#00ffff")

	callBack := UserService{
		structInterface: varName,
	}
	return &callBack
}

func (s *UserService) GetDailyClosingQuote() ([]byte, error) {
	apiURL := "https://www.twse.com.tw/rwd/zh/afterTrading/MI_INDEX?response=json&_=1709118194485"
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read HTTP response body: %v", err)
		}
		return body, nil
	}

	return nil, fmt.Errorf("HTTP request failed with status: %d", response.StatusCode)
}

func (s *UserService) GetStockMarketOpeningAndClosingDates(requestAllData bool) ([]string, error) {

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

func (s *UserService) GetTheLatestOpeningDate() (string, error) {
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
