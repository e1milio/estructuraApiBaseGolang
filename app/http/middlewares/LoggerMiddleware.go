package middlewares

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// LoggerMiddleware encargado de logear los accessos a la aplicación
type LoggerMiddleware struct {
}

// Handle función que aplica el middleware
func (loggerMiddleware LoggerMiddleware) Handle(e *echo.Echo) {

	logsFile, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("No se pudo abrir o crear el archivo de logs: %v", err)
	}
	//defer logsFile.Close()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logsFile,
	}))

}
