package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSMiddleware función intermediaria para retornar el middleware CORS del framework Echo
func CORSMiddleware() echo.MiddlewareFunc {

	return middleware.CORS();

}