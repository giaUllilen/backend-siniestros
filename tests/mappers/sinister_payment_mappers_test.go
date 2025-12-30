package mappers_test

import (
	"is-public-api/application/mappers"
	"is-public-api/application/models"
	"is-public-api/application/resources"
	"testing"
)

// Test para ModelToSinisterPaymentResponse - caso básico
func TestModelToSinisterPaymentResponse_Success(t *testing.T) {
	model := &models.SinisterPayment{
		NumeroDocumento: "12345678",
		Beneficiario:    "Juan Pérez García",
		Producto:        "Seguro Vehicular",
		Moneda:          "PEN",
		MontoPagado:     "5000.00",
		Certificado:     "CERT-001-2024",
		NumSiniestro:    "SIN-2024-001",
		FechaSiniestro:  "2024-01-15",
		FechaPago:       "2024-02-01",
		Cobertura:       "Daños Propios",
	}

	res := new(resources.MapResponse)
	mappers.ModelToSinisterPaymentResponse(model, res)

	// Verificar que el response no sea nil
	if res == nil {
		t.Fatal("Expected response to be created, got nil")
	}

	// Verificar cada campo
	if (*res)["numero_documento"] != model.NumeroDocumento {
		t.Errorf("Expected numero_documento '%s', got '%s'", model.NumeroDocumento, (*res)["numero_documento"])
	}
	if (*res)["beneficiario"] != model.Beneficiario {
		t.Errorf("Expected beneficiario '%s', got '%s'", model.Beneficiario, (*res)["beneficiario"])
	}
	if (*res)["producto"] != model.Producto {
		t.Errorf("Expected producto '%s', got '%s'", model.Producto, (*res)["producto"])
	}
	if (*res)["moneda"] != model.Moneda {
		t.Errorf("Expected moneda '%s', got '%s'", model.Moneda, (*res)["moneda"])
	}
	if (*res)["monto_pagado"] != model.MontoPagado {
		t.Errorf("Expected monto_pagado '%s', got '%s'", model.MontoPagado, (*res)["monto_pagado"])
	}
	if (*res)["certificado"] != model.Certificado {
		t.Errorf("Expected certificado '%s', got '%s'", model.Certificado, (*res)["certificado"])
	}
	if (*res)["num_siniestro"] != model.NumSiniestro {
		t.Errorf("Expected num_siniestro '%s', got '%s'", model.NumSiniestro, (*res)["num_siniestro"])
	}
	if (*res)["fecha_siniestro"] != model.FechaSiniestro {
		t.Errorf("Expected fecha_siniestro '%s', got '%s'", model.FechaSiniestro, (*res)["fecha_siniestro"])
	}
	if (*res)["fecha_pago"] != model.FechaPago {
		t.Errorf("Expected fecha_pago '%s', got '%s'", model.FechaPago, (*res)["fecha_pago"])
	}
	if (*res)["cobertura"] != model.Cobertura {
		t.Errorf("Expected cobertura '%s', got '%s'", model.Cobertura, (*res)["cobertura"])
	}
}

// Test para ModelToSinisterPaymentResponse - valores vacíos
func TestModelToSinisterPaymentResponse_EmptyValues(t *testing.T) {
	model := &models.SinisterPayment{
		NumeroDocumento: "",
		Beneficiario:    "",
		Producto:        "",
		Moneda:          "",
		MontoPagado:     "",
		Certificado:     "",
		NumSiniestro:    "",
		FechaSiniestro:  "",
		FechaPago:       "",
		Cobertura:       "",
	}

	res := new(resources.MapResponse)
	mappers.ModelToSinisterPaymentResponse(model, res)

	// Verificar que todos los campos están presentes aunque sean vacíos
	expectedKeys := []string{
		"numero_documento", "beneficiario", "producto", "moneda",
		"monto_pagado", "certificado", "num_siniestro",
		"fecha_siniestro", "fecha_pago", "cobertura",
	}

	for _, key := range expectedKeys {
		if _, exists := (*res)[key]; !exists {
			t.Errorf("Expected key '%s' to exist in response", key)
		}
		if (*res)[key] != "" {
			t.Errorf("Expected empty string for key '%s', got '%s'", key, (*res)[key])
		}
	}
}

