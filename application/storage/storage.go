package storage

import (
	"is-public-api/application/colletions"
	"is-public-api/application/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const DatabaseBuc = "buc-data"
const CollectionCollaborator = "colaborador"
const CollectionSoatReturns = "soat_devoluciones"
const CollectionSoat = "soat"
const CollectionSoatHistory = "soat_historico"
const CollectionSinisterPayment = "siniestros_pagos"
const CollectionSinisterCase = "siniestros_casos"
const CollectionSubscriptionCenter = "centro_subscripcion"
const CollectionSubscriptionFrequency = "centro_subscripcion_frecuencia"

type ICollaboratorRepository interface {
	Find(txContext *models.TxContext, code string) (*colletions.Colaborador, error)
}

type ISoatReturnRepository interface {
	FindByDocument(txContext *models.TxContext, documentNumber string) (*colletions.SoatReturn, error)
}

type ISoatRepository interface {
	FindByPlate(txContext *models.TxContext, plate string) (*colletions.Soat, error)
	FindByPlateHistory(txContext *models.TxContext, plate string, dateOccurrence string) (*colletions.Soat, error)
}

type ISinisterPaymentRepository interface {
	FindByDocumentNumber(txContext *models.TxContext, documentNumber string) (*colletions.SinisterPayment, error)
}

type ISinisterCaseRepository interface {
	Save(txContext *models.TxContext, sinister map[string]interface{}) (string, error)
	UpdateOne(txContext *models.TxContext, caseNumber primitive.ObjectID, sinister map[string]interface{}) (primitive.ObjectID, error)
	FindLastCase(txContext *models.TxContext) (*colletions.SinisterCase, error)
	FindByCase(txContext *models.TxContext, caseNumber primitive.ObjectID) (*models.CasoSiniestro, error)
	FindAll(txContext *models.TxContext) ([]colletions.SinisterCase, error)
}

type ISubscriptionCenterRepository interface {
	CreateSubscriptionOptions(txContext *models.TxContext, key, option, value, version string, description *string) (string, error)
	FindSubscriptionOptionsByID(txContext *models.TxContext, key string) (string, error)
}
