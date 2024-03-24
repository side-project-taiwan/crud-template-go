package controller

import (
	"sample/internal/config"
	//"sample/internal/controller/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
)

type StockController struct {
	App        *fiber.App
	Controller *Controller
}

func NewStockController(cfg *config.Config, _fiberApp *fiber.App, db *sqlx.DB) *StockController {
	stockController := &StockController{
		App: _fiberApp,
	}

	stockController.App.Use(logger.New())
	stockController.App.Use(recover.New())
	stockController.setupRouter()

	stockController.Controller = NewController(db)

	return stockController
}

func (_stockController *StockController) setupRouter() {
	_stockController.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! indextest test1")
	})
	_stockController.App.Get("/getSignin", func(c *fiber.Ctx) error {
		return c.SendString("test signin success")
	})

	_stockController.App.Post("/signup", func(c *fiber.Ctx) error {
		return _stockController.Controller.Singup(c)
	})

	_stockController.App.Post("/signin", func(c *fiber.Ctx) error {
		return c.SendString("test signin success")
	})

	_stockController.App.Get("/insert", func(c *fiber.Ctx) error {
		return _stockController.Controller.Insert(c)
	})
}
