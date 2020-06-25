package sub

import (
	"os"
)

var (
	brokers = map[string]ISubscriber{
		"rabbitmq": &RabbitMQ{},
	}
)

type ISubscriber interface {
	Construct(queueId string, queueName string, configURL string, payload string)
	Subscribe()
}

func GetSubscriber() ISubscriber {
	adaptor := os.Getenv("ADAPTOR")

	return brokers[adaptor]
}
