package storage

import (
	"context"
	"is-public-api/application/colletions"
	"is-public-api/application/models"
	"is-public-api/helpers/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ISinisterCoverageIARepository define la interfaz del repositorio
type ISinisterCoverageIARepository interface {
	FindActiveByProductAndCoverage(txContext *models.TxContext, producto, cobertura string) (*models.SinisterCoverageIA, error)
	FindAllActive(txContext *models.TxContext) ([]models.SinisterCoverageIA, error)
}

type sinisterCoverageIARepository struct {
	database *database.MongodbHelper
}

// NewSinisterCoverageIARepository crea una nueva instancia del repositorio
func NewSinisterCoverageIARepository(helper *database.MongodbHelper) ISinisterCoverageIARepository {
	return &sinisterCoverageIARepository{database: helper}
}

// FindActiveByProductAndCoverage busca una configuración específica activa
func (repo *sinisterCoverageIARepository) FindActiveByProductAndCoverage(txContext *models.TxContext, producto, cobertura string) (*models.SinisterCoverageIA, error) {
	err := repo.database.OpenConnection()
	if err != nil {
		return nil, err
	}

	var result models.SinisterCoverageIA
	ctx := context.Background()

	filter := bson.M{
		"producto":  producto,
		"cobertura": cobertura,
		"activo":    true,
	}

	options := options.FindOne()
	err = repo.database.Database(DatabaseBuc).Collection(colletions.CollectionSinisterCoverageIA).FindOne(ctx, filter, options).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// FindAllActive obtiene todas las configuraciones activas
func (repo *sinisterCoverageIARepository) FindAllActive(txContext *models.TxContext) ([]models.SinisterCoverageIA, error) {
	err := repo.database.OpenConnection()
	if err != nil {
		return nil, err
	}

	var results []models.SinisterCoverageIA
	ctx := context.Background()

	filter := bson.M{"activo": true}
	options := options.Find().SetSort(bson.D{{Key: "producto", Value: 1}, {Key: "cobertura", Value: 1}})

	cursor, err := repo.database.Database(DatabaseBuc).Collection(colletions.CollectionSinisterCoverageIA).Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
