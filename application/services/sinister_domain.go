package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"is-public-api/application/configs"
	"is-public-api/application/endpoints"
	"is-public-api/application/mappers"
	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/helpers/apihelpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type sinisterServiceDomain struct {
	sinister            endpoints.ISinisterEndpoints
	storage             endpoints.IStorageEndpoints
	event               endpoints.IEventsEndpoints
	notifications       endpoints.INotificationEndpoints
	templateMaker       IServiceTemplateMaker
	subDomain           string
	coverageIAService   ISinisterCoverageIAService
	aiService           ISinisterAIDomain
	sinisterCaseService ISinisterCaseService
}

func NewSinisterServiceDomain(conf *configs.ConfigServer, sinister endpoints.ISinisterEndpoints, storage endpoints.IStorageEndpoints,
	event endpoints.IEventsEndpoints, notifications endpoints.INotificationEndpoints, coverageIAService ISinisterCoverageIAService, aiService ISinisterAIDomain, sinisterCaseService ISinisterCaseService) *sinisterServiceDomain {

	return &sinisterServiceDomain{
		sinister:            sinister,
		storage:             storage,
		event:               event,
		notifications:       notifications,
		templateMaker:       NewTemplateHtmlMaker(),
		subDomain:           conf.Server.SubDomain,
		coverageIAService:   coverageIAService,
		aiService:           aiService,
		sinisterCaseService: sinisterCaseService,
	}
}

