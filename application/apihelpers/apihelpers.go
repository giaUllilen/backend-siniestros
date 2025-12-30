package apihelpers

import (
	"fmt"
	"net/http"
	"strings"

	"is-public-api/application/configs"
	"is-public-api/helpers/configloader"
	"is-public-api/helpers/logger"

	"github.com/francoispqt/gojay"

	"github.com/valyala/fasthttp"
)

var serverConf = &configs.ConfigServer{}

var AllowedHeaders = strings.Join([]string{
	"Accept",
	"Origin",
	"Content-Type",
	"X-Requested-With",
	"Access-Control-Allow-Origin",
	"Access-Control-Request-Method",
	"Cache-Control",
	"Access-Control-Request-Headers",
	"Authorization",
	"X-Origin",
}, ",")

func init() {
	configloader.ReadConf(serverConf)
}

const (
	CodeOk    ResponseCode = "01"
	CodeError ResponseCode = "99"
)

type ResponseCode string

type ResponseWrapper struct {
	Code    ResponseCode
	Data    interface{}
	Message string
}

// implement MarshalerJSONObject
func (respWrapper *ResponseWrapper) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("code", string(respWrapper.Code))
	enc.StringKeyOmitEmpty("message", respWrapper.Message)
	enc.AddInterfaceKey("data", respWrapper.Data)
}

func (respWrapper *ResponseWrapper) IsNil() bool {
	return respWrapper == nil
}

func RenderError(ctx *fasthttp.RequestCtx, err error, statusCode int, response ...interface{}) {

	log := logger.GetLogger()
	defer logger.End(log)

	log.Error(fmt.Sprintf("%v", err))

	responseOut := ResponseWrapper{
		Code:    CodeError,
		Message: err.Error(),
	}

	if len(response) > 0 {
		responseOut.Data = response[0]
	}

	enc := gojay.BorrowEncoder(ctx.Response.BodyWriter())
	defer enc.Release()

	if err := enc.EncodeObject(&responseOut); err != nil {
		fmt.Printf("ERROR Marshal Response %v :", err)
		RenderError(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", AllowedHeaders)
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetConnectionClose()
}

func RenderSuccess(ctx *fasthttp.RequestCtx, body interface{}, message string) {

	responseOut := ResponseWrapper{Code: CodeOk, Message: message, Data: body}

	enc := gojay.BorrowEncoder(ctx.Response.BodyWriter())
	defer enc.Release()

	if err := enc.EncodeObject(&responseOut); err != nil {
		fmt.Printf("ERROR Marshal Response %v :", err)
		RenderError(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", AllowedHeaders)
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetConnectionClose()
}

func RenderErrorWithData(ctx *fasthttp.RequestCtx, body interface{}, message string) {

	responseOut := ResponseWrapper{Code: CodeError, Message: message, Data: body}

	enc := gojay.BorrowEncoder(ctx.Response.BodyWriter())
	defer enc.Release()

	if err := enc.EncodeObject(&responseOut); err != nil {
		fmt.Printf("ERROR Marshal Response %v :", err)
		RenderError(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", AllowedHeaders)
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	ctx.SetConnectionClose()
}

func RenderServiceResponse(ctx *fasthttp.RequestCtx, status int, body []byte) {

	ctx.Response.BodyWriter().Write(body)

	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", AllowedHeaders)
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(status)
	ctx.SetConnectionClose()
}

func RenderFile(ctx *fasthttp.RequestCtx, status int, fileName string, data []byte, contentType string) {
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	ctx.Response.Header.Add("Content-Description", "File Transfer")
	ctx.Response.Header.Add("Content-Disposition", "attachment; filename="+fileName)
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	ctx.SetContentType(contentType)
	ctx.SetStatusCode(status)
	ctx.Write(data)
	ctx.SetConnectionClose()
}

func RenderBytesResponse(ctx *fasthttp.RequestCtx, status int, body []byte) {

	ctx.Response.BodyWriter().Write(body)

	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", AllowedHeaders)
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(status)
	ctx.SetConnectionClose()
}
