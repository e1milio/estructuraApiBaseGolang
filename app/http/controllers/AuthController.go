// Package controllers contiene todos los controladores http de la aplicación
package controllers

import (
	"estructuraapibasegolang/app/services/auth"
	"estructuraapibasegolang/app/services/env"
	"estructuraapibasegolang/app/services/response"

	"github.com/labstack/echo/v4"
)

// AuthController
type AuthController struct {
}

// Login
func (authController AuthController) Login(c echo.Context) error {

	// Obtengo los parámetros para el login
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Verifico que las credenciales sean correctas
	if username != env.Get().AuthUser || password != env.Get().AuthPassword {
		return response.New(c).Code(401).Send()
	}

	// Obtengo el token
	token, err := auth.CreateToken(username)
	if err != nil {
		return response.New(c).Code(500).Data(err.Error()).Send()
	}

	// Creo la respuesta
	rst := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	return response.New(c).Data(rst).SendOk()
}

// TestToken
func (authController AuthController) TestToken(c echo.Context) error {

	auth_token := auth.GetToken(c.Get("auth_token"))

	return response.New(c).Data(auth_token).SendOk()
}
