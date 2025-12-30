# Integraciones

Este documento describe todas las integraciones con APIs HTTP externas utilizadas en el backend de siniestros.

## APIs HTTP Externas

| Dependencia | Propósito | Focalpoint |
|------------|-----------|------------|
| API de Siniestros | Gestión completa del ciclo de vida de siniestros: creación, consulta, actualización, eliminación y gestión de documentos | |
| API de Storage | Almacenamiento y carga de archivos y documentos relacionados con siniestros | |
| API de Notificaciones | Envío de correos electrónicos para notificaciones del sistema | |
| API de Event Log | Registro centralizado de eventos y auditoría del sistema | |
| API de Qualitat Asesoría | Consulta de información de siniestros y beneficiarios desde sistema externo | |
| Cloud Function (Subscription Center) | Procesamiento de suscripciones y operaciones serverless | |
| API de GenAI/Diagnóstico | Servicios de inteligencia artificial para análisis y generación de diagnósticos automáticos | |

## Detalles de Integración

### 1. API de Siniestros

**Base URL:** Configurado en `${URL_SINISTER_API}`

**Endpoints utilizados:**

#### 1.1. Guardar Siniestro
- **Endpoint:** `POST /api/v1/siniestros/saveSiniestro`
- **Propósito:** Crear o actualizar información de un siniestro
- **Implementación:** `application/endpoints/sinister_api.go` - método `Save()`
- **Request Body:** Objeto con datos del siniestro (MapRequest)
- **Response:** Objeto ResponseSinister con información del siniestro guardado

#### 1.2. Obtener Siniestro
- **Endpoint:** `POST /api/v1/siniestros/getSiniestro`
- **Propósito:** Consultar un siniestro por número de solicitud
- **Implementación:** `application/endpoints/sinister_api.go` - método `FindByCaseNumber()`
- **Request Body:** 
  ```json
  {
    "numero_solicitud": "CIS_98746"
  }
  ```
- **Response:** Información completa del siniestro

#### 1.3. Agregar Documentos
- **Endpoint:** `POST /api/v1/siniestros/agregarDocumentos`
- **Propósito:** Añadir documentos a un siniestro existente
- **Implementación:** `application/endpoints/sinister_api.go` - método `AddDocument()`
- **Request Body:** Objeto con información del documento y referencia al siniestro

#### 1.4. Listar Siniestros (Historial)
- **Endpoint:** `POST /api/v1/siniestros/listSiniestros`
- **Propósito:** Obtener historial de siniestros según criterios de búsqueda
- **Implementación:** `application/endpoints/sinister_api.go` - método `FindByCaseHistory()`
- **Request Body:** Objeto SinisterHistoryRequest con filtros
- **Response:** Array de objetos SinisterHistory

#### 1.5. Actualizar Observación IA
- **Endpoint:** `POST /api/v1/siniestros/updateObservationIA`
- **Propósito:** Actualizar observaciones generadas por inteligencia artificial
- **Implementación:** `application/endpoints/sinister_api.go` - método `UpdateObservationIA()`
- **Request Body:** Objeto ObservationIARequest

#### 1.6. Eliminar Siniestro
- **Endpoint:** `DELETE /api/v1/siniestros/siniestro`
- **Propósito:** Eliminar un siniestro por número de caso
- **Implementación:** `application/endpoints/sinister_api.go` - método `Delete()`
- **Request Body:** 
  ```json
  {
    "numero_caso": "CIS_98746"
  }
  ```

---

### 2. API de Storage

**Base URL:** Configurado en `${URL_SINISTER_STORAGE}`

#### 2.1. Cargar Documento
- **Endpoint:** `POST /api/v1/documento/upload`
- **Propósito:** Subir archivos al sistema de almacenamiento
- **Implementación:** `application/endpoints/storage_api.go` - método `Upload()`
- **Content-Type:** `multipart/form-data`
- **Request:** 
  - Form data con metadatos del archivo
  - Archivo binario
- **Response:** Información del archivo subido (URL, ID, metadatos)

---

### 3. API de Notificaciones

**Base URL:** Configurado en `${URL_NOTIFICATIONS}`

#### 3.1. Enviar Correo Único
- **Endpoint:** `POST /v2/email/single/send?product=ZONA_SEGURA`
- **Propósito:** Enviar correos electrónicos individuales
- **Implementación:** `application/endpoints/notification_api.go` - método `SendMail()`
- **Content-Type:** `application/json` (enviado como multipart)
- **Request Body:** Objeto MapRequest con datos del correo
- **Parámetros Query:** `product=ZONA_SEGURA`

---

### 4. API de Event Log

**Base URL:** Configurado en `${URI_EVENT_LOG}`

#### 4.1. Registrar Evento
- **Endpoint:** `POST {URI_EVENT_LOG}`
- **Propósito:** Registrar eventos del sistema para auditoría y trazabilidad
- **Implementación:** `application/endpoints/event_api.go` - método `AddEvent()`
- **Autenticación:** API Key en header `X-Api-Key`
  - Clave estándar: `${API_KEY_EVENT_LOG}`
  - Clave WhatsApp: `${API_KEY_EVENT_WSP_LOG}` (usado cuando origin = "WHATSAPP-API")
- **Request Body:** Datos del evento desde TxContext.Event
- **Headers:**
  - `X-Api-Key`: API Key correspondiente según el origen

---

### 5. API de Qualitat Asesoría

**Base URL:** `https://qualitatasesoria.com/siniestros`

