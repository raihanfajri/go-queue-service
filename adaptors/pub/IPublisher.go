package pub

import "os"

var (
	brokers = map[string]IPublisher{
		"rabbitmq": &RabbitMQ{},
	}
)

type IPublisher interface {
	Construct(queueId string, queueName string, configURL string, payload string)
	Publish() (string, error)
}

func GetPublisher() IPublisher {
	adaptor := os.Getenv("ADAPTOR")

	return brokers[adaptor]
}
