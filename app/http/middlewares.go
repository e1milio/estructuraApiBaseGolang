package http

import (
	"estructuraapibasegolang/app/http/middlewares"
	"github.com/labstack/echo/v4"
)

// Middlewares carga todos los middlewares de la aplicación
func Middlewares(e *echo.Echo) {
	
	e.Use(middlewares.RecoverMiddleware())
	e.Use(middlewares.LoggerMiddleware())
	e.Use(middlewares.CORSMiddleware())

}