func (service *sinisterServiceDomain) Save(txContext *models.TxContext, request map[string]interface{}, coverages []map[string]interface{}, attachments []*multipart.FileHeader) ([]interface{}, error) {
	eventData := make(map[string]interface{})
	track := make(map[string]interface{})
	startTime := time.Now()

	eventData["Etapa"] = "Start"
	eventData["environment"] = txContext.Environment
	asegurado, okAsegurado := request["asegurado"].(map[string]interface{})
	solicitante, okSolicitante := request["solicitante"].(map[string]interface{})
	beneficiarios := request["beneficiarios"].([]interface{})
	observaciones_ai := request["observations_ai"]

	caseNumbers := make([]interface{}, len(coverages))

	if !okSolicitante || !okAsegurado {
		return nil, errors.New("BAD_REQUEST")
	}
	txContext.Event["numeroDocumento"] = asegurado["nroDocumento"]
	txContext.Event["categoria"] = "Funcionalidad"
	txContext.Event["nombre"] = "Registro de Siniestros ZP"
	eventData["request"] = request
	eventData["coberturas"] = coverages
	eventData["grupoProducto"] = request["tipoPoliza"]

	fechasIncapacidad, validFechas := request["fechasIncapacidad"].([]interface{})
	if !validFechas {
		fechasIncapacidad = make([]interface{}, 0)
	}
	tipoPoliza, ok := request["tipoPoliza"].(string)
	if !ok {
		txContext.Event.AddTimeTotal(startTime)
		eventData["error"] = "BAD_REQUEST tipoPoliza incorrecto"
		txContext.Event.Append(eventData)
		go service.event.AddEvent(txContext)
		return nil, errors.New("BAD_REQUEST")
	}

	docType := "DNI"
	if len(solicitante["nroDocumento"].(string)) != 8 {
		docType = "C/E"
	}
	date := time.Now().Add(time.Hour * -5)
	now := date.Format("02/01/2006")
	nowIso := date.Format("02-01-2006")
	nowDate := date.Format("02012006")
	occurrenceDate, _ := time.Parse("2006-01-02", fmt.Sprint(request["fechaOcurrencia"]))
	declaration, _ := request["declaracionJurada"].(bool)
	author := strings.TrimSpace(fmt.Sprint(solicitante["nombres"], " ", solicitante["apellidoPaterno"], " ", solicitante["apellidoMaterno"]))

	fechas := make([]interface{}, len(fechasIncapacidad))

	for index, value := range fechasIncapacidad {
		fecha := value.(map[string]interface{})
		starDateFormat := ""
		endDateFormat := ""
		starDate, err := time.Parse("2006-01-02", fmt.Sprint(fecha["fechaInicioIncapacidad"]))
		if err == nil {
			starDateFormat = starDate.Format("02-01-2006")
		}
		endDate, err := time.Parse("2006-01-02", fmt.Sprint(fecha["fechaFinIncapacidad"]))
		if err == nil {
			endDateFormat = endDate.Format("02-01-2006")
		}
		fechas[index] = map[string]interface{}{
			"FECHA_INICIO_INCAPACIDAD": starDateFormat,
			"FECHA_FIN_INCAPACIDAD":    endDateFormat,
		}

	}

	for x, coverage := range coverages {
		eventData["Etapa"] = "SinisterApiRequest_" + strconv.Itoa(x)
		requestSinister := map[string]interface{}{
			"APELLIDO_PATERNO_SOL":        solicitante["apellidoPaterno"], // Solicitante
			"APELLIDO_MATERNO_SOL":        solicitante["apellidoMaterno"], // Solicitante
			"AUTOR":                       author,                         // Solicitante
			"CANAL_DE_INGRESO":            "",                             // vacio
			"CODIGO_GLOBAL":               "",                             // vacio
			"CORREO_TERCERO":              solicitante["correo"],          // Solicitante
			"DEPARTAMENTO_SOLICITANTE":    "",                             // vacio
			"DESCRIPCION":                 request["narracion"],           // Detalle del siniestro
			"DIAS_TRANSCURRIDOS":          "",                             // vacio
			"DIRECCION":                   "",                             // vacio
			"DIRECCION_SOLICITANTE":       "",                             // vacio
			"DISTRITO_SOLICITANTE":        "",                             // vacio
			"ESTADO_DICTAMEN":             "pendiente",                    // Siempre pendiente
			"FECHA_CREACION_REGISTRO":     "",                             // vacio
			"FECHA_DICTAMEN":              "",                             // vacio
			"FECHA_ENTREGA_CLIENTE":       "",                             // vacio
			"FECHA_ENVIO":                 "",                             // vacio
			"FECHA_ENVIO_AUDITORIA":       "",                             // vacio
			"FECHA_ENVIO_COMERCIALIZADOR": "",                             // vacio
			"FECHA_MODIFICACION":          "",                             // vacio
			"FECHA_MODIFICACION_REGISTRO": "",                             // vacio
			"FECHA_PROCESO":               "",                             // vacio
			"FECHA_RECEPCION_AUDITORIA":   "",                             // vacio
			"FECHAS_INCAPACIDAD":          fechas,
			"FECHA_SOLICITUD":             nowIso,                      // "dd-mm-yyyy"
			"ID_FUENTE":                   10,                          // Siempre 10
			"ID_GRUPO_PRODUCTO":           0,                           // vacio
			"ID_INCIDENTE":                "",                          // vacio
			"ID_PERSONA":                  "",                          // vacio
			"INTERESES_MORATORIOS":        "",                          // vacio
			"LINEA_NEGOCIO":               "",                          // vacio
			"MEDIO_PACTADO_RESPUESTA":     "",                          // vacio
			"MONEDA":                      "",                          // vacio
			"MONTO_PAGADO":                "",                          // vacio
			"MOTIVO":                      "",                          // vacio
			"MOTIVO_OBSERVACION":          "",                          // vacio
			"MOTIVO_RECHAZO":              "",                          // vacio
			"NOMBRE_PRODUCTO_HM":          "",                          // vacio
			"NOMBRES_SOL":                 solicitante["nombres"],      // Solicitante
			"NUM_FILA":                    "",                          // vacio
			"NUMERO_CARTA":                "",                          // vacio
			"NUMERO_CASO":                 "",                          // vacio
			"NUMERO_CERTIFICADO":          "",                          // vacio
			"NUMERO_CORRELATIVO_ULTIMUS":  "",                          // vacio
			"NUMERO_DOCUMENTO":            solicitante["nroDocumento"], // Solicitante
			"NUMERO_SINIESTRO":            "",                          // vacio
			"ORIGEN":                      txContext.Origin,            // vacio
			"POLIZA":                      "",                          // vacio
			"PRIMER_NOMBRE":               solicitante["nombres"],      // vacio
			"PRODUCTO":                    tipoPoliza,                  // vacio (Solo se tiene el tipo de producto)
			"PROPIETARIO":                 author,                      // vacio
			"PROVINCIA_SOLICITANTE":       "",                          // vacio
			"SEGUNDO_NOMBRE":              "",                          // Solicitante
			"SUBCATEGORIA":                coverage["coverage"],        // Cobertura
			"TELEFONO_TERCERO":            solicitante["celular"],      // Solicitante
			"TIPO_CONTACTO_CLIENTE":       "",                          // vacio
			"TIPO_DOCUMENTO":              docType,                     // Solicitante
			"TIPO_DOCUMENTO_REC":          "",                          // vacio
			"TIPO_PERSONA":                "",                          // vacio
			"TITULO":                      "",                          // vacio
			"USUARIO_DICTAMINIO":          "",                          // vacio,
			"FECHA_OCURRENCIA":            occurrenceDate.Format("02-01-2006"),
			"MONTO_SOLICITADO":            fmt.Sprint(request["montoSolicitado"]),
			"NOMBRE_PAGO_GASTOS":          fmt.Sprint(request["pagador"]),
			"DECLARACION_JURADA":          declaration,
			"FECHA_INICIO_INCAPACIDAD":    "",
			"FECHA_FIN_INCAPACIDAD":       "",
		}

		beneficiaries := make([]map[string]string, len(beneficiarios))
		dta := "DNI"
		if len(asegurado["nroDocumento"].(string)) != 8 {
			dta = "C/E"
		}
		insured := map[string]string{
			"APELLIDO_PATERNO": asegurado["apellidoPaterno"].(string),
			"APELLIDO_MATERNO": asegurado["apellidoMaterno"].(string),
			"PRIMER_NOMBRE":    asegurado["nombres"].(string),
			"SEGUNDO_NOMBRE":   "",
			"TIPO_DOCUMENTO":   dta,
			"NUMERO_DOCUMENTO": asegurado["nroDocumento"].(string),
		}
		eventData["asegurado"] = insured

		for i, beneficiario := range beneficiarios {
			v := beneficiario.(map[string]interface{})
			dt := "DNI"
			if len(v["nroDocumento"].(string)) != 8 {
				dt = "C/E"
			}
			banco, _ := v["banco"].(string)
			tipoCuenta, _ := v["tipoCuenta"].(string)
			beneficiaries[i] = map[string]string{
				"BENEFICIARIO_SOLICITUD":     "",                         // vacio
				"APELLIDO_PATERNO":           "",                         // vacio
				"APELLIDO_MATERNO":           "",                         // vacio
				"TIPO_DOCUMENTO":             dt,                         // vacio
				"NUMERO_DOCUMENTO":           v["nroDocumento"].(string), // vacio
				"NOMBRES":                    "",                         // vacio
				"NOMBRE_TITULAR":             v["titular"].(string),      // opcional
				"MONTO_PAGADO":               "",                         // opcional
				"METODO_PAGO":                v["metodoPago"].(string),   // opcional
				"INTERESES_MORATORIOS":       "",                         // opcional
				"NUMERO_CUENTA":              v["nroCuenta"].(string),    // opcional
				"MONEDA_CUENTA":              v["moneda"].(string),       // opcional
				"BANCO":                      banco,                      // opcional
				"NUMERO_INCIDENTE_ULTIMUS":   "",                         // opcional
				"TIPO_CUENTA":                tipoCuenta,                 // opcional
				"NUMERO_CORRELATIVO_ULTIMUS": "",                         // opcional
				"ESTADO":                     "ACTIVO",                   // ACTIVO
			}
		}

		requestApi := map[string]interface{}{
			"siniestro":        requestSinister,
			"asegurado":        insured,
			"beneficiarios":    beneficiaries,
			"observaciones_ai": observaciones_ai,
		}

		eventData["Etapa"] = "SinisterApiCall_" + strconv.Itoa(x)
		startTimeEndpoint := time.Now()
		statusCode, responseSinister, err := service.sinister.Save(txContext, mappers.MapStringInterfaceToMapRequest(requestApi))
		txContext.Event.AddTimeEndpoint("SinisterApi", startTimeEndpoint)
		track["SinisterApi_"+strconv.Itoa(x)] = txContext.LastStageData.Pop()
		if err != nil {
			fmt.Println(statusCode, " - Error al guardar siniestro: ", err)
			eventData["error"] = err
			eventData["tracking"] = track
			txContext.Event.Append(eventData)
			txContext.Event.AddTimeTotal(startTime)
			go service.event.AddEvent(txContext)
			return nil, errors.New("INTERNAL_ERROR")
		}

		if responseSinister.Code != "01" {
			fmt.Println(statusCode, " - Error al guardar siniestro: ", statusCode, responseSinister)
			eventData["error"] = "error al registrar en el api de siniestros"
			eventData["tracking"] = track
			txContext.Event.Append(eventData)
			txContext.Event.AddTimeTotal(startTime)
			go service.event.AddEvent(txContext)
			return caseNumbers, errors.New("INTERNAL_ERROR")
		}

		contactID := fmt.Sprint("CIS_", responseSinister.Results.PersonDocID, "_", responseSinister.Results.CaseValueID, "_", nowDate)
		ticketNumber := fmt.Sprint("CIS_", responseSinister.Results.PersonDocID, "_", responseSinister.Results.CaseValueID)
		caseID := responseSinister.Results.CaseID
		caseNumbers[x] = responseSinister.Results.CaseID
		if eventData["casos"] == nil {
			eventData["casos"] = make([]map[string]interface{}, 0)
		}
		eventData["casos"] = append(eventData["casos"].([]map[string]interface{}), map[string]interface{}{"caso": caseNumbers[x], "cobertura": coverage["coverage"]})

		fmt.Println("Caso generado: ", responseSinister.Results.CaseID)
		documents, _ := coverage["documents"].([]interface{})
		additional, _ := coverage["additional"].([]interface{})

		for k, v := range append(documents, additional...) {
			document, documentOk := v.(map[string]interface{})
			if !documentOk {
				eventData["error"] = err
				eventData["tracking"] = track
				txContext.Event.Append(eventData)
				txContext.Event.AddTimeTotal(startTime)
				go service.event.AddEvent(txContext)
				return caseNumbers, errors.New("FILE_METADATA_PARSE_ERROR")
			}
			indexFloat, _ := document["index"].(float64)
			index := int(indexFloat)
			filename, _ := document["filename"].(string)
			if index < len(attachments) && strings.TrimSpace(filename) != "" {
				attachments[index].Filename = fmt.Sprint(caseID, "_", attachments[index].Filename)
				form := map[string]string{
					"app":         "siniestros",
					"contact_id":  contactID,
					"numero_caso": ticketNumber,
					"folder":      contactID,
				}

				fileType := apihelpers.DetectContentType(attachments[index])
				buf, err := apihelpers.ReadFile(attachments[index])
				// fmt.Println(index+1, ".- original mimeType: ", fileType)
				if fileType == "application/octet-stream" {
					fileType = "application/pdf"
				}
				fmt.Println(index+1, ".-  fileType: ", fileType)
				startTimeEndpoint = time.Now()
				eventData["Etapa"] = "StorageServiceCall_" + strconv.Itoa(k)
				// Sinister Add Document
				responseUpload, err := service.storage.Upload(txContext, mappers.MapStringToMapRequest(form), apihelpers.File{
					Key:         "file",
					Name:        attachments[index].Filename,
					File:        buf,
					ContentType: fileType,
					Size:        attachments[index].Size,
				})
				log.Println(responseUpload)
				txContext.Event.AddTimeEndpoint("StorageServiceCall_"+strconv.Itoa(k), startTimeEndpoint)
				track["StorageServiceCall_"+strconv.Itoa(k)] = txContext.LastStageData.Pop()
				if err != nil {
					// apihelpers.RenderError(ctx, errors.New("FILE_ERROR_REQUEST_API"), fasthttp.StatusBadRequest)
					fmt.Printf("🚨 ERROR STORAGE UPLOAD - Archivo: %s, Error: %v\n", attachments[index].Filename, err)
					eventData["error"] = err
					// Delete Sinister by numberCase
					startTimeEndpoint = time.Now()
					service.sinister.Delete(txContext, caseID)
					txContext.Event.AddTimeEndpoint("SinisterApiDelete_"+strconv.Itoa(k), startTimeEndpoint)
					track["SinisterApiDelete_"+strconv.Itoa(k)] = txContext.LastStageData.Pop()
					eventData["tracking"] = track
					txContext.Event.AddTimeTotal(startTime)
					txContext.Event.Append(eventData)
					go service.event.AddEvent(txContext)

					return caseNumbers, errors.New("FILE_ERROR_REQUEST_API")
				}

				// Sinister Add Document
				sinisterDoc := &resources.MapRequest{
					"app":         "siniestros",
					"contact_id":  contactID,
					"numero_caso": ticketNumber,
					"filename":    attachments[index].Filename,
				}
				startTimeEndpoint = time.Now()
				eventData["Etapa"] = "SinisterAddDocumentCall_" + strconv.Itoa(k)
				_, _, err = service.sinister.AddDocument(txContext, sinisterDoc)
				track["SinisterAddDocumentCall_"+strconv.Itoa(k)] = txContext.LastStageData.Pop()
				txContext.Event.AddTimeEndpoint("SinisterAddDocumentCall_"+strconv.Itoa(k), startTimeEndpoint)
			}
		}

		eventData["Etapa"] = "MakeTemplate"
		template, err := service.templateMaker.Make(txContext, mappers.SendEmailToMapRequest(service.subDomain, tipoPoliza, caseID, now, solicitante, asegurado))

		mailData := map[string]interface{}{
			"title":       "Registro de solicitud de cobertura de siniestros",
			"subject":     strings.Split(solicitante["nombres"].(string), " ")[0] + ", se ha registrado tu solicitud de cobertura de siniestros",
			"htmlContent": template.String(),
			"to": map[string]string{
				"name":  solicitante["nombres"].(string),
				"email": solicitante["correo"].(string),
			},
		}
		if service.subDomain != "www" {
			mailData["subject"] = fmt.Sprint("[TEST] ", mailData["subject"])
		}
		startTimeEndpoint = time.Now()
		eventData["Etapa"] = "SendMail"
		_, err = service.notifications.SendMail(txContext, mappers.MapStringInterfaceToMapRequest(mailData))
		txContext.Event.AddTimeEndpoint("SendMail", startTimeEndpoint)
		track["SendMailApi"] = txContext.LastStageData.Pop()
		eventData["Etapa"] = "Final"
		eventData["tracking"] = track
		txContext.Event.Append(eventData)
		txContext.Event.AddTimeTotal(startTime)
		go service.event.AddEvent(txContext)
	}

	// Verificar combinaciones producto-cobertura configuradas en siniestro_coberturas_ia
	for i, coverage := range coverages {
		if coverageType, ok := coverage["coverage"].(string); ok {
			isValidCombination, err := service.coverageIAService.IsValidCombination(txContext, tipoPoliza, coverageType)
			if err != nil {
				fmt.Printf("⚠️ Error al validar combinación %s-%s: %v\n", tipoPoliza, coverageType, err)
				continue
			}

			if isValidCombination {
				fmt.Printf("🔍 LOG ESPECIAL - Producto: %s, Cobertura: %s (Configurado en siniestro_coberturas_ia)\n", tipoPoliza, coverageType)

				// Obtener el caseNumber correspondiente a esta cobertura
				var caseNumber string
				if i < len(caseNumbers) {
					if caseNum, ok := caseNumbers[i].(string); ok {
						caseNumber = caseNum
					}
				}

				// Agregar la validación con IA para documentos de esta cobertura
				service.processDocumentsWithIA(txContext, coverage, tipoPoliza, coverageType, attachments, caseNumber)
			}
		}
	}

	return caseNumbers, nil
}

