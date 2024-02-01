package use_cases

import (
	"Belobetty-Starter/pkg/queue"
	"Belobetty-Starter/pkg/queue/dto"
	"encoding/json"
)

const urlRabbitMQ = "guest@localhost:5672"

type SenderQueue struct {
	producer queue.Producer
}

func NewSenderQueue(keyTopic string) (*SenderQueue, error) {
	p, err := queue.NewRabbitMQ(keyTopic, urlRabbitMQ)
	if err != nil {
		return nil, err
	}
	return &SenderQueue{producer: p}, nil
}

func (s *SenderQueue) Exec(body entityGeneric, user, action string) error {

	err := body.Validate()
	if err != nil {
		return err
	}

	messageDto := dto.NewMessageOut(body, user, action)

	message, err := json.Marshal(messageDto)
	if err != nil {
		return err
	}

	err = s.producer.SendMessage(message)
	if err != nil {
		return err
	}
	return nil
}
