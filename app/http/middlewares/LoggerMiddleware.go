package middlewares

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// LoggerMiddleware función intermediaria para retornar el middleware Logger del framework Echo
func LoggerMiddleware() echo.MiddlewareFunc {

	logsFile, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("No se pudo abrir o crear el archivo de logs: %v", err)
	}
	//defer logsFile.Close()

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logsFile,
	})

}