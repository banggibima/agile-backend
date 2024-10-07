package rabbitmq

import (
	"github.com/banggibima/agile-backend/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Client(config *config.Config) (*amqp.Connection, error) {
	url := config.RabbitMQ.URL

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
