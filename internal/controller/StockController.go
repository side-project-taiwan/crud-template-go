package controller

import (
	"sample/internal/config"
	"sample/internal/model"
	"sample/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
)

type StockController struct {
	App     *fiber.App
	Service *service.Service
}

func NewStockController(cfg *config.Config, app *fiber.App, db *sqlx.DB) *StockController {

	// 初始化 StockController，并添加 Fiber 中间件
	sc := &StockController{
		App:     app,
		Service: service.NewService(db),
	}

	sc.App.Use(logger.New())
	sc.App.Use(recover.New())
	sc.setupRoutes()

	return sc
}

func (sc *StockController) setupRoutes() {
	// 设置路由处理程序
	sc.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! indextest___23423131")
	})

	sc.App.Get("/getSignin", func(c *fiber.Ctx) error {
		return c.SendString("test signin success")
	})

	sc.App.Post("/signup", sc.Signup)
	sc.App.Post("/signin", func(c *fiber.Ctx) error {
		return c.SendString("test signin success")
	})

	sc.App.Get("/insert", sc.Insert)
}

func (sc *StockController) Signup(c *fiber.Ctx) error {
	// 处理注册请求
	req := new(model.SignupRequest)
	if err := c.BodyParser(req); err != nil {
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
	return c.JSON(res)
}

func (sc *StockController) Insert(c *fiber.Ctx) error {
	return sc.Service.InsertService()
}
