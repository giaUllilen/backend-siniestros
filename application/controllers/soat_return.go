package controllers

import (
	"errors"

	"is-public-api/application/apihelpers"
	"is-public-api/application/mappers"
	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/application/services"
	"is-public-api/helpers/logger"

	"github.com/valyala/fasthttp"
)

type SoatReturnHandler struct {
	finder services.ISoatReturnFinder
}

func NewSoatReturnHandler(finder services.ISoatReturnFinder) *SoatReturnHandler {
	return &SoatReturnHandler{finder: finder}
}

func (controller *SoatReturnHandler) GetController(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SoatReturnHandler", "GetSoatReturn")
	defer logger.End(log)

	documentNumber := ctx.QueryArgs().Peek("document")
	if documentNumber == nil {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	soat, err := controller.finder.FindByDocument(txContext, string(documentNumber))
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusBadRequest)
		return
	}
	response := new(resources.MapResponse)

	mappers.ModelToSoatReturnResponse(soat, response)
	apihelpers.RenderSuccess(ctx, response, "success")
}
