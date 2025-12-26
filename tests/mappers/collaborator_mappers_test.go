package mappers_test

import (
	"is-public-api/application/mappers"
	"is-public-api/application/models"
	"is-public-api/application/resources"
	"testing"
	"time"
)

// Test para ModelToCollaboratorResponse - colaborador activo con DNI
func TestModelToCollaboratorResponse_ActiveWithDNI(t *testing.T) {
	fechaIngreso := time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC)
	
	model := &models.Collaborator{
		IdPersona:             "P001",
		Nombres:               "Juan Carlos",
		ApellidoPaterno:       "Pérez",
		ApellidoMaterno:       "González",
		Apellidos:             "Pérez González",
		Puesto:                "Desarrollador Senior",
		DescripcionPuesto:     "Desarrollo de Software",
		CodigoArea:            "TI-001",
		DescripcionArea:       "Tecnología de la Información",
		TipoDocumento:         "DNI",
		CodigoTipoDocumento:   "01",
		DocumentoIdentidad:    "12345678",
		CodigoIS:              "IS-001",
		FechaIngreso:          fechaIngreso,
		FechaCese:             time.Time{},
		Estado:                "Activo",
		CodigoVicepresidencia: "VP-001",
		Vicepresidencia:       "Operaciones",
	}

	res := new(resources.MapResponse)
	mappers.ModelToCollaboratorResponse(model, res)

	// Verificar estado
	if (*res)["estado"] != "Activo" {
		t.Errorf("Expected estado 'Activo', got '%v'", (*res)["estado"])
	}

	// Verificar todos los campos para colaborador activo
	if (*res)["id_persona"] != model.IdPersona {
		t.Errorf("Expected id_persona '%s', got '%v'", model.IdPersona, (*res)["id_persona"])
	}
	if (*res)["nombres"] != model.Nombres {
		t.Errorf("Expected nombres '%s', got '%v'", model.Nombres, (*res)["nombres"])
	}
	if (*res)["apellido_paterno"] != model.ApellidoPaterno {
		t.Errorf("Expected apellido_paterno '%s', got '%v'", model.ApellidoPaterno, (*res)["apellido_paterno"])
	}
	if (*res)["tipo_documento"] != "DNI" {
		t.Errorf("Expected tipo_documento 'DNI', got '%v'", (*res)["tipo_documento"])
	}
	if (*res)["numero_documento"] != model.DocumentoIdentidad {
		t.Errorf("Expected numero_documento '%s', got '%v'", model.DocumentoIdentidad, (*res)["numero_documento"])
	}
}

// Test para ModelToCollaboratorResponse - colaborador activo con CE
func TestModelToCollaboratorResponse_ActiveWithCE(t *testing.T) {
	fechaIngreso := time.Date(2019, 5, 10, 0, 0, 0, 0, time.UTC)
	
	model := &models.Collaborator{
		IdPersona:           "P002",
		Nombres:             "John",
		ApellidoPaterno:     "Smith",
		Estado:              "Activo",
		CodigoTipoDocumento: "02",
		DocumentoIdentidad:  "000123456",
		FechaIngreso:        fechaIngreso,
	}

	res := new(resources.MapResponse)
	mappers.ModelToCollaboratorResponse(model, res)

	// Verificar que el tipo de documento sea CE
	if (*res)["tipo_documento"] != "CE" {
		t.Errorf("Expected tipo_documento 'CE', got '%v'", (*res)["tipo_documento"])
	}
}

// Test para ModelToCollaboratorResponse - colaborador inactivo
func TestModelToCollaboratorResponse_Inactive(t *testing.T) {
	model := &models.Collaborator{
		IdPersona:      "P003",
		Nombres:        "María",
		Estado:         "Cesado",
		FechaIngreso:   time.Now(),
	}

	res := new(resources.MapResponse)
	mappers.ModelToCollaboratorResponse(model, res)

	// Verificar que solo el estado está presente
	if (*res)["estado"] != "Cesado" {
		t.Errorf("Expected estado 'Cesado', got '%v'", (*res)["estado"])
	}

	// Verificar que otros campos NO están presentes para colaborador inactivo
	if _, exists := (*res)["id_persona"]; exists {
		t.Error("Expected id_persona to not exist for inactive collaborator")
	}
}

// Test para ModelToCollaboratorResponse - estado en minúsculas
func TestModelToCollaboratorResponse_LowercaseActive(t *testing.T) {
	model := &models.Collaborator{
		IdPersona:    "P004",
		Nombres:      "Pedro",
		Estado:       "activo",
		FechaIngreso: time.Now(),
	}

	res := new(resources.MapResponse)
	mappers.ModelToCollaboratorResponse(model, res)

	// Verificar que los campos están presentes
	if _, exists := (*res)["id_persona"]; !exists {
		t.Error("Expected id_persona to exist for lowercase 'activo' collaborator")
	}
}

