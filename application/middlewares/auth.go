package middlewares

import (
	"fmt"

	"is-public-api/application/apihelpers"
	"is-public-api/application/models"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorizedHandler func(*fasthttp.RequestCtx, *models.TxContext)

func Trace(authorizedHandler AuthorizedHandler) fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {

		defer func() {
			if err := recover(); err != nil {
				apihelpers.RenderError(ctx, fmt.Errorf("%+v", err), fasthttp.StatusInternalServerError)
			}
		}()

		transactionId := primitive.NewObjectID().Hex()
		txContext := &models.TxContext{
			TransactionID: transactionId,
			ClientIp:      ctx.RemoteIP().DefaultMask().String(),
			Event:         make(map[string]interface{}),
			LastStageData: make(map[string]interface{}),
		}

		authorizedHandler(ctx, txContext)
		return
	}
}
