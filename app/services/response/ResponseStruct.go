package response

// ResponseStruct estructura de la respuesta
type ResponseStruct struct {
	Data     interface{} `json:"data"`
	Metadata Metadata    `json:"metadata"`
	State    State       `json:"state"`
}

// Metadata subestructura de la respuesta
type Metadata struct {
	Request  interface{} `json:"request"`
	Origin   string      `json:"origin"`
	Paginado interface{} `json:"paginado"`
}

// state subestructura de la respuesta
type State struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}