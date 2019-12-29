// Package controllers contiene todos los controladores http de la aplicación
package controllers

import (
	"net/http"

	"estructuraapibasegolang/app/services/env"
	"estructuraapibasegolang/app/services/response"
	"github.com/labstack/echo/v4"
)

// AppInfoController brinda información general sobre la aplicación
type AppInfoController struct {
}

// GetStatus retorna un estado de ok si la aplicación está corriendo
func (appInfoController AppInfoController) GetStatus(c echo.Context) error {

	return c.NoContent(http.StatusOK)

}

// GetVersion retorna la versión de la aplicación
func (appInfoController AppInfoController) GetVersion(c echo.Context) error {

	// Creo la respuesta
	rst := struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}{
		Name:    env.Get().AppName,
		Version: env.Get().AppVersion,
	}

	return response.New(c).Data(rst).SendOk()
}