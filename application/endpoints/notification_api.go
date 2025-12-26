package endpoints

import (
	"encoding/json"
	"errors"
	"strings"

	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/helpers/apihelpers"
)

type notificationEndpoints struct {
	conf *apihelpers.Service
}

func NewNotificationEndpoints(conf *apihelpers.EndpointConfig) *notificationEndpoints {
	return &notificationEndpoints{conf: &apihelpers.Service{
		Path:     conf.Notification.Path,
		Endpoint: conf.Notification.Endpoint,
	}}
}

func (service *notificationEndpoints) SendMail(txContext *models.TxContext, request *resources.MapRequest) ([]byte, error) {
	url := service.conf.Path + service.conf.Endpoint
	bytes, err := json.Marshal(request)

	if err != nil {
		txContext.LastStageData["error"] = err
		txContext.LastStageData["errorType"] = "ENDPOINT_ERROR_PARSE_REQUEST"
		return nil, errors.New("ENDPOINT_ERROR_PARSE_REQUEST")
	}
	statusCode, responseBytes, err := apihelpers.HttpMultipartRequest(url, "POST", strings.NewReader(string(bytes)), "")
	txContext.LastStageData["statusCode"] = statusCode
	txContext.LastStageData["response"] = string(responseBytes)
	if err != nil {
		// apihelpers.RenderError(ctx, errors.New("FILE_ERROR_REQUEST_API"), fasthttp.StatusBadRequest)
		txContext.LastStageData["error"] = err
		txContext.LastStageData["errorType"] = "ENDPOINT_ERROR_MAKE_REQUEST"
		return responseBytes, errors.New("ENDPOINT_ERROR_MAKE_REQUEST")
	}
	return responseBytes, nil
}
