package use_cases

import (
	"Belobetty-Starter/internal/git/dto"
	"Belobetty-Starter/pkg/queue"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
)

const url = "guest@localhost:5672"

type SenderQueue struct {
	producer queue.Producer
}

func NewSenderQueue(keyTopic string) (*SenderQueue, error) {
	p, err := queue.NewRabbitMQ(keyTopic, url)
	if err != nil {
		return nil, err
	}
	return &SenderQueue{producer: p}, nil
}

func (s *SenderQueue) Exec(msg *dto.MessageOut) error {
	log.Infof("Test %s", msg)
	message, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	log.Info(message)

	err = s.producer.SendMessage(message)
	if err != nil {
		return err
	}
	return nil
}
