package config


import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

var MyRabbitMQ *RabbitMQ

func NewRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return
	}
	ch, err := conn.Channel()
	if err != nil {
		return
	}

	MyRabbitMQ = &RabbitMQ{
		conn:    conn,
		channel: ch,
	}

}

func (r *RabbitMQ) Publish(queueName string, body string) error {
	_, err := r.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return r.channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}

func (r *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	_, err := r.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return r.channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		err := r.channel.Close()
		if err != nil {
			return
		}
	}
	if r.conn != nil {
		err := r.conn.Close()
		if err != nil {
			return
		}
	}
}

func StartConsumer(queueName string) {

	messages, err := MyRabbitMQ.Consume(queueName)

	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("Consumer started")
	for d := range messages {
		log.Printf("Received message: %s\n", d.Body)

		go func() {

			fmt.Println("Message: ", string(d.Body))

		}()
	}

}