package main

import (
	eventslistener "github.com/anant-sharma/go-blockchain-events/controller/v1/events-listener"
	"github.com/anant-sharma/go-blockchain-events/models"
)

func main() {

	models.InitModels()
	eventslistener.InitListening()

}
