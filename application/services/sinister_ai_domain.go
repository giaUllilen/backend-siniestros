package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"is-public-api/application/configs"
	"is-public-api/application/endpoints"
	"is-public-api/application/mappers"
	"is-public-api/application/models"
	"is-public-api/application/resources"
	helpers_ai "is-public-api/helpers/apihelpers"
	"is-public-api/helpers/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type sinisterAIServiceDomain struct {
	sinister             endpoints.ISinisterEndpoints
	storage              endpoints.IStorageEndpoints
	event                endpoints.IEventsEndpoints
	notifications        endpoints.INotificationEndpoints
	templateMaker        IServiceTemplateMaker
	subDomain            string
	conf                 *helpers_ai.EndpointHelper
	qualitat             endpoints.IQualitatEndpoints
	soat                 ISoatFinder
	sinisterCase         ISinisterCaseService
	domainPortalSinister ISinisterServiceDomain
}

func NewSinisterAIServiceDomain(conf *configs.ConfigServer, sinister endpoints.ISinisterEndpoints,
	storage endpoints.IStorageEndpoints, event endpoints.IEventsEndpoints, notifications endpoints.INotificationEndpoints,
	confHelper *helpers_ai.EndpointHelper, qualitat endpoints.IQualitatEndpoints, soat ISoatFinder,
	sinisterCase ISinisterCaseService, domainPortalSinister ISinisterServiceDomain,
) ISinisterAIDomain {
	return &sinisterAIServiceDomain{
		sinister:             sinister,
		storage:              storage,
		event:                event,
		notifications:        notifications,
		templateMaker:        NewTemplateHtmlMaker(),
		subDomain:            conf.Server.SubDomain,
		conf:                 confHelper,
		qualitat:             qualitat,
		soat:                 soat,
		sinisterCase:         sinisterCase,
		domainPortalSinister: domainPortalSinister,
	}
}

