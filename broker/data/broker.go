package data

import "github.com/streadway/amqp"

type Broker struct {
	conn *amqp.Connection
}

func NewBroker(conn *amqp.Connection) *Broker {
	return &Broker{conn: conn}
}
