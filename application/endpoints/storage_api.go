package endpoints

import (
	"errors"

	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/helpers/apihelpers"
)

type storageEndpoints struct {
	conf *apihelpers.Service
}

func NewStorageEndpoints(conf *apihelpers.EndpointConfig) *storageEndpoints {
	return &storageEndpoints{conf: &apihelpers.Service{
		Path:     conf.Storage.Path,
		Endpoint: conf.Storage.Endpoint,
	}}
}

func (service *storageEndpoints) Upload(txContext *models.TxContext, form *resources.MapString, file apihelpers.File) ([]byte, error) {
	url := service.conf.Path + service.conf.Endpoint
	// txContext.LastStageData["url"] = url
	txContext.LastStageData["file"] = file.Name
	txContext.LastStageData["size"] = file.Size
	txContext.LastStageData["type"] = file.ContentType

	contentType, body, err := apihelpers.MakeMultipart(*form, file)
	if err != nil {
		// apihelpers.RenderError(ctx, errors.New("FILE_ERROR_MAKE_REQUEST"), fasthttp.StatusBadRequest)
		txContext.LastStageData["error"] = err
		txContext.LastStageData["errorType"] = "FILE_ERROR_MAKE_REQUEST"
		return nil, errors.New("FILE_ERROR_MAKE_REQUEST")
	}

	statusCode, responseBytes, err := apihelpers.HttpMultipartRequest(url, "POST", body, contentType)
	txContext.LastStageData["statusCode"] = statusCode
	txContext.LastStageData["response"] = string(responseBytes)
	if err != nil {
		// apihelpers.RenderError(ctx, errors.New("FILE_ERROR_REQUEST_API"), fasthttp.StatusBadRequest)
		txContext.LastStageData["error"] = err
		txContext.LastStageData["errorType"] = "FILE_ERROR_REQUEST_API"
		return responseBytes, errors.New("FILE_ERROR_REQUEST_API")
	}
	return responseBytes, nil
}
