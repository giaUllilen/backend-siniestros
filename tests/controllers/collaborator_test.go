package controllers

import (
	"errors"
	"is-public-api/application/models"
	"testing"
	_ "is-public-api/tests" // Importar setup de variables de entorno
)

// Mock del servicio ICollaboratorFinder
type MockCollaboratorFinder struct {
	FindFunc func(txContext *models.TxContext, code string) (*models.Collaborator, error)
}

func (m *MockCollaboratorFinder) Find(txContext *models.TxContext, code string) (*models.Collaborator, error) {
	if m.FindFunc != nil {
		return m.FindFunc(txContext, code)
	}
	return nil, errors.New("not implemented")
}

// Test para NewCollaboratorHandler
func TestNewCollaboratorHandler(t *testing.T) {
	mockFinder := &MockCollaboratorFinder{}
	handler := NewCollaboratorHandler(mockFinder)
	
	if handler == nil {
		t.Error("Expected handler to be created, got nil")
	}
}

// Test para verificar que el handler puede ser creado con un finder válido
func TestCollaboratorHandler_WithValidFinder(t *testing.T) {
	collaborator := &models.Collaborator{
		IdPersona:          "P001",
		Nombres:            "Juan",
		ApellidoPaterno:    "Pérez",
		DocumentoIdentidad: "12345678",
		Estado:             "Activo",
	}
	
	mockFinder := &MockCollaboratorFinder{
		FindFunc: func(txContext *models.TxContext, code string) (*models.Collaborator, error) {
			return collaborator, nil
		},
	}
	
	handler := NewCollaboratorHandler(mockFinder)
	
	if handler == nil {
		t.Fatal("Expected handler to be created")
	}
}

// Test para verificar que el handler maneja errores del finder
func TestCollaboratorHandler_FinderError(t *testing.T) {
	expectedError := errors.New("collaborator not found")
	
	mockFinder := &MockCollaboratorFinder{
		FindFunc: func(txContext *models.TxContext, code string) (*models.Collaborator, error) {
			return nil, expectedError
		},
	}
	
	handler := NewCollaboratorHandler(mockFinder)
	if handler == nil {
		t.Error("Expected handler to be created")
	}
}

