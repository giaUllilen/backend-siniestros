package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"is-public-api/application/apihelpers"
	"is-public-api/application/mappers"
	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/application/services"
	"is-public-api/helpers/logger"

	"github.com/valyala/fasthttp"
)

type SinisterHandler struct {
	service services.ISinisterService
	domain  services.ISinisterServiceDomain
}

func NewSinisterPaymentHandler(service services.ISinisterService, domain services.ISinisterServiceDomain) *SinisterHandler {
	return &SinisterHandler{service: service, domain: domain}
}

func (controller *SinisterHandler) GetController(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SinisterHandler", "GetSinisterPayment")
	defer logger.End(log)

	documentNumber := ctx.QueryArgs().Peek("document")
	if documentNumber == nil {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	payment, err := controller.service.FindByDocumentNumber(txContext, string(documentNumber))
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusBadRequest)
		return
	}
	response := new(resources.MapResponse)

	mappers.ModelToSinisterPaymentResponse(payment, response)
	apihelpers.RenderSuccess(ctx, response, "success")
}

func (controller *SinisterHandler) PostController(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SinisterSaveHandler", "SaveSinister")
	defer logger.End(log)
	form, err := ctx.MultipartForm()
	if err != nil {
		apihelpers.RenderError(ctx, errors.New("FILE_BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	originBytes := ctx.Request.Header.Peek("X-Origin")
	txContext.Origin = "Web"
	if originBytes != nil {
		txContext.Origin = string(originBytes)
	}

	attachments := form.File["attachments"]

	request := make(map[string]interface{})
	data := ctx.FormValue("sinister")
	txContext.Environment = fmt.Sprint(string(ctx.FormValue("environment")))

	err = json.Unmarshal(data, &request)

	data = ctx.FormValue("attachments_metadata")
	var coverages []map[string]interface{}
	err = json.Unmarshal(data, &coverages)

	response, err := controller.domain.Save(txContext, request, coverages, attachments)
	if err != nil {
		apihelpers.RenderError(ctx, err, 500, &resources.MapResponse{"cases": response})
		return
	}
	apihelpers.RenderSuccess(ctx, &resources.MapResponse{"cases": response}, "ok")
}

func (controller *SinisterHandler) GetSinister(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SinisterSaveHandler", "GetSinister")
	defer logger.End(log)

	caseID := ctx.UserValue("case_id").(string)
	if caseID == "" {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	// Nota: Se cambió para enviar caseID como string (ej: "CIS_98746") directamente al API externo
	// En lugar de convertirlo a ObjectID de MongoDB, ya que el endpoint externo espera el número de caso como string
	// Código anterior (comentado):
	// objectIDCase, err := primitive.ObjectIDFromHex(caseID)
	// if err != nil {
	//     apihelpers.RenderError(ctx, err, fasthttp.StatusNotAcceptable)
	//     return
	// }
	// statusCode, body, _ := controller.domain.FindByCaseNumber(txContext, objectIDCase)

	statusCode, body, _ := controller.domain.FindByCaseNumber(txContext, caseID)
	apihelpers.RenderServiceResponse(ctx, fasthttp.StatusOK, body)
	fmt.Println(statusCode, "Caso encontrado: ", caseID)

}
