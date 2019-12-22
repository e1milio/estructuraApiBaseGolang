package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CORSMiddleware función intermediaria para retornar el middleware CORS del framework Echo
func CORSMiddleware() echo.MiddlewareFunc {

	return middleware.CORS();

}