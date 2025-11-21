package web

import (
	"github.com/gracchi-stdio/goaat/internal/services"
	"github.com/labstack/echo/v4"
)

type HandlerFunc func(svc *services.Services, c echo.Context) error

func With(svc *services.Services, fn HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return fn(svc, c)
	}
}