// processDocumentsWithIA procesa documentos con IA para coberturas válidas
func (service *sinisterServiceDomain) processDocumentsWithIA(txContext *models.TxContext, coverage map[string]interface{}, tipoPoliza, coverageType string, attachments []*multipart.FileHeader, caseNumber string) {
	// Extraer solo documentos principales (no adicionales)
	documents, documentsOk := coverage["documents"].([]interface{})

	// Solo procesar documentos principales
	allDocuments := make([]interface{}, 0)
	if documentsOk {
		allDocuments = append(allDocuments, documents...)
	}

	// Ordenar documentos por prioridad: primero "Denuncia Policial", luego "Certificado médico"
	orderedDocuments := make([]interface{}, 0)
	var denunciaPolicial []interface{}
	var certificadoMedico []interface{}
	var otrosDocumentos []interface{}

	// Clasificar documentos por tipo
	for _, doc := range allDocuments {
		document, ok := doc.(map[string]interface{})
		if !ok {
			continue
		}

		documentName, nameOk := document["name"].(string)
		if !nameOk {
			continue
		}

		switch documentName {
		case "Denuncia Policial":
			denunciaPolicial = append(denunciaPolicial, doc)
		case "Certificado médico":
			certificadoMedico = append(certificadoMedico, doc)
		case "Descanso medico":
			certificadoMedico = append(certificadoMedico, doc) // Tratar "Descanso medico" como certificado médico
		default:
			otrosDocumentos = append(otrosDocumentos, doc)
		}
	}

	// Agregar en orden de prioridad
	orderedDocuments = append(orderedDocuments, denunciaPolicial...)
	orderedDocuments = append(orderedDocuments, certificadoMedico...)
	orderedDocuments = append(orderedDocuments, otrosDocumentos...)

	fmt.Printf("📋 Procesando %d documentos en orden de prioridad (Denuncia Policial → Certificado médico → Otros)\n", len(orderedDocuments))

	// Procesar documentos de forma secuencial sincronizada
	// El primer análisis (Denuncia Policial) debe completarse antes del segundo
	service.processDocumentsSequentially(txContext, orderedDocuments, tipoPoliza, coverageType, attachments, caseNumber)
}

