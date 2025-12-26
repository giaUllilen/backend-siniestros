package endpoints

import (
	"encoding/json"
	"log"

	"is-public-api/application/models"
	"is-public-api/application/resources"
	restful "is-public-api/helpers/apihelpers"
)

type eventsEndpoint struct {
	conf *restful.EventLog
}

func NewEventsEndpoint(conf *restful.EndpointConfig) *eventsEndpoint {
	return &eventsEndpoint{conf: &restful.EventLog{
		Uri:       conf.EventLog.Uri,
		ApiKey:    conf.EventLog.ApiKey,
		ApiKeyWsp: conf.EventLog.ApiKeyWsp,
	}}
}

func (service *eventsEndpoint) AddEvent(txContext *models.TxContext) (int, []byte, error) {
	url := service.conf.Uri
	body := resources.MapRequest(txContext.Event)
	bytes, err := json.Marshal(body)
	log.Println("[AddEvent]: Init process --> url: ", url)
	if err != nil {
		log.Println("[AddEvent] error gojay.MarshalJSONObject: ", err)
		bytes, _ = json.Marshal(body)
	}
	apiKey := service.conf.ApiKey
	if txContext.Origin == "WHATSAPP-API" {
		apiKey = service.conf.ApiKeyWsp
	}
	statusCode, bytes, err := restful.HttpRequest(url, "POST", bytes, map[string]string{"X-Api-Key": apiKey})
	log.Println("[AddEvent]: ", statusCode, " Response: ", string(bytes))
	return statusCode, bytes, err
}
