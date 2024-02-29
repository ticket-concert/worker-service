package handlers

import (
	"worker-service/configs"
	"worker-service/internal/modules/worker"
	"worker-service/internal/pkg/log"

	kafkaConfluent "worker-service/internal/pkg/kafka/confluent"
)

func InitWorkerEventConflHandler(wc worker.UsecaseCommand) {
	topicKcr := "concert-create-bank-ticket"
	kcr, _ := kafkaConfluent.NewConsumer(kafkaConfluent.GetConfig().GetKafkaConfig(configs.GetConfig().ServiceName, true), log.GetLogger())
	kcr.SetHandler(NewWorkerEventConsumer(wc))
	kcr.Subscribe(topicKcr)

	topicKut := "concert-update-online-bank-ticket"
	kut, _ := kafkaConfluent.NewConsumer(kafkaConfluent.GetConfig().GetKafkaConfig(configs.GetConfig().ServiceName, true), log.GetLogger())
	kut.SetHandler(NewWorkerEventConsumer(wc))
	kut.Subscribe(topicKut)

}
