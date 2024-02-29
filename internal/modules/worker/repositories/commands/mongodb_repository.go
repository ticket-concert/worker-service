package commands

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
)

type commandMongodbRepository struct {
	mongoDb mongodb.Collections
	logger  log.Logger
}

func NewCommandMongodbRepository(mongodb mongodb.Collections, log log.Logger) worker.MongodbRepositoryCommand {
	return &commandMongodbRepository{
		mongoDb: mongodb,
		logger:  log,
	}
}

func (c commandMongodbRepository) InsertManyTicketCollection(ctx context.Context, collection string, ticket []entity.BankTicket) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	go func() {
		var documentsInsert []interface{}
		for _, v := range ticket {
			documentsInsert = append(documentsInsert, v)
		}
		resp := <-c.mongoDb.InsertMany(mongodb.InsertMany{
			CollectionName: collection,
			Documents:      documentsInsert,
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (c commandMongodbRepository) DeleteOneOrder(ctx context.Context, ticketNumber string) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	go func() {
		resp := <-c.mongoDb.DeleteOne(mongodb.DeleteOne{
			CollectionName: "order",
			Filter: bson.M{
				"ticketNumber": ticketNumber,
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (c commandMongodbRepository) UpdateOneBankTicket(ctx context.Context, payload request.UpdateBankTicketRequest) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	go func() {
		resp := <-c.mongoDb.UpdateOne(mongodb.UpdateOne{
			CollectionName: "bank-ticket",
			Filter: bson.M{
				"ticketNumber": payload.TicketNumber,
			},
			Document: bson.M{
				"isUsed":        false,
				"userId":        "",
				"queueId":       "",
				"paymentStatus": "",
				"price":         payload.Price,
				"updatedAt":     time.Now(),
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (c commandMongodbRepository) UpdateOnePayment(ctx context.Context, paymentId string) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	go func() {
		resp := <-c.mongoDb.UpdateOne(mongodb.UpdateOne{
			CollectionName: "payment-history",
			Filter: bson.M{
				"paymentId": paymentId,
			},
			Document: bson.M{
				"isValidPayment": false,
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (c commandMongodbRepository) UpdateOnlineTicketConfig(ctx context.Context, payload request.UpdateOnlineTicketConfigReq) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	go func() {
		resp := <-c.mongoDb.UpdateOne(mongodb.UpdateOne{
			CollectionName: "online-ticket-config",
			Filter: bson.M{
				"tag":                       payload.Tag,
				"countryList.countryNumber": payload.CountryNumber,
			},
			Document: bson.M{
				"countryList.$.countryCode": payload.CountryCode,
				"updatedAt":                 time.Now(),
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (c commandMongodbRepository) UpdateTicketDetailByTag(ctx context.Context, payload request.UpdateTicketDetailReq) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	go func() {
		resp := <-c.mongoDb.UpdateOne(mongodb.UpdateOne{
			CollectionName: "ticket-detail",
			Filter: bson.M{
				"tag":          payload.Tag,
				"ticketType":   payload.TicketType,
				"country.code": payload.CountryCode,
			},
			Document: bson.M{
				"totalQuota":     payload.TotalQuota,
				"totalRemaining": payload.TotalRemaining,
				"updatedAt":      time.Now(),
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}

func (c commandMongodbRepository) UpdateTicketDetailById(ctx context.Context, payload request.UpdateTicketDetailByIdReq) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	go func() {
		resp := <-c.mongoDb.UpdateOne(mongodb.UpdateOne{
			CollectionName: "ticket-detail",
			Filter: bson.M{
				"ticketId": payload.TicketId,
			},
			Document: bson.M{
				"totalRemaining": payload.TotalRemaining,
				"updatedAt":      time.Now(),
			},
		}, ctx)
		output <- resp
		close(output)
	}()

	return output
}
