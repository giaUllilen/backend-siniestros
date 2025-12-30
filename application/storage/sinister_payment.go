package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"is-public-api/application/colletions"
	"is-public-api/application/models"
	"is-public-api/helpers/database"
)

type sinisterPaymentRepository struct {
	database *database.MongodbHelper
}

func NewSinisterPaymentRepository(helper *database.MongodbHelper) *sinisterPaymentRepository {
	return &sinisterPaymentRepository{database: helper}
}

func (repo *sinisterPaymentRepository) FindByDocumentNumber(txContext *models.TxContext, documentNumber string) (*colletions.SinisterPayment, error) {
	err := repo.database.OpenConnection()
	var Return = new(colletions.SinisterPayment)
	if err != nil {
		return Return, err
	}
	ctx := context.Background()
	options := new(options.FindOneOptions)

	options.SetSort(bson.D{{"fecha_pago_time", -1}})
	filter := bson.M{"numero_documento": documentNumber}
	err = repo.database.Database("buc-data").Collection(CollectionSinisterPayment).FindOne(ctx, filter, options).Decode(&Return)
	if err != nil {
		return Return, err
	}

	return Return, nil
}