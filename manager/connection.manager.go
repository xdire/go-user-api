package manager

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connection *gorm.DB

func GetDefaultConnection() (*gorm.DB, error) {
	if connection != nil {
		return connection, nil
	}
	return nil, errors.New("connection wasn't established, please create new connection first")
}

func ConnectDefault() {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=root dbname=user-api password=localpassword sslmode=disable")
	if err != nil {
		fmt.Printf("%v", err)
		panic("Cannot connect to default Postgres instance")
	}
	connection = db
}

func DisconnectDefault()  {
	if connection != nil {
		_ = connection.Close()
	}
}