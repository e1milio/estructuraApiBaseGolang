package http

import (
	"estructuraapibasegolang/app/http/routes"
	"github.com/labstack/echo/v4"
)

// Routes carga las rutas de la aplicación
func Routes(e *echo.Echo) {

	routes.StaticRoutes(e)
	routes.AppInfoRoutes(e)
	routes.AuthRoutes(e)

}