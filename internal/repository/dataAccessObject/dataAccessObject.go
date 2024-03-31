package dataAccessObject

import (
	"database/sql"
	"errors"
	"fmt"
)

type DataAccessObject interface {
	CreateUser(name string, email string) error
}

type DAO struct {
	DB *sql.DB
}

func (dao *DAO) CreateUser(name string, email string) error {
	if name == "" {
		return errors.New("name is required")
	}
	if email == "" {
		return errors.New("email is required")
	}

	fmt.Println("User created successfully")

	return nil
}
