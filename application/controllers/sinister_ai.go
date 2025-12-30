package controllers

import (
	"errors"
	"is-public-api/application/apihelpers"
	"is-public-api/application/mappers"
	"is-public-api/application/models"
	"is-public-api/application/resources"
	"is-public-api/application/services"
	"is-public-api/helpers/logger"

	"github.com/francoispqt/gojay"
	"github.com/valyala/fasthttp"
)

type SinisterAIHandler struct {
	service  services.ISinisterService
	sinister services.ISinisterCaseService
	domain   services.ISinisterAIDomain
}

func NewSinisterAIHandler(service services.ISinisterService, sinister services.ISinisterCaseService,
	domain services.ISinisterAIDomain) *SinisterAIHandler {
	return &SinisterAIHandler{service: service, sinister: sinister, domain: domain}
}

func (controller *SinisterAIHandler) SinisterSaveHandlerWithGenAI(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SinisterAIHandler", "SinisterSaveHandlerWithGenAI")
	defer logger.End(log)

	request := &resources.SinisterRequest{}

	err := gojay.Unmarshal(ctx.PostBody(), request)
	if err != nil {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	originBytes := ctx.Request.Header.Peek("X-Origin")
	txContext.Origin = "Web"
	if originBytes != nil {
		txContext.Origin = string(originBytes)
	}

	response, err := controller.domain.Save(txContext, mappers.RequestToSinister(request))
	if err != nil {
		apihelpers.RenderError(ctx, err, 500, &resources.MapResponse{"cases": response})
		return
	}
	apihelpers.RenderSuccess(ctx, &resources.MapResponse{"cases": response}, "ok")
}

func (controller *SinisterAIHandler) SinisterAnalyzeDocHandler(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SinisterAIHandler", "SinisterAnalyzeDocHandler")
	defer logger.End(log)

	log.InfoWith().String("request", string(ctx.PostBody()))
	request := &resources.SinisterCase{}
	bytes := ctx.PostBody()
	err := gojay.Unmarshal(bytes, request)
	if err != nil {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	result, err := controller.domain.AnalyzeDocument(txContext, request)
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusInternalServerError)
		return
	}

	apihelpers.RenderSuccess(ctx, result, "success")
}

func (controller *SinisterAIHandler) UploadObservedDocument(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SinisterAIHandler", "UploadObservedDocument")
	defer logger.End(log)

	request := &resources.SinisterCase{}
	bytes := ctx.PostBody()
	err := gojay.Unmarshal(bytes, request)
	if err != nil {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	err = controller.domain.UploadObservedDocument(txContext, request)
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusInternalServerError)
		return
	}

	apihelpers.RenderSuccess(ctx, nil, "Documento observado actualizado correctamente")
}

func (controller *SinisterAIHandler) GetByCase(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SinisterAIHandler", "GetByCase")
	defer logger.End(log)

	caseID := ctx.UserValue("case_id").(string)
	if caseID == "" {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	sinister, err := controller.domain.GetByCase(txContext, caseID)
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusInternalServerError)
		return
	}

	apihelpers.RenderSuccess(ctx, sinister, "success")
}

func (controller *SinisterAIHandler) GetAll(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SinisterAIHandler", "GetAll")
	defer logger.End(log)

	cases, err := controller.domain.GetAll(txContext)
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusInternalServerError)
		return
	}

	apihelpers.RenderSuccess(ctx, cases, "success")
}

func (controller *SinisterAIHandler) SinisterPreviousCases(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "SinisterAIHandler", "SinisterPreviousCases")
	defer logger.End(log)

	doc := ctx.UserValue("doc").(string)
	if doc == "" {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	result, err := controller.domain.GetPreviousCases(txContext, doc)
	if err != nil {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	apihelpers.RenderSuccess(ctx, result, "success")
}
