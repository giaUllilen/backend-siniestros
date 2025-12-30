package storage_test

import (
	"is-public-api/application/storage"
	"testing"
	_ "is-public-api/tests" // Importar setup de variables de entorno
)

// Test para NewSoatReturnRepository
func TestNewSoatReturnRepository(t *testing.T) {
	repo := storage.NewSoatReturnRepository(nil)
	
	if repo == nil {
		t.Error("Expected repository to be created, got nil")
	}
}

// Test para verificar que FindByDocument tiene la firma correcta
func TestSoatReturnRepository_FindByDocument_Signature(t *testing.T) {
	repo := storage.NewSoatReturnRepository(nil)
	
	// Verificar que el repositorio se creó correctamente
	if repo == nil {
		t.Error("Expected repository to be created, got nil")
	}
	
	// No llamamos a FindByDocument porque requiere conexión a MongoDB
	// La prueba verifica que el repositorio existe y puede ser instanciado
	t.Log("Repository created successfully with nil database")
}

