package sender

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

const (
	domain    = "education"
	subdomain = "person"
)

type EventSender interface {
	Send(person *model.PersonEvent) error
	Close() error
}

var _ EventSender = &eventSender{}

type eventSender struct {
	topic    string
	producer sarama.SyncProducer
}

func NewEventSender(server string) (*eventSender, error) {
	const op = "NewEventSender"

	producer, err := sarama.NewSyncProducer([]string{server}, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: can't create producer: %w", op, err)
	}
	return &eventSender{
		topic:    domain + "-" + subdomain,
		producer: producer,
	}, nil
}

func (s *eventSender) Send(event *model.PersonEvent) error {
	const op = "eventSender.Send"

	requestID := uuid.New().String()

	value, err := json.Marshal(event)
	if err != nil {
		log.Error().Str("op", op).Err(err).Msg("can't marshal event")
		return fmt.Errorf("%s: can't marshal event: %w", op, err)
	}

	msg := &sarama.ProducerMessage{
		Topic: s.topic,
		Key:   sarama.StringEncoder(requestID),
		Value: sarama.ByteEncoder(value),
	}

	log.Debug().Str("op", op).Msgf("send: %s", msg)

	_, _, err = s.producer.SendMessage(msg)
	if err != nil {
		log.Error().Str("op", op).Err(err).Msg("can't send message")
		return fmt.Errorf("%s: can't send message: %w", op, err)
	}

	return nil
}

func (s *eventSender) Close() error {
	return s.producer.Close()
}
