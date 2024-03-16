package queries_test

import (
	"context"
	"testing"
	"worker-service/internal/modules/worker"
	"worker-service/internal/modules/worker/models/entity"
	"worker-service/internal/modules/worker/models/request"
	mongoRQ "worker-service/internal/modules/worker/repositories/queries"
	"worker-service/internal/pkg/databases/mongodb"
	"worker-service/internal/pkg/helpers"
	mocks "worker-service/mocks/pkg/databases/mongodb"
	mocklog "worker-service/mocks/pkg/log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type QueryTestSuite struct {
	suite.Suite
	mockMongodb *mocks.Collections
	mockLogger  *mocklog.Logger
	repository  worker.MongodbRepositoryQuery
	ctx         context.Context
}

func (suite *QueryTestSuite) SetupTest() {
	suite.mockMongodb = new(mocks.Collections)
	suite.mockLogger = &mocklog.Logger{}
	suite.repository = mongoRQ.NewQueryMongodbRepository(
		suite.mockMongodb,
		suite.mockLogger,
	)
	suite.ctx = context.Background()
}

func TestCommandTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}

func (suite *QueryTestSuite) TestFindOneTicketDetail() {

	req := mongodb.FindOne{
		Result:         &entity.TicketDetail{},
		CollectionName: "ticket-detail",
		Filter: bson.M{
			"eventId": "", "ticketId": "",
		},
	}
	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("FindOne", req, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindOneTicketDetail(suite.ctx, request.CreateTicketReq{})
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "FindOne", req, mock.Anything)
}

func (suite *QueryTestSuite) TestFindOneTicketDetailByTag() {

	req := mongodb.FindOne{
		Result:         &entity.TicketDetail{},
		CollectionName: "ticket-detail",
		Filter: bson.M{
			"tag":          "",
			"ticketType":   "",
			"country.code": "",
		},
	}
	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("FindOne", req, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindOneTicketDetailByTag(suite.ctx, request.TicketDetailByTagReq{})
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "FindOne", req, mock.Anything)
}

func (suite *QueryTestSuite) TestFindOnlineTicketConfigByTag() {

	req := mongodb.FindOne{
		Result:         &entity.OnlineTicketConfig{},
		CollectionName: "online-ticket-config",
		Filter: bson.M{
			"tag": "",
		},
	}
	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("FindOne", req, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindOnlineTicketConfigByTag(suite.ctx, "")
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "FindOne", req, mock.Anything)
}

func (suite *QueryTestSuite) TestFindOneLastTicket() {

	req := mongodb.FindOne{
		Result:         &entity.BankTicket{},
		CollectionName: "mock",
		Filter: bson.M{
			"countryCode": "code",
			"ticketType":  "mock",
			"eventId":     "mock",
		},
		Sort: &mongodb.Sort{
			FieldName: "seatNumber",
			By:        mongodb.SortDescending,
		},
	}
	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("FindOne", req, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindOneLastTicket(suite.ctx, "code", "mock", "mock", "mock")
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "FindOne", req, mock.Anything)
}

func (suite *QueryTestSuite) TestFindAllExpirePayment() {

	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("FindAllData", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindAllExpirePayment(suite.ctx)
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "FindAllData", mock.Anything, mock.Anything)
}

func (suite *QueryTestSuite) TestFindBankTicketByTicketNumber() {

	req := mongodb.FindOne{
		Result:         &entity.BankTicket{},
		CollectionName: "bank-ticket",
		Filter: bson.M{
			"ticketNumber": "1",
		},
	}
	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("FindOne", req, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindBankTicketByTicketNumber(suite.ctx, "1")
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "FindOne", req, mock.Anything)
}

func (suite *QueryTestSuite) TestFindOneTicketDetailById() {

	req := mongodb.FindOne{
		Result:         &entity.TicketDetail{},
		CollectionName: "ticket-detail",
		Filter: bson.M{
			"ticketId": "1",
		},
	}
	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("FindOne", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindOneTicketDetailById(suite.ctx, "1")
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "FindOne", req, mock.Anything)
}

func (suite *QueryTestSuite) TestFindTotalAvalailableTicket() {

	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("Aggregate", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindTotalAvalailableTicket(suite.ctx, "tag")
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "Aggregate", mock.Anything, mock.Anything)
}

func (suite *QueryTestSuite) TestFindAllExpireBankTicket() {

	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("FindAllData", mock.Anything, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindAllExpireBankTicket(suite.ctx)
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "FindAllData", mock.Anything, mock.Anything)
}

func (suite *QueryTestSuite) TestFindPaymentByTicketNumber() {

	req := mongodb.FindOne{
		Result:         &entity.PaymentHistory{},
		CollectionName: "payment-history",
		Filter: bson.M{
			"ticket.ticketNumber": "1",
			"isValidPayment":      true,
		},
	}
	// Mock FindOne
	expectedResult := make(chan helpers.Result)
	suite.mockMongodb.On("FindOne", req, mock.Anything).Return((<-chan helpers.Result)(expectedResult))

	// Act
	result := suite.repository.FindPaymentByTicketNumber(suite.ctx, "1")
	// Asset
	assert.NotNil(suite.T(), result, "Expected a result")

	// Simulate receiving a result from the channel
	go func() {
		expectedResult <- helpers.Result{Data: "result not nil", Error: nil}
		close(expectedResult)
	}()

	// Wait for the goroutine to complete
	<-result

	// Assert FindOne
	suite.mockMongodb.AssertCalled(suite.T(), "FindOne", req, mock.Anything)
}
