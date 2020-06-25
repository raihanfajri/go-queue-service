package sub

import (
	"encoding/json"
	"fmt"

	rabbitmq "../config"
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

func (p *RabbitMQ) Subscribe() {

	forever := make(chan bool)

	qName := "queue"
	r := &rabbitmq.Config{}
	r.Initialize()

	ch, _, _ := r.Connect(qName)

	defer ch.Close()

	msgs, _ := ch.Consume(
		qName, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	go func() {
		for d := range msgs {
			json.Unmarshal(d.Body, p)
			fmt.Println("Received a message", p)
		}
	}()

	<-forever
}
