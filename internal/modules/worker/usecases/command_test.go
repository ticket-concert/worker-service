package usecases_test

import (
	"context"
	"testing"
	"worker-service/internal/modules/worker"
	"worker-service/internal/pkg/errors"
	"worker-service/internal/pkg/helpers"

	"worker-service/internal/modules/worker/models/entity"
	"worker-service/internal/modules/worker/models/request"
	uc "worker-service/internal/modules/worker/usecases"
	mockcert "worker-service/mocks/modules/worker"
	mocklog "worker-service/mocks/pkg/log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CommandUsecaseTestSuite struct {
	suite.Suite
	mockWorkerRepositoryQuery   *mockcert.MongodbRepositoryQuery
	mockWorkerRepositoryCommand *mockcert.MongodbRepositoryCommand
	mockLogger                  *mocklog.Logger
	usecase                     worker.UsecaseCommand
	ctx                         context.Context
}

func (suite *CommandUsecaseTestSuite) SetupTest() {
	suite.mockWorkerRepositoryQuery = &mockcert.MongodbRepositoryQuery{}
	suite.mockWorkerRepositoryCommand = &mockcert.MongodbRepositoryCommand{}
	suite.mockLogger = &mocklog.Logger{}
	suite.ctx = context.Background()
	suite.usecase = uc.NewCommandUsecase(
		suite.mockWorkerRepositoryQuery,
		suite.mockWorkerRepositoryCommand,
		suite.mockLogger,
	)
}

func TestCommandUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(CommandUsecaseTestSuite))
}

// Helper function to create a channel
func mockChannel(result helpers.Result) <-chan helpers.Result {
	responseChan := make(chan helpers.Result)

	go func() {
		responseChan <- result
		close(responseChan)
	}()

	return responseChan
}

