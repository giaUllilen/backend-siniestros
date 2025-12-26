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

type CollaboratorHandler struct {
	finder services.ICollaboratorFinder
}

func NewCollaboratorHandler(finder services.ICollaboratorFinder) *CollaboratorHandler {
	return &CollaboratorHandler{finder: finder}
}

func (controller *CollaboratorHandler) GetController(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "CollaboratorHandler", "PostCollaborator")
	defer logger.End(log)

	code := ctx.QueryArgs().Peek("code")
	if code == nil {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	collaborator, err := controller.finder.Find(txContext, string(code))
	if err != nil {
		apihelpers.RenderError(ctx, err, fasthttp.StatusBadRequest)
		return
	}
	response := new(resources.MapResponse)

	mappers.ModelToCollaboratorResponse(collaborator, response)
	apihelpers.RenderSuccess(ctx, response, "success")
}
