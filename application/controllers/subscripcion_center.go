package controllers

import (
	"errors"
	"time"

	"github.com/francoispqt/gojay"
	"is-public-api/application/apihelpers"
	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/application/services"
	"is-public-api/helpers/logger"

	"github.com/valyala/fasthttp"
)

type SubscriptionCenterHandler struct {
	service services.ISubscriptionCenter
	cloudFunction services.ICloudFunctionSubscriptionCenter
}

func NewSubscriptionCenterHandler(service services.ISubscriptionCenter, cloudFunction services.ICloudFunctionSubscriptionCenter) *SubscriptionCenterHandler {
	return &SubscriptionCenterHandler{service: service, cloudFunction: cloudFunction}
}

func (controller *SubscriptionCenterHandler) CreateSubscriptionCenterController(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SubscriptionCenterHandler", "CreateSubscriptionCenterController")
	defer logger.End(log)
	request := make(resources.MapRequest)
	err := gojay.UnmarshalJSONObject(ctx.PostBody(), &request)
	if err != nil {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}
	key, _ := request["key"].(string)
	option, _ := request["option"].(string)
	value, _ := request["value"].(string)
	version, _ := request["type"].(string)
	if version == "" {
		version = "v1"
	}
	description, _ := request["description"].(string)
	id, err := controller.service.CreateSubscriptionCenterOptions(txContext, key, option, value, version, &description)
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusBadRequest)
		return
	}
	request["type"] = version
	request["created_at"] = time.Now()
	err = controller.cloudFunction.CreateSubscriptionCenterOptions(txContext, request)
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusInternalServerError)
		return
	}
	if version == "v2" {
		go func() {

		}()
	}
	response := make(resources.MapResponse)
	response["id"] = id
	apihelpers.RenderSuccess(ctx, response, "success")
}

func (controller *SubscriptionCenterHandler) GetSubscriptionCenterController(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SubscriptionCenterHandler", "CreateSubscriptionCenterController")
	defer logger.End(log)

	key := ctx.UserValue("key").(string)
	if key == "" {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	id, err := controller.service.FindSubscriptionCenterOptionsByID(txContext, key)
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusBadRequest)
		return
	}
	response := make(resources.MapResponse)
	response["id"] = id
	apihelpers.RenderSuccess(ctx, response, "success")
}
