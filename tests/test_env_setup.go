package tests

import (
	"os"
)

// SetupTestEnvironment configura las variables de entorno necesarias para las pruebas
func init() {
	// Configuración del servidor
	os.Setenv("SUBDOMAIN", "test")
	os.Setenv("KEY_AES_SHA", "test-aes-key-12345678901234567890123456789012")

	// Configuración de MongoDB
	os.Setenv("URI_DB_MONGO", "mongodb://localhost:27017")
	os.Setenv("NAME_DB_MONGO", "test_db")
	os.Setenv("USR_DB_MONGO", "test_user")
	os.Setenv("PWD_DB_MONGO", "test_password")

	// Configuración de servicios externos
	os.Setenv("URI_CLOUD_FUNCTION", "http://localhost:8080/cloud-function")
	os.Setenv("URL_SINISTER_API", "http://localhost:8080/sinister-api")
	os.Setenv("URL_SINISTER_STORAGE", "http://localhost:8080/storage")
	os.Setenv("URL_NOTIFICATIONS", "http://localhost:8080/notifications")
	
	// Configuración de Qualitat
	os.Setenv("QUALITAT_USER", "test_user")
	os.Setenv("QUALITAT_PASS", "test_pass")
	
	// Configuración de GenAI
	os.Setenv("SERVICES_SINISTER_API_GENAI_KEY", "test-genai-key")
	os.Setenv("SERVICES_SINISTER_API_GENAI_URL", "http://localhost:8080/genai")
	os.Setenv("PROMT_DP", "test prompt dp")
	os.Setenv("PROMT_DM", "test prompt dm")
	os.Setenv("PROMT_DICTAMEN", "test prompt dictamen")
	os.Setenv("API_DIAGNOSTIC", "http://localhost:8080/diagnostic")
	
	// Configuración de event log
	os.Setenv("URI_EVENT_LOG", "http://localhost:8080/event-log")
	os.Setenv("API_KEY_EVENT_LOG", "test-event-log-key")
	os.Setenv("API_KEY_EVENT_WSP_LOG", "test-wsp-log-key")
}

