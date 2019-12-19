package http

import (
	"estructuraapibasegolang/app/http/middlewares"
	"github.com/labstack/echo"
)

// Middlewares crea los middlewares de la aplicación
func Middlewares(e *echo.Echo) {

	// Se declaran los middlewares que se usrán
	var loggerMiddleware middlewares.LoggerMiddleware
	var recoverMiddleware middlewares.RecoverMiddleware
	var corsMiddleware middlewares.CORSMiddleware

	loggerMiddleware.Handle(e)  // Middleware para registrar logs
	recoverMiddleware.Handle(e) // Middleware para recuperación de errores
	corsMiddleware.Handle(e)    // Middleware CORS

}
