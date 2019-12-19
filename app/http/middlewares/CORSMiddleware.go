package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CORSMiddleware encargado de logear los accessos a la aplicación
type CORSMiddleware struct {
}

// Handle función que aplica el middleware
func (CORSMiddleware CORSMiddleware) Handle(e *echo.Echo) {

	e.Use(middleware.CORS())

}