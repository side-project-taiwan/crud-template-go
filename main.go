package main

import (
	"sample/internal/controller"
	_ "sample/internal/service"

	_ "github.com/joho/godotenv/autoload"
)

//_ "sample/internal/database"

func main() {

	//_ = controller.InitializeController()
	_ = controller.InitializeController()
}
