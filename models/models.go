package models

import (
	"log"

	config "github.com/anant-sharma/go-blockchain-config"
	"github.com/go-bongo/bongo"
)

var connection *bongo.Connection

// InitModels to initialize
func InitModels() {

	config := config.GetConfig()

	mongoConfig := &bongo.Config{
		ConnectionString: config.MongoDB.ConnectionString,
		Database:         config.MongoDB.Database,
	}

	conn, err := bongo.Connect(mongoConfig)

	if err != nil {
		log.Fatal(err)
	}

	connection = conn
	log.Println("[*] Database Connection Established Successfully.")
}
