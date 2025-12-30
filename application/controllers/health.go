package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"is-public-api/application/apihelpers"
	"is-public-api/helpers/logger"

	"github.com/francoispqt/onelog"
	"github.com/valyala/fasthttp"
)

type HealthHandler struct {
}

func NewHealthController() *HealthHandler {
	return &HealthHandler{}
}

func (controller *HealthHandler) Health(ctx *fasthttp.RequestCtx) {
	log := logger.GetLogger()
	defer logger.End(log)

	log.Hook(func(entry onelog.Entry) {
		entry.String("time", time.Now().Format(time.RFC3339))
		entry.String("caller", "HealthHandler")
	})

	log.Info("check_health")

	responseOut := make(map[string]interface{})
	enc := json.NewEncoder(ctx.Response.BodyWriter())
	now := time.Now()

	err := checkStatus()
	if err != nil {

		responseOut["status"] = "DOWN"
		responseOut["time"] = now.Format("02/01/2006 3:04:05 PM")
		responseOut["error"] = err.Error()

		if err := enc.Encode(&responseOut); err != nil {
			fmt.Printf("ERROR Marshal Response %v :", err)
			apihelpers.RenderError(ctx, err, http.StatusInternalServerError)
			return
		}

		ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Add("Access-Control-Allow-Headers", apihelpers.AllowedHeaders)
		ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
		ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetConnectionClose()
		return
	}

	responseOut["status"] = "UP"
	responseOut["time"] = now.Format("02/01/2006 3:04:05 PM")

	if err := enc.Encode(&responseOut); err != nil {
		fmt.Printf("ERROR Marshal Response %v :", err)
		apihelpers.RenderError(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", apihelpers.AllowedHeaders)
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetConnectionClose()
}

func (controller *HealthHandler) DateTime(ctx *fasthttp.RequestCtx) {
	log := logger.GetLogger()
	defer logger.End(log)

	log.Hook(func(entry onelog.Entry) {
		entry.String("time", time.Now().Format(time.RFC3339))
		entry.String("caller", "DateTimeHandler")
	})

	responseOut := make(map[string]interface{})
	enc := json.NewEncoder(ctx.Response.BodyWriter())
	now := time.Now()
	responseOut["time"] = now.Add(-5*time.Hour).Format("2006-01-02 15:04:05")

	if err := enc.Encode(&responseOut); err != nil {
		fmt.Printf("ERROR Marshal Response %v :", err)
		apihelpers.RenderError(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", apihelpers.AllowedHeaders)
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetConnectionClose()
}

// Check anything that ensure the correct status  api
// example: check db connection
func checkStatus() error {

	return nil
}
