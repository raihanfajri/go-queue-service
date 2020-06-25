package rabbitmq

import (
	"net"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type Config struct {
	user     string
	password string
	host     string
	port     string
}

func (c *Config) Initialize() {
	c.user = os.Getenv("RABBITMQ_USER")
	c.password = os.Getenv("RABBITMQ_PASSWORD")
	c.host = os.Getenv("RABBITMQ_HOST")
	c.port = os.Getenv("RABBITMQ_PORT")
}

func (c *Config) Connect(qName string) (*amqp.Channel, string, error) {
	ch := &amqp.Channel{}

	conenction := "amqp://" + c.user + ":" + c.password + "@" + c.host + ":" + c.port

	conn, err := amqp.DialConfig(conenction, amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 2*time.Second)
		},
	})

	if err != nil {
		return ch, "Failed to create rabbitmq connection", err
	}

	ch, err = conn.Channel()

	if err != nil {
		return ch, "Failed to create rabbitmq channel", err
	}

	_, err = ch.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return ch, "Failed to declare queue", err
	}

	return ch, "Success", err
}
