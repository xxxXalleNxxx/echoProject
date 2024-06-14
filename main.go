package main

import (
	"echoProject/handlers"
	"github.com/labstack/echo"
	"net/http"
)

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "popaaaaaaaaaaaaaa")
}
func main() {
	e := echo.New()
	e.GET("/health-check", handlers.HealtCheckhandler())
	e.GET("/posts", handlers.HealtCheckhandler())
	e.GET("/post/:id", handlers.HealtCheckhandler())

	e.Start(":8080")
}