func (service *sinisterAIServiceDomain) Save(txContext *models.TxContext, sinister *models.Sinister) ([]interface{}, error) {
	eventData := make(map[string]interface{})
	track := make(map[string]interface{})
	startTime := time.Now()
	eventData["Etapa"] = "Start"
	eventData["environment"] = txContext.Environment
	caseNumbers := make([]interface{}, len(sinister.Documentos))

	txContext.Event["numeroDocumento"] = sinister.Asegurado.NroDocumento
	txContext.Event["categoria"] = "Funcionalidad"
	txContext.Event["nombre"] = "Registro de Siniestros ZP"
	eventData["request"] = sinister

	tipoPoliza := sinister.TipoPoliza

	docType := "DNI"
	if len(sinister.Solicitante.NroDocumento) != 8 {
		docType = "C/E"
	}
	date := time.Now().Add(time.Hour * -5)
	// now := date.Format("02/01/2006")
	nowIso := date.Format("02-01-2006")
	nowDate := date.Format("02012006")
	occurrenceDate, _ := time.Parse("2006-01-02", fmt.Sprint(sinister.FechaOcurrencia))
	declaration := sinister.DeclaracionJurada
	author := strings.TrimSpace(fmt.Sprint(sinister.Solicitante.Nombres, " ", sinister.Solicitante.ApellidoPaterno, " ", sinister.Solicitante.ApellidoMaterno))

	fechas := make([]interface{}, len(sinister.FechasIncapacidad))

	for index, fecha := range sinister.FechasIncapacidad {
		starDateFormat := ""
		endDateFormat := ""
		starDate, err := time.Parse("2006-01-02", fmt.Sprint(fecha.FechaInicioIncapacidad))
		if err == nil {
			starDateFormat = starDate.Format("02-01-2006")
		}
		endDate, err := time.Parse("2006-01-02", fmt.Sprint(fecha.FechaFinIncapacidad))
		if err == nil {
			endDateFormat = endDate.Format("02-01-2006")
		}
		fechas[index] = map[string]interface{}{
			"FECHA_INICIO_INCAPACIDAD": starDateFormat,
			"FECHA_FIN_INCAPACIDAD":    endDateFormat,
		}

	}

	for x, coverage := range sinister.Documentos {
		eventData["Etapa"] = "SinisterApiRequest_" + strconv.Itoa(x)
		requestSinister := map[string]interface{}{
			"APELLIDO_PATERNO_SOL":        sinister.Solicitante.ApellidoPaterno, // Solicitante
			"APELLIDO_MATERNO_SOL":        sinister.Solicitante.ApellidoMaterno, // Solicitante
			"AUTOR":                       author,                               // Solicitante
			"CANAL_DE_INGRESO":            "",                                   // vacio
			"CODIGO_GLOBAL":               "",                                   // vacio
			"CORREO_TERCERO":              sinister.Solicitante.Correo,          // Solicitante
			"DEPARTAMENTO_SOLICITANTE":    "",                                   // vacio
			"DESCRIPCION":                 sinister.Narracion,                   // Detalle del siniestro
			"DIAS_TRANSCURRIDOS":          "",                                   // vacio
			"DIRECCION":                   "",                                   // vacio
			"DIRECCION_SOLICITANTE":       "",                                   // vacio
			"DISTRITO_SOLICITANTE":        "",                                   // vacio
			"ESTADO_DICTAMEN":             "pendiente",                          // Siempre pendiente
			"FECHA_CREACION_REGISTRO":     "",                                   // vacio
			"FECHA_DICTAMEN":              "",                                   // vacio
			"FECHA_ENTREGA_CLIENTE":       "",                                   // vacio
			"FECHA_ENVIO":                 "",                                   // vacio
			"FECHA_ENVIO_AUDITORIA":       "",                                   // vacio
			"FECHA_ENVIO_COMERCIALIZADOR": "",                                   // vacio
			"FECHA_MODIFICACION":          "",                                   // vacio
			"FECHA_MODIFICACION_REGISTRO": "",                                   // vacio
			"FECHA_PROCESO":               "",                                   // vacio
			"FECHA_RECEPCION_AUDITORIA":   "",                                   // vacio
			"FECHAS_INCAPACIDAD":          fechas,
			"FECHA_SOLICITUD":             nowIso,                            // "dd-mm-yyyy"
			"ID_FUENTE":                   10,                                // Siempre 10
			"ID_GRUPO_PRODUCTO":           0,                                 // vacio
			"ID_INCIDENTE":                "",                                // vacio
			"ID_PERSONA":                  "",                                // vacio
			"INTERESES_MORATORIOS":        "",                                // vacio
			"LINEA_NEGOCIO":               "",                                // vacio
			"MEDIO_PACTADO_RESPUESTA":     "",                                // vacio
			"MONEDA":                      "",                                // vacio
			"MONTO_PAGADO":                "",                                // vacio
			"MOTIVO":                      "",                                // vacio
			"MOTIVO_OBSERVACION":          "",                                // vacio
			"MOTIVO_RECHAZO":              "",                                // vacio
			"NOMBRE_PRODUCTO_HM":          "",                                // vacio
			"NOMBRES_SOL":                 sinister.Solicitante.Nombres,      // Solicitante
			"NUM_FILA":                    "",                                // vacio
			"NUMERO_CARTA":                "",                                // vacio
			"NUMERO_CASO":                 "",                                // vacio
			"NUMERO_CERTIFICADO":          "",                                // vacio
			"NUMERO_CORRELATIVO_ULTIMUS":  "",                                // vacio
			"NUMERO_DOCUMENTO":            sinister.Solicitante.NroDocumento, // Solicitante
			"NUMERO_SINIESTRO":            "",                                // vacio
			"ORIGEN":                      txContext.Origin,                  // vacio
			"POLIZA":                      "",                                // vacio
			"PRIMER_NOMBRE":               sinister.Solicitante.Nombres,      // vacio
			"PRODUCTO":                    tipoPoliza,                        // vacio (Solo se tiene el tipo de producto)
			"PROPIETARIO":                 author,                            // vacio
			"PROVINCIA_SOLICITANTE":       "",                                // vacio
			"SEGUNDO_NOMBRE":              "",                                // Solicitante
			"SUBCATEGORIA":                coverage.Coverage,                 // Cobertura
			"TELEFONO_TERCERO":            sinister.Solicitante.Celular,      // Solicitante
			"TIPO_CONTACTO_CLIENTE":       "",                                // vacio
			"TIPO_DOCUMENTO":              docType,                           // Solicitante
			"TIPO_DOCUMENTO_REC":          "",                                // vacio
			"TIPO_PERSONA":                "",                                // vacio
			"TITULO":                      "",                                // vacio
			"USUARIO_DICTAMINIO":          "",                                // vacio,
			"FECHA_OCURRENCIA":            occurrenceDate.Format("02-01-2006"),
			"MONTO_SOLICITADO":            sinister.MontoSolicitado,
			"NOMBRE_PAGO_GASTOS":          sinister.Pagador,
			"DECLARACION_JURADA":          declaration,
			"FECHA_INICIO_INCAPACIDAD":    "",
			"FECHA_FIN_INCAPACIDAD":       "",
		}

		beneficiaries := make([]map[string]string, len(sinister.Beneficiarios))
		dta := "DNI"
		if len(sinister.Asegurado.NroDocumento) != 8 {
			dta = "C/E"
		}
		insured := map[string]string{
			"APELLIDO_PATERNO": sinister.Asegurado.ApellidoPaterno,
			"APELLIDO_MATERNO": sinister.Asegurado.ApellidoMaterno,
			"PRIMER_NOMBRE":    sinister.Asegurado.Nombres,
			"SEGUNDO_NOMBRE":   "",
			"TIPO_DOCUMENTO":   dta,
			"NUMERO_DOCUMENTO": sinister.Asegurado.NroDocumento,
		}
		eventData["asegurado"] = insured

		for i, beneficiario := range sinister.Beneficiarios {
			dt := "DNI"
			if len(beneficiario.NroDocumento) != 8 {
				dt = "C/E"
			}
			beneficiaries[i] = map[string]string{
				"BENEFICIARIO_SOLICITUD":     "",                        // vacio
				"APELLIDO_PATERNO":           "",                        // vacio
				"APELLIDO_MATERNO":           "",                        // vacio
				"TIPO_DOCUMENTO":             dt,                        // vacio
				"NUMERO_DOCUMENTO":           beneficiario.NroDocumento, // vacio
				"NOMBRES":                    "",                        // vacio
				"NOMBRE_TITULAR":             beneficiario.Titular,      // opcional
				"MONTO_PAGADO":               "",                        // opcional
				"METODO_PAGO":                beneficiario.MetodoPago,   // opcional
				"INTERESES_MORATORIOS":       "",                        // opcional
				"NUMERO_CUENTA":              beneficiario.NroCuenta,    // opcional
				"MONEDA_CUENTA":              beneficiario.Moneda,       // opcional
				"BANCO":                      beneficiario.Banco,        // opcional
				"NUMERO_INCIDENTE_ULTIMUS":   "",                        // opcional
				"TIPO_CUENTA":                beneficiario.TipoCuenta,   // opcional
				"NUMERO_CORRELATIVO_ULTIMUS": "",                        // opcional
				"ESTADO":                     "ACTIVO",                  // ACTIVO
			}
		}

		requestApi := map[string]interface{}{
			"siniestro":     requestSinister,
			"asegurado":     insured,
			"beneficiarios": beneficiaries,
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
		eventData["casos"] = append(eventData["casos"].([]map[string]interface{}), map[string]interface{}{"caso": caseNumbers[x], "cobertura": coverage.Coverage})

		fmt.Println("Caso generado: ", responseSinister.Results.CaseID)

		for k, document := range append(coverage.Documents, coverage.Additional...) {
			document.Filename = fmt.Sprint(caseID, "_", document.Filename)
			form := map[string]string{
				"app":         "siniestros",
				"contact_id":  contactID,
				"numero_caso": ticketNumber,
				"folder":      contactID,
			}

			file, fileType, size := &bytes.Buffer{}, "", int64(0) // Descargar archivo y obtener FileType
			fmt.Println("[", k, "] fileType: ", fileType)
			startTimeEndpoint = time.Now()
			eventData["Etapa"] = "StorageServiceCall_" + strconv.Itoa(k)
			// Sinister Add Document
			responseUpload, err := service.storage.Upload(txContext, mappers.MapStringToMapRequest(form), helpers_ai.File{
				Key:         "file",
				Name:        document.Filename,
				File:        file,
				ContentType: fileType,
				Size:        size,
			})
			log.Println(responseUpload)
			txContext.Event.AddTimeEndpoint("StorageServiceCall_"+strconv.Itoa(k), startTimeEndpoint)
			track["StorageServiceCall_"+strconv.Itoa(k)] = txContext.LastStageData.Pop()
			if err != nil {
				// apihelpers.RenderError(ctx, errors.New("FILE_ERROR_REQUEST_API"), fasthttp.StatusBadRequest)
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
				"filename":    document.Filename,
			}
			startTimeEndpoint = time.Now()
			eventData["Etapa"] = "SinisterAddDocumentCall_" + strconv.Itoa(k)
			_, _, err = service.sinister.AddDocument(txContext, sinisterDoc)
			track["SinisterAddDocumentCall_"+strconv.Itoa(k)] = txContext.LastStageData.Pop()
			txContext.Event.AddTimeEndpoint("SinisterAddDocumentCall_"+strconv.Itoa(k), startTimeEndpoint)
		}
	}
	return caseNumbers, nil
}

// AnalyzeWebIA ejecuta el análisis de documentos de forma asíncrona
func (service *sinisterAIServiceDomain) AnalyzeWebIA(txContext *models.TxContext, request *resources.SinisterCase, callback func(resources.MapResponse, error)) {
	log := logger.GetLoggerHooks(txContext, "sinisterAIServiceDomain", "AnalyzeWebIA")
	defer logger.End(log)

	// Ejecutar análisis de forma asíncrona
	go func() {
		// Crear un nuevo contexto para la goroutine
		asyncTxContext := &models.TxContext{
			Event:       txContext.Event,
			Origin:      txContext.Origin,
			Environment: txContext.Environment,
		}

		// Recuperación de panics en la goroutine
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("🚨 Panic recuperado en AnalyzeWebIA: %v\n", r)
				if callback != nil {
					callback(nil, fmt.Errorf("error interno en análisis asíncrono: %v", r))
				}
			}
		}()

		fmt.Printf("🔄 Iniciando análisis asíncrono para documento: %s (Caso: %s)\n",
			request.Documento.Name, request.Caso)

		// Llamar al método AnalyzeDocument
		result, err := service.AnalyzeDocumentWeb(asyncTxContext, request, request.Documento.FileURL)

		if err != nil {
			fmt.Printf("❌ Error en análisis asíncrono: %v\n", err)
		} else {
			fmt.Printf("✅ Análisis asíncrono completado exitosamente para caso: %s\n", request.Caso)
		}

		// Ejecutar callback si se proporcionó
		if callback != nil {
			callback(result, err)
		}
	}()

	fmt.Printf("🚀 Análisis asíncrono iniciado para caso: %s\n", request.Caso)
}