Esta integración realiza web scraping con autenticación sobre el sistema de Qualitat Asesoría.

#### 5.1. Proceso de Autenticación y Consulta
- **Implementación:** `application/endpoints/qualitat.go`
- **Autenticación:** 
  - Usuario: `${QUALITAT_USER}`
  - Contraseña: `${QUALITAT_PASS}`

**Flujo de integración:**

1. **Login Inicial**
   - **Endpoint:** `GET https://qualitatasesoria.com/siniestros/Login/Login.aspx`
   - **Propósito:** Obtener cookies de sesión

2. **Validar Login**
   - **Endpoint:** `POST https://qualitatasesoria.com/siniestros/Login/ValidarLogin.aspx`
   - **Propósito:** Autenticar usuario
   - **Form Data:**
     - `txtusuario`: Usuario de Qualitat
     - `txtclave`: Contraseña de Qualitat

3. **Guardar Configuración**
   - **Endpoint:** `POST https://qualitatasesoria.com/siniestros/General/GuardarConfiguracion.aspx`
   - **Propósito:** Establecer filtros de búsqueda (cliente, riesgo, ramo)
   - **Form Data:**
     - `id_perfil`: 2
     - `confcliente`: 1
     - `ddlconfriesgo`: 2
     - `ddlconframo`: 4

4. **Buscar Beneficiarios**
   - **Endpoint:** `POST https://qualitatasesoria.com/siniestros/General/ListadoBeneficiario.aspx`
   - **Propósito:** Buscar beneficiario por DNI o nombre
   - **Form Data:**
     - `txtdni`: Documento de identidad
     - `txtdatos`: Nombre completo (usado como fallback)
     - `modo`: C (Consulta)
     - `pagina`: 1
     - `regxpag`: 10
   - **Response:** HTML con tabla de beneficiarios

5. **Consultar Siniestros**
   - **Endpoint:** `POST https://qualitatasesoria.com/siniestros/Consultas/ListadoConsultaSiniestroCli.aspx`
   - **Propósito:** Obtener siniestros del beneficiario
   - **Form Data:**
     - `idbeneficiario`: ID del beneficiario
     - `txtbeneficiario`: Datos concatenados del beneficiario
     - Parámetros de paginación
   - **Response:** HTML con tabla de siniestros parseada a array de QualitatResponse

**Método Principal:**
- `QualitatStart(Document, InjuredCompleteName)` - Ejecuta todo el flujo

**Estructura de Respuesta (QualitatResponse):**
```go
{
    FechaSiniestro      string
    NroSiniestroCliente string
    NroSiniestro        string
    NroPoliza           string
    Placa               string
    NroCaso             string
    Nombres             string
    ApPaterno           string
    ApMaterno           string
    DNI                 string
    CentroMedico        string
    TipoSiniestro       string
    Recepcion           string
    Ocupante            string
    Fallecido           string
    FechaFallecimiento  string
    EstadoSiniestro     string
}
```

**Características especiales:**
- Utiliza cookie jar para mantener sesión
- Headers personalizados para simular navegador
- Parseo de HTML usando `golang.org/x/net/html`
- Extracción de datos mediante regex y análisis de DOM

---

### 6. Cloud Function (Subscription Center)

**Base URL:** Configurado en `${URI_CLOUD_FUNCTION}`

#### 6.1. Procesamiento de Suscripciones
- **Propósito:** Ejecutar funciones serverless relacionadas con el centro de suscripciones
- **Implementación:** `application/services/cloud_functio_sub_center.go`
- **Nota:** Los detalles específicos de endpoints dependen de la configuración de Cloud Functions

---

### 7. API de GenAI/Diagnóstico

**Configuración:**
- **URL Base:** `${SERVICES_SINISTER_API_GENAI_URL}`
- **API Key:** `${SERVICES_SINISTER_API_GENAI_KEY}`

#### 7.1. Servicios de IA
- **Propósito:** Procesamiento de inteligencia artificial para:
  - Análisis de documentos médicos (Parte Diario - DP)
  - Análisis de documentos médicos (Dictamen Médico - DM)
  - Generación de dictámenes automáticos
  - Análisis de diagnósticos médicos

**Prompts Configurados:**
- **Prompt DP:** `${PROMT_DP}` - Para análisis de Parte Diario
- **Prompt DM:** `${PROMT_DM}` - Para análisis de Dictamen Médico
- **Prompt Dictamen:** `${PROMT_DICTAMEN}` - Para generación de dictámenes

**API de Diagnóstico:** `${API_DIAGNOSTIC}` - Endpoint específico para análisis diagnósticos

---

## Notas de Implementación

### Manejo de Errores
Todas las integraciones implementan registro de errores en `TxContext.LastStageData` con los siguientes campos:
- `error`: Objeto de error
- `errorType`: Tipo de error descriptivo
- `statusCode`: Código HTTP de respuesta
- `response`: Respuesta del API (cuando está disponible)

### Headers Comunes
- Content-Type: Generalmente `application/json` o `application/x-www-form-urlencoded`
- User-Agent: Configurado para Qualitat para simular navegador Chrome
- X-Api-Key: Usado en API de Event Log

### Librerías Utilizadas
- `is-public-api/helpers/apihelpers`: Utilidades para requests HTTP
- `github.com/francoispqt/gojay`: Serialización JSON de alto rendimiento
- `golang.org/x/net/html`: Parseo de HTML para Qualitat

### Configuración
Todas las URLs y credenciales se configuran mediante variables de entorno referenciadas en `config/conf.yaml`.

