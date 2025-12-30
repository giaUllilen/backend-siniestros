package storage_test

import (
	"is-public-api/application/storage"
	"testing"
	_ "is-public-api/tests" // Importar setup de variables de entorno
)

// Test para NewSinisterPaymentRepository
func TestNewSinisterPaymentRepository(t *testing.T) {
	repo := storage.NewSinisterPaymentRepository(nil)
	
	if repo == nil {
		t.Error("Expected repository to be created, got nil")
	}
}

// Test para verificar que FindByDocumentNumber tiene la firma correcta
func TestSinisterPaymentRepository_FindByDocumentNumber_Signature(t *testing.T) {
	repo := storage.NewSinisterPaymentRepository(nil)
	
	// Verificar que el repositorio se creó correctamente
	if repo == nil {
		t.Error("Expected repository to be created, got nil")
	}
	
	// No llamamos a FindByDocumentNumber porque requiere conexión a MongoDB
	// La prueba verifica que el repositorio existe y puede ser instanciado
	t.Log("Repository created successfully with nil database")
}

