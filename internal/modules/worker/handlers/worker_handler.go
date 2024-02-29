package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"worker-service/internal/modules/worker"
	"worker-service/internal/modules/worker/models/request"
	"worker-service/internal/pkg/log"

	kafkaPkgConfluent "worker-service/internal/pkg/kafka/confluent"

	k "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type WorkerEventHandler struct {
	WorkerUsecaseCommand worker.UsecaseCommand
}

func NewWorkerEventConsumer(wc worker.UsecaseCommand) kafkaPkgConfluent.ConsumerHandler {
	return &WorkerEventHandler{
		WorkerUsecaseCommand: wc,
	}
}

func (w WorkerEventHandler) CreateBankTicket(message *k.Message, topic string) {
	log.GetLogger().Info(context.Background(), string(message.Value), fmt.Sprintf("Topic: %v Partition: %v - Offset: %v", *message.TopicPartition.Topic, message.TopicPartition.Partition, message.TopicPartition.Offset.String()))

	var msg request.CreateTicketReq
	if err := json.Unmarshal(message.Value, &msg); err != nil {
		log.GetLogger().Error(context.Background(), err.Error(), string(message.Value))
		return
	}
	if _, err := w.WorkerUsecaseCommand.CreateBankTicket(context.Background(), msg); err != nil {
		log.GetLogger().Error(context.Background(), err.Error(), string(message.Value))
		return
	}
	return
}

func (w WorkerEventHandler) UpdateOnlineBankTicket(message *k.Message, topic string) {
	log.GetLogger().Info(context.Background(), string(message.Value), fmt.Sprintf("Topic: %v Partition: %v - Offset: %v", *message.TopicPartition.Topic, message.TopicPartition.Partition, message.TopicPartition.Offset.String()))

	var msg request.CreateOnlineTicketReq
	if err := json.Unmarshal(message.Value, &msg); err != nil {
		log.GetLogger().Error(context.Background(), err.Error(), string(message.Value))
		return
	}

	resp, err := w.WorkerUsecaseCommand.CreateOnlineBankTicket(context.Background(), msg)
	if err != nil {
		log.GetLogger().Error(context.Background(), err.Error(), string(message.Value))
		return
	}
	if resp != nil {
		log.GetLogger().Info(context.Background(), *resp, string(message.Value))
		return
	}
	return
}
