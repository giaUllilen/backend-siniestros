package services

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/helpers/apihelpers"
	"is-public-api/helpers/logger"

	"github.com/francoispqt/gojay"
)

type cloudFunctionSubCenterService struct {
	conf *apihelpers.EndpointHelper
}

func NewCloudFunctionSubscriptionCenterService(conf *apihelpers.EndpointHelper) ICloudFunctionSubscriptionCenter {
	return &cloudFunctionSubCenterService{
		conf: conf,
	}
}

func (service *cloudFunctionSubCenterService) CreateSubscriptionCenterOptions(txContext *models.TxContext, data resources.MapRequest) error {
	log := logger.GetLoggerHooks(txContext, "cloudFunctionSubCenterService", "CreateSubscriptionCenterOptions")
	bytes, err := gojay.MarshalJSONObject(data)
	req, err := http.NewRequest("POST", service.conf.Conf.CloudFunction.Uri, strings.NewReader(string(bytes)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Logger.Info("cloudFunctionSubCenterService response --> " + string(body))
	if res.StatusCode != 200 {
		return errors.New("no enviado a cloud function")
	}
	return nil
}
