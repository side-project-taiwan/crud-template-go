package http

import (
	"sample/internal/config"
	"sample/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
)

type API struct {
	App         *fiber.App
	_controller *controller.Controller
}

func NewAPI(cfg *config.Config, app *fiber.App, db *sqlx.DB) *API {

	api := new(API)
	api.App = app

	api.App.Use(logger.New())
	api.App.Use(recover.New())
	api.setupRouter()

	api._controller = controller.NewController(db)

	return api
}

func (api *API) setupRouter() {

	api.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.App.Post("/signup", func(c *fiber.Ctx) error {
		return api._controller.Singup(c)
	})

	api.App.Post("/signin", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.App.Get("/insert", func(c *fiber.Ctx) error {
		return api._controller.Insert(c)
	})

}
