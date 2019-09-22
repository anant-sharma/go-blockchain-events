package eventslistener

import (
	"encoding/json"
	"log"

	"github.com/anant-sharma/go-blockchain-events/models"
	"github.com/anant-sharma/go-blockchain/controller/v1/blockchain"
	"github.com/anant-sharma/go-blockchain/controller/v1/pubsub"
	"github.com/mitchellh/mapstructure"

	config "github.com/anant-sharma/go-blockchain-config"
	"github.com/anant-sharma/go-utils/mq"
	"github.com/streadway/amqp"
)

func createQueue(exchange string, queue string) mq.MQ {
	_mq := mq.NewMQ()

	// Connect to MQ
	_mq.Connect(config.GetConfig().MQConnectionString)

	// Create Channel with MQ
	_mq.CreateChannel()

	// Create Exchange
	_mq.CreateExchange(exchange, "fanout", true)

	// Create Queue
	_mq.CreateQueue(queue, true, false, false, false, amqp.Table{})
	log.Printf("[*] Queue %s Created.", queue)

	// Bind Queue With Exchange
	_mq.BindQueueWithExchange(queue, exchange, "", amqp.Table{})
	log.Printf("[*] Queue %s Bound To Exchange %s.", queue, exchange)

	// Create Exchange Log Queue
	_mq.CreateQueue(exchange+"-logs", true, false, false, false, amqp.Table{
		"x-message-ttl": int(86400),
	})

	// Bind Exchange Log Queue
	_mq.BindQueueWithExchange(exchange+"-logs", exchange, "", amqp.Table{
		"x-message-ttl": int64(86400),
	})

	return _mq
}

// InitListening - Start
func InitListening() {

	_mq := createQueue("bc.msg.exchange", "bc.events.queue")

	// Add Event Subscribers
	messages := _mq.EstablishWorker("bc.events.queue")

	for d := range messages {
		if d.Body != nil {
			var msg pubsub.Message
			json.Unmarshal(d.Body, &msg)

			switch msg.Event {
			case pubsub.PubSubEvents.TransactionCreated:
				{
					var transaction blockchain.Transaction
					mapstructure.Decode(msg.Data, &transaction)
					models.SaveTransaction(transaction)
				}
			}

			log.Printf("[x] %s", d.Body)
		}
	}
}
