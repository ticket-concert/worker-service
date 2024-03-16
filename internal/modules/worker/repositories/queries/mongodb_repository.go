package queries

import (
	"context"
	"time"
	"worker-service/internal/modules/worker"
	"worker-service/internal/modules/worker/models/entity"
	"worker-service/internal/modules/worker/models/request"
	"worker-service/internal/pkg/databases/mongodb"
	wrapper "worker-service/internal/pkg/helpers"
	"worker-service/internal/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type queryMongodbRepository struct {
	mongoDb mongodb.Collections
	logger  log.Logger
}

func NewQueryMongodbRepository(mongodb mongodb.Collections, log log.Logger) worker.MongodbRepositoryQuery {
	return &queryMongodbRepository{
		mongoDb: mongodb,
		logger:  log,
	}
}

func (q queryMongodbRepository) FindOneTicketDetail(ctx context.Context, payload request.CreateTicketReq) <-chan wrapper.Result {
	var ticketDetail entity.TicketDetail
	output := make(chan wrapper.Result)

	go func() {
		resp := <-q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &ticketDetail,
			CollectionName: "ticket-detail",
			Filter: bson.M{
				"ticketId": payload.TicketId,
				"eventId":  payload.EventId,
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (q queryMongodbRepository) FindOneTicketDetailByTag(ctx context.Context, payload request.TicketDetailByTagReq) <-chan wrapper.Result {
	var ticketDetail entity.TicketDetail
	output := make(chan wrapper.Result)

	go func() {
		resp := <-q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &ticketDetail,
			CollectionName: "ticket-detail",
			Filter: bson.M{
				"tag":          payload.Tag,
				"ticketType":   payload.TicketType,
				"country.code": payload.CountryCode,
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (q queryMongodbRepository) FindOnlineTicketConfigByTag(ctx context.Context, tag string) <-chan wrapper.Result {
	var ticketConfig entity.OnlineTicketConfig
	output := make(chan wrapper.Result)

	go func() {
		resp := <-q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &ticketConfig,
			CollectionName: "online-ticket-config",
			Filter: bson.M{
				"tag": tag,
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (q queryMongodbRepository) FindOneLastTicket(ctx context.Context, countryCode string, ticketType string, eventId string, collectionName string) <-chan wrapper.Result {
	var ticket entity.BankTicket
	output := make(chan wrapper.Result)

	go func() {
		resp := <-q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &ticket,
			CollectionName: collectionName,
			Filter: bson.M{
				"countryCode": countryCode,
				"ticketType":  ticketType,
				"eventId":     eventId,
			},
			Sort: &mongodb.Sort{
				FieldName: "seatNumber",
				By:        mongodb.SortDescending,
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (q queryMongodbRepository) FindAllExpirePayment(ctx context.Context) <-chan wrapper.Result {
	var payment []entity.PaymentHistory
	output := make(chan wrapper.Result)

	count := 15
	then := time.Now().Add(time.Duration(-count) * time.Minute)

	go func() {
		resp := <-q.mongoDb.FindAllData(mongodb.FindAllData{
			Result:         &payment,
			CollectionName: "payment-history",
			Filter: bson.M{
				"payment.transactionStatus": "pending",
				"isValidPayment":            true,
				"expiryTime": bson.M{
					"$lte": primitive.NewDateTimeFromTime(then),
				},
			},
			Sort: &mongodb.Sort{
				FieldName: "createdAt",
				By:        mongodb.SortAscending,
			},
			Page: 1,
			Size: 100,
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (q queryMongodbRepository) FindBankTicketByTicketNumber(ctx context.Context, ticketNumber string) <-chan wrapper.Result {
	var ticket entity.BankTicket
	output := make(chan wrapper.Result)

	go func() {
		resp := <-q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &ticket,
			CollectionName: "bank-ticket",
			Filter: bson.M{
				"ticketNumber": ticketNumber,
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (q queryMongodbRepository) FindOneTicketDetailById(ctx context.Context, id string) <-chan wrapper.Result {
	var ticketDetail entity.TicketDetail
	output := make(chan wrapper.Result)

	go func() {
		resp := <-q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &ticketDetail,
			CollectionName: "ticket-detail",
			Filter: bson.M{
				"ticketId": id,
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (q queryMongodbRepository) FindTotalAvalailableTicket(ctx context.Context, tag string) <-chan wrapper.Result {
	var ticket []entity.AggregateTotalTicket
	output := make(chan wrapper.Result)

	go func() {
		resp := <-q.mongoDb.Aggregate(mongodb.Aggregate{
			Result:         &ticket,
			CollectionName: "ticket-detail",
			Filter: []bson.M{
				{
					"$match": bson.M{
						"tag":        tag,
						"ticketType": bson.M{"$ne": "Online"},
					},
				},
				{
					"$group": bson.M{
						"_id":                  "$country.code",
						"countryName":          bson.M{"$first": "$country.name"},
						"totalAvailableTicket": bson.M{"$sum": "$totalRemaining"},
						"totalTicket":          bson.M{"$sum": "$totalQuota"},
					},
				},
				{
					"$sort": bson.M{
						"totalAvailableTicket": 1,
					},
				},
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (q queryMongodbRepository) FindAllExpireBankTicket(ctx context.Context) <-chan wrapper.Result {
	var bankTicket []entity.BankTicket
	output := make(chan wrapper.Result)

	count := 15
	then := time.Now().Add(time.Duration(-count) * time.Minute)

	go func() {
		resp := <-q.mongoDb.FindAllData(mongodb.FindAllData{
			Result:         &bankTicket,
			CollectionName: "bank-ticket",
			Filter: bson.M{
				"paymentStatus": "pending",
				"updatedAt": bson.M{
					"$lte": primitive.NewDateTimeFromTime(then),
				},
			},
			Sort: &mongodb.Sort{
				FieldName: "createdAt",
				By:        mongodb.SortAscending,
			},
			Page: 1,
			Size: 100,
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (q queryMongodbRepository) FindPaymentByTicketNumber(ctx context.Context, ticketNumber string) <-chan wrapper.Result {
	var payment entity.PaymentHistory
	output := make(chan wrapper.Result)

	go func() {
		resp := <-q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &payment,
			CollectionName: "payment-history",
			Filter: bson.M{
				"ticket.ticketNumber": ticketNumber,
				"isValidPayment":      true,
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}
