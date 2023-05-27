package main

import (
	"main/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", controller.GetUser)

}
