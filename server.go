package main

import (
	"helloshuya-web-api/entities"
	"helloshuya-web-api/infrastructure"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	infrastructure.NewDB()

	// get test
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Shuya!")
	})

	e.GET("/post", GetTest)

	e.Logger.Fatal(e.Start(":1323"))
}

func GetTest(c echo.Context) error {
	u := new(entities.Post)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
