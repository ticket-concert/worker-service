package handlers_test

import (
	"worker-service/internal/modules/worker/handlers"
	mockcert "worker-service/mocks/modules/worker"
	mocklog "worker-service/mocks/pkg/log"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EventHandlerTestSuite struct {
	suite.Suite
	workerUsecaseCommand *mockcert.UsecaseCommand
	mockLogger           *mocklog.Logger
}

func (suite *EventHandlerTestSuite) SetupTest() {
	suite.workerUsecaseCommand = new(mockcert.UsecaseCommand)
	suite.mockLogger = &mocklog.Logger{}
}

func (suite *WorkerHandlerTestSuite) TestInitWorkerEventConflHandler() {
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)
	suite.workerUsecaseCommand.On("CreateBankTicket", mock.Anything, mock.Anything).Return(nil)
	suite.workerUsecaseCommand.On("CreateOnlineBankTicket", mock.Anything, mock.Anything).Return(nil)
	handlers.InitWorkerEventConflHandler(suite.workerUsecaseCommand, suite.mockLogger)
}
