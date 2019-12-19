package env

// Env estructura variable de entorno
type Env struct {
	AppName    string `json:"APP_NAME"`
	AppVersion string `json:"APP_VERSION"`
	AppPort    string `json:"APP_PORT"`
	URLAPIExterna    string `json:"URL_API_EXTERNA"`
	DB         []DBConfig `json:"DB"`
}

// DBConfig subestructura de env
type DBConfig struct {
	Name   string `json:"NAME"`
	Connection string `json:"CONNECTION"`
	Host       string `json:"HOST"`
	Port       string `json:"PORT"`
	DataBase   string `json:"DATABASE"`
	UserName   string `json:"USERNAME"`
	Password   string `json:"PASSWORD"`
}

