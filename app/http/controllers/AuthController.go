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
func (AuthController) Login(c echo.Context) error {

	// Obtengo los parámetros para el login
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Verifico que las credenciales sean correctas
	if username != env.Get().AuthUser || password != env.Get().AuthPassword {
		return response.New(c).Code(401).HideRequest().Send()
	}

	// Obtengo el token
	token, err := auth.CreateToken(username)
	if err != nil {
		return response.New(c).Code(500).Data(err.Error()).HideRequest().Send()
	}

	// Creo la respuesta
	rst := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	return response.New(c).Data(rst).HideRequest().Send()
}

// TestToken verify the validity of a token
func (AuthController) TestToken(c echo.Context) error {

	authToken := auth.GetToken(c.Get("auth_token"))

	return response.New(c).Data(authToken).Send()
}