package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Logmiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] requestID:${id} ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
}
