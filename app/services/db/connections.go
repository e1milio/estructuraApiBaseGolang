package db

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var once sync.Once
var connections map[string]*gorm.DB

// GetConnections retorna un map de conecciones
func GetConnections() map[string]*gorm.DB {

	once.Do(func() { // Función nativa que hace que el código se ejecute una sola vez

		connections = make(map[string]*gorm.DB)

	})

	return connections

}