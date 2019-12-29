package routes

import "github.com/labstack/echo/v4"

// StaticRoutes crea las rutas estáticas de la aplicación
func StaticRoutes(e *echo.Echo) {

	e.Static("/", "public") // Servir archivos estáticos públicos

}