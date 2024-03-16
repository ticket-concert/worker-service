package usecases

import (
	"context"
	"fmt"
	"time"
	"worker-service/internal/modules/worker"
	"worker-service/internal/modules/worker/models/dto"
	"worker-service/internal/modules/worker/models/entity"
	"worker-service/internal/modules/worker/models/request"
	"worker-service/internal/pkg/constants"
	"worker-service/internal/pkg/errors"
	"worker-service/internal/pkg/log"

	"github.com/google/uuid"
	"go.elastic.co/apm"
)

type commandUsecase struct {
	workerRepositoryQuery   worker.MongodbRepositoryQuery
	workerRepositoryCommand worker.MongodbRepositoryCommand
	logger                  log.Logger
}

func NewCommandUsecase(wrq worker.MongodbRepositoryQuery, wrc worker.MongodbRepositoryCommand, log log.Logger) worker.UsecaseCommand {
	return commandUsecase{
		workerRepositoryQuery:   wrq,
		workerRepositoryCommand: wrc,
		logger:                  log,
	}
}

func (c commandUsecase) CreateBankTicket(origCtx context.Context, payload request.CreateTicketReq) (*string, error) {
	domain := "workerUsecase-CreateBankTicket"
	span, ctx := apm.StartSpanOptions(origCtx, domain, "function", apm.SpanOptions{
		Start:  time.Now(),
		Parent: apm.TraceContext{},
	})
	defer span.End()

	ticketDetailData := <-c.workerRepositoryQuery.FindOneTicketDetail(ctx, payload)
	if ticketDetailData.Error != nil {
		return nil, ticketDetailData.Error
	}
	if ticketDetailData.Data == nil {
		msg := "Price not found"
		return nil, errors.NotFound(msg)
	}

	ticketDetail, ok := ticketDetailData.Data.(*entity.TicketDetail)
	if !ok {
		return nil, errors.InternalServerError("cannot parsing data")
	}

	state := 1
	collection := "bank-ticket"
	lastTicket := <-c.workerRepositoryQuery.FindOneLastTicket(ctx, ticketDetail.Country.Code, ticketDetail.TicketType, ticketDetail.EventId, collection)
	if lastTicket.Error != nil {
		return nil, lastTicket.Error
	}
	if lastTicket.Data != nil {
		ticket, ok := lastTicket.Data.(*entity.BankTicket)
		if !ok {
			return nil, errors.InternalServerError("cannot parsing data")
		}
		state = ticket.SeatNumber + 1
	}

	var results = make([]entity.BankTicket, 0)
	counter := ticketDetail.TotalQuota
	for i := state; i <= counter; i++ {
		// Create a ticket map and append it to results
		ticket := entity.BankTicket{
			TicketNumber: uuid.NewString(),
			SeatNumber:   i,
			IsUsed:       false,
			TicketId:     ticketDetail.TicketId,
			EventId:      ticketDetail.EventId,
			CountryCode:  ticketDetail.Country.Code,
			Price:        ticketDetail.TicketPrice,
			TicketType:   ticketDetail.TicketType,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		results = append(results, ticket)
	}

	respTicket := <-c.workerRepositoryCommand.InsertManyTicketCollection(ctx, collection, results)
	if respTicket.Error != nil {
		return nil, respTicket.Error
	}

	rs := "Success create bank ticket"
	return &rs, nil
}

func (c commandUsecase) UpdateAllExpiryPayment(origCtx context.Context) (*string, error) {
	domain := "workerUsecase-UpdateAllExpiryPayment"
	span, ctx := apm.StartSpanOptions(origCtx, domain, "function", apm.SpanOptions{
		Start:  time.Now(),
		Parent: apm.TraceContext{},
	})
	defer span.End()

	paymentData := <-c.workerRepositoryQuery.FindAllExpirePayment(ctx)
	if paymentData.Error != nil {
		return nil, paymentData.Error
	}

	if paymentData.Data == nil {
		return nil, errors.BadRequest("payment not found")
	}

	payments, ok := paymentData.Data.(*[]entity.PaymentHistory)
	if !ok {
		return nil, errors.InternalServerError("cannot parsing data payment")
	}

	result := "Success delete expired payment ticket"
	if len(*payments) == 0 {
		result = "Expiry payment ticket empty"
		c.logger.Info(ctx, result, payments)
		return &result, nil
	}

	for _, p := range *payments {
		ticketNumber := p.Ticket.TicketNumber
		ticketDetailData := <-c.workerRepositoryQuery.FindOneTicketDetailById(ctx, p.Ticket.TicketId)
		if ticketDetailData.Error != nil {
			return nil, ticketDetailData.Error
		}

		if ticketDetailData.Data == nil {
			return nil, errors.BadRequest("ticket not found")
		}

		ticketDetail, ok := ticketDetailData.Data.(*entity.TicketDetail)
		if !ok {
			return nil, errors.InternalServerError("cannot parsing data ticket")
		}

		c.logger.Info(ctx, "Payment Expired", p)

		updatePaymentResp := <-c.workerRepositoryCommand.UpdateOnePayment(ctx, p.PaymentId)
		if updatePaymentResp.Error != nil {
			return nil, updatePaymentResp.Error
		}

		c.logger.Info(ctx, "Deleted Ticket Order", ticketNumber)

		deleteOrderResp := <-c.workerRepositoryCommand.DeleteOneOrder(ctx, ticketNumber)
		if deleteOrderResp.Error != nil {
			return nil, deleteOrderResp.Error
		}

		bankTicketReq := request.UpdateBankTicketRequest{
			TicketNumber: ticketNumber,
			Price:        ticketDetail.TicketPrice,
		}
		bankTicketResp := <-c.workerRepositoryCommand.UpdateOneBankTicket(ctx, bankTicketReq)
		if bankTicketResp.Error != nil {
			return nil, bankTicketResp.Error
		}

		totalRemaining := ticketDetail.TotalRemaining + 1
		if totalRemaining > ticketDetail.TotalQuota {
			return nil, errors.BadRequest("totalRemaining full")
		}
		ticketDetailReq := request.UpdateTicketDetailByIdReq{
			TicketId:       ticketDetail.TicketId,
			TotalRemaining: totalRemaining,
		}
		ticketDetailResp := <-c.workerRepositoryCommand.UpdateTicketDetailById(ctx, ticketDetailReq)
		if ticketDetailResp.Error != nil {
			return nil, ticketDetailResp.Error
		}
	}

	return &result, nil

}

func (c commandUsecase) UpdateAllExpiryBankTicket(origCtx context.Context) (*string, error) {
	domain := "workerUsecase-UpdateAllExpiryBankTicket"
	span, ctx := apm.StartSpanOptions(origCtx, domain, "function", apm.SpanOptions{
		Start:  time.Now(),
		Parent: apm.TraceContext{},
	})
	defer span.End()

	bankTicketData := <-c.workerRepositoryQuery.FindAllExpireBankTicket(ctx)
	if bankTicketData.Error != nil {
		return nil, bankTicketData.Error
	}

	if bankTicketData.Data == nil {
		return nil, errors.BadRequest("bank ticket not found")
	}

	bankTickets, ok := bankTicketData.Data.(*[]entity.BankTicket)
	if !ok {
		return nil, errors.InternalServerError("cannot parsing data bank ticket")
	}

	result := "Success delete expired bank ticket"
	if len(*bankTickets) == 0 {
		result = "Expiry bank ticket empty"
		c.logger.Info(ctx, result, bankTickets)
		return &result, nil
	}

	for _, b := range *bankTickets {
		ticketNumber := b.TicketNumber
		paymentData := <-c.workerRepositoryQuery.FindPaymentByTicketNumber(ctx, ticketNumber)
		if paymentData.Error != nil {
			return nil, paymentData.Error
		}
		if paymentData.Data != nil {
			c.logger.Info(ctx, "Skip ticketNumber: ", b)
			continue
		}

		ticketDetailData := <-c.workerRepositoryQuery.FindOneTicketDetailById(ctx, b.TicketId)
		if ticketDetailData.Error != nil {
			return nil, ticketDetailData.Error
		}

		if ticketDetailData.Data == nil {
			return nil, errors.BadRequest("ticket not found")
		}

		ticketDetail, ok := ticketDetailData.Data.(*entity.TicketDetail)
		if !ok {
			return nil, errors.InternalServerError("cannot parsing data ticket")
		}

		c.logger.Info(ctx, "Bank Ticket Expired", b)

		bankTicketReq := request.UpdateBankTicketRequest{
			TicketNumber: ticketNumber,
			Price:        ticketDetail.TicketPrice,
		}
		bankTicketResp := <-c.workerRepositoryCommand.UpdateOneBankTicket(ctx, bankTicketReq)
		if bankTicketResp.Error != nil {
			return nil, bankTicketResp.Error
		}

		totalRemaining := ticketDetail.TotalRemaining + 1
		if totalRemaining > ticketDetail.TotalQuota {
			return nil, errors.BadRequest("totalRemaining full")
		}

		ticketDetailReq := request.UpdateTicketDetailByIdReq{
			TicketId:       ticketDetail.TicketId,
			TotalRemaining: totalRemaining,
		}
		ticketDetailResp := <-c.workerRepositoryCommand.UpdateTicketDetailById(ctx, ticketDetailReq)
		if ticketDetailResp.Error != nil {
			return nil, ticketDetailResp.Error
		}
	}

	return &result, nil

}

func (c commandUsecase) CreateOnlineBankTicket(origCtx context.Context, payload request.CreateOnlineTicketReq) (*string, error) {
	domain := "workerUsecase-CreateOnlineBankTicket"
	span, ctx := apm.StartSpanOptions(origCtx, domain, "function", apm.SpanOptions{
		Start:  time.Now(),
		Parent: apm.TraceContext{},
	})
	defer span.End()

	ticketConfigData := <-c.workerRepositoryQuery.FindOnlineTicketConfigByTag(ctx, payload.Tag)
	if ticketConfigData.Error != nil {
		return nil, ticketConfigData.Error
	}
	if ticketConfigData.Data == nil {
		msg := "ticket config not found"
		return nil, errors.NotFound(msg)
	}

	ticketConfig, ok := ticketConfigData.Data.(*entity.OnlineTicketConfig)
	if !ok {
		return nil, errors.InternalServerError("cannot parsing data")
	}

	totalTicketAvailable := <-c.workerRepositoryQuery.FindTotalAvalailableTicket(ctx, payload.Tag)
	if totalTicketAvailable.Error != nil {
		return nil, totalTicketAvailable.Error
	}

	if totalTicketAvailable.Data == nil {
		return nil, errors.BadRequest("ticket not found")
	}

	ticketAvailable, ok := totalTicketAvailable.Data.(*[]entity.AggregateTotalTicket)
	if !ok {
		return nil, errors.InternalServerError("cannot parsing data ticket")
	}

	var percentage1, percentage2, percentage3, percentage4 int
	for _, v := range ticketConfig.CountryList {
		if v.CountryNumber == 1 {
			percentage1 = v.Percentage
		}
		if v.CountryNumber == 2 {
			percentage2 = v.Percentage
		}
		if v.CountryNumber == 3 {
			percentage3 = v.Percentage
		}
		if v.CountryNumber == 4 {
			percentage4 = v.Percentage
		}
	}
	countries := make([]dto.CountryQuota, 0)
	for i, v := range *ticketAvailable {
		isSkiped := false
		for _, t := range ticketConfig.CountryList {
			if v.Id == t.CountryCode {
				isSkiped = true
			}
		}
		if isSkiped {
			continue
		}
		if i == 0 {
			if v.TotalAvailableTicket != 0 {
				return nil, errors.BadRequest("offline ticket still ready")
			}
			countries = append(countries, dto.CountryQuota{
				CountryCode:   v.Id,
				CountryNumber: i + 1,
				TotalQuota:    ticketConfig.TotalQuota * percentage1 / 100,
			})
		}
		if i == 1 {
			countries = append(countries, dto.CountryQuota{
				CountryCode:   v.Id,
				CountryNumber: i + 1,
				TotalQuota:    ticketConfig.TotalQuota * percentage2 / 100,
			})
		}
		if i == 2 {
			countries = append(countries, dto.CountryQuota{
				CountryCode:   v.Id,
				CountryNumber: i + 1,
				TotalQuota:    ticketConfig.TotalQuota * percentage3 / 100,
			})
		}
		if i == 3 {
			countries = append(countries, dto.CountryQuota{
				CountryCode:   v.Id,
				CountryNumber: i + 1,
				TotalQuota:    ticketConfig.TotalQuota * percentage4 / 100,
			})
		}
	}

	if len(countries) == 0 {
		result := "update online ticket empty"
		return &result, nil
	}

	for _, country := range countries {
		state := 1
		collection := "bank-ticket"
		ticketDetailReq := request.TicketDetailByTagReq{
			Tag:         payload.Tag,
			CountryCode: country.CountryCode,
			TicketType:  constants.Online,
		}

		ticketDetailData := <-c.workerRepositoryQuery.FindOneTicketDetailByTag(ctx, ticketDetailReq)
		if ticketDetailData.Error != nil {
			return nil, ticketDetailData.Error
		}

		if ticketDetailData.Data == nil {
			return nil, errors.BadRequest("ticket detail not found")
		}

		ticketDetail, ok := ticketDetailData.Data.(*entity.TicketDetail)
		if !ok {
			return nil, errors.InternalServerError("cannot parsing data ticket")
		}
		lastTicket := <-c.workerRepositoryQuery.FindOneLastTicket(ctx, country.CountryCode, constants.Online, ticketDetail.EventId, collection)
		if lastTicket.Error != nil {
			return nil, lastTicket.Error
		}
		if lastTicket.Data != nil {
			ticket, ok := lastTicket.Data.(*entity.BankTicket)
			if !ok {
				return nil, errors.InternalServerError("cannot parsing data")
			}
			state = ticket.SeatNumber + 1
		}

		fmt.Println("lastTicket.Data: ", lastTicket.Data)

		fmt.Println(state)

		var results = make([]entity.BankTicket, 0)
		counter := country.TotalQuota
		for i := state; i <= counter; i++ {
			// Create a ticket map and append it to results
			ticket := entity.BankTicket{
				TicketNumber: uuid.NewString(),
				SeatNumber:   i,
				IsUsed:       false,
				TicketId:     ticketDetail.TicketId,
				EventId:      ticketDetail.EventId,
				CountryCode:  country.CountryCode,
				Price:        ticketDetail.TicketPrice,
				TicketType:   ticketDetail.TicketType,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
			results = append(results, ticket)
		}

		respTicket := <-c.workerRepositoryCommand.InsertManyTicketCollection(ctx, collection, results)
		if respTicket.Error != nil {
			return nil, respTicket.Error
		}

		updateTicketConfigReq := request.UpdateOnlineTicketConfigReq{
			Tag:           payload.Tag,
			CountryNumber: country.CountryNumber,
			CountryCode:   country.CountryCode,
		}
		respConfig := <-c.workerRepositoryCommand.UpdateOnlineTicketConfig(ctx, updateTicketConfigReq)
		if respConfig.Error != nil {
			c.logger.Error(ctx, "Failed UpdateOnlineTicketConfig", respConfig)
			return nil, respConfig.Error
		}
		c.logger.Info(ctx, "Success UpdateOnlineTicketConfig", respConfig)

		updateTicketDetailReq := request.UpdateTicketDetailReq{
			Tag:            payload.Tag,
			TicketType:     constants.Online,
			CountryCode:    country.CountryCode,
			TotalQuota:     country.TotalQuota,
			TotalRemaining: (country.TotalQuota - ticketDetail.TotalQuota) + ticketDetail.TotalRemaining,
		}

		respTicketDetail := <-c.workerRepositoryCommand.UpdateTicketDetailByTag(ctx, updateTicketDetailReq)
		if respTicketDetail.Error != nil {
			c.logger.Error(ctx, "Failed UpdateTicketDetailByTag", respTicketDetail)
			return nil, respTicketDetail.Error
		}
		c.logger.Info(ctx, "Success UpdateTicketDetailByTag", respTicketDetail)

	}

	rs := "Success create bank ticket online"
	return &rs, nil
}
