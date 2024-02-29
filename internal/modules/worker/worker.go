package worker

import (
	"context"
	"worker-service/internal/modules/worker/models/entity"
	"worker-service/internal/modules/worker/models/request"
	wrapper "worker-service/internal/pkg/helpers"
)

type UsecaseCommand interface {
	CreateBankTicket(origCtx context.Context, payload request.CreateTicketReq) (*string, error)
	UpdateAllExpiryPayment(origCtx context.Context) (*string, error)
	CreateOnlineBankTicket(origCtx context.Context, payload request.CreateOnlineTicketReq) (*string, error)
	UpdateAllExpiryBankTicket(origCtx context.Context) (*string, error)
}

type MongodbRepositoryQuery interface {
	FindOneTicketDetail(ctx context.Context, payload request.CreateTicketReq) <-chan wrapper.Result
	FindOneLastTicket(ctx context.Context, countryCode string, ticketType string, eventId string, collectionName string) <-chan wrapper.Result
	FindAllExpirePayment(ctx context.Context) <-chan wrapper.Result
	FindAllExpireBankTicket(ctx context.Context) <-chan wrapper.Result
	FindBankTicketByTicketNumber(ctx context.Context, ticketNumber string) <-chan wrapper.Result
	FindOneTicketDetailById(ctx context.Context, id string) <-chan wrapper.Result
	FindOnlineTicketConfigByTag(ctx context.Context, tag string) <-chan wrapper.Result
	FindTotalAvalailableTicket(ctx context.Context, tag string) <-chan wrapper.Result
	FindOneTicketDetailByTag(ctx context.Context, payload request.TicketDetailByTagReq) <-chan wrapper.Result
	FindPaymentByTicketNumber(ctx context.Context, ticketNumber string) <-chan wrapper.Result
}

type MongodbRepositoryCommand interface {
	InsertManyTicketCollection(ctx context.Context, collection string, ticket []entity.BankTicket) <-chan wrapper.Result
	DeleteOneOrder(ctx context.Context, ticketNumber string) <-chan wrapper.Result
	UpdateOneBankTicket(ctx context.Context, payload request.UpdateBankTicketRequest) <-chan wrapper.Result
	UpdateOnePayment(ctx context.Context, paymentId string) <-chan wrapper.Result
	UpdateOnlineTicketConfig(ctx context.Context, payload request.UpdateOnlineTicketConfigReq) <-chan wrapper.Result
	UpdateTicketDetailByTag(ctx context.Context, payload request.UpdateTicketDetailReq) <-chan wrapper.Result
	UpdateTicketDetailById(ctx context.Context, payload request.UpdateTicketDetailByIdReq) <-chan wrapper.Result
}
