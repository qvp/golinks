package main

import "golinks/internal/worker"

func main() {
	conf := worker.ConsumerConfig{
		QueueName: "link_images",
	}
	worker.ConsumeQueue(conf, worker.LinkImagesHandler)
}
