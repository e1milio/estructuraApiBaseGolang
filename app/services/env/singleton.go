package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var once sync.Once // Lo utilizamos para construir un singleton

var env *Env // Variable Una sola instancia

// Get retorna una instancia de Env
func Get() *Env {

	once.Do(func() { // Función nativa que hace que el código se ejecute una sola vez

		// Lee el archivo con las variables de entorno
		fileEnv, err := ioutil.ReadFile("env.json")
		if err != nil {
			log.Fatalf("No se pudo abrir el archivo de entorno: %v", err)
		}

		json.Unmarshal(fileEnv, &env)

	})

	return env

}