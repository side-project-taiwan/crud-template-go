package routes

import (
	"net/http"
	"sample/internal/util"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Routes *routes
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
