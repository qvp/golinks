package main

import (
	"golinks/internal/common"
	"golinks/internal/worker"
)

func main() {
	common.InitEnv()

	config := worker.ConsumerConfig{
		ConsumerName: "link_images",
		QueueName:    "link_images",
		WorkersCount: common.GetEnvInt("WORKERS_COUNT", 1),
	}
	worker.ConsumeQueue(config, worker.LinkImagesHandler)
}
