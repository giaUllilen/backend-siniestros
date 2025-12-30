package helpers_test

import (
	"bytes"
	"is-public-api/helpers/apihelpers"
	"mime/multipart"
	"testing"
)

// Helper para crear un multipart.FileHeader de prueba
func createTestFileHeader(filename string, content []byte) *multipart.FileHeader {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	
	part, _ := writer.CreateFormFile("file", filename)
	part.Write(content)
	writer.Close()
	
	reader := multipart.NewReader(body, writer.Boundary())
	form, _ := reader.ReadForm(10 << 20)
	
	if files, ok := form.File["file"]; ok && len(files) > 0 {
		return files[0]
	}
	return nil
}

// Test para DetectContentType - archivo de texto
func TestDetectContentType_TextFile(t *testing.T) {
	content := []byte("Este es un archivo de texto plano")
	fileHeader := createTestFileHeader("test.txt", content)
	
	if fileHeader == nil {
		t.Fatal("Failed to create test file header")
	}
	
	contentType := apihelpers.DetectContentType(fileHeader)
	
	if contentType == "" {
		t.Error("Expected content type to be detected")
	}
}

// Test para DetectContentType - archivo PDF
func TestDetectContentType_PDFFile(t *testing.T) {
	content := []byte("%PDF-1.4\n%some pdf content")
	fileHeader := createTestFileHeader("test.pdf", content)
	
	if fileHeader == nil {
		t.Fatal("Failed to create test file header")
	}
	
	contentType := apihelpers.DetectContentType(fileHeader)
	
	if contentType != "application/pdf" {
		t.Errorf("Expected content type 'application/pdf', got '%s'", contentType)
	}
}

// Test para ReadFile - archivo pequeño
func TestReadFile_SmallFile(t *testing.T) {
	expectedContent := []byte("Contenido de prueba pequeño")
	fileHeader := createTestFileHeader("small.txt", expectedContent)
	
	if fileHeader == nil {
		t.Fatal("Failed to create test file header")
	}
	
	buffer, err := apihelpers.ReadFile(fileHeader)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if buffer == nil {
		t.Fatal("Expected buffer, got nil")
	}
	
	actualContent := buffer.Bytes()
	if !bytes.Equal(expectedContent, actualContent) {
		t.Errorf("Expected content '%s', got '%s'", string(expectedContent), string(actualContent))
	}
}

// Test para ReadFile - archivo grande
func TestReadFile_LargeFile(t *testing.T) {
	expectedContent := bytes.Repeat([]byte("A"), 5000)
	fileHeader := createTestFileHeader("large.txt", expectedContent)
	
	if fileHeader == nil {
		t.Fatal("Failed to create test file header")
	}
	
	buffer, err := apihelpers.ReadFile(fileHeader)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if buffer == nil {
		t.Fatal("Expected buffer, got nil")
	}
	
	actualContent := buffer.Bytes()
	if !bytes.Equal(expectedContent, actualContent) {
		t.Errorf("Expected content length %d, got %d", len(expectedContent), len(actualContent))
	}
}

// Test para ReadFile - archivo vacío
func TestReadFile_EmptyFile(t *testing.T) {
	expectedContent := []byte{}
	fileHeader := createTestFileHeader("empty.txt", expectedContent)
	
	if fileHeader == nil {
		t.Fatal("Failed to create test file header")
	}
	
	buffer, err := apihelpers.ReadFile(fileHeader)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if buffer == nil {
		t.Fatal("Expected buffer, got nil")
	}
	
	actualContent := buffer.Bytes()
	if len(actualContent) != 0 {
		t.Errorf("Expected empty content, got %d bytes", len(actualContent))
	}
}

