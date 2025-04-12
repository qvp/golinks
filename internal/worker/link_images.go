package worker

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"

	"golinks/internal/db"
	"golinks/internal/parser"
)

func LinkImagesHandler(d amqp.Delivery) error {
	var msg DebeziumMessage
	err := json.Unmarshal(d.Body, &msg)
	if err != nil {
		log.Printf("Message unmarshal fail %s", err)
		return d.Nack(false, false)
	}

	linkID := msg.Payload.After["id"].(float64)
	url := msg.Payload.After["url"].(string)
	page, err := parser.LoadHtml(url)
	if err != nil {
		log.Printf("unable to load url %s: %v", url, err)
		return d.Nack(false, false)
	}

	images, err := parser.GetImagesFromHtml(page)
	if err != nil {
		log.Printf("unable to parse page from url %s: %v", url, err)
		return d.Nack(false, false)
	}

	if len(images) > 0 {
		err = db.Q.SaveLinkImagesTx(context.Background(), int(linkID), images)
		if err != nil {
			log.Printf("unable to save images for link %d: %v", linkID, err)
			return d.Nack(false, false)
		}
	}

	return d.Ack(false)
}
