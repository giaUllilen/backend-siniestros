# Variables de entorno

| Variable | Descripción | Valores |
|----------|-------------|---------|
| `SUBDOMAIN` | Subdominio del servidor de la aplicación | `dev`, `qa`, `prod` |
| `KEY_AES_SHA` | Clave AES para encriptación de datos sensibles | `abc123def456...` (cadena hexadecimal) |
| `URI_DB_MONGO` | URI de conexión a la base de datos MongoDB | `mongodb://localhost:27017` |
| `NAME_DB_MONGO` | Nombre de la base de datos MongoDB | `siniestros_dev`, `siniestros_prod` |
| `USR_DB_MONGO` | Usuario para autenticación en MongoDB | `admin`, `db_user` |
| `PWD_DB_MONGO` | Contraseña para autenticación en MongoDB | `password123` |
| `URI_CLOUD_FUNCTION` | URI de Google Cloud Function para suscripciones | `https://us-central1-project.cloudfunctions.net/function` |
| `URL_SINISTER_API` | URL base del API de siniestros | `https://api.example.com` |
| `SERVICES_SINISTER_API_GENAI_KEY` | API Key para el servicio de GenAI | `AIzaSyD...` |
| `SERVICES_SINISTER_API_GENAI_URL` | URL del API de GenAI para procesamiento de IA | `https://generativelanguage.googleapis.com` |
| `PROMT_DP` | ID del prompt DP para el análisis de IA | `prompt_dp_001` |
| `PROMT_DM` | ID del prompt DM para el análisis de IA | `prompt_dm_001` |
| `PROMT_DICTAMEN` | ID del prompt de dictamen para el análisis de IA | `prompt_dictamen_001` |
| `API_DIAGNOSTIC` | URL del API de diagnóstico | `https://diagnostic.api.example.com` |
| `URL_SINISTER_STORAGE` | URL del servicio de almacenamiento de documentos | `https://storage.example.com` |
| `URL_NOTIFICATIONS` | URL del servicio de notificaciones por email | `https://notifications.example.com` |
| `QUALITAT_USER` | Usuario para autenticación en Qualitat | `qualitat_user` |
| `QUALITAT_PASS` | Contraseña para autenticación en Qualitat | `qualitat_pass123` |
| `URI_EVENT_LOG` | URI del servicio de registro de eventos | `https://eventlog.example.com` |
| `API_KEY_EVENT_LOG` | API Key para el servicio de Event Log | `event_log_key_123` |
| `API_KEY_EVENT_WSP_LOG` | API Key para el servicio de Event WSP Log | `event_wsp_key_123` |

