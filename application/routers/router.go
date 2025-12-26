package routers

import (
	"is-public-api/application/apihelpers"
	"is-public-api/application/container"
	"is-public-api/application/controllers"
	"is-public-api/application/middlewares"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type pathBuilder func(string) string

type httpRouter struct {
	contextPath               string
	healthHandler             *controllers.HealthHandler
	collaboratorHandler       *controllers.CollaboratorHandler
	soatReturnHandler         *controllers.SoatReturnHandler
	sinisterPaymentHandler    *controllers.SinisterHandler
	sinisterAIHandler         *controllers.SinisterAIHandler
	subscriptionCenterHandler *controllers.SubscriptionCenterHandler
	documentHandler           *controllers.DocumentHandler
	handleCORS                func(fasthttp.RequestHandler) fasthttp.RequestHandler
}

func NewHttpRouter(contextPath string) *httpRouter {
	return &httpRouter{
		contextPath:               contextPath,
		healthHandler:             container.HealthHandler(),
		collaboratorHandler:       container.CollaboratorHandler(),
		soatReturnHandler:         container.SoatReturnHandler(),
		sinisterPaymentHandler:    container.SinisterPaymentHandler(),
		sinisterAIHandler:         container.SinisterAIHandler(),
		subscriptionCenterHandler: container.SubscriptionCenterHandler(),
		documentHandler:           container.DocumentHandler(),
	}
}

func (httpRouter *httpRouter) Handler() fasthttp.RequestHandler {

	path_v1 := httpRouter.pathVersion("v1.0")
	path_v2 := httpRouter.pathVersion("v2.0")

	router := fasthttprouter.New()
	router.RedirectTrailingSlash = true

	router.GET(path_v1("/info"), httpRouter.healthHandler.Health)
	router.GET(path_v1("/datetime"), httpRouter.healthHandler.DateTime)
	router.GET(path_v1("/photocheck"), middlewares.Trace(httpRouter.collaboratorHandler.GetController))
	router.GET(path_v1("/soat/returns"), middlewares.Trace(httpRouter.soatReturnHandler.GetController))
	router.GET(path_v1("/subscription-center/:key"), middlewares.Trace(httpRouter.subscriptionCenterHandler.GetSubscriptionCenterController))
	router.POST(path_v1("/subscription-center"), middlewares.Trace(httpRouter.subscriptionCenterHandler.CreateSubscriptionCenterController))

	router.GET(path_v1("/sinister/payments"), middlewares.Trace(httpRouter.sinisterPaymentHandler.GetController))
	router.POST(path_v1("/sinister/save"), middlewares.Trace(httpRouter.sinisterPaymentHandler.PostController))
	router.POST(path_v2("/sinister"), middlewares.Trace(httpRouter.sinisterAIHandler.SinisterSaveHandlerWithGenAI))

	router.POST(path_v2("/sinister/doc"), middlewares.Trace(httpRouter.sinisterAIHandler.SinisterAnalyzeDocHandler))
	router.POST(path_v2("/sinister/doc/observed"), middlewares.Trace(httpRouter.sinisterAIHandler.UploadObservedDocument))
	router.GET(path_v2("/sinister/previous/:doc"), middlewares.Trace(httpRouter.sinisterAIHandler.SinisterPreviousCases))

	// Deprecated: OCR + GenAI
	// router.POST(path_v2("/sinister/analyze/doc"), middlewares.Trace(httpRouter.sinisterAIHandler.SinisterAnalyzeDocHandlerWithGenAI))

	// New: Only GenAI
	router.GET(path_v2("/sinister"), middlewares.Trace(httpRouter.sinisterAIHandler.GetAll))
	router.GET(path_v2("/sinister/case/:case_id"), middlewares.Trace(httpRouter.sinisterAIHandler.GetByCase))
	router.GET(path_v1("/sinister/get/:case_id"), middlewares.Trace(httpRouter.sinisterPaymentHandler.GetSinister))
	router.GET(path_v1("/sinister/document"), middlewares.Trace(httpRouter.documentHandler.Download))

	if httpRouter.handleCORS != nil {
		return httpRouter.handleCORS(router.Handler)
	}

	return router.Handler
}

func (httpRouter *httpRouter) EnableCORS(origins string) *httpRouter {

	httpRouter.handleCORS = func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {

		return func(ctx *fasthttp.RequestCtx) {

			if string(ctx.Method()) == fasthttp.MethodOptions {
				ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
				ctx.Response.Header.Add("Access-Control-Allow-Headers", apihelpers.AllowedHeaders)
				ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
				ctx.Response.Header.Add("Access-Control-Allow-Origin", origins)
				ctx.SetStatusCode(fasthttp.StatusOK)
				ctx.SetConnectionClose()
				return
			}

			handler(ctx)
		}
	}

	return httpRouter
}

func (httpRouter *httpRouter) pathVersion(version string) pathBuilder {
	return path(httpRouter.contextPath + "/" + version)
}

func path(basePath string) pathBuilder {
	return func(path string) string {
		return basePath + path
	}
}
