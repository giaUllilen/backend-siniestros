package storage

import (
	"context"
	"fmt"
	"time"

	"is-public-api/application/colletions"
	"is-public-api/application/models"
	"is-public-api/helpers/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type soatRepository struct {
	database *database.MongodbHelper
}

func NewSoatRepository(helper *database.MongodbHelper) *soatRepository {
	return &soatRepository{database: helper}
}

func (repo *soatRepository) FindByPlate(txContext *models.TxContext, plate string) (*colletions.Soat, error) {
	err := repo.database.OpenConnection()
	var Return = new(colletions.Soat)
	if err != nil {
		return Return, err
	}
	ctx := context.Background()
	options := new(options.FindOneOptions)

	filter := bson.M{"placa": plate}
	err = repo.database.Database("buc-data").Collection(CollectionSoat).FindOne(ctx, filter, options).Decode(&Return)
	if err != nil {
		return Return, err
	}

	return Return, nil
}

// FindByPlateHistory busca un registro SOAT histórico por placa y fecha de ocurrencia
// plate: número de placa del vehículo
// dateOccurrence: fecha de ocurrencia en formato YYYY-MM-DD
// Retorna el registro SOAT encontrado o error si no se encuentra o hay problemas
func (repo *soatRepository) FindByPlateHistory(txContext *models.TxContext, plate string, dateOccurrence string) (*colletions.Soat, error) {
	// Validación de parámetros
	if plate == "" {
		return nil, fmt.Errorf("el número de placa no puede estar vacío")
	}
	if dateOccurrence == "" {
		return nil, fmt.Errorf("la fecha de ocurrencia no puede estar vacía")
	}

	// Abrir conexión a la base de datos
	if err := repo.database.OpenConnection(); err != nil {
		return nil, fmt.Errorf("error al abrir conexión a la base de datos: %w", err)
	}

	// Parsear fecha
	const dateLayout = "2006-01-02"
	parsedTime, err := time.Parse(dateLayout, dateOccurrence[:10])
	if err != nil {
		return nil, fmt.Errorf("error al parsear la fecha '%s': %w", dateOccurrence, err)
	}

	// Convertir a ISODate para MongoDB
	isoDate := primitive.NewDateTimeFromTime(parsedTime)

	// Construir filtro de búsqueda
	filter := bson.M{
		"placa": plate,
		"$and": []bson.M{
			{"fechaInicio": bson.M{"$lte": isoDate}},
			{"fechaFin": bson.M{"$gte": isoDate}},
		},
	}

	// Realizar búsqueda
	ctx := context.Background()
	result := new(colletions.Soat)
	err = repo.database.Database("buc-data").Collection(CollectionSoatHistory).FindOne(ctx, filter).Decode(result)

	// Manejo de errores específicos
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("no se encontró SOAT para la placa %s en la fecha %s", plate, dateOccurrence)
	}
	if err != nil {
		return nil, fmt.Errorf("error al buscar SOAT: %w", err)
	}

	return result, nil
}
