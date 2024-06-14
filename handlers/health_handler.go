package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func HealtCheckhandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK!")

}
