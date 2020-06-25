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
		return "Error parsing YML", err
	}

	queueID, err := helpers.CreateQueueID()

	if err != nil {
		return "Failed to generate Queue ID", err
	}

	publisher := pub.GetPublisher()

	publisher.Construct(queueID.String(), yml.QueueName, url, message)

	sentMessage, err := publisher.Publish()

	return sentMessage, err
}
