package main

import "golinks/internal/worker"

func main() {
	config := worker.ConsumerConfig{
		ConsumerName: "link_images",
		QueueName:    "link_images",
		WorkersCount: 20, // todo env
	}
	worker.ConsumeQueue(config, worker.LinkImagesHandler)
}
