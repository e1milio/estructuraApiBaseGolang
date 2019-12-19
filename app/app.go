package app

import (
	"estructuraapibasegolang/app/http"
	"estructuraapibasegolang/app/services/env"
	"github.com/labstack/echo"
)

// Boot inicia la api
func Boot() {

	e := echo.New() // Se instancia el framework Echo

	http.Middlewares(e) // Se cargan los Middlewares

	http.Routes(e) // Se cargan las rutas

	e.Start(":" + env.Get().AppPort) // Se inica el servidor

}
