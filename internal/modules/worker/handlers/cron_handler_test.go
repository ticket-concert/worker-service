// user_http_handler_test.go

package handlers_test

import (
	"testing"
	"worker-service/internal/modules/worker/handlers"
	mockcert "worker-service/mocks/modules/worker"
	mocklog "worker-service/mocks/pkg/log"

	"github.com/stretchr/testify/suite"
)

type CronHandlerTestSuite struct {
	suite.Suite

	cUC         *mockcert.UsecaseCommand
	cLog        *mocklog.Logger
	cronHandler *handlers.CronHttpHandler
}

func (suite *CronHandlerTestSuite) SetupTestCron() {
	suite.cUC = new(mockcert.UsecaseCommand)
	suite.cLog = &mocklog.Logger{}
	suite.cronHandler = &handlers.CronHttpHandler{
		WorkerUsecaseCommand: suite.cUC,
		Logger:               suite.cLog,
	}
	handlers.InitCronHandler(suite.cUC, suite.cLog)
}

func TestCronHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CronHandlerTestSuite))
}

// func (suite *CronHandlerTestSuite) TestUpdateAllExpiryPayment() {
// 	// rs := "result"
// 	// suite.cUC.On("UpdateAllExpiryPayment", context.Background()).Return(&rs, nil)
// 	suite.cLog.On("Info", mock.Anything, mock.Anything, mock.Anything).Return(nil)
// 	suite.cLog.On("Error", mock.Anything, mock.Anything, mock.Anything).Return(nil)

// 	suite.cronHandler.UpdateAllExpiryPayment()
// }
