package worker

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumerConfig struct {
	ConsumerName string
	QueueName    string
	WorkersCount int
}

func ConsumeQueue(config ConsumerConfig, handler func(d amqp.Delivery) error) {
	if config.QueueName == "" {
		log.Panicf("Empty queue name")
	}
	if config.WorkersCount == 0 {
		config.WorkersCount = 10
	}

	conn, err := amqp.Dial("amqp://root:root@localhost:5672/vhost") // fixme
	failOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("Failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(config.QueueName, true, false, false, false, nil)
	failOnError("Failed to declare a queue", err)

	err = ch.Qos(1, 0, false)
	failOnError("Failed to set QoS", err)

	messages, err := ch.Consume(q.Name, config.ConsumerName, false, false, false, false, nil)
	failOnError("Failed to register a consumer", err)

	semaphore := make(chan struct{}, config.WorkersCount)
	var forever chan struct{}

	go func() {
		for delivery := range messages {
			semaphore <- struct{}{}

			go func(d amqp.Delivery) {
				err := handler(d)
				if err != nil {
					log.Printf("handle message error: %s", err)
				}
				<-semaphore

			}(delivery)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")

	<-forever
}

func failOnError(message string, err error) {
	if err != nil {
		log.Panicf("%s: %s", message, err)
	}
}