func (suite *CommandUsecaseTestSuite) TestCreateBankTicket() {
	payload := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetail", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	_, err := suite.usecase.CreateBankTicket(suite.ctx, payload)
	assert.NoError(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateBankTicketErrCompleted() {
	payload := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  0,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetail", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	_, err := suite.usecase.CreateBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateBankTicketErrDetail() {
	payload := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: errors.BadRequest("error"),
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetail", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	_, err := suite.usecase.CreateBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateBankTicketErrNilDetail() {
	payload := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}

	mockTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetail", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	_, err := suite.usecase.CreateBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateBankTicketErrParseDetail() {
	payload := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.Country{
			Name: "name",
		},
		Error: nil,
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetail", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	_, err := suite.usecase.CreateBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateBankTicketErrBank() {
	payload := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: errors.BadRequest("error"),
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetail", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	_, err := suite.usecase.CreateBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateBankTicketErrParseBank() {
	payload := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockBankTicket := helpers.Result{
		Data: &entity.Country{
			Name: "name",
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetail", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	_, err := suite.usecase.CreateBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateBankTicketErrInsert() {
	payload := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetail", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	_, err := suite.usecase.CreateBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPayment() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.NoError(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrHistory() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: errors.BadRequest("error"),
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrNilHistory() {
	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrParseHistory() {
	mockPaymentHistory := helpers.Result{
		Data: &entity.PaymentHistory{
			PaymentId:      "id",
			IsValidPayment: true,
			Ticket: &entity.Ticket{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrEmpty() {
	mockPaymentHistory := helpers.Result{
		Data:  &[]entity.PaymentHistory{},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.NoError(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrTicket() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: errors.BadRequest("error"),
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrNilTicket() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrParseTicket() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.Country{
			Name: "name",
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrUpdatePayment() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrDeleteOrder() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrUpdateBank() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrTotalRemaining() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:       "id",
			TotalQuota:     10,
			TicketPrice:    40,
			TotalRemaining: 10,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryPaymentErrUpdateTicketDetail() {
	mockPaymentHistory := helpers.Result{
		Data: &[]entity.PaymentHistory{
			{
				PaymentId:      "id",
				IsValidPayment: true,
				Ticket: &entity.Ticket{
					TicketNumber: "1",
					TicketId:     "id",
				},
			},
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdatePayment := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockDeleteOrder := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpirePayment", mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOnePayment", mock.Anything, mock.Anything).Return(mockChannel(mockUpdatePayment))
	suite.mockWorkerRepositoryCommand.On("DeleteOneOrder", mock.Anything, mock.Anything).Return(mockChannel(mockDeleteOrder))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryPayment(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicket() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.NoError(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErr() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: errors.BadRequest("error"),
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrNil() {
	mockBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrParse() {
	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "1",
			TicketId:     "id",
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrEmpty() {
	mockBankTicket := helpers.Result{
		Data:  &[]entity.BankTicket{},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.NoError(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrPayment() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketExistPayment() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data: &entity.PaymentHistory{
			PaymentId: "id",
		},
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.NoError(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrDetail() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: errors.BadRequest("error"),
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrNilDetail() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrParseDetail() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.Country{
			Name: "name",
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrUpdate() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrTotalRemaining() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:       "id",
			TotalQuota:     10,
			TicketPrice:    40,
			TotalRemaining: 10,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestUpdateAllExpiryBankTicketErrUpdateDetail() {
	mockBankTicket := helpers.Result{
		Data: &[]entity.BankTicket{
			{
				TicketNumber: "1",
				TicketId:     "id",
			},
		},
		Error: nil,
	}

	mockPaymentHistory := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTicketDetail := helpers.Result{
		Data: &entity.TicketDetail{
			TicketId:    "id",
			TotalQuota:  10,
			TicketPrice: 40,
			Country: entity.Country{
				Code: "code",
			},
		},
		Error: nil,
	}

	mockUpdateBankTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	suite.mockWorkerRepositoryQuery.On("FindAllExpireBankTicket", mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryQuery.On("FindPaymentByTicketNumber", mock.Anything, mock.Anything).Return(mockChannel(mockPaymentHistory))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockTicketDetail))
	suite.mockWorkerRepositoryCommand.On("UpdateOneBankTicket", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateBankTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailById", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.UpdateAllExpiryBankTicket(suite.ctx)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicket() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.NoError(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrConfig() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: errors.BadRequest("error"),
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrNil() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrParse() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.Country{
			Name: "name",
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrTotal() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
		Error: errors.BadRequest("error"),
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrNilTotal() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrParseTotal() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 1,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &entity.AggregateTotalTicket{
			Id:                   "1",
			TotalAvailableTicket: 0,
			TotalTicket:          100,
		},
		Error: nil,
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketSkip() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "c1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "c2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.NoError(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrExist() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 7,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrEmpty() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.NoError(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrDetail() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: errors.BadRequest("error"),
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrNilDetail() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data:  nil,
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrParseDetail() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.Country{
					Name: "name",
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrLast() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: errors.BadRequest("error"),
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrParseLast() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.Country{
			Name: "name",
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrInsert() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrUpdateOnline() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}

func (suite *CommandUsecaseTestSuite) TestCreateOnlineBankTicketErrUpdateDetail() {
	payload := request.CreateOnlineTicketReq{
		Tag:         "tag",
		CountryCode: "code",
	}

	mockOnlineTicketConfig := helpers.Result{
		Data: &entity.OnlineTicketConfig{
			Tag:        "tag",
			TotalQuota: 1000,
			CountryList: []entity.CountryList{
				{
					CountryNumber: 1,
					Percentage:    40,
					CountryCode:   "c1",
				},
				{
					CountryNumber: 2,
					Percentage:    20,
					CountryCode:   "c2",
				},
				{
					CountryNumber: 3,
					Percentage:    20,
					CountryCode:   "c3",
				},
				{
					CountryNumber: 4,
					Percentage:    20,
					CountryCode:   "c3",
				},
			},
		},
		Error: nil,
	}

	mockTotalAvailableTicket := helpers.Result{
		Data: &[]entity.AggregateTotalTicket{
			{
				Id:                   "1",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "2",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "3",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
			{
				Id:                   "4",
				TotalAvailableTicket: 0,
				TotalTicket:          100,
			},
		},
	}

	mockBankTicket := helpers.Result{
		Data: &entity.BankTicket{
			TicketNumber: "5",
			SeatNumber:   5,
		},
		Error: nil,
	}

	mockInsertManyTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateOnlineTicket := helpers.Result{
		Data:  nil,
		Error: nil,
	}

	mockUpdateTicketDetail := helpers.Result{
		Data:  nil,
		Error: errors.BadRequest("error"),
	}

	mockFindOneTicketDetailByTag := func(ctx context.Context, payload request.TicketDetailByTagReq) <-chan helpers.Result {
		responseChan := make(chan helpers.Result)

		go func() {
			responseChan <- helpers.Result{
				Data: &entity.TicketDetail{
					TicketId:       "id",
					EventId:        "id",
					Tag:            "tag",
					TicketType:     "Gold",
					TicketPrice:    40,
					TotalQuota:     100,
					TotalRemaining: 0,
					Country: entity.Country{
						Code: "code",
						Name: "name",
					},
				},
				Error: nil,
			}
			close(responseChan)
		}()

		return responseChan
	}

	suite.mockWorkerRepositoryQuery.On("FindOnlineTicketConfigByTag", mock.Anything, mock.Anything).Return(mockChannel(mockOnlineTicketConfig))
	suite.mockWorkerRepositoryQuery.On("FindTotalAvalailableTicket", mock.Anything, mock.Anything).Return(mockChannel(mockTotalAvailableTicket))
	suite.mockWorkerRepositoryQuery.On("FindOneTicketDetailByTag", mock.Anything, mock.Anything).Return(mockFindOneTicketDetailByTag)
	suite.mockWorkerRepositoryQuery.On("FindOneLastTicket", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockBankTicket))
	suite.mockWorkerRepositoryCommand.On("InsertManyTicketCollection", mock.Anything, mock.Anything, mock.Anything).Return(mockChannel(mockInsertManyTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateOnlineTicketConfig", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateOnlineTicket))
	suite.mockWorkerRepositoryCommand.On("UpdateTicketDetailByTag", mock.Anything, mock.Anything).Return(mockChannel(mockUpdateTicketDetail))
	suite.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything)
	suite.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything)

	_, err := suite.usecase.CreateOnlineBankTicket(suite.ctx, payload)
	assert.Error(suite.T(), err)
}
