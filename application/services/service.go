package services

import (
	"bytes"
	"mime/multipart"

	"is-public-api/application/models"
	"is-public-api/application/resources"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICollaboratorFinder interface {
	Find(txContext *models.TxContext, code string) (*models.Collaborator, error)
}

type ISoatReturnFinder interface {
	FindByDocument(txContext *models.TxContext, documentNumber string) (*models.SoatReturn, error)
}

type ISoatFinder interface {
	FindByPlate(txContext *models.TxContext, plate string) (*models.Soat, error)
	FindByPlateHistory(txContext *models.TxContext, plate string, dateOccurrence string) (*models.Soat, error)
}

type ISinisterService interface {
	FindByDocumentNumber(txContext *models.TxContext, documentNumber string) (*models.SinisterPayment, error)
}

type ISinisterCaseService interface {
	Save(txContext *models.TxContext, sinister map[string]interface{}) (string, error)
	UpdateOne(txContext *models.TxContext, caseNumber primitive.ObjectID, sinister map[string]interface{}) (primitive.ObjectID, error)
	FindByCaseNumber(txContext *models.TxContext, caseNumber primitive.ObjectID) (*models.CasoSiniestro, error)
	FindAll(txContext *models.TxContext) (resources.ListArray, error)
}

type ISinisterServiceDomain interface {
	Save(txContext *models.TxContext, request map[string]interface{}, coverages []map[string]interface{}, attachments []*multipart.FileHeader) ([]interface{}, error)
	// FindByCaseNumber - Cambiado de primitive.ObjectID a string para soportar números de caso como "CIS_98746"
	// Anterior: FindByCaseNumber(txContext *models.TxContext, caseNumber primitive.ObjectID) (int, []byte, error)
	FindByCaseNumber(txContext *models.TxContext, caseNumber string) (int, []byte, error)
	FindByCaseHistory(txContext *models.TxContext, documentNumber string) (int, []models.SinisterHistory, error)
}

type ISinisterAIDomain interface {
	Save(txContext *models.TxContext, request *models.Sinister) ([]interface{}, error)
	SaveSiniesterPolice(numeroDocumento string, soatIS *models.Soat, lesionado models.Lesionado, documento resources.Document, dataResponseGenIA models.DataDenunciaPolicialResponseGenIA, txContext *models.TxContext, qualitat models.DataQualitat, dataObservation *models.ObservationWrapper) (string, error)
	AnalyzeDocument(txContext *models.TxContext, request *resources.SinisterCase) (resources.MapResponse, error)
	AnalyzeDocumentWeb(txContext *models.TxContext, request *resources.SinisterCase, filebase64Data string) (resources.MapResponse, error)
	AnalyzeWebIA(txContext *models.TxContext, request *resources.SinisterCase, callback func(resources.MapResponse, error))
	AnalyzeWebIASimple(txContext *models.TxContext, request *resources.SinisterCase)
	UploadObservedDocument(txContext *models.TxContext, request *resources.SinisterCase) error
	GetByCase(txContext *models.TxContext, caseID string) (*models.CasoSiniestro, error)
	GetAll(txContext *models.TxContext) (resources.ListArray, error)
	GetPreviousCases(txContext *models.TxContext, doc string) (resources.MapResponse, error)
}

type IServiceTemplateMaker interface {
	Make(txContext *models.TxContext, data map[string]interface{}) (*bytes.Buffer, error)
}

type IStorageServiceDomain interface {
	Upload(txContext *models.TxContext, caseNumber primitive.ObjectID) (int, []byte, error)
}

type ISubscriptionCenter interface {
	CreateSubscriptionCenterOptions(txContext *models.TxContext, key, option, value, version string, description *string) (string, error)
	FindSubscriptionCenterOptionsByID(txContext *models.TxContext, key string) (string, error)
}

type ICloudFunctionSubscriptionCenter interface {
	CreateSubscriptionCenterOptions(txContext *models.TxContext, data resources.MapRequest) error
}
