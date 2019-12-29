package response

import (
	"github.com/labstack/echo/v4"
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
			Code:        404,
			Status:      response.getStatusToCode(404),
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

// SendOk envia una respuesta con codigo 200
func (response *Response) SendOk() error {

	return response.Code(200).Send()

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
	response.Estructura.State.Status = response.getStatusToCode(code)

	return response

}

// etStatusToCode retorna el status
func (response *Response) getStatusToCode(code int) string {

	var status string

	switch code {
	case 200:
		status = "OK"
	case 400:
		status = "BAD_REQUEST"
	case 401:
		status = "UNAUTHORIZED"
	case 403:
		status = "FORBIDDEN"
	case 404:
		status = "NOT_FOUND"
	case 500:
		status = "INTERNAL_SERVER_ERROR"
	default:
		status = "ERROR"
	}

	return status
}