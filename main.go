package main

import (
	"log"
	"sample/internal/config"
	"sample/internal/controller"
	"sample/internal/util"

	"github.com/gin-gonic/gin"
)

func main() {

	_config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := util.NewDB(_config) //
	if err != nil {
		log.Fatal(err)
	}

	ginInstance := gin.Default()

	_ = controller.InitializeController(_config, ginInstance, db)

	if err := ginInstance.Run(":" + _config.PORT); err != nil {
		panic(err)
	}

	//log.Fatal(api2.FiberAppstruct.Listen(_config.PORT))
	//f := fiber.New()
	//api := http.NewAPI(_config, f, db)
	//log.Fatal(api2.FiberAppstruct.Listen(_config.PORT))
}
