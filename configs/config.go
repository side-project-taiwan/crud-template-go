package configs

import (
	"os"
	"sample/internal/util"
)

var (
	DB_TYPE string
	DB_URL  string
	PORT    string
)

func init() {
	util.PrintLogWithColor("Enter configsInit log=>"+os.Getenv("DB_URL"), "#00ffff")

	DB_TYPE = os.Getenv("DB_TYPE")
	DB_URL = os.Getenv("DB_URL")
	PORT = os.Getenv("PORT")
}

// type Config struct {
// 	DB_TYPE string
// 	DB_URL  string
// 	PORT    string
// }

// func NewConfig() (*Config, error) {
// 	return &Config{
// 		DB_TYPE: os.Getenv("DB_TYPE"),
// 		DB_URL:  os.Getenv("DB_URL"),
// 		PORT:    os.Getenv("PORT"),
// 	}, nil
// }
