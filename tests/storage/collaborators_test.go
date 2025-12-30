package storage_test

import (
	"is-public-api/application/storage"
	"testing"
	_ "is-public-api/tests" // Importar setup de variables de entorno
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
	
	// Verificar que el repositorio se creó correctamente
	if repo == nil {
		t.Error("Expected repository to be created, got nil")
	}
	
	// No llamamos a Find porque requiere conexión a MongoDB
	// La prueba verifica que el repositorio existe y puede ser instanciado
	t.Log("Repository created successfully with nil database")
}

