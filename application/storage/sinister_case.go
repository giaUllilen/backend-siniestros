package storage

import (
	"context"
	"fmt"
	"is-public-api/application/colletions"
	"is-public-api/application/models"
	"is-public-api/helpers/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type sinisterCaseRepository struct {
	database *database.MongodbHelper
}

func NewSinisterCaseRepository(helper *database.MongodbHelper) *sinisterCaseRepository {
	return &sinisterCaseRepository{database: helper}
}

func (repo *sinisterCaseRepository) Save(txContext *models.TxContext, sinister map[string]interface{}) (string, error) {
	/* lastCase, err := repo.FindLastCase(txContext)
	if err != nil && err != mongo.ErrNoDocuments {
		return "", err
	} */
	/* var id []string
	if lastCase == nil || lastCase.Case == "" {
		id = []string{"CIS", "0"}
	} else {
		id = strings.Split(lastCase.Case, "_")
	} */
	sinister["created_date"] = time.Now()
	//caseNumber, _ := strconv.Atoi(id[1])
	//newCaseID := fmt.Sprintf("CIS_%05d", caseNumber+1)
	//sinister["case"] = newCaseID

	ctx := context.Background()
	options := new(options.InsertOneOptions)
	result, err := repo.database.Database(DatabaseBuc).Collection(CollectionSinisterCase).InsertOne(ctx, sinister, options)
	if err != nil {
		return "", err
	}

	// Intentar convertir el interface{} a primitive.ObjectID
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("el ID insertado no es un ObjectID válido")
	}
	// Convertir el ObjectID a string usando el método Hex()
	return objectID.Hex(), nil

}

func (repo *sinisterCaseRepository) UpdateOne(txContext *models.TxContext, caseNumber primitive.ObjectID, sinister map[string]interface{}) (primitive.ObjectID, error) {
	ctx := context.Background()

	filter := bson.M{"_id": caseNumber}
	sinister["updated_date"] = time.Now()
	update := bson.M{"$set": sinister}
	options := options.Update().SetUpsert(true)

	_, err := repo.database.Database(DatabaseBuc).Collection(CollectionSinisterCase).UpdateOne(ctx, filter, update, options)
	return caseNumber, err
}

func (repo *sinisterCaseRepository) FindLastCase(txContext *models.TxContext) (*colletions.SinisterCase, error) {
	err := repo.database.OpenConnection()
	var Return = new(colletions.SinisterCase)
	if err != nil {
		return Return, err
	}
	ctx := context.Background()
	options := new(options.FindOneOptions)

	options.SetSort(bson.D{{Key: "case", Value: -1}})
	err = repo.database.Database(DatabaseBuc).Collection(CollectionSinisterCase).FindOne(ctx, bson.M{}, options).Decode(&Return)
	if err != nil {
		return Return, err
	}

	return Return, nil
}

func (repo *sinisterCaseRepository) FindByCase(txContext *models.TxContext, caseNumber primitive.ObjectID) (*models.CasoSiniestro, error) {
	err := repo.database.OpenConnection()
	var sinisterCase = new(models.CasoSiniestro)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	filter := bson.M{"_id": caseNumber}
	options := new(options.FindOneOptions)
	options.SetSort(bson.D{{Key: "_id", Value: -1}})
	err = repo.database.Database(DatabaseBuc).Collection(CollectionSinisterCase).FindOne(ctx, filter, options).Decode(&sinisterCase)
	if err != nil {
		return nil, err
	}

	return sinisterCase, nil
}

func (repo *sinisterCaseRepository) FindAll(txContext *models.TxContext) ([]colletions.SinisterCase, error) {
	err := repo.database.OpenConnection()
	var cases []colletions.SinisterCase
	if err != nil {
		return cases, err
	}
	ctx := context.Background()
	filter := bson.M{}
	options := new(options.FindOptions)
	options.SetSort(bson.D{{Key: "_id", Value: -1}})
	cursor, err := repo.database.Database(DatabaseBuc).Collection(CollectionSinisterCase).Find(ctx, filter, options)
	if err != nil {
		return cases, err
	}

	err = cursor.All(ctx, &cases)
	if err != nil {
		return nil, err
	}
	return cases, nil
}
