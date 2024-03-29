package configs

import "os"

var (
	DB_TYPE string
	DB_URL  string
	PORT    string
)

func init() {
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
