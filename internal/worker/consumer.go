package worker

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumerConfig struct {
	QueueName     string
	PrefetchCount int
	PrefetchSize  int
}

var configDefault = ConsumerConfig{
	QueueName:     "",
	PrefetchCount: 1,
	PrefetchSize:  0,
}

func ConsumeQueue(config ConsumerConfig, handler func(d amqp.Delivery)) {

	//todo use configDefault

	conn, err := amqp.Dial("amqp://root:root@localhost:5672/vhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name,
		"golinks_worker",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Println("Received a message: %s", d.Body) // fixme

			handler(d)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
