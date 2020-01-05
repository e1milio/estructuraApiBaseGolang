package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Response encargado enviar respuestas
type Response struct {
	Context  echo.Context
	Estructura ResponseStruct
}

// Crea y retorna una instancia de Response
func New(c echo.Context) *Response {

	var response Response

	response.Context = c

	// Creo la respuesta por defecto
	params, _ :=c.FormParams()
	response.Estructura = ResponseStruct{
		Data: "",
		Metadata: Metadata{
			Request:  params,
			Origin:   c.Request().RequestURI,
			Paginado: "",
		},
		State: State{
			Code:        200,
			Status:      http.StatusText(200),
			Message:     "",
			Description: "",
		},
	}

	return &response
}

// Send Envia la respuesta
func (response *Response) Send() error {

	// Envio la respuesta
	return response.Context.JSON(response.Estructura.State.Code, response.Estructura)

}

// Data setea la data de la respuesta
func (response *Response) Data(data interface{}) *Response {

	response.Estructura.Data = data

	return response

}

// Message setea el mensaje de la respuesta
func (response *Response) Message(message string) *Response {

	response.Estructura.State.Message = message

	return response

}

// Description setea el mensaje de la respuesta
func (response *Response) Description(description string) *Response {

	response.Estructura.State.Description = description

	return response

}

// Code setea el código de respuesta
func (response *Response) Code(code int) *Response {

	response.Estructura.State.Code = code
	response.Estructura.State.Status = http.StatusText(code)

	return response

}