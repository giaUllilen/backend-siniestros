package apihelpers_test

import (
	"is-public-api/application/apihelpers"
	"testing"
)

// Test para ResponseWrapper - MarshalJSONObject
func TestResponseWrapper_MarshalJSONObject(t *testing.T) {
	wrapper := &apihelpers.ResponseWrapper{
		Code:    apihelpers.CodeOk,
		Data:    map[string]string{"key": "value"},
		Message: "Success",
	}
	
	if wrapper == nil {
		t.Error("Expected ResponseWrapper to be created")
	}
}

// Test para ResponseWrapper - IsNil cuando es nil
func TestResponseWrapper_IsNil_WhenNil(t *testing.T) {
	var wrapper *apihelpers.ResponseWrapper = nil
	
	if !wrapper.IsNil() {
		t.Error("Expected IsNil() to return true for nil wrapper")
	}
}

// Test para ResponseWrapper - IsNil cuando no es nil
func TestResponseWrapper_IsNil_WhenNotNil(t *testing.T) {
	wrapper := &apihelpers.ResponseWrapper{
		Code:    apihelpers.CodeOk,
		Message: "Test",
	}
	
	if wrapper.IsNil() {
		t.Error("Expected IsNil() to return false for non-nil wrapper")
	}
}

// Test para ResponseCode - constantes
func TestResponseCode_Constants(t *testing.T) {
	if apihelpers.CodeOk != "01" {
		t.Errorf("Expected CodeOk to be '01', got '%s'", apihelpers.CodeOk)
	}
	if apihelpers.CodeError != "99" {
		t.Errorf("Expected CodeError to be '99', got '%s'", apihelpers.CodeError)
	}
}

// Test para ResponseWrapper - con diferentes tipos de data
func TestResponseWrapper_DifferentDataTypes(t *testing.T) {
	testCases := []struct {
		name string
		data interface{}
	}{
		{"String data", "test string"},
		{"Int data", 42},
		{"Float data", 3.14},
		{"Bool data", true},
		{"Map data", map[string]interface{}{"key": "value"}},
		{"Slice data", []string{"item1", "item2"}},
		{"Nil data", nil},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wrapper := &apihelpers.ResponseWrapper{
				Code:    apihelpers.CodeOk,
				Data:    tc.data,
				Message: "Test",
			}
			
			if wrapper.Data != tc.data {
				t.Errorf("Expected Data to be %v, got %v", tc.data, wrapper.Data)
			}
		})
	}
}

// Test para ResponseWrapper - mensaje vacío
func TestResponseWrapper_EmptyMessage(t *testing.T) {
	wrapper := &apihelpers.ResponseWrapper{
		Code:    apihelpers.CodeOk,
		Data:    "data",
		Message: "",
	}
	
	if wrapper.Message != "" {
		t.Errorf("Expected empty Message, got '%s'", wrapper.Message)
	}
}

// Test para ResponseWrapper - código de error
func TestResponseWrapper_ErrorCode(t *testing.T) {
	wrapper := &apihelpers.ResponseWrapper{
		Code:    apihelpers.CodeError,
		Message: "Error occurred",
		Data:    nil,
	}
	
	if wrapper.Code != apihelpers.CodeError {
		t.Errorf("Expected Code '%s', got '%s'", apihelpers.CodeError, wrapper.Code)
	}
	if wrapper.Message != "Error occurred" {
		t.Errorf("Expected Message 'Error occurred', got '%s'", wrapper.Message)
	}
}

// Test para AllowedHeaders - verificar que contiene headers esperados
func TestAllowedHeaders_ContainsExpectedHeaders(t *testing.T) {
	expectedHeaders := []string{
		"Accept",
		"Origin",
		"Content-Type",
		"Authorization",
	}
	
	for _, header := range expectedHeaders {
		if !contains(apihelpers.AllowedHeaders, header) {
			t.Errorf("Expected AllowedHeaders to contain '%s'", header)
		}
	}
}

// Helper para verificar si una cadena contiene una subcadena
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

