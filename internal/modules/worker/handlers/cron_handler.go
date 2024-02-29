package handlers

import (
	"context"
	"time"
	"worker-service/internal/modules/worker"
	"worker-service/internal/pkg/log"

	"github.com/robfig/cron/v3"
)

type CronHttpHandler struct {
	WorkerUsecaseCommand worker.UsecaseCommand
	Logger               log.Logger
}

func InitCronHandler(wuc worker.UsecaseCommand, log log.Logger) {
	handler := &CronHttpHandler{
		WorkerUsecaseCommand: wuc,
		Logger:               log,
	}

	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	// scheduler := cron.New(cron.WithLocation(jakartaTime), cron.WithSeconds())
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	scheduler.AddFunc("*/5 * * * *", handler.UpdateAllExpiryPayment)
	scheduler.AddFunc("*/10 * * * *", handler.UpdateAllExpiryBankTicket)

	go scheduler.Start()
}

func (c CronHttpHandler) UpdateAllExpiryPayment() {
	ctx := context.Background()
	resp, err := c.WorkerUsecaseCommand.UpdateAllExpiryPayment(ctx)
	if err != nil {
		c.Logger.Error(ctx, "error UpdateAllExpiryPayment", err.Error())
	}
	if resp != nil {
		c.Logger.Info(ctx, *resp, "success UpdateAllExpiryPayment")
	}

}

func (c CronHttpHandler) UpdateAllExpiryBankTicket() {
	ctx := context.Background()
	resp, err := c.WorkerUsecaseCommand.UpdateAllExpiryBankTicket(ctx)
	if err != nil {
		c.Logger.Error(ctx, "error UpdateAllExpiryBankTicket", err.Error())
	}
	if resp != nil {
		c.Logger.Info(ctx, *resp, "success UpdateAllExpiryBankTicket")
	}

}
