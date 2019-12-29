package routes

import (
	"estructuraapibasegolang/app/http/controllers"
	"github.com/labstack/echo/v4"
)

// AppInfoRoutes crea las rutas de la aplicación
func AppInfoRoutes(e *echo.Echo) {

	// Se declaran los controladores que se usarán
	var appInfoController controllers.AppInfoController

	e.GET("/", appInfoController.GetStatus) // Ruta para conocer el status de la aplicación

	e.GET("/version", appInfoController.GetVersion) // Ruta para conocer la versión de la aplicación

}