package helpers

import "github.com/google/uuid"

func CreateQueueID() (uuid.UUID, error) {

	uuid, err := uuid.NewRandom()

	return uuid, err
}
