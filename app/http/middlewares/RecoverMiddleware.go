package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RecoverMiddleware encargado de logear los accessos a la aplicación
type RecoverMiddleware struct {
}

// Handle función que aplica el middleware
func (RecoverMiddleware RecoverMiddleware) Handle(e *echo.Echo) {

	e.Use(middleware.Recover())

}
