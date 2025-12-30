package controllers

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"is-public-api/application/apihelpers"
	"is-public-api/application/models"
	"is-public-api/helpers/logger"
	"net/http"
)

type DocumentHandler struct {
}

func NewDocumentHandler() *DocumentHandler {
	return &DocumentHandler{}
}

func (controller *DocumentHandler) Download(ctx *fasthttp.RequestCtx, txContext *models.TxContext) {
	log := logger.GetLoggerHooks(txContext, "DocumentHandler", "Download")
	defer logger.End(log)
	url := string(ctx.QueryArgs().Peek("url"))
	if url == "" {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}
	file := string(ctx.QueryArgs().Peek("file"))
	if file == "" {
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprint(url), nil)

	if err != nil {
		fmt.Println(err)
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		apihelpers.RenderError(ctx, errors.New("BAD_REQUEST"), fasthttp.StatusBadRequest)
	}

	apihelpers.RenderFile(ctx, res.StatusCode, file, result, "")
}
