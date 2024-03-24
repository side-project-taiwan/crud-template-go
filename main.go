package main

import (
	"log"
	"sample/internal/config"
	"sample/internal/controller"
	"sample/internal/util"

	"github.com/gofiber/fiber/v2"
)

func main() {

	_config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := util.NewDB(_config)
	if err != nil {
		log.Fatal(err)
	}

	f := fiber.New()
	//api := http.NewAPI(_config, f, db)

	api2 := controller.InitializeController(_config, f, db)

	log.Fatal(api2.App.Listen(_config.PORT))
}
