package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RecoverMiddleware función intermediaria para retornar el middleware Recover del framework Echo
func RecoverMiddleware() echo.MiddlewareFunc {

	return middleware.Recover();

}