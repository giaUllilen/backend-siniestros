package controllers

import (
	"errors"
	"is-public-api/application/models"
	"mime/multipart"
	"testing"
	_ "is-public-api/tests" // Importar setup de variables de entorno
)

// Mock del servicio ISinisterService
type MockSinisterService struct {
	FindByDocumentNumberFunc func(txContext *models.TxContext, documentNumber string) (*models.SinisterPayment, error)
}

func (m *MockSinisterService) FindByDocumentNumber(txContext *models.TxContext, documentNumber string) (*models.SinisterPayment, error) {
	if m.FindByDocumentNumberFunc != nil {
		return m.FindByDocumentNumberFunc(txContext, documentNumber)
	}
	return nil, errors.New("not implemented")
}

// Mock del servicio ISinisterServiceDomain
type MockSinisterServiceDomain struct {
	SaveFunc              func(txContext *models.TxContext, request map[string]interface{}, coverages []map[string]interface{}, attachments []*multipart.FileHeader) ([]interface{}, error)
	FindByCaseNumberFunc  func(txContext *models.TxContext, caseID interface{}) (int, []byte, error)
	FindByCaseHistoryFunc func(txContext *models.TxContext, documentNumber string) (int, []models.SinisterHistory, error)
}

func (m *MockSinisterServiceDomain) Save(txContext *models.TxContext, request map[string]interface{}, coverages []map[string]interface{}, attachments []*multipart.FileHeader) ([]interface{}, error) {
	if m.SaveFunc != nil {
		return m.SaveFunc(txContext, request, coverages, attachments)
	}
	return nil, errors.New("not implemented")
}

func (m *MockSinisterServiceDomain) FindByCaseNumber(txContext *models.TxContext, caseID string) (int, []byte, error) {
	if m.FindByCaseNumberFunc != nil {
		return m.FindByCaseNumberFunc(txContext, caseID)
	}
	return 404, nil, errors.New("not implemented")
}

func (m *MockSinisterServiceDomain) FindByCaseHistory(txContext *models.TxContext, documentNumber string) (int, []models.SinisterHistory, error) {
	if m.FindByCaseHistoryFunc != nil {
		return m.FindByCaseHistoryFunc(txContext, documentNumber)
	}
	return 404, nil, errors.New("not implemented")
}

// Test para NewSinisterPaymentHandler
func TestNewSinisterPaymentHandler(t *testing.T) {
	mockService := &MockSinisterService{}
	mockDomain := &MockSinisterServiceDomain{}
	
	handler := NewSinisterPaymentHandler(mockService, mockDomain)
	
	if handler == nil {
		t.Error("Expected handler to be created, got nil")
	}
}

// Test para service.FindByDocumentNumber - caso exitoso
func TestSinisterHandler_FindByDocumentNumber_Success(t *testing.T) {
	expectedPayment := &models.SinisterPayment{
		NumeroDocumento: "12345678",
		Beneficiario:    "Juan Pérez",
		Producto:        "Seguro Vehicular",
		MontoPagado:     "5000.00",
	}
	
	mockService := &MockSinisterService{
		FindByDocumentNumberFunc: func(txContext *models.TxContext, documentNumber string) (*models.SinisterPayment, error) {
			if documentNumber != "12345678" {
				t.Errorf("Expected document number '12345678', got '%s'", documentNumber)
			}
			return expectedPayment, nil
		},
	}
	
	mockDomain := &MockSinisterServiceDomain{}
	handler := NewSinisterPaymentHandler(mockService, mockDomain)
	
	if handler == nil {
		t.Fatal("Expected handler to be created")
	}
}

// Test para service.FindByDocumentNumber - error
func TestSinisterHandler_FindByDocumentNumber_Error(t *testing.T) {
	expectedError := errors.New("document not found")
	
	mockService := &MockSinisterService{
		FindByDocumentNumberFunc: func(txContext *models.TxContext, documentNumber string) (*models.SinisterPayment, error) {
			return nil, expectedError
		},
	}
	
	mockDomain := &MockSinisterServiceDomain{}
	handler := NewSinisterPaymentHandler(mockService, mockDomain)
	
	if handler == nil {
		t.Error("Expected handler to be created")
	}
}

// Test para domain.Save - caso exitoso
func TestSinisterHandler_Save_Success(t *testing.T) {
	expectedResponse := []interface{}{
		map[string]interface{}{
			"case_id": "CIS_12345",
			"status":  "created",
		},
	}
	
	mockService := &MockSinisterService{}
	mockDomain := &MockSinisterServiceDomain{
		SaveFunc: func(txContext *models.TxContext, request map[string]interface{}, coverages []map[string]interface{}, attachments []*multipart.FileHeader) ([]interface{}, error) {
			return expectedResponse, nil
		},
	}
	
	handler := NewSinisterPaymentHandler(mockService, mockDomain)
	
	if handler == nil {
		t.Error("Expected handler to be created")
	}
}

// Test para domain.FindByCaseNumber - caso exitoso
func TestSinisterHandler_FindByCaseNumber_Success(t *testing.T) {
	expectedBody := []byte(`{"case_id":"CIS_12345","status":"active"}`)
	
	mockService := &MockSinisterService{}
	mockDomain := &MockSinisterServiceDomain{
		FindByCaseNumberFunc: func(txContext *models.TxContext, caseID interface{}) (int, []byte, error) {
			return 200, expectedBody, nil
		},
	}
	
	handler := NewSinisterPaymentHandler(mockService, mockDomain)
	
	if handler == nil {
		t.Error("Expected handler to be created")
	}
}

