package main

import (
	"fmt"
	"hello/fizzbuzz"
	"hello/oscar"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	e.GET("/fizzbuzz/:number", FizzBuzzHandle)
	e.GET("/token", TokenHandle)
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
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("panic: %s", r),
			})
		}
	}()
	tokenString := c.Request().Header.Get("Authorization")[7:]

	//validation token
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Err sign method %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Token invalid")
	}

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

func TokenHandle(c echo.Context) error {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "sutin injitt"
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
