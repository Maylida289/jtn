package main

import (
	"jtn/database"
	"jtn/routes"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	database.Init()
	e := routes.Init()
	e.Validator = &CustomValidator{
		Validator: validator.New(),
	}

	e.Logger.Fatal(e.Start(":8002"))
}