// AnalyzeWebIASimple ejecuta el análisis de documentos de forma asíncrona sin callback
// Útil cuando no necesitas procesar el resultado inmediatamente
func (service *sinisterAIServiceDomain) AnalyzeWebIASimple(txContext *models.TxContext, request *resources.SinisterCase) {
	service.AnalyzeWebIA(txContext, request, nil)
}

func (service *sinisterAIServiceDomain) AnalyzeDocument(txContext *models.TxContext, request *resources.SinisterCase) (resources.MapResponse, error) {
	log := logger.GetLoggerHooks(txContext, "sinisterAIServiceDomain", "AnalyzeDocument")
	defer logger.End(log)

	resp, err := http.Get(request.Documento.FileURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Leer los bytes del cuerpo de la respuesta
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Convertir a Base64 el archivo
	filebase64Data := base64.StdEncoding.EncodeToString(data)

	return service.handleDocumentRequest(txContext, filebase64Data, request)
}

// AnalyzeDocumentWeb analiza documentos que ya vienen en base64 desde el request
func (service *sinisterAIServiceDomain) AnalyzeDocumentWeb(txContext *models.TxContext, request *resources.SinisterCase, filebase64Data string) (resources.MapResponse, error) {
	log := logger.GetLoggerHooks(txContext, "sinisterAIServiceDomain", "AnalyzeDocumentWeb")
	defer logger.End(log)

	fmt.Printf("📄 Analizando documento web: %s (Base64 length: %d)\n", request.Documento.Name, len(filebase64Data))

	// Directamente usar los datos base64 que vienen en el request
	return service.handleDocumentRequest(txContext, filebase64Data, request)
}

func (service *sinisterAIServiceDomain) handleDocumentRequest(txContext *models.TxContext, filebase64Data string, request *resources.SinisterCase) (resources.MapResponse, error) {
	fmt.Printf("🔍 handleDocumentRequest - Documento: %s, Filename: %s\n", request.Documento.Name, request.Documento.Filename)

	fileMetadata := map[string]string{
		"data":      filebase64Data,
		"file_name": request.Documento.Filename,
		"extension": request.Documento.FileURL[strings.LastIndex(request.Documento.FileURL, ".")+1:],
	}

	fmt.Printf("🔍 handleDocumentRequest - Extension detectada: %s\n", fileMetadata["extension"])

	// avanzar de acuerdo al documento que se va a analizar
	switch request.Documento.Name {
	case "Denuncia Policial":
		fileJSON, err := service.parseFileGenAI(service.conf.Conf.Sinister.IdPromtDP, fileMetadata)
		if err != nil {
			fmt.Printf("❌ Error en parseFileGenAI para Denuncia Policial: %v\n", err)
			return nil, err
		}

		if fileJSON == nil {
			return resources.MapResponse{"document": "empty"}, nil
		}

		result, err := service.handleDenunciaPolicial(txContext, fileJSON, request)
		return result, err

	case "Descanso medico", "Certificado médico":
		fileJSON, err := service.parseFileGenAI(service.conf.Conf.Sinister.IdPromtDM, fileMetadata)
		if err != nil {
			fmt.Printf("❌ Error en parseFileGenAI para %s: %v\n", request.Documento.Name, err)
			return nil, err
		}

		if fileJSON == nil {
			return resources.MapResponse{"document": "empty"}, nil
		}

		result, err := service.handleDescansoMedico(txContext, fileJSON, request)
		return result, err
	default:
		return nil, errors.New("unsupported document type")
	}
}

func (service *sinisterAIServiceDomain) handleDenunciaPolicial(txContext *models.TxContext, fileJSON []byte, request *resources.SinisterCase) (resources.MapResponse, error) {
	var (
		dpData                  models.DataDenunciaPolicialResponseGenIA
		lesionado               models.Lesionado
		soatIS                  *models.Soat
		listPlates, listInjured resources.ListArray
	)

	// Init Observations
	var dpObservation = &models.ObservationWrapper{
		Data: models.ObservationDataAnswer{
			Answer: models.ObservationAnswer{
				Dictamen: models.ObservationDictamen{
					Justificacion: "",
					Status:        "",
				},
				Observations: []models.ObservationData{},
			},
		},
	}

	// Parse DP Data
	if err := json.Unmarshal(fileJSON, &dpData); err != nil {
		fmt.Println("Error deserializando JSON:", err)
		return nil, err
	}

	/*
		Validación de denuncia policial:
		(1) Antigüedad desde fecha de ocurrencia menor a 2 años
	*/
	fechaOcurrencia, err := time.Parse("2006-01-02 15:04:05.000Z", dpData.Answer.FechaOcurrencia)
	if err == nil {
		dpData.Answer.DiasTranscurridos = int32(time.Since(fechaOcurrencia).Hours() / 24)
	}
	if dpData.Answer.DiasTranscurridos >= 721 {
		dpObservation.Data.Answer.Observations = append(dpObservation.Data.Answer.Observations, models.ObservationData{
			Classification: "observacion",
			Text:           "Siniestro tiene 2 o mas años de antiguedad en la denuncia policial",
		})
	}

	// Recorrer vehiculos y sacar datos del mismo y del lesionado
	for _, v := range dpData.Answer.Vehiculos {
		listPlates = append(listPlates, v.Placa)
		if soatIS == nil {
			soatIS, _ = service.soat.FindByPlateHistory(txContext, v.Placa, dpData.Answer.FechaOcurrencia)
		}
		for _, l := range v.Lesionados {
			for _, d := range l.ListaDocumentos {
				listInjured = append(listInjured, d.Numero)
				if d.Numero == request.NumeroDocumento {
					lesionado = l
				}
			}
		}
	}

	/*
		Validación de denuncia policial:
		(2) Documento de lesionado reportado en siniestro presente
	*/
	if len(lesionado.Nombres) == 0 {
		dpObservation.Data.Answer.Observations = append(dpObservation.Data.Answer.Observations, models.ObservationData{
			Classification: "observacion",
			Text:           "Lesionado no se encuentra reportado en siniestro en la denuncia policial",
		})
	}

	/*
		Validación de denuncia policial:
		(3) Alguna placa del siniestro cuenta con SOAT vigente
	*/

	if soatIS == nil {
		dpObservation.Data.Answer.Observations = append(dpObservation.Data.Answer.Observations, models.ObservationData{
			Classification: "observacion",
			Text:           "Ninguna placa del siniestro cuenta con SOAT vigente en la denuncia policial",
		})
	}

	fmt.Println("soatIS", soatIS)

	dpData.ObserveDocuments = &models.ObserveDocuments{}
	injuredCompleteName := strings.ToUpper(fmt.Sprint(lesionado.Nombres, " ", lesionado.ApellidoPaterno, " ", lesionado.ApellidoMaterno))

	// Buscar en qualitat los casos registrados por numero de documento
	qualitatResponse, _ := service.qualitat.QualitatStart(request.NumeroDocumento, injuredCompleteName)
	dataQualitat := models.DataQualitat{}
	if qualitatResponse != nil {
		for _, itemQualitad := range *qualitatResponse {
			// Parsear ambas fechas al tipo time.Time
			dateQualitat, _ := time.Parse("02/01/2006", itemQualitad.FechaSiniestro)
			dateDP, _ := time.Parse("2006-01-02", dpData.Answer.FechaOcurrencia[:10])
			dateDPFormated := dateDP.Format("02/01/2006")
			dateQualitatFormated := dateQualitat.Format("02/01/2006")
			if dateQualitatFormated == dateDPFormated {
				dataQualitat = models.DataQualitat{
					EstadoSiniestro:        strings.ToLower(itemQualitad.EstadoSiniestro),
					CentroMedico:           strings.ToLower(itemQualitad.CentroMedico),
					NumeroSiniestro:        strings.ToLower(itemQualitad.NroSiniestro),
					NumeroSiniestroCliente: strings.ToLower(itemQualitad.NroSiniestroCliente),
					NumeroPoliza:           strings.ToLower(itemQualitad.NroPoliza),
					TipoSiniestro:          strings.ToLower(itemQualitad.TipoSiniestro),
				}
			}
		}
	}

	// validar si es menor de edad y registrar como observado el siniestro
	if lesionado.Edad < 18 {
		dpData.ObserveDocuments.DocumentMenor = &models.ObserveDocument{
			Nombre:     "Documento del menor",
			Requerido:  true,
			Completado: false,
		}
		dpData.ObserveDocuments.DocumentPadre = &models.ObserveDocument{
			Nombre:     "Documento del padre",
			Requerido:  true,
			Completado: false,
		}
		dpData.ObserveDocuments.DocumentMadre = &models.ObserveDocument{
			Nombre:     "Documento de la madre",
			Requerido:  true,
			Completado: false,
		}
	}

	/*
		Validación de denuncia policial:
		(4) Interviniente presente
		(5) FDO presente
		(6) FIRMAS y SELLOS presente
	*/
	if !dpData.Answer.Validacion.FirmasSellos || (len(dpData.Answer.Validacion.FDO) <= 0 && len(dpData.Answer.Validacion.Interviniente) <= 0) {
		dpData.ObserveDocuments.DP = &models.ObserveDocument{
			Nombre:     "Denuncia policial",
			Requerido:  true,
			Completado: false,
		}

		if !dpData.Answer.Validacion.FirmasSellos {
			dpObservation.Data.Answer.Observations = append(dpObservation.Data.Answer.Observations, models.ObservationData{
				Classification: "observacion",
				Text:           "Denuncia policial sin firma o sello",
			})
		}

		if len(dpData.Answer.Validacion.FDO) <= 0 && len(dpData.Answer.Validacion.Interviniente) <= 0 {
			dpObservation.Data.Answer.Observations = append(dpObservation.Data.Answer.Observations, models.ObservationData{
				Classification: "observacion",
				Text:           "Denuncia policial sin FDO o Interviniente",
			})
		}
	}

	dpData.Caso, err = service.SaveSiniesterPolice(request.NumeroDocumento, soatIS, lesionado, request.Documento, dpData, txContext, dataQualitat, dpObservation)
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(dpData)
	var result map[string]interface{}
	json.Unmarshal(jsonData, &result)
	return resources.MapResponse(result), nil
}

func (service *sinisterAIServiceDomain) handleDescansoMedico(txContext *models.TxContext, fileJSON []byte, request *resources.SinisterCase) (resources.MapResponse, error) {
	var dmData models.DataDescansoMedicoResponseGenIA

	// Init Sinister Case
	objectIDCase, err := primitive.ObjectIDFromHex(request.Caso)
	if err != nil {
		return nil, err
	}

	sinisterCase, err := service.sinisterCase.FindByCaseNumber(txContext, objectIDCase)
	if err != nil {
		return nil, err
	}
	if sinisterCase == nil {
		return nil, fmt.Errorf("caso no encontrado")
	}

	// Parse DM Data
	if err := json.Unmarshal(fileJSON, &dmData); err != nil {
		return nil, err
	}

	// Init Lesionado
	var lesionado models.LesionadoDM

	// Si NumeroDocumento está vacío, tomar el primer lesionado disponible
	if strings.TrimSpace(request.NumeroDocumento) == "" && len(dmData.Answer.Reports) > 0 {
		lesionado = dmData.Answer.Reports[0].Lesionado
	} else {
		// Buscar lesionado por NumeroDocumento
		for _, report := range dmData.Answer.Reports {
			if report.Lesionado.NumeroDocumento == request.NumeroDocumento {
				lesionado = report.Lesionado
				break
			}
		}
	}

	/*
		Validación de descanso medico:
		(1) Documento de lesionado reportado en siniestro presente
	*/
	if len(lesionado.Nombres) == 0 {
		// Verificar que sinisterCase.Observation no sea nil
		if sinisterCase.Observation == nil {
			return resources.MapResponse{"error": "sinisterCase.Observation is nil"}, nil
		}
		if sinisterCase.Observation.Data.Answer.Observations == nil {
			sinisterCase.Observation.Data.Answer.Observations = make([]models.ObservationData, 0)
		}

		sinisterCase.Observation.Data.Answer.Observations = append(sinisterCase.Observation.Data.Answer.Observations, models.ObservationData{
			Classification: "observacion",
			Text:           "Lesionado no se encuentra en descanso médico",
		})
	}

	/*
		Validación de descanso medico:
		(2) Firma presente
		(3) Sello presente
	*/
	// Asegurar que las estructuras de observación estén inicializadas
	if sinisterCase.Observation == nil {
		return resources.MapResponse{"error": "sinisterCase.Observation is nil"}, nil
	}

	// Como Data y Answer son structs (no punteros), solo validamos que Observations no sea nil
	if sinisterCase.Observation.Data.Answer.Observations == nil {
		sinisterCase.Observation.Data.Answer.Observations = make([]models.ObservationData, 0)
	}

	dmData.ObserveDocuments = &models.ObserveDocuments{}
	for _, v := range dmData.Answer.Reports {
		// Validar que v.Medico no esté vacío antes de acceder a sus campos
		if v.Medico == (models.MedicoDM{}) {
			continue
		}

		if !v.Medico.Firma || !v.Medico.Sello {
			dmData.ObserveDocuments.DM = &models.ObserveDocument{
				Nombre:     "Descanso medico",
				Requerido:  true,
				Completado: false,
			}

			if !v.Medico.Firma {
				sinisterCase.Observation.Data.Answer.Observations = append(sinisterCase.Observation.Data.Answer.Observations, models.ObservationData{
					Classification: "observacion",
					Text:           "Descanso medico sin firma",
				})
			}

			if !v.Medico.Sello {
				sinisterCase.Observation.Data.Answer.Observations = append(sinisterCase.Observation.Data.Answer.Observations, models.ObservationData{
					Classification: "observacion",
					Text:           "Descanso medico sin sello",
				})
			}
		}
	}

	var dmDataObservation = &models.ObservationWrapper{
		Data: models.ObservationDataAnswer{
			Answer: models.ObservationAnswer{
				Dictamen: models.ObservationDictamen{
					Justificacion: "",
					Status:        "",
				},
				Observations: []models.ObservationData{},
			},
		},
	}

	// Consultar siniestros históricos
	_, historicalSinisters, err := service.domainPortalSinister.FindByCaseHistory(txContext, request.NumeroDocumento)
	if err != nil {
		fmt.Println("Error al consultar siniestros históricos", err)
	}

	/*
		Validación de descanso medico:
		(4) Fecha de inicio incapacidad reportada es diferente a fechas de inicio de incapacidad históricas
		(5) Fecha de fin incapacidad reportada es diferente a fechas de fin de incapacidad históricas
	*/
	var foundDuplicate = false
	txContext.Event["historical_sinisters"] = historicalSinisters
	for _, v := range historicalSinisters {
		// if strings.Contains(strings.ToLower(v.EstadoDictamen), "aprobado") || strings.Contains(strings.ToLower(v.EstadoDictamen), "pendiente") {
		if strings.Contains(strings.ToLower(v.EstadoDictamen), "aprobado") {
			// Validar que FechasIncapacidad tenga al menos un elemento
			if len(v.FechasIncapacidad) == 0 {
				fmt.Printf("⚠️ handleDescansoMedico - Caso histórico %s no tiene fechas de incapacidad, saltando\n", v.NumeroCaso)
				continue
			}

			for _, vReportDM := range dmData.Answer.Reports {
				formattedDateInit := formatToYYYYMMDD(v.FechasIncapacidad[0].FechaInicioIncapacidad)
				formattedDateEnd := formatToYYYYMMDD(v.FechasIncapacidad[0].FechaFinIncapacidad)
				if formattedDateInit == vReportDM.FechaInicioIncapacidad {
					fmt.Printf("⚠️ handleDescansoMedico - Agregando observación: fecha inicio duplicada\n")
					if sinisterCase.Observation != nil && sinisterCase.Observation.Data.Answer.Observations != nil {
						sinisterCase.Observation.Data.Answer.Observations = append(sinisterCase.Observation.Data.Answer.Observations, models.ObservationData{
							Classification: "observacion",
							Text:           "Se encontró caso " + v.NumeroCaso + " con fecha de inicio de incapacidad igual a la fecha de inicio de incapacidad reportada",
						})
					}
					foundDuplicate = true
					break
				}
				if formattedDateEnd == vReportDM.FechaFinIncapacidad {
					fmt.Printf("⚠️ handleDescansoMedico - Agregando observación: fecha fin duplicada\n")
					if sinisterCase.Observation != nil && sinisterCase.Observation.Data.Answer.Observations != nil {
						sinisterCase.Observation.Data.Answer.Observations = append(sinisterCase.Observation.Data.Answer.Observations, models.ObservationData{
							Classification: "observacion",
							Text:           "Se encontró caso " + v.NumeroCaso + " con fecha de fin de incapacidad igual a la fecha de fin de incapacidad reportada",
						})
					}
					foundDuplicate = true
					break
				}
			}
		}
		if foundDuplicate {
			break
		}
	}

	// Obtener relación de diagnóstico final
	var diagnosticoFinal interface{}
	if len(dmData.Answer.Reports) > 0 && dmData.Answer.Reports[0].Lesionado != (models.LesionadoDM{}) {
		diagnosticoFinal, _ = service.apiDiagnostic(dmData.Answer.Reports[0].Lesionado.DiagnosticoMedico)
	} else {
		fmt.Printf("⚠️ handleDescansoMedico - No hay reports disponibles o lesionado vacío para obtener diagnóstico\n")
		diagnosticoFinal = "No disponible"
	}

	var stadoQualitatSiniestro = "No se encotro el estado de qualitat"
	// Validamos si DataQualitat no está vacío y EstadoSiniestro no es una cadena vacía
	if sinisterCase.DataQualitat != (models.DataQualitat{}) && sinisterCase.DataQualitat.EstadoSiniestro != "" && sinisterCase.DataQualitat.EstadoSiniestro != "null" {
		stadoQualitatSiniestro = sinisterCase.DataQualitat.EstadoSiniestro
	}

	// Validar campos antes de crear el payload
	var denunciaData interface{} = "No disponible"
	if sinisterCase.Denuncia.Caso != "" || sinisterCase.Denuncia.AnswerID != "" {
		denunciaData = sinisterCase.Denuncia
	}

	var lesionadoData interface{} = "No disponible"
	if sinisterCase.Lesionado.Nombres != "" || sinisterCase.Lesionado.ApellidoPaterno != "" {
		lesionadoData = sinisterCase.Lesionado
	}

	payloadDictamen := map[string]interface{}{
		"prompt_id":     service.conf.Conf.Sinister.IdPromtDictamen,
		"response_type": "application/json",
		"parameters": map[string]interface{}{
			"denuncia":         stringify(denunciaData),
			"certificado":      stringify(dmData),
			"injured":          stringify(lesionadoData),
			"qualitat_status":  stringify(stadoQualitatSiniestro),
			"diagnostic_final": stringify(diagnosticoFinal),
			"historic_cases":   stringify(historicalSinisters),
		},
	}

	responseObservation, _ := apiGenAI(service.conf.Conf.Sinister.UrlGenAI, service.conf.Conf.Sinister.TokenGenAI, payloadDictamen)

	// Volver a codificar el map en JSON
	jsonBytesObservation, err := json.Marshal(responseObservation)
	if err != nil {
		log.Fatalf("Error al codificar map a JSON: %v", err)
	}

	err = json.Unmarshal(jsonBytesObservation, &dmDataObservation)
	if err != nil {
		log.Fatalf("Error al decodificar JSON en struct: %v", err)
	}

	// DICTAMEN FINAL
	jsonData, _ := json.Marshal(dmData)
	var result map[string]interface{}
	json.Unmarshal(jsonData, &result)

	fmt.Printf("🔍 handleDescansoMedico - Verificando observaciones finales...\n")

	// Verificar que las estructuras no sean nil antes del acceso final
	if sinisterCase.Observation != nil &&
		sinisterCase.Observation.Data.Answer.Observations != nil &&
		len(sinisterCase.Observation.Data.Answer.Observations) > 0 {

		fmt.Printf("✅ handleDescansoMedico - Encontradas %d observaciones, procesando dictamen...\n",
			len(sinisterCase.Observation.Data.Answer.Observations))

		// Dictamen con Observaciones por checklist + IA
		sinisterCase.Observation.Data.Answer.Dictamen.Justificacion = "Observaciones encontradas"
		sinisterCase.Observation.Data.Answer.Dictamen.Status = "OBSERVADO"

		for _, v := range dmDataObservation.Data.Answer.Observations {
			sinisterCase.Observation.Data.Answer.Observations = append(sinisterCase.Observation.Data.Answer.Observations, models.ObservationData{
				Classification: v.Classification,
				Text:           v.Text,
			})
		}

		sinister := map[string]interface{}{
			"certificado":           dmData,
			"certificado_documento": request.Documento,
			"observaciones":         sinisterCase.Observation,
		}

		fmt.Printf("🔍 handleDescansoMedico - Actualizando caso con observaciones del checklist...\n")
		service.sinisterCase.UpdateOne(txContext, objectIDCase, sinister)

		fmt.Printf("✅ handleDescansoMedico - Proceso completado exitosamente con observaciones\n")
		return resources.MapResponse(result), nil
	} else {
		fmt.Printf("🔍 handleDescansoMedico - Sin observaciones del checklist, usando solo IA...\n")

		// Dictamen con Observaciones solo por IA
		sinister := map[string]interface{}{
			"certificado":           dmData,
			"certificado_documento": request.Documento,
			"observaciones":         dmDataObservation,
		}

		fmt.Printf("🔍 handleDescansoMedico - Actualizando caso solo con observaciones IA...\n")
		service.sinisterCase.UpdateOne(txContext, objectIDCase, sinister)

		fmt.Printf("✅ handleDescansoMedico - Proceso completado exitosamente sin observaciones del checklist\n")
		return resources.MapResponse(result), nil
	}
}

func (service *sinisterAIServiceDomain) parseFileGenAI(promptID string, file map[string]string) ([]byte, error) {
	payload := map[string]interface{}{
		"prompt_id":     promptID,
		"response_type": "application/json",
		"parameters":    map[string]interface{}{},
		"file":          file,
	}
	responseGenAI, err := apiGenAI(service.conf.Conf.Sinister.UrlGenAI, service.conf.Conf.Sinister.TokenGenAI, payload)
	if err != nil {
		return nil, err
	}

	if responseGenAI["data"] == nil {
		return nil, err
	}

	if responseGenAI["status"] != "OK" {
		return nil, err
	}

	data := responseGenAI["data"].(map[string]interface{})
	fileJSON, _ := json.Marshal(data)
	return fileJSON, err
}

func (service *sinisterAIServiceDomain) UploadObservedDocument(txContext *models.TxContext, request *resources.SinisterCase) error {
	log := logger.GetLoggerHooks(txContext, "sinisterAIServiceDomain", "UploadObservedDocument")
	defer logger.End(log)

	objectIDCase, err := primitive.ObjectIDFromHex(request.Caso)
	if err != nil {
		return err
	}

	// Buscar el caso
	sinisterCase, err := service.sinisterCase.FindByCaseNumber(txContext, objectIDCase)
	if err != nil {
		return err
	}

	// Validar que existan los ObserveDocuments
	if sinisterCase.ObserveDocuments == nil {
		return errors.New("ObserveDocuments no inicializado")
	}

	// Actualizar el path del documento correspondiente
	switch request.Documento.Name {
	case "DocumentMenor":
		sinisterCase.ObserveDocuments.DocumentMenor.Path = request.Documento.FileURL
		sinisterCase.ObserveDocuments.DocumentMenor.Completado = true
	case "DocumentPadre":
		sinisterCase.ObserveDocuments.DocumentPadre.Path = request.Documento.FileURL
		sinisterCase.ObserveDocuments.DocumentPadre.Completado = true
	case "DocumentMadre":
		sinisterCase.ObserveDocuments.DocumentMadre.Path = request.Documento.FileURL
		sinisterCase.ObserveDocuments.DocumentMadre.Completado = true
	}

	// Actualizar el registro en la base de datos
	update := map[string]interface{}{
		"denuncia.observe_documents": sinisterCase.ObserveDocuments,
	}
	_, err = service.sinisterCase.UpdateOne(txContext, objectIDCase, update)
	return err
}

func (service *sinisterAIServiceDomain) GetByCase(txContext *models.TxContext, caseID string) (*models.CasoSiniestro, error) {
	log := logger.GetLoggerHooks(txContext, "sinisterAIServiceDomain", "GetByCase")
	defer logger.End(log)

	// Convertir el string a ObjectID
	objectIDCase, err := primitive.ObjectIDFromHex(caseID)
	if err != nil {
		return nil, err
	}

	sinister, err := service.sinisterCase.FindByCaseNumber(txContext, objectIDCase)
	return sinister, err
}

func (service *sinisterAIServiceDomain) GetAll(txContext *models.TxContext) (resources.ListArray, error) {
	log := logger.GetLoggerHooks(txContext, "sinisterAIServiceDomain", "GetAll")
	defer logger.End(log)

	cases, err := service.sinisterCase.FindAll(txContext)
	return cases, err
}

func (service *sinisterAIServiceDomain) GetPreviousCases(txContext *models.TxContext, doc string) (resources.MapResponse, error) {
	log := logger.GetLoggerHooks(txContext, "sinisterAIServiceDomain", "GetPreviousCases")
	defer logger.End(log)

	// Consultar siniestros históricos
	statusCode, historicalSinisters, err := service.domainPortalSinister.FindByCaseHistory(txContext, doc)
	if err != nil {
		return nil, errors.New("BAD_REQUEST")
	}

	fmt.Println("Estado http: ", statusCode)

	txContext.Event["historical_sinisters"] = historicalSinisters
	// Convertir []models a ListArray
	response := make(resources.ListArray, 0, len(historicalSinisters))
	for _, s := range historicalSinisters {
		response = append(response, s)
	}

	// Convierte el struct a JSON y luego a map
	jsonBytes, _ := json.Marshal(response[0])
	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)
	// Solo devolver el primer objeto si existe
	if len(response) > 0 {
		return resources.MapResponse{"case": result}, nil
	}
	return resources.MapResponse{"message": "success"}, nil
}

