package main

import (
	"hello/oscar"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Echo MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router
	e.GET("oscarmale", OsCarMaleHandle)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func OsCarMaleHandle(c echo.Context) error {
	result := oscar.ActorWhoGotMoreThanOne("../oscar/oscar_age_male.csv")
	return c.JSON(http.StatusOK, result)
}