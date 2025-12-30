package services_test

import (
	"errors"
	"is-public-api/application/colletions"
	"is-public-api/application/models"
	"is-public-api/application/services"
	"testing"
)

// Mock del repositorio ISinisterPaymentRepository
type MockSinisterPaymentRepository struct {
	FindByDocumentNumberFunc func(txContext *models.TxContext, documentNumber string) (*colletions.SinisterPayment, error)
}

func (m *MockSinisterPaymentRepository) FindByDocumentNumber(txContext *models.TxContext, documentNumber string) (*colletions.SinisterPayment, error) {
	if m.FindByDocumentNumberFunc != nil {
		return m.FindByDocumentNumberFunc(txContext, documentNumber)
	}
	return nil, errors.New("not implemented")
}

// Test para NewSinisterPaymentFinder
func TestNewSinisterPaymentFinder(t *testing.T) {
	mockRepo := &MockSinisterPaymentRepository{}
	service := services.NewSinisterPaymentFinder(mockRepo)

	if service == nil {
		t.Error("Expected service to be created, got nil")
	}
}

// Test para FindByDocumentNumber - caso exitoso
func TestFindByDocumentNumber_Success(t *testing.T) {
	expectedPayment := &colletions.SinisterPayment{
		NumeroDocumento: "12345678",
		Beneficiario:    "Juan Perez",
		Producto:        "Seguro Vehicular",
		Moneda:          "PEN",
		MontoPagado:     "5000.00",
		Certificado:     "CERT-001",
		NumSiniestro:    "SIN-001",
		FechaSiniestro:  "2024-01-15",
		FechaPago:       "2024-02-01",
		Cobertura:       "Daños Propios",
	}

	mockRepo := &MockSinisterPaymentRepository{
		FindByDocumentNumberFunc: func(txContext *models.TxContext, documentNumber string) (*colletions.SinisterPayment, error) {
			if documentNumber != "12345678" {
				t.Errorf("Expected document number '12345678', got '%s'", documentNumber)
			}
			return expectedPayment, nil
		},
	}

	service := services.NewSinisterPaymentFinder(mockRepo)
	txContext := &models.TxContext{
		TransactionID: "test-tx-001",
	}

	result, err := service.FindByDocumentNumber(txContext, "12345678")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	// Verificar todos los campos
	if result.NumeroDocumento != expectedPayment.NumeroDocumento {
		t.Errorf("Expected NumeroDocumento '%s', got '%s'", expectedPayment.NumeroDocumento, result.NumeroDocumento)
	}
	if result.Beneficiario != expectedPayment.Beneficiario {
		t.Errorf("Expected Beneficiario '%s', got '%s'", expectedPayment.Beneficiario, result.Beneficiario)
	}
	if result.Producto != expectedPayment.Producto {
		t.Errorf("Expected Producto '%s', got '%s'", expectedPayment.Producto, result.Producto)
	}
	if result.Moneda != expectedPayment.Moneda {
		t.Errorf("Expected Moneda '%s', got '%s'", expectedPayment.Moneda, result.Moneda)
	}
	if result.MontoPagado != expectedPayment.MontoPagado {
		t.Errorf("Expected MontoPagado '%s', got '%s'", expectedPayment.MontoPagado, result.MontoPagado)
	}
}

// Test para FindByDocumentNumber - caso de error del repositorio
func TestFindByDocumentNumber_RepositoryError(t *testing.T) {
	expectedError := errors.New("database connection error")

	mockRepo := &MockSinisterPaymentRepository{
		FindByDocumentNumberFunc: func(txContext *models.TxContext, documentNumber string) (*colletions.SinisterPayment, error) {
			return nil, expectedError
		},
	}

	service := services.NewSinisterPaymentFinder(mockRepo)
	txContext := &models.TxContext{
		TransactionID: "test-tx-002",
	}

	result, err := service.FindByDocumentNumber(txContext, "12345678")

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

// Test para FindByDocumentNumber - documento no encontrado
func TestFindByDocumentNumber_NotFound(t *testing.T) {
	notFoundError := errors.New("document not found")

	mockRepo := &MockSinisterPaymentRepository{
		FindByDocumentNumberFunc: func(txContext *models.TxContext, documentNumber string) (*colletions.SinisterPayment, error) {
			return nil, notFoundError
		},
	}

	service := services.NewSinisterPaymentFinder(mockRepo)
	txContext := &models.TxContext{
		TransactionID: "test-tx-003",
	}

	result, err := service.FindByDocumentNumber(txContext, "99999999")

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if result != nil {
		t.Error("Expected nil result for not found document")
	}
}

// Test para FindByDocumentNumber - validación de mapeo de campos
func TestFindByDocumentNumber_FieldMapping(t *testing.T) {
	expectedPayment := &colletions.SinisterPayment{
		NumeroDocumento: "11111111",
		Beneficiario:    "Maria Lopez",
		Producto:        "SOAT",
		Moneda:          "USD",
		MontoPagado:     "1500.00",
		Certificado:     "CERT-999",
		NumSiniestro:    "SIN-999",
		FechaSiniestro:  "2024-06-10",
		FechaPago:       "2024-06-25",
		Cobertura:       "Muerte Accidental",
	}

	mockRepo := &MockSinisterPaymentRepository{
		FindByDocumentNumberFunc: func(txContext *models.TxContext, documentNumber string) (*colletions.SinisterPayment, error) {
			return expectedPayment, nil
		},
	}

	service := services.NewSinisterPaymentFinder(mockRepo)
	txContext := &models.TxContext{}

	result, err := service.FindByDocumentNumber(txContext, "11111111")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verificar que todos los campos se mapearon correctamente
	fields := []struct {
		name     string
		expected string
		actual   string
	}{
		{"NumeroDocumento", expectedPayment.NumeroDocumento, result.NumeroDocumento},
		{"Beneficiario", expectedPayment.Beneficiario, result.Beneficiario},
		{"Producto", expectedPayment.Producto, result.Producto},
		{"Moneda", expectedPayment.Moneda, result.Moneda},
		{"MontoPagado", expectedPayment.MontoPagado, result.MontoPagado},
		{"Certificado", expectedPayment.Certificado, result.Certificado},
		{"NumSiniestro", expectedPayment.NumSiniestro, result.NumSiniestro},
		{"FechaSiniestro", expectedPayment.FechaSiniestro, result.FechaSiniestro},
		{"FechaPago", expectedPayment.FechaPago, result.FechaPago},
		{"Cobertura", expectedPayment.Cobertura, result.Cobertura},
	}

	for _, field := range fields {
		if field.expected != field.actual {
			t.Errorf("Field %s: expected '%s', got '%s'", field.name, field.expected, field.actual)
		}
	}
}

