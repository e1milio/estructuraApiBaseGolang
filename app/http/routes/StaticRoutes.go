package routes

import "github.com/labstack/echo"

// StaticRoutes crea las rutas de la aplicación
func StaticRoutes(e *echo.Echo) {

	e.Static("/", "public") // Servir archivos estáticos públicos

}
