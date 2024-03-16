package commands_test

import (
	"context"
	"testing"
	"worker-service/internal/modules/worker"
	"worker-service/internal/modules/worker/models/entity"
	"worker-service/internal/modules/worker/models/request"
	mongoRC "worker-service/internal/modules/worker/repositories/commands"
	"worker-service/internal/pkg/helpers"
	mocks "worker-service/mocks/pkg/databases/mongodb"
	mocklog "worker-service/mocks/pkg/log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CommandTestSuite struct {
	suite.Suite
	mockMongodb *mocks.Collections
	mockLogger  *mocklog.Logger
	repository  worker.MongodbRepositoryCommand
	ctx         context.Context
}

func (suite *CommandTestSuite) SetupTest() {
	suite.mockMongodb = new(mocks.Collections)
	suite.mockLogger = &mocklog.Logger{}
	suite.repository = mongoRC.NewCommandMongodbRepository(
		suite.mockMongodb,
		suite.mockLogger,
	)
	suite.ctx = context.Background()
}

func TestCommandTestSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}

func (suite *CommandTestSuite) TestInsertManyTicketCollection() {

	// Mock UpsertOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("InsertMany", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.InsertManyTicketCollection(suite.ctx, mock.Anything, []entity.BankTicket{
		{
			TicketNumber: "1",
		},
	})
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert UpsertOne
	suite.mockMongodb.AssertCalled(suite.T(), "InsertMany", mock.Anything, mock.Anything)
}

func (suite *CommandTestSuite) TestDeleteOneOrder() {

	// Mock UpsertOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("DeleteOne", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.DeleteOneOrder(suite.ctx, mock.Anything)
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert UpsertOne
	suite.mockMongodb.AssertCalled(suite.T(), "DeleteOne", mock.Anything, mock.Anything)
}

func (suite *CommandTestSuite) TestUpdateOneBankTicket() {

	// Mock UpsertOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("UpdateOne", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.UpdateOneBankTicket(suite.ctx, request.UpdateBankTicketRequest{})
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert UpsertOne
	suite.mockMongodb.AssertCalled(suite.T(), "UpdateOne", mock.Anything, mock.Anything)
}

func (suite *CommandTestSuite) TestUpdateOnePayment() {

	// Mock UpsertOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("UpdateOne", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.UpdateOnePayment(suite.ctx, "id")
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert UpsertOne
	suite.mockMongodb.AssertCalled(suite.T(), "UpdateOne", mock.Anything, mock.Anything)
}

func (suite *CommandTestSuite) TestUpdateOnlineTicketConfig() {

	// Mock UpsertOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("UpdateOne", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.UpdateOnlineTicketConfig(suite.ctx, request.UpdateOnlineTicketConfigReq{})
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert UpsertOne
	suite.mockMongodb.AssertCalled(suite.T(), "UpdateOne", mock.Anything, mock.Anything)
}

func (suite *CommandTestSuite) TestUpdateTicketDetailByTag() {

	// Mock UpsertOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("UpdateOne", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.UpdateTicketDetailByTag(suite.ctx, request.UpdateTicketDetailReq{})
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert UpsertOne
	suite.mockMongodb.AssertCalled(suite.T(), "UpdateOne", mock.Anything, mock.Anything)
}

func (suite *CommandTestSuite) TestUpdateTicketDetailById() {

	// Mock UpsertOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("UpdateOne", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.UpdateTicketDetailById(suite.ctx, request.UpdateTicketDetailByIdReq{})
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert UpsertOne
	suite.mockMongodb.AssertCalled(suite.T(), "UpdateOne", mock.Anything, mock.Anything)
}
