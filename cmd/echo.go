package main

import (
	"hello/fizzbuzz"
	"hello/oscar"
	"net/http"
	"strconv"

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
	e.GET("/oscarmale", OsCarMaleHandle)
	// e.GET("/fizzbuzz/:number", FizzBuzzHandle)
	e.POST("/fizzbuzz", postFizzBuzzHandle)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func OsCarMaleHandle(c echo.Context) error {
	result := oscar.ActorWhoGotMoreThanOne("../oscar/oscar_age_male.csv")
	return c.JSON(http.StatusOK, result)
}

func FizzBuzzHandle(c echo.Context) error {
	numberString := c.Param("number") // param value
	n, _ := strconv.Atoi(numberString)
	return c.JSON(http.StatusOK, n)
}

func postFizzBuzzHandle(c echo.Context) error {
	var req map[string]int
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	type fizzBuzzResponse struct {
		Number   int    `json: number`
		FizzBuzz string `json: fizzbuzz`
	}

	return c.JSON(http.StatusOK, fizzBuzzResponse{
		Number:   req["number"],
		FizzBuzz: fizzbuzz.Say(req["number"]),
	})
}
