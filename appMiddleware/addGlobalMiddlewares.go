package appMiddleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AddGlobalMiddlewares(e *echo.Echo) {
	// Log middleware
	//e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// AddTrailingSlashMiddleware: Ensure all request has '/' suffix
	e.Use(middleware.AddTrailingSlash())
}
