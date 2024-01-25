package routes

import (
	"jtn/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello guys !")
	})

	e.GET("/All/NoHP", controller.GetData)
	e.POST("/Create/NoHP", controller.CreateData)
	return e
}
