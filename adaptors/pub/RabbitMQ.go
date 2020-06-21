package pub

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

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
	conn, err := connect()

	if err != nil {
		fmt.Println(err)

		return "Failed to create rabbitmq connection", err
	}

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)

		return "Failed to create rabbitmq channel", err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		p.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		fmt.Println(err)

		return "Failed to declare queue", err
	}

	body, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		fmt.Println(err)

		return "Failed to publish rabbitmq message", err
	}

	return "Success sent message to rabbitmq. Message : " + string(body), nil
}

func connect() (*amqp.Connection, error) {
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	conenction := "amqp://" + user + ":" + password + "@" + host + ":" + port

	conn, err := amqp.DialConfig(conenction, amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 2*time.Second)
		},
	})

	return conn, err
}
