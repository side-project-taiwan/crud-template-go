package controller

import (
	"encoding/json"
	"net/http"

	"sample/internal/config"
	"sample/internal/model"
	"sample/internal/service"
	"sample/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type StockController struct {
	//FiberAppstruct  *fiber.App
	ginEngineStruct *gin.Engine
	Service         *service.Service
	StocksService   *service.StocksService
}

func InitializeController(cfg *config.Config, _ginEngine *gin.Engine, db *sqlx.DB) *StockController {
	sc := &StockController{
		//FiberAppstruct:  _fiberApp,
		ginEngineStruct: _ginEngine,
		Service:         service.NewService(db),
		StocksService:   service.InitStocksService(db),
	}

	// sc.FiberAppstruct.Use(logger.New())
	// sc.FiberAppstruct.Use(recover.New())
	sc.setupRoutes()
	return sc
}

func (sc *StockController) setupRoutes() {
	sc.ginEngineStruct.GET("/ping", func(_ginCTX *gin.Context) {
		_ginCTX.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	sc.ginEngineStruct.GET("/", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "Hello, World! ----index 20240326")
	})

	sc.ginEngineStruct.GET("/getSignin", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "test signin success")
	})

	sc.ginEngineStruct.GET("/dailyClosingQuote", sc.dailyClosingQuote)

	sc.ginEngineStruct.GET("/theLatestOpeningDate", sc.theLatestOpeningDate)

	sc.ginEngineStruct.GET("/insert", sc.Insert)

	// 注册 POST 路由
	sc.ginEngineStruct.POST("/signup", sc.Signup)

	sc.ginEngineStruct.GET("/stockMarketOpeningAndClosingDates", sc.GetStockMarketOpeningAndClosingDates)

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

func (sc *StockController) Signup(_ginCTX *gin.Context) {
	// 处理注册请求
	req := new(model.SignupRequest)
	if err := _ginCTX.BindJSON(req); err != nil {
		_ginCTX.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sc.Service.SignupService(req); err != nil {
		_ginCTX.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回响应
	res := &model.SignupResponse{
		Email: req.Email,
		Name:  req.Name,
	}
	_ginCTX.JSON(http.StatusOK, res)
}

func (sc *StockController) Insert(c *gin.Context) {
	err := sc.Service.InsertService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Insertion successful"})
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
