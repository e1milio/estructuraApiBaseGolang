package routes

import (
	"estructuraapibasegolang/app/http/controllers"
	"estructuraapibasegolang/app/http/middlewares"
	"github.com/labstack/echo/v4"
)

// AuthRoutes crea las rutas auth
func AuthRoutes(e *echo.Echo) {

	// Se declaran los controladores que se usarán
	var authController controllers.AuthController

	e.POST("/auth/login", authController.Login) // Ruta para logear un usuario

	group := e.Group("/auth/test")
	group.Use(middlewares.JWTMiddleware())
	group.POST("", authController.TestToken) // Ruta para conocer el status del token

}