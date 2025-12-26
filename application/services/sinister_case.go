package services

import (
	"is-public-api/application/mappers"
	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/application/storage"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type sinisterCaseService struct {
	sinisterCaseRepository storage.ISinisterCaseRepository
}

func NewSinisterCaseService(repo storage.ISinisterCaseRepository) *sinisterCaseService {
	return &sinisterCaseService{sinisterCaseRepository: repo}
}

func (service *sinisterCaseService) Save(txContext *models.TxContext, sinister map[string]interface{}) (string, error) {
	res, err := service.sinisterCaseRepository.Save(txContext, sinister)
	return res, err
}

func (service *sinisterCaseService) UpdateOne(txContext *models.TxContext, caseNumber primitive.ObjectID, sinister map[string]interface{}) (primitive.ObjectID, error) {
	res, err := service.sinisterCaseRepository.UpdateOne(txContext, caseNumber, sinister)
	return res, err
}

func (service *sinisterCaseService) FindByCaseNumber(txContext *models.TxContext, caseNumber primitive.ObjectID) (*models.CasoSiniestro, error) {
	c, err := service.sinisterCaseRepository.FindByCase(txContext, caseNumber)
	return c, err
}

func (service *sinisterCaseService) FindAll(txContext *models.TxContext) (resources.ListArray, error) {
	cases, err := service.sinisterCaseRepository.FindAll(txContext)

	result := make(resources.ListArray, len(cases))
	for i, c := range cases {
		result[i] = mappers.SinisterToMapResponse(c)
	}
	return result, err
}
