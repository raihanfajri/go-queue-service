package pubservice

import (
	"../adaptors/pub"
	"../helpers"
)

type PubService struct{}

func (p *PubService) PublishMessage(message string, url string) (string, error) {

	filename := "../routes/" + url + ".yml"

	yml := helpers.RouteYML{}

	_, err := helpers.ReadYML(filename, &yml)

	if err != nil {
		return "Error parsing", err
	}

	publisher := pub.GetPublisher()

	publisher.Construct("2", yml.QueueName, url, message)

	sentMessage, err := publisher.Publish()

	return sentMessage, err
}
