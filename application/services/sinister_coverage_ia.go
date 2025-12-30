package services

import (
	"is-public-api/application/models"
	"is-public-api/application/storage"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

// ISinisterCoverageIAService define la interfaz del servicio
type ISinisterCoverageIAService interface {
	IsValidCombination(txContext *models.TxContext, producto, cobertura string) (bool, error)
	GetAllActiveCombinations(txContext *models.TxContext) ([]models.SinisterCoverageIA, error)
}

type sinisterCoverageIAService struct {
	repository storage.ISinisterCoverageIARepository
}

// NewSinisterCoverageIAService crea una nueva instancia del servicio
func NewSinisterCoverageIAService(repo storage.ISinisterCoverageIARepository) ISinisterCoverageIAService {
	return &sinisterCoverageIAService{repository: repo}
}

// IsValidCombination verifica si la combinación producto-cobertura está configurada y activa
func (service *sinisterCoverageIAService) IsValidCombination(txContext *models.TxContext, producto, cobertura string) (bool, error) {
	// Normalizar strings para comparación case-insensitive
	producto = strings.TrimSpace(producto)
	cobertura = strings.TrimSpace(cobertura)

	if producto == "" || cobertura == "" {
		return false, nil
	}

	// Buscar la combinación en la base de datos
	_, err := service.repository.FindActiveByProductAndCoverage(txContext, producto, cobertura)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró la combinación, es válido pero no está configurada
			return false, nil
		}
		// Error de base de datos
		return false, err
	}

	// Se encontró la combinación y está activa
	return true, nil
}

// GetAllActiveCombinations obtiene todas las combinaciones activas
func (service *sinisterCoverageIAService) GetAllActiveCombinations(txContext *models.TxContext) ([]models.SinisterCoverageIA, error) {
	return service.repository.FindAllActive(txContext)
}
