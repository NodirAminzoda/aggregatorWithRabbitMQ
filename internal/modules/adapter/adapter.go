package adapter

import (
	"aggreagtor/internal/config"
	"aggreagtor/internal/domain/precheck"
	"aggreagtor/pkg/zerolog"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

type Adapter struct {
	cfg       *config.Adapter
	log       *zerolog.Logger
	responses sync.Map // Map to store response channels by CorrelationId
}

func New(cfg *config.Adapter, log *zerolog.Logger) *Adapter {
	adapter := &Adapter{
		cfg: cfg,
		log: log,
	}

	go adapter.startConsumer()

	return adapter
}

type IAdapter interface {
	PreCheck(request precheck.Request, requestID string) (*precheck.Response, error)
}

func (a *Adapter) startConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		a.log.Error(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		a.log.Error(fmt.Sprintf("Failed to open a channel", err))
		return
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"response", // Shared response queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		a.log.Error(fmt.Sprintf("Failed to start consuming", err))
		return
	}

	for d := range msgs {
		correlationID := d.CorrelationId
		if ch, ok := a.responses.Load(correlationID); ok {
			responseChan := ch.(chan amqp.Delivery)
			responseChan <- d
			close(responseChan) // Close the channel after sending the message
			a.responses.Delete(correlationID)
		}
	}
}

func (a *Adapter) PreCheck(request precheck.Request, requestID string) (*precheck.Response, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		a.log.ErrorWithPrefix(requestID, "unable to open connection to RabbitMQ server", err)
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		a.log.ErrorWithPrefix(requestID, "Unable to open a channel", err)
		return nil, err
	}
	defer ch.Close()

	// Declare request queue
	_, err = ch.QueueDeclare(
		"request",
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)

	if err != nil {
		a.log.ErrorWithPrefix(requestID, "Failed to declare request queue", err)
		return nil, err
	}

	correlationID := uuid.New().String()

	// Response channel for the current request
	responseChan := make(chan amqp.Delivery, 1)
	a.responses.Store(correlationID, responseChan)

	body, err := json.Marshal(request)
	if err != nil {
		a.log.ErrorWithPrefix(requestID, "can't marshal request", err)
		return nil, err
	}

	// Publish request
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"",
		"request",
		false,
		false,
		amqp.Publishing{
			Body:          body,
			ContentType:   "application/json",
			CorrelationId: correlationID,
			ReplyTo:       "response",
		},
	)

	if err != nil {
		a.log.ErrorWithPrefix(requestID, "failed publishing messages", err)
		return nil, err
	}

	// Wait for response
	select {
	case msg := <-responseChan:
		var result precheck.Response
		if err := json.Unmarshal(msg.Body, &result); err != nil {
			a.log.ErrorWithPrefix(requestID, "can't unmarshal response", err)
			return nil, err
		}
		a.log.InfoWithPrefix(requestID, "Response body", fmt.Sprintf("%+v", result))
		return &result, nil
	case <-ctx.Done():
		a.log.ErrorWithPrefix(requestID, "timeout waiting for response", ctx.Err())
		return nil, ctx.Err()
	}
}
