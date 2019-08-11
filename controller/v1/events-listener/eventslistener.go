package eventslistener

import (
	"github.com/anant-sharma/go-blockchain-events/models"

	utils "github.com/anant-sharma/go-blockchain/common"
	"github.com/anant-sharma/go-blockchain/controller/v1/blockchain"
	"github.com/anant-sharma/go-blockchain/controller/v1/pubsub"
	"github.com/mitchellh/mapstructure"
)

// InitListening - Start
func InitListening() {

	// Init PubSub
	pubsub.NewPubSub("bc.msg.exchange", utils.GenerateUUID())

	// Add Event Subscribers
	messages := pubsub.Subscribe()

	for msg := range messages {

		switch msg.Event {
		case pubsub.PubSubEvents.TransactionCreated:
			{
				var transaction blockchain.Transaction
				mapstructure.Decode(msg.Data, &transaction)
				models.SaveTransaction(transaction)
			}
		}
	}
}
