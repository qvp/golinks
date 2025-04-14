package worker

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"golinks/internal/db"
	"golinks/internal/parser"
)

// LinkImagesHandler handle message from rabbitmq
func LinkImagesHandler(d amqp.Delivery) error {
	var msg DebeziumMessage
	err := json.Unmarshal(d.Body, &msg)
	if err != nil {
		return nackOnError(d, "Message unmarshal fail", err)
	}

	linkID := msg.Payload.After["id"].(float64)
	url := msg.Payload.After["url"].(string)
	page, err := parser.LoadHtml(url)
	if err != nil {
		return nackOnError(d, fmt.Sprintf("unable to load url %s", url), err)
	}

	images, err := parser.GetImagesFromHtml(page)
	if err != nil {
		return nackOnError(d, fmt.Sprintf("unable to parse page from url %s", url), err)
	}

	if len(images) > 0 {
		err = db.Q.SaveLinkImagesTx(context.Background(), int(linkID), images)
		if err != nil {
			return nackOnError(d, fmt.Sprintf("unable to save images for link %f", linkID), err)
		}
	}

	return d.Ack(false)
}