/*
	Guardar Casos
*/

// Utilidad para convertir JSON a string
func (service *sinisterAIServiceDomain) SaveSiniesterPolice(
	numeroDocumento string,
	soatIS *models.Soat,
	lesionado models.Lesionado,
	documento resources.Document,
	dataResponseGenIA models.DataDenunciaPolicialResponseGenIA,
	txContext *models.TxContext,
	qualitat models.DataQualitat,
	dataObservation *models.ObservationWrapper,
) (string, error) {

	sinister := map[string]interface{}{
		"denuncia":           dataResponseGenIA,
		"denuncia_documento": documento,
		"lesionado":          lesionado,
		"numero_documento":   numeroDocumento,
		"qualitat":           qualitat,
		"observaciones":      dataObservation,
	}

	// Manejar el caso cuando soatIS es nil
	if soatIS != nil {
		sinister["placa"] = soatIS.Placa
		sinister["soat"] = map[string]interface{}{
			"numero_poliza": soatIS.NumeroPoliza,
			"fecha_inicio":  soatIS.FechaInicio,
			"fecha_fin":     soatIS.FechaFin,
			"estado":        soatIS.Estado,
			"coberturas":    soatIS.Coberturas,
		}
	} else {
		sinister["placa"] = ""
		sinister["soat"] = map[string]interface{}{
			"numero_poliza": "",
			"fecha_inicio":  time.Time{},
			"fecha_fin":     time.Time{},
			"estado":        "",
			"coberturas":    nil,
		}
	}

	caso, err := service.sinisterCase.Save(txContext, sinister)
	return caso, err
}

