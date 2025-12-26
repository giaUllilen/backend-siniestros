package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"is-public-api/application/colletions"
	"is-public-api/application/models"
	"is-public-api/helpers/database"
)

type subscriptionCenterRepository struct {
	database *database.MongodbHelper
}

func NewSubscriptionCenterRepository(helper *database.MongodbHelper) ISubscriptionCenterRepository {
	return &subscriptionCenterRepository{database: helper}
}

func (repo *subscriptionCenterRepository) CreateSubscriptionOptions(txContext *models.TxContext, key, option, value, version string, description *string) (string, error) {
	var document colletions.SubscriptionCenter
	err := repo.database.OpenConnection()
	if err != nil {
		return "", err
	}
	document.CreatedAt = time.Now()
	document.Key = key
	document.Option = option
	document.Value = value
	document.Type = version
	document.Description = description
	ctx := context.Background()
	options := new(options.InsertOneOptions)

	result, err := repo.database.Database("buc-data").Collection(CollectionSubscriptionCenter).InsertOne(ctx, document, options)
	if err != nil {
		return "", err
	}

	return strings.Split(fmt.Sprintf("%v", result.InsertedID), "\"")[1], nil
}

func (repo *subscriptionCenterRepository) FindSubscriptionOptionsByID(txContext *models.TxContext, id string) (string, error) {
	var subscriptionOptions []colletions.SubscriptionCenter
	err := repo.database.OpenConnection()
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	options := new(options.FindOptions)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	cursor, err := repo.database.Database("buc-data").Collection(CollectionSubscriptionFrequency).Find(ctx, filter, options)
	if err != nil {
		return "", err
	}

	err = cursor.All(ctx, &subscriptionOptions)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (repo *soatReturnRepository) GetAllActiveFrequencies(txContext *models.TxContext) ([]colletions.SubscriptionCenterFrequency, error) {
	var optDownOptions []colletions.SubscriptionCenterFrequency
	err := repo.database.OpenConnection()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	options := new(options.FindOptions)

	options.SetSort(bson.D{{"ordinal", 1}})

	filter := bson.M{"enabled": true}
	cursor, err := repo.database.Database("buc-data").Collection(CollectionSubscriptionFrequency).Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &optDownOptions)
	if err != nil {
		return nil, err
	}
	return optDownOptions, nil
}