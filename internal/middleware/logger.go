package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ColorfulLogger returns a colorful request logging middleware for Echo
func ColorfulLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		LogMethod:   true,
		LogLatency:  true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			statusColor := "\033[32m" // Green
			if v.Status >= 400 && v.Status < 500 {
				statusColor = "\033[33m" // Yellow
			} else if v.Status >= 500 {
				statusColor = "\033[31m" // Red
			}

			methodColor := "\033[36m" // Cyan

			fmt.Printf("%s%-7s\033[0m | %s%3d\033[0m | %13v | %s | %s\n",
				methodColor, v.Method,
				statusColor, v.Status,
				v.Latency,
				v.RemoteIP,
				v.URI,
			)

			if v.Error != nil {
				fmt.Printf("\033[31mError: %v\033[0m\n", v.Error)
			}
			return nil
		},
	})
}
