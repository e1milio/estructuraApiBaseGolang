package db

import (
	"sync"

	"estructuraapibasegolang/app/services/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// lock mutex
var lock = &sync.Mutex{}

// Open Crea y retorna una instancia de gorm.DB
func Open() *gorm.DB {

	db, _ := openDBConfig(env.Get().DB[0])
	return db
}

// OpenWithError Crea y retorna una instancia de gorm.DB junto a un error
func OpenWithError() (*gorm.DB, error) {

	db, err := openDBConfig(env.Get().DB[0])
	return db, err
}

// OpenDB Crea y retorna una instancia de gorm.DB
func OpenDB(dbName string) *gorm.DB {

	for _, element := range env.Get().DB {

		if element.Name == dbName {

			db, _ := openDBConfig(element)
			return db
		}
	}

	return nil
}

// OpenDBWithError Crea y retorna una instancia de gorm.DB junto a un error
func OpenDBWithError(dbName string) (*gorm.DB, error) {

	for _, element := range env.Get().DB {

		if element.Name == dbName {

			db, err := openDBConfig(element)
			return db, err
		}
	}

	return nil, *new(error)
}

// OpenDB Crea y retorna una instancia de gorm.DB
func openDBConfig(dbConfig env.DBConfig) (*gorm.DB, error) {

	// Declaro las variables que voy a retornar
	var db *gorm.DB 
	var err error

	// Obtengo las coneciones que ya fueron abiertas
	connections := GetConnections()

	if _, ok := connections[dbConfig.Name]; ok {

		db = connections[dbConfig.Name]

	} else {

		lock.Lock()         //Bloqueo la función
		defer lock.Unlock() //Quito el bloqueo cuando termino la ejecución de la función

		args := ""

		switch dbConfig.Connection {

		case "mysql":
			args = dbConfig.UserName + ":" +
				dbConfig.Password +
				"@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")" +
				"/" + dbConfig.DataBase +
				"?loc=Local"

		default:
			return nil, *new(error)
		}

		// Abro una conección
		db, err = gorm.Open(dbConfig.Connection, args)
		if err == nil {
			// Guardo la conección abierta en el map conecciones
			connections[dbConfig.Name] = db
		}
		
	}

	return db, err
}