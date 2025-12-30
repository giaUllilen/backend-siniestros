package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"is-public-api/application/colletions"
	"is-public-api/application/models"
	"is-public-api/helpers/database"
)

type customerRepository struct {
	database *database.MongodbHelper
}

func NewCustomerRepository(helper *database.MongodbHelper) *customerRepository {
	return &customerRepository{database: helper}
}

func (repo *customerRepository) Find(txContext *models.TxContext, code string) (*colletions.Colaborador, error) {
	err := repo.database.OpenConnection()
	var photocheck = new(colletions.Colaborador)
	if err != nil {
		return photocheck, err
	}
	ctx := context.Background()
	options := new(options.FindOneOptions)

	filter := bson.M{"estado": "ACTIVO", "documentoIdentidad": code}
	err = repo.database.Database("buc-data").Collection(CollectionCollaborator).FindOne(ctx, filter, options).Decode(&photocheck)
	if err != nil {
		return photocheck, err
	}

	return photocheck, nil
}