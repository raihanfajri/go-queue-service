package pub

import (
	"encoding/json"

	rabbitmq "../config"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	QueueID   string `json:"queue_id"`
	QueueName string `json:"queue_name"`
	ConfigURL string `json:"config_url"`
	Payload   string `json:"payload"`
}

func (p *RabbitMQ) Construct(queueId string, queueName string, configURL string, payload string) {
	p.QueueID = queueId
	p.QueueName = queueName
	p.ConfigURL = configURL
	p.Payload = payload
}

func (p *RabbitMQ) Publish() (string, error) {
	qName := "queue"
	r := &rabbitmq.Config{}
	r.Initialize()

	ch, m, err := r.Connect(qName)

	if err != nil {
		return m, err
	}

	defer ch.Close()

	body, _ := json.Marshal(p)

	err = ch.Publish(
		"",    // exchange
		qName, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		return "Failed to publish rabbitmq message", err
	}

	return "Success sent message to rabbitmq. Message : " + string(body), nil
}