// processDocumentsSequentially procesa documentos de forma secuencial sincronizada
// donde cada análisis espera a que termine el anterior y puede usar su resultado
func (service *sinisterServiceDomain) processDocumentsSequentially(txContext *models.TxContext, documents []interface{}, tipoPoliza, coverageType string, attachments []*multipart.FileHeader, caseNumber string) {
	var previousAnalysisID string                // ID del análisis anterior para pasar al siguiente
	var analysisResults []map[string]interface{} // Capturar todos los resultados de análisis

	// Capturar caseNumber para usar en la goroutine
	casNum := caseNumber

	go func() {
		// Crear un nuevo contexto para la goroutine principal
		asyncTxContext := &models.TxContext{
			Event:       txContext.Event,
			Origin:      txContext.Origin,
			Environment: txContext.Environment,
		}

		// Recuperación de panics en la goroutine principal
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("🚨 Panic recuperado en processDocumentsSequentially: %v\n", r)
				fmt.Printf("📍 Stack trace completo:\n%s\n", debug.Stack())
			}
		}()

		// Procesar cada documento secuencialmente
		for i, doc := range documents {
			document, ok := doc.(map[string]interface{})
			if !ok {
				continue
			}

			filename, filenameOk := document["filename"].(string)
			documentName, nameOk := document["name"].(string)
			indexFloat, indexOk := document["index"].(float64)

			if !filenameOk || !nameOk || !indexOk || strings.TrimSpace(filename) == "" {
				continue
			}

			// Filtrar solo documentos específicos
			if documentName != "Denuncia Policial" && documentName != "Certificado médico" && documentName != "Descanso medico" {
				continue
			}

			index := int(indexFloat)

			// Verificar que el índice sea válido y obtener el archivo en base64
			var fileBase64 string
			if index < len(attachments) {
				buf, err := apihelpers.ReadFile(attachments[index])
				if err != nil {
					fmt.Printf("❌ Error al leer archivo %s: %v\n", filename, err)
					continue
				}
				fileBase64 = base64.StdEncoding.EncodeToString(buf.Bytes())
			} else {
				fmt.Printf("❌ Índice de archivo inválido para %s: %d\n", filename, index)
				continue
			}

			// Crear el request para AnalyzeDocumentWeb
			// Para el primer documento, usar un ID generado localmente
			// Para documentos posteriores, usar el ID del análisis anterior si está disponible
			var caseID string
			if i == 0 || previousAnalysisID == "" {
				caseID = fmt.Sprintf("%s_%s_IA_%d", tipoPoliza, coverageType, time.Now().Unix())
			} else {
				caseID = previousAnalysisID // Usar el ID del análisis anterior
			}

			// Obtener el número de documento del contexto del evento
			numeroDocumento := ""
			if numDoc, exists := txContext.Event["numeroDocumento"]; exists {
				numeroDocumento = fmt.Sprintf("%v", numDoc)
			}

			sinisterCase := &resources.SinisterCase{
				Caso:            caseID,
				NumeroDocumento: numeroDocumento,
				Documento: resources.Document{
					Name:     documentName,
					Filename: filename,
					FileURL:  fmt.Sprintf("%s/documents/%s", service.subDomain, filename),
				},
			}

			fmt.Printf("🔄 [%d] Iniciando análisis SECUENCIAL para documento: %s (Caso: %s, NumeroDocumento: %s)\n",
				i+1, documentName, caseID, numeroDocumento)

			if i > 0 && previousAnalysisID != "" {
				fmt.Printf("📎 [%d] Usando ID del análisis anterior: %s\n", i+1, previousAnalysisID)
				sinisterCase.PreviousAnalysisID = previousAnalysisID
			}

			// Llamar al método AnalyzeDocumentWeb de forma SÍNCRONA (sin goroutine)
			result, err := service.aiService.AnalyzeDocumentWeb(asyncTxContext, sinisterCase, fileBase64)

			// Agregar logs para ver la respuesta
			fmt.Printf("📋 [%d] Respuesta de AnalyzeDocumentWeb para %s:\n", i+1, documentName)
			if result != nil {
				fmt.Printf("✅ [%d] Resultado exitoso: %+v\n", i+1, result)
			} else {
				fmt.Printf("⚠️ [%d] Resultado es nil\n", i+1)
			}

			if err != nil {
				fmt.Printf("❌ [%d] Error en análisis IA para %s-%s (documento: %s): %v\n",
					i+1, tipoPoliza, coverageType, documentName, err)
				continue // Continuar con el siguiente documento aunque falle uno
			}

			// Capturar el resultado para usar en updateObservationIA
			if result != nil {
				analysisResultItem := map[string]interface{}{
					"documento": documentName,
					"resultado": result,
					"caso_id":   previousAnalysisID,
				}

				// Si obtuvimos datos de buc-data, agregarlos al resultado
				if sinisterCase.PreviousAnalysisID != "" {
					fmt.Printf("🔍 [%d] Obteniendo registro de buc-data con ID: %s\n", i+1, sinisterCase.PreviousAnalysisID)
					objectID, err := primitive.ObjectIDFromHex(sinisterCase.PreviousAnalysisID)
					if err == nil {
						bucDataRecord, err := service.sinisterCaseService.FindByCaseNumber(asyncTxContext, objectID)
						if err == nil && bucDataRecord != nil {
							analysisResultItem["buc_data_record"] = bucDataRecord
							fmt.Printf("📦 [%d] Datos de buc-data agregados al resultado para %s\n", i+1, documentName)
						}
					}
				}

				analysisResults = append(analysisResults, analysisResultItem)
			}

			// Procesar el resultado y extraer el ID para el siguiente análisis
			if result != nil {
				// Extraer el campo "caso" del resultado (está en el nivel raíz, no dentro de "data")

				if caso, casoExists := result["caso"]; casoExists {
					previousAnalysisID = fmt.Sprintf("%v", caso)
					if _, err := primitive.ObjectIDFromHex(previousAnalysisID); err != nil {
						// Usar el caseID generado localmente como fallback
						previousAnalysisID = caseID
					}
				} else {
					// Fallback: usar el caseID generado localmente
					previousAnalysisID = caseID
				}
			}

		}

		// Al finalizar el procesamiento, enviar los resultados a updateObservationIA
		fmt.Printf("🏁 Procesamiento completado. Enviando %d resultados a updateObservationIA\n", len(analysisResults))
		fmt.Printf("📊 Contenido completo de analysisResults:\n")
		for i, result := range analysisResults {
			fmt.Printf("📋 [%d] analysisResults[%d]:\n", i+1, i)
			fmt.Printf("    📄 Documento: %v\n", result["documento"])
			fmt.Printf("    🆔 Caso ID: %v\n", result["caso_id"])
			fmt.Printf("    📝 Resultado completo: %+v\n", result["resultado"])
			fmt.Printf("    ─────────────────────────────────\n")
		}
		service.updateObservationIAWithResults(asyncTxContext, tipoPoliza, coverageType, analysisResults, casNum)
	}()
}