// Test para ModelToSinisterPaymentResponse - verificar cantidad de campos
func TestModelToSinisterPaymentResponse_FieldCount(t *testing.T) {
	model := &models.SinisterPayment{
		NumeroDocumento: "12345678",
		Beneficiario:    "Test User",
		Producto:        "Test Product",
		Moneda:          "USD",
		MontoPagado:     "1000",
		Certificado:     "CERT-001",
		NumSiniestro:    "SIN-001",
		FechaSiniestro:  "2024-01-01",
		FechaPago:       "2024-01-10",
		Cobertura:       "Test Coverage",
	}

	res := new(resources.MapResponse)
	mappers.ModelToSinisterPaymentResponse(model, res)

	expectedFieldCount := 10
	actualFieldCount := len(*res)

	if actualFieldCount != expectedFieldCount {
		t.Errorf("Expected %d fields in response, got %d", expectedFieldCount, actualFieldCount)
	}
}

// Test para ModelToSinisterPaymentResponse - caracteres especiales
func TestModelToSinisterPaymentResponse_SpecialCharacters(t *testing.T) {
	model := &models.SinisterPayment{
		NumeroDocumento: "12345678",
		Beneficiario:    "José María O'Connor",
		Producto:        "Seguro & Protección",
		Moneda:          "S/.",
		MontoPagado:     "5,000.50",
		Certificado:     "CERT-001/2024",
		NumSiniestro:    "SIN-2024-001-A",
		FechaSiniestro:  "15/01/2024",
		FechaPago:       "01-02-2024",
		Cobertura:       "Daños \"Propios\"",
	}

	res := new(resources.MapResponse)
	mappers.ModelToSinisterPaymentResponse(model, res)

	// Verificar que los caracteres especiales se preservan
	if (*res)["beneficiario"] != model.Beneficiario {
		t.Errorf("Expected beneficiario '%s', got '%s'", model.Beneficiario, (*res)["beneficiario"])
	}
	if (*res)["producto"] != model.Producto {
		t.Errorf("Expected producto '%s', got '%s'", model.Producto, (*res)["producto"])
	}
	if (*res)["cobertura"] != model.Cobertura {
		t.Errorf("Expected cobertura '%s', got '%s'", model.Cobertura, (*res)["cobertura"])
	}
}

// Test para ModelToSinisterPaymentResponse - respuesta preexistente
func TestModelToSinisterPaymentResponse_OverwriteExisting(t *testing.T) {
	model := &models.SinisterPayment{
		NumeroDocumento: "87654321",
		Beneficiario:    "María López",
		Producto:        "SOAT",
		Moneda:          "USD",
		MontoPagado:     "2000.00",
		Certificado:     "CERT-999",
		NumSiniestro:    "SIN-999",
		FechaSiniestro:  "2024-06-15",
		FechaPago:       "2024-07-01",
		Cobertura:       "Muerte Accidental",
	}

	// Crear una respuesta con datos previos
	res := new(resources.MapResponse)
	*res = resources.MapResponse{
		"numero_documento": "00000000",
		"old_field":        "should be removed",
	}

	mappers.ModelToSinisterPaymentResponse(model, res)

	// Verificar que se sobrescribió el campo existente
	if (*res)["numero_documento"] != model.NumeroDocumento {
		t.Errorf("Expected numero_documento to be overwritten to '%s', got '%s'", model.NumeroDocumento, (*res)["numero_documento"])
	}

	// Verificar que el campo antiguo ya no existe
	if _, exists := (*res)["old_field"]; exists {
		t.Error("Expected old_field to be removed from response")
	}

	// Verificar que el nuevo campo está presente
	if _, exists := (*res)["beneficiario"]; !exists {
		t.Error("Expected beneficiario field to exist in response")
	}
}

// Test para ModelToSinisterPaymentResponse - montos con diferentes formatos
func TestModelToSinisterPaymentResponse_DifferentAmountFormats(t *testing.T) {
	testCases := []struct {
		name        string
		montoPagado string
	}{
		{"Decimal con dos dígitos", "1000.00"},
		{"Decimal con tres dígitos", "1000.000"},
		{"Sin decimales", "1000"},
		{"Con comas", "1,000.00"},
		{"Cero", "0"},
		{"Cero con decimales", "0.00"},
		{"Número negativo", "-500.00"},
		{"Número grande", "999999999.99"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := &models.SinisterPayment{
				NumeroDocumento: "12345678",
				MontoPagado:     tc.montoPagado,
			}

			res := new(resources.MapResponse)
			mappers.ModelToSinisterPaymentResponse(model, res)

			if (*res)["monto_pagado"] != tc.montoPagado {
				t.Errorf("Expected monto_pagado '%s', got '%s'", tc.montoPagado, (*res)["monto_pagado"])
			}
		})
	}
}