/*
	Invocaciones de APIs
*/

func apiGenAI(url string, token string, body map[string]interface{}) (map[string]interface{}, error) {
	method := "POST"
	bytesPayload, _ := json.Marshal(body)
	payload := strings.NewReader(string(bytesPayload))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal(result, &response)
	return response, err
}

func (service *sinisterAIServiceDomain) apiDiagnostic(searchInput string) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"user":         "user",
		"base_id":      "diagnostic-table-1536",
		"search_input": searchInput,
		"params": []string{
			"diagnostico",
			"leve",
			"moderado",
			"severo",
		},
		"offset": 0,
		"limit":  1,
	}

	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", service.conf.Conf.Sinister.ApiDiagnostic+"/v1/llm/search", strings.NewReader(string(bytesPayload)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+service.conf.Conf.Sinister.TokenGenAI)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal(result, &response)
	return response, err
}

/*
	Utilidades
*/

// Parsear y formatear la fecha
func formatToYYYYMMDD(dateStr string) string {
	if dateStr == "" {
		return ""
	}
	parsed, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return dateStr // Si falla, retorna el original
	}
	return parsed.Format("2006-01-02")
}

// Utilidad para convertir JSON a string
func stringify(data interface{}) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Error al convertir JSON:", err)
		return ""
	}
	return string(jsonData)
}
