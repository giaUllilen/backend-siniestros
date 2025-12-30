package endpoints

import (
	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/helpers/apihelpers"
)

type IEventsEndpoints interface {
	AddEvent(txContext *models.TxContext) (int, []byte, error)
}

type IStorageEndpoints interface {
	Upload(txContext *models.TxContext, form *resources.MapString, files apihelpers.File) ([]byte, error)
}

type INotificationEndpoints interface {
	SendMail(txContext *models.TxContext, request *resources.MapRequest) ([]byte, error)
}

type ISinisterEndpoints interface {
	Save(txContext *models.TxContext, body *resources.MapRequest) (int, *resources.ResponseSinister, error)
	AddDocument(txContext *models.TxContext, body *resources.MapRequest) (int, []byte, error)
	// FindByCaseNumber - Cambiado de primitive.ObjectID a string para enviar el número de caso directamente (ej: "CIS_98746")
	// Anterior: FindByCaseNumber(txContext *models.TxContext, caseNumber primitive.ObjectID) (int, []byte, error)
	FindByCaseNumber(txContext *models.TxContext, caseNumber string) (int, []byte, error)
	Delete(txContext *models.TxContext, caseNumber string) (int, []byte, error)
	FindByCaseHistory(txContext *models.TxContext, request models.SinisterHistoryRequest) (int, []models.SinisterHistory, error)
	UpdateObservationIA(txContext *models.TxContext, request models.ObservationIARequest) (int, []byte, error)
}

type IQualitatEndpoints interface {
	QualitatStart(Document string, InjuredCompleteName string) (*[]models.QualitatResponse, error)
}
