package worker

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"

	"golinks/internal/db"
	"golinks/internal/link"
)

func LinkImagesHandler(d amqp.Delivery) {
	var msg DebeziumMessage
	err := json.Unmarshal(d.Body, &msg)
	failOnError(err, "Message unmarshal fail")

	linkIDFloat := msg.Payload.After["id"].(float64)
	url := msg.Payload.After["url"].(string)
	page, err := link.LoadHtml(url)
	if err != nil {
		fmt.Printf("%s: %s", "Unable to load url", err)
		d.Nack(false, false)
		return
	}

	images, err := link.GetImagesFromHtml(page)
	if err != nil {
		fmt.Printf("%s: %s", "Unable to parse page from url", err)
		d.Nack(false, false)
		return
	}
	// todo change status, dont save empty images list
	err = db.Func.LinkImageAddMultiple(context.Background(), int(linkIDFloat), images)
	if err != nil {
		fmt.Printf("%s: %s", "Unable to save images", err)
		d.Nack(false, false)
		return
	}

	d.Ack(false)
}
