// user_http_handler_test.go

package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"worker-service/internal/modules/worker/handlers"
	"worker-service/internal/modules/worker/models/request"
	"worker-service/internal/pkg/errors"
	mockcert "worker-service/mocks/modules/worker"
	mocklog "worker-service/mocks/pkg/log"
	mockredis "worker-service/mocks/pkg/redis"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
)

type WorkerHttpHandlerTestSuite struct {
	suite.Suite

	cUC       *mockcert.UsecaseCommand
	cLog      *mocklog.Logger
	validator *validator.Validate
	cRedis    *mockredis.Collections
	handler   *handlers.WorkerHttpHandler
	app       *fiber.App
}

func (suite *WorkerHttpHandlerTestSuite) SetupTest() {
	suite.cUC = new(mockcert.UsecaseCommand)
	suite.cLog = new(mocklog.Logger)
	suite.validator = validator.New()
	suite.cRedis = new(mockredis.Collections)
	suite.handler = &handlers.WorkerHttpHandler{
		WorkerUsecaseCommand: suite.cUC,
		Logger:               suite.cLog,
		Validator:            suite.validator,
	}
	suite.app = fiber.New()
	handlers.InitWorkerHttpHandler(suite.app, suite.cUC, suite.cLog, suite.cRedis)
}

func TestUserHttpHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerHttpHandlerTestSuite))
}

func (suite *WorkerHttpHandlerTestSuite) TestCreateBankTicket() {
	rs := "result"
	suite.cUC.On("CreateBankTicket", mock.Anything, mock.Anything).Return(&rs, nil)
	suite.cLog.On("Info", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	reqM := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}
	requestBody, _ := json.Marshal(reqM)
	req := httptest.NewRequest(fiber.MethodPost, "/v1/ticket", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	ctx := suite.app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().SetRequestURI("/v1/ticket")
	ctx.Request().Header.SetMethod(fiber.MethodPost)
	ctx.Request().Header.SetContentType("application/json")
	ctx.Request().SetBody(requestBody)

	err := suite.handler.CreateBankTicket(ctx)
	assert.Nil(suite.T(), err)
}

func (suite *WorkerHttpHandlerTestSuite) TestCreateBankTicketErrBodyParser() {
	rs := "result"
	suite.cUC.On("CreateBankTicket", mock.Anything, mock.Anything).Return(&rs, nil)
	suite.cLog.On("Info", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	reqM := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}
	requestBody, _ := json.Marshal(reqM)
	req := httptest.NewRequest(fiber.MethodPost, "/v1/ticket", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	ctx := suite.app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().SetRequestURI("/v1/ticket")
	ctx.Request().Header.SetMethod(fiber.MethodPost)
	ctx.Request().Header.SetContentType("application/json")
	// ctx.Request().SetBody(requestBody)

	err := suite.handler.CreateBankTicket(ctx)
	assert.Nil(suite.T(), err)
}

func (suite *WorkerHttpHandlerTestSuite) TestCreateBankTicketErrValidate() {
	rs := "result"
	suite.cUC.On("CreateBankTicket", mock.Anything, mock.Anything).Return(&rs, nil)
	suite.cLog.On("Info", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	reqM := request.CreateTicketReq{}
	requestBody, _ := json.Marshal(reqM)
	req := httptest.NewRequest(fiber.MethodPost, "/v1/ticket", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	ctx := suite.app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().SetRequestURI("/v1/ticket")
	ctx.Request().Header.SetMethod(fiber.MethodPost)
	ctx.Request().Header.SetContentType("application/json")
	ctx.Request().SetBody(requestBody)

	err := suite.handler.CreateBankTicket(ctx)
	assert.Nil(suite.T(), err)
}

func (suite *WorkerHttpHandlerTestSuite) TestCreateBankTicketErr() {
	suite.cUC.On("CreateBankTicket", mock.Anything, mock.Anything).Return(nil, errors.BadRequest("error"))
	suite.cLog.On("Info", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	reqM := request.CreateTicketReq{
		TicketId: "id",
		EventId:  "id",
	}
	requestBody, _ := json.Marshal(reqM)
	req := httptest.NewRequest(fiber.MethodPost, "/v1/ticket", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	ctx := suite.app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().SetRequestURI("/v1/ticket")
	ctx.Request().Header.SetMethod(fiber.MethodPost)
	ctx.Request().Header.SetContentType("application/json")
	ctx.Request().SetBody(requestBody)

	err := suite.handler.CreateBankTicket(ctx)
	assert.Nil(suite.T(), err)
}
