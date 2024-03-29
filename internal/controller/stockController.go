package controller

import (
	"encoding/json"
	"net/http"
	"sample/internal/service"
	"sample/internal/util"

	"github.com/gin-gonic/gin"
)

type StockController struct {
	ginEngineStruct *gin.Engine
	StocksService   *service.StocksService
}

func InitializeController() *StockController {
	sc := &StockController{
		StocksService: &service.StocksService{},
	}
	sc.setStockRoutes()
	//setStockRoutes()
}
func (sc *StockController) setStockRoutes() {
	util.PrintLog("Enter setupRoutes log")

	gin_Instance := service.GetHttpServiceInstance().Gin_Instance
	gin_Instance.GET("/ping", func(_ginCTX *gin.Context) {
		_ginCTX.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	gin_Instance.GET("/", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "Hello, World! ----index 20240326")
	})

	gin_Instance.GET("/getSignin", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "test signin success")
	})

	gin_Instance.GET("/dailyClosingQuote", sc.dailyClosingQuote)

	gin_Instance.GET("/theLatestOpeningDate", sc.theLatestOpeningDate)

	gin_Instance.GET("/stockMarketOpeningAndClosingDates", sc.GetStockMarketOpeningAndClosingDates)

}

func (sc *StockController) GetStockMarketOpeningAndClosingDates(_ginCTX *gin.Context) {
	util.PrintLog("Enter GetStockMarketOpeningAndClosingDates log")

	dates, err := sc.StocksService.GetStockMarketOpeningAndClosingDates(true)
	if err != nil {
		_ginCTX.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := struct {
		Dates []string `json:"dates"`
	}{
		Dates: dates,
	}
	_ginCTX.JSON(http.StatusOK, response)
}

func (sc *StockController) theLatestOpeningDate(_ginCTX *gin.Context) {
	util.PrintLog("Enter GetTheLatestOpeningDate log")

	dates, err := sc.StocksService.GetTheLatestOpeningDate()
	if err != nil {
		_ginCTX.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := struct {
		Dates string `json:"dates"`
	}{
		Dates: dates,
	}

	_ginCTX.JSON(http.StatusOK, response)

}

func (sc *StockController) dailyClosingQuote(c *gin.Context) {
	util.PrintLog("dailyClosingQuote")

	dailyQuote, err := sc.StocksService.GetDailyClosingQuote()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var jsonResponse interface{}
	if err := json.Unmarshal(dailyQuote, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}
