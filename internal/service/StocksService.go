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

type HolidayScheduleResponse struct {
	Data [][]string `json:"data"`
}

// 获取股市开盘和收盘日期
func (s *StocksService) GetStockMarketOpeningAndClosingDates(requestAllData bool) ([]string, error) {
	apiURL := fmt.Sprintf("https://www.twse.com.tw/rwd/zh/holidaySchedule/holidaySchedule?response=json&_=%d", time.Now().Unix())

	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %d", response.StatusCode)
	}

	var responseBody HolidayScheduleResponse
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	if len(responseBody.Data) == 0 {
		return nil, fmt.Errorf("response data is empty")
	}

	if !requestAllData {
		var dates []string
		for _, item := range responseBody.Data {
			dates = append(dates, item[0])
		}
		return dates, nil
	}

	// 处理 requestAllData 为 true 的情况
	var dates []string
	for _, item := range responseBody.Data {
		dates = append(dates, item[0])
	}
	return dates, nil
}
