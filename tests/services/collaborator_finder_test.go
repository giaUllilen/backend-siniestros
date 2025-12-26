package services_test

import (
	"encoding/base64"
	"errors"
	"is-public-api/application/colletions"
	"is-public-api/application/models"
	"is-public-api/application/services"
	"testing"
)

// Mock del repositorio ICollaboratorRepository
type MockCollaboratorRepository struct {
	FindFunc func(txContext *models.TxContext, code string) (*colletions.Colaborador, error)
}

func (m *MockCollaboratorRepository) Find(txContext *models.TxContext, code string) (*colletions.Colaborador, error) {
	if m.FindFunc != nil {
		return m.FindFunc(txContext, code)
	}
	return nil, errors.New("not implemented")
}

// Test para NewCollaboratorFinder
func TestNewCollaboratorFinder(t *testing.T) {
	mockRepo := &MockCollaboratorRepository{}
	service := services.NewCollaboratorFinder(mockRepo)

	if service == nil {
		t.Error("Expected service to be created, got nil")
	}
}

// Test para Find - caso exitoso
func TestCollaboratorFind_Success(t *testing.T) {
	expectedCollaborator := &colletions.Colaborador{
		IdPersona:            "P001",
		Nombres:              "Juan Carlos",
		ApellidoPaterno:      "Pérez",
		ApellidoMaterno:      "González",
		Apellidos:            "Pérez González",
		Puesto:               "Desarrollador Senior",
		DescripcionPuesto:    "Desarrollo de Software",
		CodigoArea:           "TI-001",
		DescripcionArea:      "Tecnología de la Información",
		TipoDocumento:        "DNI",
		CodigoTipoDocumento:  "1",
		DocumentoIdentidad:   "12345678",
		CodigoIS:             "IS-001",
		FechaIngreso:         "2020-01-15",
		FechaCese:            "",
		Estado:               "Activo",
		CodigoVicepresidencia: "VP-001",
		Vicepresidencia:      "Operaciones",
	}

	// Código sin codificar
	decodedCode := "EMP12345"
	// Codificar en base64
	encodedCode := base64.StdEncoding.EncodeToString([]byte(decodedCode))

	mockRepo := &MockCollaboratorRepository{
		FindFunc: func(txContext *models.TxContext, code string) (*colletions.Colaborador, error) {
			if code != decodedCode {
				t.Errorf("Expected decoded code '%s', got '%s'", decodedCode, code)
			}
			return expectedCollaborator, nil
		},
	}

	service := services.NewCollaboratorFinder(mockRepo)
	txContext := &models.TxContext{
		TransactionID: "test-tx-001",
	}

	result, err := service.Find(txContext, encodedCode)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	// Verificar campos
	if result.IdPersona != expectedCollaborator.IdPersona {
		t.Errorf("Expected IdPersona '%s', got '%s'", expectedCollaborator.IdPersona, result.IdPersona)
	}
	if result.Nombres != expectedCollaborator.Nombres {
		t.Errorf("Expected Nombres '%s', got '%s'", expectedCollaborator.Nombres, result.Nombres)
	}
	if result.DocumentoIdentidad != expectedCollaborator.DocumentoIdentidad {
		t.Errorf("Expected DocumentoIdentidad '%s', got '%s'", expectedCollaborator.DocumentoIdentidad, result.DocumentoIdentidad)
	}
	if result.Estado != expectedCollaborator.Estado {
		t.Errorf("Expected Estado '%s', got '%s'", expectedCollaborator.Estado, result.Estado)
	}
}

// Test para Find - código base64 inválido
func TestCollaboratorFind_InvalidBase64(t *testing.T) {
	mockRepo := &MockCollaboratorRepository{}
	service := services.NewCollaboratorFinder(mockRepo)
	txContext := &models.TxContext{
		TransactionID: "test-tx-002",
	}

	// Código base64 inválido
	invalidCode := "!!!invalid-base64!!!"

	result, err := service.Find(txContext, invalidCode)

	if err == nil {
		t.Error("Expected error for invalid base64, got nil")
	}

	if result != nil {
		t.Errorf("Expected nil result for invalid base64, got %v", result)
	}
}

// Test para Find - error del repositorio
func TestCollaboratorFind_RepositoryError(t *testing.T) {
	expectedError := errors.New("database connection error")
	encodedCode := base64.StdEncoding.EncodeToString([]byte("EMP12345"))

	mockRepo := &MockCollaboratorRepository{
		FindFunc: func(txContext *models.TxContext, code string) (*colletions.Colaborador, error) {
			return nil, expectedError
		},
	}

	service := services.NewCollaboratorFinder(mockRepo)
	txContext := &models.TxContext{
		TransactionID: "test-tx-003",
	}

	result, err := service.Find(txContext, encodedCode)

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if err != expectedError {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

// Test para Find - colaborador no encontrado
func TestCollaboratorFind_NotFound(t *testing.T) {
	notFoundError := errors.New("collaborator not found")
	encodedCode := base64.StdEncoding.EncodeToString([]byte("INVALID"))

	mockRepo := &MockCollaboratorRepository{
		FindFunc: func(txContext *models.TxContext, code string) (*colletions.Colaborador, error) {
			return nil, notFoundError
		},
	}

	service := services.NewCollaboratorFinder(mockRepo)
	txContext := &models.TxContext{}

	result, err := service.Find(txContext, encodedCode)

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if result != nil {
		t.Error("Expected nil result for not found collaborator")
	}
}

// Test para Find - código vacío base64 válido
func TestCollaboratorFind_EmptyCode(t *testing.T) {
	emptyCode := base64.StdEncoding.EncodeToString([]byte(""))

	mockRepo := &MockCollaboratorRepository{
		FindFunc: func(txContext *models.TxContext, code string) (*colletions.Colaborador, error) {
			if code != "" {
				t.Errorf("Expected empty code, got '%s'", code)
			}
			return nil, errors.New("code cannot be empty")
		},
	}

	service := services.NewCollaboratorFinder(mockRepo)
	txContext := &models.TxContext{}

	result, err := service.Find(txContext, emptyCode)

	if err == nil {
		t.Error("Expected error for empty code, got nil")
	}

	if result != nil {
		t.Error("Expected nil result for empty code")
	}
}

