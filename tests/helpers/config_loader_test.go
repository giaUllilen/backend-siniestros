package helpers_test

import (
	"is-public-api/helpers/configloader"
	"testing"
)

// Test para GetVal - valor válido int
func TestGetVal_ValidInt(t *testing.T) {
	value := 42
	defaultValue := 0
	
	result := configloader.GetVal(value, defaultValue)
	
	if result != value {
		t.Errorf("Expected %d, got %v", value, result)
	}
}

// Test para GetVal - valor cero int (debería devolver default)
func TestGetVal_ZeroInt(t *testing.T) {
	value := 0
	defaultValue := 100
	
	result := configloader.GetVal(value, defaultValue)
	
	if result != defaultValue {
		t.Errorf("Expected default value %d, got %v", defaultValue, result)
	}
}

// Test para GetVal - valor válido string
func TestGetVal_ValidString(t *testing.T) {
	value := "test value"
	defaultValue := "default"
	
	result := configloader.GetVal(value, defaultValue)
	
	if result != value {
		t.Errorf("Expected '%s', got '%v'", value, result)
	}
}

// Test para GetVal - string vacío (debería devolver default)
func TestGetVal_EmptyString(t *testing.T) {
	value := ""
	defaultValue := "default value"
	
	result := configloader.GetVal(value, defaultValue)
	
	if result != defaultValue {
		t.Errorf("Expected default value '%s', got '%v'", defaultValue, result)
	}
}

// Test para GetVal - valor nil (debería devolver default)
func TestGetVal_NilValue(t *testing.T) {
	var value interface{} = nil
	defaultValue := "default"
	
	result := configloader.GetVal(value, defaultValue)
	
	if result != defaultValue {
		t.Errorf("Expected default value '%s', got '%v'", defaultValue, result)
	}
}

// Test para GetVal - valor false (debería devolver default)
func TestGetVal_FalseBool(t *testing.T) {
	value := false
	defaultValue := true
	
	result := configloader.GetVal(value, defaultValue)
	
	if result != defaultValue {
		t.Errorf("Expected default value %v, got %v", defaultValue, result)
	}
}

// Test para GetVal - valor true (debería devolver el valor)
func TestGetVal_TrueBool(t *testing.T) {
	value := true
	defaultValue := false
	
	result := configloader.GetVal(value, defaultValue)
	
	if result != value {
		t.Errorf("Expected %v, got %v", value, result)
	}
}

// Test para GetVal - número negativo
func TestGetVal_NegativeNumber(t *testing.T) {
	value := -42
	defaultValue := 0
	
	result := configloader.GetVal(value, defaultValue)
	
	if result != value {
		t.Errorf("Expected %d, got %v", value, result)
	}
}

