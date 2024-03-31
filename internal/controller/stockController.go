package controller

import (
	"encoding/json"
	"net/http"
	"sample/internal/service"
	"sample/internal/util"

	"github.com/gin-gonic/gin"
)

type StockController struct {
	StocksService *service.StocksService
}

func InitializeStockController(stocksService *service.StocksService, gin_Instance *gin.Engine) *StockController {

	util.PrintLogWithColor("Enter InitializeStockController", "#ff0000")
	sc := &StockController{
		StocksService: stocksService,
	}
	sc.setAssignRoutes(gin_Instance)
	return sc
}

func (sc *StockController) setAssignRoutes(gin_Instance *gin.Engine) {

	util.PrintLogWithColor("Enter setAssignRoutes", "#00F500")

	//gin_Instance := gin.Default()
	gin_Instance.GET("/", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "Hello, World! ----index 20240326")
	})
	gin_Instance.GET("/ping", func(_ginCTX *gin.Context) {
		_ginCTX.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	gin_Instance.GET("/getSignin", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "test signin success")
	})

	gin_Instance.GET("/dailyClosingQuote", sc.dailyClosingQuote)

	gin_Instance.GET("/theLatestOpeningDate", sc.theLatestOpeningDate)

	gin_Instance.GET("/stockMarketOpeningAndClosingDates", sc.GetStockMarketOpeningAndClosingDates)

}

func (sc *StockController) GetStockMarketOpeningAndClosingDates(_ginCTX *gin.Context) {
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
	util.PrintLogWithColor("Enter GetTheLatestOpeningDate log")

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