// updateObservationIAWithResults envía los datos de buc-data directamente
func (service *sinisterServiceDomain) updateObservationIAWithResults(txContext *models.TxContext, tipoPoliza, coverageType string, analysisResults []map[string]interface{}, caseNumber string) {
	if len(analysisResults) == 0 {
		fmt.Printf("⚠️ No hay resultados de análisis para enviar a updateObservationIA\n")
		return
	}

	// Usar el caseNumber pasado como parámetro
	numeroCaso := caseNumber

	// Validar que tenemos un número de caso válido
	if numeroCaso == "" {
		fmt.Printf("⚠️ No se proporcionó número de caso, omitiendo envío de observaciones IA\n")
		return
	}

	fmt.Printf("📊 Procesando %d resultados con datos de buc-data para caso: %s\n", len(analysisResults), numeroCaso)

	// Buscar y extraer dictamen y observations del buc_data_record
	var observacionesIA models.ObservacionesIA
	var found bool

	for i, analysisResult := range analysisResults {
		documento := analysisResult["documento"].(string)

		// Verificar si tenemos datos de buc-data
		if bucDataRecord, hasBucData := analysisResult["buc_data_record"]; hasBucData {
			fmt.Printf("📋 [%d] Extrayendo dictamen y observations de buc-data para %s\n", i+1, documento)

			// Convertir bucDataRecord a la estructura correcta
			if casoSiniestro, ok := bucDataRecord.(*models.CasoSiniestro); ok && casoSiniestro != nil {
				// Extraer dictamen y observations de observaciones.data.answer
				if casoSiniestro.Observation != nil &&
					casoSiniestro.Observation.Data.Answer.Dictamen.Justificacion != "" {

					fmt.Printf("🎯 [%d] Extrayendo dictamen y observations de observaciones.data.answer\n", i+1)

					// Crear la estructura ObservacionesIA con los datos extraídos
					observacionesIA = models.ObservacionesIA{
						Dictamen: models.DictamenIA{
							Justificacion: casoSiniestro.Observation.Data.Answer.Dictamen.Justificacion,
							Status:        casoSiniestro.Observation.Data.Answer.Dictamen.Status,
						},
						Observations: make([]models.ObservationIA, len(casoSiniestro.Observation.Data.Answer.Observations)),
					}

					// Copiar las observaciones
					for j, obs := range casoSiniestro.Observation.Data.Answer.Observations {
						observacionesIA.Observations[j] = models.ObservationIA{
							Classification: obs.Classification,
							Text:           obs.Text,
						}
					}

					found = true
					fmt.Printf("✅ [%d] Se extrajeron dictamen y %d observaciones del buc-data\n", i+1, len(observacionesIA.Observations))
					break // Usar el primer buc-data válido encontrado
				}
			}
		}
	}

	if !found {
		fmt.Printf("⚠️ No se encontraron observaciones.data.answer válidas en ningún buc-data\n")
		return
	}

	// Crear el request con dictamen y observations extraídos del buc-data
	observationRequest := models.ObservationIARequest{
		NumeroCaso:      numeroCaso,
		ObservacionesIA: observacionesIA,
	}

	// Llamar al endpoint de forma asíncrona
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("🚨 Panic recuperado en updateObservationIAWithResults: %v\n", r)
			}
		}()

		fmt.Printf("📤 Enviando dictamen y observations extraídos del buc-data para %s-%s (Caso: %s)\n", tipoPoliza, coverageType, numeroCaso)
		fmt.Printf("📋 Request completo: %+v\n", observationRequest)

		statusCode, response, err := service.sinister.UpdateObservationIA(txContext, observationRequest)

		if err != nil {
			fmt.Printf("❌ Error al enviar observaciones IA: %v\n", err)
		} else {
			fmt.Printf("✅ Observaciones IA enviadas exitosamente - StatusCode: %d\n", statusCode)
			fmt.Printf("📋 Respuesta del servidor: %s\n", string(response))
		}
	}()
}

