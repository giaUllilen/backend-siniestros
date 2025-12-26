package endpoints

import (
	"encoding/json"
	"log"

	"is-public-api/application/models"
	"is-public-api/application/resources"
	restful "is-public-api/helpers/apihelpers"

	"github.com/francoispqt/gojay"
)

type sinisterEndpoints struct {
	conf *restful.Sinister
}

func NewSinisterEndpoints(conf *restful.EndpointConfig) *sinisterEndpoints {
	return &sinisterEndpoints{conf: &restful.Sinister{
		Path:                conf.Sinister.Path,
		Get:                 conf.Sinister.Get,
		Save:                conf.Sinister.Save,
		Document:            conf.Sinister.Document,
		History:             conf.Sinister.History,
		UpdateObservationIA: conf.Sinister.UpdateObservationIA,
	}}
}

func (service *sinisterEndpoints) Save(txContext *models.TxContext, body *resources.MapRequest) (int, *resources.ResponseSinister, error) {
	url := service.conf.Path + service.conf.Save
	bytes, err := json.Marshal(body)

	if err != nil {
		log.Println("[SinisterSave] error gojay.MarshalJSONObject: ", err)
		bytes, _ = json.Marshal(body)
	}
	// Inicializar LastStageData si es nil
	if txContext.LastStageData == nil {
		txContext.LastStageData = make(models.Event)
	}

	txContext.LastStageData["request"] = *body
	response := &resources.ResponseSinister{}
	// "http://apisiniestro:2244/api/v1/siniestros/saveSiniestro"
	statusCode, err := restful.HttpRequestToStruct(url, "POST", bytes, response)
	txContext.LastStageData["statusCode"] = statusCode
	txContext.LastStageData["response"] = response
	if err != nil {
		txContext.LastStageData["error"] = err
		txContext.LastStageData["errorType"] = "RequestServerError"
		log.Println("[SinisterSave] error request: ", err)
		return statusCode, nil, err
	}
	return statusCode, response, err
}

func (service *sinisterEndpoints) AddDocument(txContext *models.TxContext, body *resources.MapRequest) (int, []byte, error) {
	url := service.conf.Path + service.conf.Document
	bytes, err := gojay.MarshalJSONObject(body)

	log.Println("[SinisterAddDocument]: Init process --> url: ", url, " request: ", string(bytes))
	if err != nil {
		log.Println("[SinisterAddDocument] error gojay.MarshalJSONObject: ", err)
		bytes, err = json.Marshal(body)
	}
	statusCode, bytes, err := restful.HttpRequest(url, "POST", bytes)

	// Inicializar LastStageData si es nil
	if txContext.LastStageData == nil {
		txContext.LastStageData = make(models.Event)
	}

	if bytes != nil {
		txContext.LastStageData["response"] = string(bytes)
	}
	txContext.LastStageData["statusCode"] = statusCode
	if err != nil {
		txContext.LastStageData["error"] = err
		txContext.LastStageData["errorType"] = "RequestServerError"
		log.Println("[SinisterSave] error request: ", err)
		return statusCode, nil, err
	}
	return statusCode, bytes, err
}

// FindByCaseNumber - Modificado para recibir string en lugar de ObjectID
// El API externo espera el número de caso como string (ej: "CIS_98746")
// Anterior:
// func (service *sinisterEndpoints) FindByCaseNumber(txContext *models.TxContext, caseNumber primitive.ObjectID) (int, []byte, error) {
//     request := resources.MapRequest{"numero_solicitud": caseNumber}
//     log.Println("[SinisterGet]: Init process --> url: ", url, " request: {\"numero_solicitud\": \""+caseNumber.Hex()+"\"}")
func (service *sinisterEndpoints) FindByCaseNumber(txContext *models.TxContext, caseNumber string) (int, []byte, error) {
	url := service.conf.Path + service.conf.Get
	request := resources.MapRequest{"numero_solicitud": caseNumber}
	bytes, err := gojay.MarshalJSONObject(request)
	log.Println("[SinisterGet]: Init process --> url: ", url, " request: {\"numero_solicitud\": \""+caseNumber+"\"}")
	if err != nil {
		log.Println("[SinisterGet] error gojay.MarshalJSONObject: ", err)
		bytes, err = json.Marshal(request)
	}
	statusCode, bytes, err := restful.HttpRequest(url, "POST", bytes)
	if err != nil {
		log.Println("[SinisterGet] error request: ", err)
		return statusCode, nil, err
	}
	return statusCode, bytes, nil
}

func (service *sinisterEndpoints) Delete(txContext *models.TxContext, caseNumber string) (int, []byte, error) {
	url := service.conf.Path + service.conf.Delete
	request := resources.MapRequest{"numero_caso": caseNumber}
	bytes, err := gojay.MarshalJSONObject(request)
	log.Println("[SinisterDelete]: Init process --> url: ", url, " request: {\"numero_caso\": \""+caseNumber+"\"}")
	if err != nil {
		log.Println("[SinisterDelete] error gojay.MarshalJSONObject: ", err)
		bytes, _ = json.Marshal(request)
	}
	statusCode, bytes, err := restful.HttpRequest(url, "DELETE", bytes)

	// Inicializar LastStageData si es nil
	if txContext.LastStageData == nil {
		txContext.LastStageData = make(models.Event)
	}

	if bytes != nil {
		txContext.LastStageData["response"] = string(bytes)
	}
	txContext.LastStageData["statusCode"] = statusCode
	if err != nil {
		txContext.LastStageData["error"] = err
		txContext.LastStageData["errorType"] = "RequestServerError"
		log.Println("[SinisterDelete] error request: ", err)
		return statusCode, nil, err
	}
	return statusCode, bytes, nil
}

func (service *sinisterEndpoints) FindByCaseHistory(txContext *models.TxContext, request models.SinisterHistoryRequest) (int, []models.SinisterHistory, error) {
	url := service.conf.Path + service.conf.History

	bytes, err := json.Marshal(request)
	if err != nil {
		log.Println("[ListSiniestros] error al serializar request:", err)
		return 0, nil, err
	}
	statusCode, respBytes, err := restful.HttpRequest(url, "POST", bytes)
	if err != nil {
		log.Println("[ListSiniestros] error en la petición:", err)
		return statusCode, nil, err
	}

	var response []models.SinisterHistory
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		log.Println("[ListSiniestros] error al deserializar respuesta:", err)
		return statusCode, nil, err
	}

	return statusCode, response, nil
}

func (service *sinisterEndpoints) UpdateObservationIA(txContext *models.TxContext, request models.ObservationIARequest) (int, []byte, error) {
	url := service.conf.Path + service.conf.UpdateObservationIA
	bytes, err := json.Marshal(request)

	if err != nil {
		log.Println("[UpdateObservationIA] error marshaling request: ", err)
		return 500, nil, err
	}

	// Inicializar LastStageData si es nil
	if txContext.LastStageData == nil {
		txContext.LastStageData = make(models.Event)
	}

	txContext.LastStageData["request"] = request
	statusCode, response, err := restful.HttpRequest(url, "POST", bytes)
	txContext.LastStageData["statusCode"] = statusCode
	txContext.LastStageData["response"] = response

	if err != nil {
		txContext.LastStageData["error"] = err
		txContext.LastStageData["errorType"] = "RequestServerError"
		log.Println("[UpdateObservationIA] error request: ", err)
	}

	return statusCode, response, err
}
