package handlers_test

import (
	"testing"
	"worker-service/internal/modules/worker/handlers"
	"worker-service/internal/pkg/errors"
	mockcert "worker-service/mocks/modules/worker"
	mocklog "worker-service/mocks/pkg/log"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type WorkerHandlerTestSuite struct {
	suite.Suite
	workerUsecaseCommand *mockcert.UsecaseCommand
	handler              *handlers.WorkerEventHandler
	mockLogger           *mocklog.Logger
}

func (suite *WorkerHandlerTestSuite) SetupTest() {
	suite.workerUsecaseCommand = new(mockcert.UsecaseCommand)
	suite.mockLogger = &mocklog.Logger{}
	suite.handler = &handlers.WorkerEventHandler{
		WorkerUsecaseCommand: suite.workerUsecaseCommand,
		Logger:               suite.mockLogger,
	}
	handlers.NewWorkerEventConsumer(suite.workerUsecaseCommand, suite.mockLogger)
}

func TestWorkerHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerHandlerTestSuite))
}

func (suite *WorkerHandlerTestSuite) TestCreateBankTicket() {
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)
	suite.workerUsecaseCommand.On("CreateBankTicket", mock.Anything, mock.Anything).Return(nil, nil)
	topic := "topic"
	msg := kafka.Message{
		Value: []byte(`{"ticketId": "154fb7e0-07b8-43ba-898c-2de889f272ef",
		"eventId": "1659cc60-0523-429e-96e3-a54cdb1898d0"}`),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: 1,
			Offset:    kafka.OffsetBeginning,
		},
	}
	suite.handler.CreateBankTicket(&msg, topic)
}

func (suite *WorkerHandlerTestSuite) TestCreateBankTicketErr() {
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)
	suite.workerUsecaseCommand.On("CreateBankTicket", mock.Anything, mock.Anything).Return(nil, errors.BadRequest("error"))
	topic := "topic"
	msg := kafka.Message{
		Value: []byte(`{"ticketId": "154fb7e0-07b8-43ba-898c-2de889f272ef",
		"eventId": "1659cc60-0523-429e-96e3-a54cdb1898d0"}`),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: 1,
			Offset:    kafka.OffsetBeginning,
		},
	}
	suite.handler.CreateBankTicket(&msg, topic)
}

func (suite *WorkerHandlerTestSuite) TestCreateBankTicketErrParse() {
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)
	suite.workerUsecaseCommand.On("CreateBankTicket", mock.Anything, mock.Anything).Return(nil, nil)
	topic := "topic"
	msg := kafka.Message{
		Value: []byte(""),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: 1,
			Offset:    kafka.OffsetBeginning,
		},
	}
	suite.handler.CreateBankTicket(&msg, topic)
}

func (suite *WorkerHandlerTestSuite) TestUpdateOnlineBankTicket() {
	topic := "topic"
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)
	suite.workerUsecaseCommand.On("CreateOnlineBankTicket", mock.Anything, mock.Anything).Return(&topic, nil)
	msg := kafka.Message{
		Value: []byte(`{"tag": "tag", "countryCode": "code"}`),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: 1,
			Offset:    kafka.OffsetBeginning,
		},
	}
	suite.handler.UpdateOnlineBankTicket(&msg, topic)
}

func (suite *WorkerHandlerTestSuite) TestUpdateOnlineBankTicketNil() {
	topic := "topic"
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)
	suite.workerUsecaseCommand.On("CreateOnlineBankTicket", mock.Anything, mock.Anything).Return(nil, nil)
	msg := kafka.Message{
		Value: []byte(`{"tag": "tag", "countryCode": "code"}`),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: 1,
			Offset:    kafka.OffsetBeginning,
		},
	}
	suite.handler.UpdateOnlineBankTicket(&msg, topic)
}

func (suite *WorkerHandlerTestSuite) TestUpdateOnlineBankTicketErr() {
	topic := "topic"
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)
	suite.workerUsecaseCommand.On("CreateOnlineBankTicket", mock.Anything, mock.Anything).Return(&topic, errors.BadRequest("error"))
	msg := kafka.Message{
		Value: []byte(`{"tag": "tag", "countryCode": "code"}`),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: 1,
			Offset:    kafka.OffsetBeginning,
		},
	}
	suite.handler.UpdateOnlineBankTicket(&msg, topic)
}

func (suite *WorkerHandlerTestSuite) TestUpdateOnlineBankTicketErrParse() {
	topic := "topic"
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)
	suite.workerUsecaseCommand.On("CreateOnlineBankTicket", mock.Anything, mock.Anything).Return(&topic, nil)
	msg := kafka.Message{
		Value: []byte("test"),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: 1,
			Offset:    kafka.OffsetBeginning,
		},
	}
	suite.handler.UpdateOnlineBankTicket(&msg, topic)
}
