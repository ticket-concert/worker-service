package handlers

import (
	"worker-service/internal/modules/worker"
	"worker-service/internal/modules/worker/models/request"
	"worker-service/internal/pkg/errors"
	"worker-service/internal/pkg/helpers"
	"worker-service/internal/pkg/log"
	"worker-service/internal/pkg/redis"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type WorkerHttpHandler struct {
	WorkerUsecaseCommand worker.UsecaseCommand
	Logger               log.Logger
	Validator            *validator.Validate
}

func InitWorkerHttpHandler(app *fiber.App, wuc worker.UsecaseCommand, log log.Logger, redisClient redis.Collections) {
	handler := &WorkerHttpHandler{
		WorkerUsecaseCommand: wuc,
		Logger:               log,
		Validator:            validator.New(),
	}
	// middlewares := middlewares.NewMiddlewares(redisClient)
	route := app.Group("/api/worker")

	route.Post("/v1/ticket", handler.CreateBankTicket)
}

func (w WorkerHttpHandler) CreateBankTicket(c *fiber.Ctx) error {
	req := new(request.CreateTicketReq)
	if err := c.BodyParser(req); err != nil {
		return helpers.RespError(c, w.Logger, errors.BadRequest("bad request"))
	}

	if err := w.Validator.Struct(req); err != nil {
		return helpers.RespError(c, w.Logger, errors.BadRequest(err.Error()))
	}
	resp, err := w.WorkerUsecaseCommand.CreateBankTicket(c.Context(), *req)
	if err != nil {
		return helpers.RespCustomError(c, w.Logger, err)
	}
	return helpers.RespSuccess(c, w.Logger, resp, "Create bank ticket success")
}
