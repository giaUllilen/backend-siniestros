package container

import (
	"sync"

	"is-public-api/application/configs"
	"is-public-api/application/controllers"
	"is-public-api/application/endpoints"
	"is-public-api/application/services"
	"is-public-api/application/storage"
	restful "is-public-api/helpers/apihelpers"
	"is-public-api/helpers/configloader"
	"is-public-api/helpers/database"
)

var syncDbHelper sync.Once

var mongoDbHelper *database.MongodbHelper

var endpointsConfig *restful.EndpointConfig

func getEndpointsConfig() *restful.EndpointConfig {
	if endpointsConfig == nil {
		conf := &configs.EndpointsConfig{}
		configloader.ReadConf(conf)
		endpointsConfig = &restful.EndpointConfig{
			CloudFunction: restful.Endpoint{
				Uri: conf.Services.CloudFunction.Uri,
			},
			EventLog: restful.EventLog{
				Uri:       conf.Services.EventLog.Uri,
				ApiKey:    conf.Services.EventLog.ApiKey,
				ApiKeyWsp: conf.Services.EventLog.ApiKeyWsp,
			},
			Sinister: restful.Sinister{
				Path:                conf.Services.Sinister.Path,
				Save:                conf.Services.Sinister.Save,
				Get:                 conf.Services.Sinister.Get,
				Document:            conf.Services.Sinister.Document,
				History:             conf.Services.Sinister.History,
				Delete:              conf.Services.Sinister.Delete,
				UrlGenAI:            conf.Services.Sinister.UrlGenAI,
				TokenGenAI:          conf.Services.Sinister.TokenGenAI,
				IdPromtDP:           conf.Services.Sinister.IdPromtDP,
				IdPromtDM:           conf.Services.Sinister.IdPromtDM,
				IdPromtDictamen:     conf.Services.Sinister.IdPromtDictamen,
				ApiDiagnostic:       conf.Services.Sinister.ApiDiagnostic,
				UpdateObservationIA: conf.Services.Sinister.UpdateObservationIA,
			},
			Storage: restful.Service{
				Path:     conf.Services.Storage.Path,
				Endpoint: conf.Services.Storage.Endpoint,
			},
			Notification: restful.Service{
				Path:     conf.Services.Notifications.Path,
				Endpoint: conf.Services.Notifications.Endpoint,
			},
			Qualitat: restful.Qualitat{
				QualitatUser: conf.Services.Qualitat.QualitatUser,
				QualitatPass: conf.Services.Qualitat.QualitatPass,
			},
		}
	}

	return endpointsConfig
}

func serverConfig() *configs.ConfigServer {
	conf := &configs.ConfigServer{}
	configloader.ReadConf(conf)
	return conf
}

func dbConfig() *database.MongoConfig {
	dbConf := &configs.MongoConfig{}
	configloader.ReadConf(dbConf)

	return &database.MongoConfig{
		Uri:             dbConf.Mongodb.Uri,
		User:            dbConf.Mongodb.Credentials.Username,
		Password:        dbConf.Mongodb.Credentials.Password,
		Database:        dbConf.Mongodb.DatabaseName,
		ApplicationName: dbConf.Mongodb.ApplicationName,
		MinPoolSize:     dbConf.Mongodb.ConnectionPool.MinSize,
		MaxPoolSize:     dbConf.Mongodb.ConnectionPool.MaxSize,
		AuthMechanism:   dbConf.Mongodb.AuthMechanism,
	}
}

func MongodbHelper() *database.MongodbHelper {
	syncDbHelper.Do(func() {
		mongodbHelper := database.NewMongodbHelper(dbConfig())
		go mongodbHelper.OpenConnection()
		mongoDbHelper = mongodbHelper
	})
	return mongoDbHelper
}

func EndpointHelper() *restful.EndpointHelper {
	return restful.NewEndpointHelper(getEndpointsConfig())
}

func HealthHandler() *controllers.HealthHandler {
	return controllers.NewHealthController()
}

func CollaboratorRepository() storage.ICollaboratorRepository {
	return storage.NewCustomerRepository(MongodbHelper())
}

func CollaboratorFinder() services.ICollaboratorFinder {
	return services.NewCollaboratorFinder(CollaboratorRepository())
}

func CollaboratorHandler() *controllers.CollaboratorHandler {
	return controllers.NewCollaboratorHandler(CollaboratorFinder())
}

func SoatReturnRepository() storage.ISoatReturnRepository {
	return storage.NewSoatReturnRepository(MongodbHelper())
}

func SoatReturnFinder() services.ISoatReturnFinder {
	return services.NewSoatReturnFinder(SoatReturnRepository())
}

