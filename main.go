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
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)
	e.GET("/posts/popa", yallo)

	e.Start(":8080")
}
