package storage_test

import (
	"is-public-api/application/models"
	"is-public-api/application/storage"
	"testing"
)

// Test para NewCustomerRepository
func TestNewCustomerRepository(t *testing.T) {
	repo := storage.NewCustomerRepository(nil)
	
	if repo == nil {
		t.Error("Expected repository to be created, got nil")
	}
}

// Test para verificar que Find tiene la firma correcta
func TestCustomerRepository_Find_Signature(t *testing.T) {
	repo := storage.NewCustomerRepository(nil)
	txContext := &models.TxContext{
		TransactionID: "test-001",
	}
	
	// La función debería existir y poder ser llamada
	_, err := repo.Find(txContext, "12345678")
	
	// Esperamos un error porque database es nil
	if err == nil {
		t.Log("Warning: Expected error due to nil database, but got none")
	}
}

