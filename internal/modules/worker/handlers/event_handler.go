package handlers

import (
	"worker-service/configs"
	"worker-service/internal/modules/worker"
	"worker-service/internal/pkg/log"

	kafkaConfluent "worker-service/internal/pkg/kafka/confluent"
)

func InitWorkerEventConflHandler(wc worker.UsecaseCommand, log log.Logger) {
	topicKcr := "concert-create-bank-ticket"
	kcr, _ := kafkaConfluent.NewConsumer(kafkaConfluent.GetConfig().GetKafkaConfig(configs.GetConfig().ServiceName, true), log)
	kcr.SetHandler(NewWorkerEventConsumer(wc, log))
	kcr.Subscribe(topicKcr)

	topicKut := "concert-update-online-bank-ticket"
	kut, _ := kafkaConfluent.NewConsumer(kafkaConfluent.GetConfig().GetKafkaConfig(configs.GetConfig().ServiceName, true), log)
	kut.SetHandler(NewWorkerEventConsumer(wc, log))
	kut.Subscribe(topicKut)

}
