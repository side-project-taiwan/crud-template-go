package controller

import (
	"encoding/json"
	"fmt"
	"sample/internal/config"
	"sample/internal/model"
	"sample/internal/service"
	"sample/internal/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
)

type StockController struct {
	App           *fiber.App
	Service       *service.Service
	StocksService *service.StocksService
}

func InitializeController(cfg *config.Config, app *fiber.App, db *sqlx.DB) *StockController {

	// 初始化 StockController，并添加 Fiber 中间件
	sc := &StockController{
		App:           app,
		Service:       service.NewService(db),
		StocksService: service.InitStocksService(db),
	}

	sc.App.Use(logger.New())
	sc.App.Use(recover.New())
	sc.setupRoutes()

	return sc
}

func (sc *StockController) setupRoutes() {
	sc.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! ----index 20240326")
	})

	sc.App.Get("/getSignin", func(c *fiber.Ctx) error {
		return c.SendString("test signin success")
	})

	sc.App.Get("/stockMarketOpeningAndClosingDates", sc.GetStockMarketOpeningAndClosingDates)

	sc.App.Get("/dailyClosingQuote", sc.dailyClosingQuote)

	sc.App.Get("/theLatestOpeningDate", sc.theLatestOpeningDate)

	sc.App.Post("/signup", sc.Signup)

	sc.App.Get("/insert", sc.Insert)
}

func (sc *StockController) Signup(ctxStruct *fiber.Ctx) error {
	// 处理注册请求
	req := new(model.SignupRequest)
	if err := ctxStruct.BodyParser(req); err != nil {
		return err
	}

	if err := sc.Service.SignupService(req); err != nil {
		return err
	}

	// 返回响应
	res := &model.SignupResponse{
		Email: req.Email,
		Name:  req.Name,
	}
	return ctxStruct.JSON(res)
}

func (sc *StockController) Insert(c *fiber.Ctx) error {
	return sc.Service.InsertService()
}

func (sc *StockController) theLatestOpeningDate(ctx *fiber.Ctx) error {
	util.PrintLog("Enter GetTheLatestOpeningDate log")

	dates, err := sc.StocksService.GetTheLatestOpeningDate()
	if err != nil {
		return err
	}

	// 构造 JSON 响应
	response := struct {
		Dates string `json:"dates"`
	}{
		Dates: dates,
	}

	// 返回 JSON 响应
	return ctx.JSON(response)
}

func (sc *StockController) dailyClosingQuote(ctx *fiber.Ctx) error {
	//util.PrintLog("This is a GetDailyClosingQuote log", true)
	util.PrintLog("dailyClosingQuote")
	dailyQuote, err := sc.StocksService.GetDailyClosingQuote()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(dailyQuote, &jsonResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}

	return ctx.JSON(jsonResponse)
}
func (sc *StockController) GetStockMarketOpeningAndClosingDates(ctx *fiber.Ctx) error {
	util.PrintLog("Enter GetStockMarketOpeningAndClosingDates log")

	dates, err := sc.StocksService.GetStockMarketOpeningAndClosingDates(true)
	if err != nil {
		return err
	}

	response := struct {
		Dates []string `json:"dates"`
	}{
		Dates: dates,
	}
	return ctx.JSON(response)
}