func SoatReturnHandler() *controllers.SoatReturnHandler {
	return controllers.NewSoatReturnHandler(SoatReturnFinder())
}

func SinisterPaymentRepository() storage.ISinisterPaymentRepository {
	return storage.NewSinisterPaymentRepository(MongodbHelper())
}

func SinisterPaymentFinder() services.ISinisterService {
	return services.NewSinisterPaymentFinder(SinisterPaymentRepository())
}

func SinisterCaseRepository() storage.ISinisterCaseRepository {
	return storage.NewSinisterCaseRepository(MongodbHelper())
}

func SinisterCaseService() services.ISinisterCaseService {
	return services.NewSinisterCaseService(SinisterCaseRepository())
}

func SinisterPaymentHandler() *controllers.SinisterHandler {
	return controllers.NewSinisterPaymentHandler(SinisterPaymentFinder(), SinisterServiceDomain())
}

func SoatRepository() storage.ISoatRepository {
	return storage.NewSoatRepository(MongodbHelper())
}

func SoatFinder() services.ISoatFinder {
	return services.NewSoatFinder(SoatRepository())
}

func GetSiniesterQualitat() endpoints.IQualitatEndpoints {
	return endpoints.NewQualitatEndpoint(getEndpointsConfig(), EndpointHelper())
}

func SinisterAIHandler() *controllers.SinisterAIHandler {
	return controllers.NewSinisterAIHandler(SinisterPaymentFinder(), SinisterCaseService(), SinisterAIDomain())
}

func DocumentHandler() *controllers.DocumentHandler {
	return controllers.NewDocumentHandler()
}

func SubscriptionCenterRepository() storage.ISubscriptionCenterRepository {
	return storage.NewSubscriptionCenterRepository(MongodbHelper())
}

func SubscriptionCenterService() services.ISubscriptionCenter {
	return services.NewSubscriptionCenterService(SubscriptionCenterRepository())
}

func CloudFunctionSubscriptionCenterService() services.ICloudFunctionSubscriptionCenter {
	return services.NewCloudFunctionSubscriptionCenterService(EndpointHelper())
}

func SubscriptionCenterHandler() *controllers.SubscriptionCenterHandler {
	return controllers.NewSubscriptionCenterHandler(SubscriptionCenterService(), CloudFunctionSubscriptionCenterService())
}

func SinisterCoverageIAService() services.ISinisterCoverageIAService {
	return services.NewSinisterCoverageIAService(SinisterCoverageIARepository())
}

func SinisterCoverageIARepository() storage.ISinisterCoverageIARepository {
	return storage.NewSinisterCoverageIARepository(MongodbHelper())
}

func SinisterServiceDomain() services.ISinisterServiceDomain {
	return services.NewSinisterServiceDomain(serverConfig(), SinisterEndpoints(), StorageEndpoints(), EventsEndpoints(), NotificationEndpoints(), SinisterCoverageIAService(), SinisterAIDomainForDomain(), SinisterCaseService())
}

func SinisterServiceDomainForAI() services.ISinisterServiceDomain {
	return services.NewSinisterServiceDomain(serverConfig(), SinisterEndpoints(), StorageEndpoints(), EventsEndpoints(), NotificationEndpoints(), SinisterCoverageIAService(), nil, SinisterCaseService())
}

func SinisterAIDomain() services.ISinisterAIDomain {
	return services.NewSinisterAIServiceDomain(serverConfig(), SinisterEndpoints(), StorageEndpoints(), EventsEndpoints(), NotificationEndpoints(), EndpointHelper(), GetSiniesterQualitat(), SoatFinder(), SinisterCaseService(), SinisterServiceDomain())
}

func SinisterAIDomainForDomain() services.ISinisterAIDomain {
	return services.NewSinisterAIServiceDomain(serverConfig(), SinisterEndpoints(), StorageEndpoints(), EventsEndpoints(), NotificationEndpoints(), EndpointHelper(), GetSiniesterQualitat(), SoatFinder(), SinisterCaseService(), SinisterServiceDomainForAI())
}

func SinisterEndpoints() endpoints.ISinisterEndpoints {
	return endpoints.NewSinisterEndpoints(getEndpointsConfig())
}

func EventsEndpoints() endpoints.IEventsEndpoints {
	return endpoints.NewEventsEndpoint(getEndpointsConfig())
}

func StorageEndpoints() endpoints.IStorageEndpoints {
	return endpoints.NewStorageEndpoints(getEndpointsConfig())
}

func NotificationEndpoints() endpoints.INotificationEndpoints {
	return endpoints.NewNotificationEndpoints(getEndpointsConfig())
}
