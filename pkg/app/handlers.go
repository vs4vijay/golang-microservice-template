package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) HealthCheck(c echo.Context) error {
	res := healthCheckResponse{
		Success: true,
	}

	return c.JSON(http.StatusOK, res)
}