// getMapKeys helper function para debugging
func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// FindByCaseNumber - Modificado para recibir caseNumber como string (ej: "CIS_98746")
// en lugar de primitive.ObjectID, ya que el endpoint externo de API siniestros 
// espera el número de caso en formato string
// Anterior:
// func (service *sinisterServiceDomain) FindByCaseNumber(txContext *models.TxContext, caseNumber primitive.ObjectID) (int, []byte, error) {
//     eventData["caso"] = caseNumber
//     statusCode, bytes, err := service.sinister.FindByCaseNumber(txContext, caseNumber)
func (service *sinisterServiceDomain) FindByCaseNumber(txContext *models.TxContext, caseNumber string) (int, []byte, error) {
	eventData := make(map[string]interface{})
	startTime := time.Now()
	var response []map[string]interface{}
	txContext.Event["nombre"] = "Consulta de Siniestro ZP"
	txContext.Event["categoria"] = "Consulta"

	eventData["caso"] = caseNumber
	statusCode, bytes, err := service.sinister.FindByCaseNumber(txContext, caseNumber)
	eventData["statusCode"] = statusCode
	txContext.Event.AddTimeTotal(startTime)
	err = json.Unmarshal(bytes, &response)
	if err == nil && len(response) > 0 {
		txContext.Event["numeroDocumento"] = response[0]["NUMERO_DOCUMENTO"]
		eventData["estado"] = response[0]["ESTADO_DICTAMEN"]
		eventData["grupoProducto"] = response[0]["NOMBRE_GRUPO_PRODUCTO"]
		eventData["semaforo"] = response[0]["SEMAFORO"]
		txContext.Event.Append(eventData)
		go service.event.AddEvent(txContext)
	}

	return statusCode, bytes, err
}

