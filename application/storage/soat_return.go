package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"is-public-api/application/colletions"
	"is-public-api/application/models"
	"is-public-api/helpers/database"
)

type soatReturnRepository struct {
	database *database.MongodbHelper
}

func NewSoatReturnRepository(helper *database.MongodbHelper) *soatReturnRepository {
	return &soatReturnRepository{database: helper}
}

func (repo *soatReturnRepository) FindByDocument(txContext *models.TxContext, documentNumber string) (*colletions.SoatReturn, error) {
	err := repo.database.OpenConnection()
	var Return = new(colletions.SoatReturn)
	if err != nil {
		return Return, err
	}
	ctx := context.Background()
	options := new(options.FindOneOptions)

	filter := bson.M{"nro_documento": documentNumber}
	err = repo.database.Database("buc-data").Collection(CollectionSoatReturns).FindOne(ctx, filter, options).Decode(&Return)
	if err != nil {
		return Return, err
	}

	return Return, nil
}