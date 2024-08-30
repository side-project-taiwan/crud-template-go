package main

import (
	"fmt"
	"log"
	"os"

	"spt/internal/utility"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "./models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	rootDir, envPath, err := utility.GetProjectRootDirAndEnvPath()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Project root directory:", rootDir)
	fmt.Println("Environment file path:", envPath)

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("GORM_GEN_CONNECTION")
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	g.UseDB(gormdb)

	g.ApplyBasic(g.GenerateModel("project")) //你自己想產的表

	g.Execute()
}