func (service *sinisterServiceDomain) FindByCaseHistory(txContext *models.TxContext, documentNumber string) (int, []models.SinisterHistory, error) {

	eventData := make(map[string]interface{})
	startTime := time.Now()
	txContext.Event["nombre"] = "Consulta de Siniestro WS"
	txContext.Event["categoria"] = "Consulta"

	eventData["numero_documento"] = documentNumber

	request := models.SinisterHistoryRequest{
		Filtro:             []string{"documento"},
		CurrentPage:        1,
		PageSize:           10,
		EstadoAprobacion:   "",
		EstadoDictamen:     "",
		FechaDesde:         "",
		FechaHasta:         "",
		FechaDictamenDesde: "",
		FechaDictamenHasta: "",
		FechaIngresoDesde:  "",
		FechaIngresoHasta:  "",
		NumeroDocumento:    documentNumber,
		NumeroTicket:       "",
		Producto:           "",
		TipoDocumento:      "DNI",
		NumeroSiniestro:    "",
	}

	statusCode, bytes, err := service.sinister.FindByCaseHistory(txContext, request)

	if err != nil {
		fmt.Printf("❌ FindByCaseHistory - Error en la consulta: %v\n", err)
	}

	eventData["statusCode"] = statusCode
	txContext.Event.AddTimeTotal(startTime)

	if len(bytes) > 0 {
		txContext.Event["numeroDocumento"] = documentNumber
		eventData["resultado"] = bytes
		txContext.Event.Append(eventData)
	}

	go service.event.AddEvent(txContext)
	return statusCode, bytes, err
}
