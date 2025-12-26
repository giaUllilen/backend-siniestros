package configs

import "is-public-api/helpers/configloader"

type EndpointsConfig struct {
	Services struct {
		CloudFunction struct {
			Uri string `yaml:"uri"`
		} `yaml:"cloud_function"`
		EventLog struct {
			Uri       string `yaml:"uri"`
			ApiKey    string `yaml:"api_key"`
			ApiKeyWsp string `yaml:"api_key_wsp"`
		} `yaml:"event_log"`
		Sinister struct {
			Path                string `yaml:"path"`
			Save                string `yaml:"save"`
			Get                 string `yaml:"get"`
			Document            string `yaml:"document"`
			History             string `yaml:"history"`
			Delete              string `yaml:"delete"`
			UrlGenAI            string `yaml:"api_genai_url"`
			TokenGenAI          string `yaml:"api_genai_key"`
			IdPromtDP           string `yaml:"promt_dp"`
			IdPromtDM           string `yaml:"promt_dm"`
			IdPromtDictamen     string `yaml:"promt_dictamen"`
			ApiDiagnostic       string `yaml:"api_diagnostic"`
			UpdateObservationIA string `yaml:"updateObservationIA"`
		} `yaml:"sinister"`
		Storage struct {
			Path     string `yaml:"path"`
			Endpoint string `yaml:"endpoint"`
		} `yaml:"storage"`
		Notifications struct {
			Path     string `yaml:"path"`
			Endpoint string `yaml:"endpoint"`
		} `yaml:"notifications"`
		Qualitat struct {
			QualitatUser string `yaml:"qualitat_user"`
			QualitatPass string `yaml:"qualitat_pass"`
		}
	} `yaml:"services"`
}

func (conf *EndpointsConfig) Merge(envCfg interface{}) configloader.ConfigurationProperties {
	envConfig := envCfg.(*EndpointsConfig)
	// Cloud Function Subscription
	conf.Services.CloudFunction.Uri = configloader.GetVal(envConfig.Services.CloudFunction.Uri, conf.Services.CloudFunction.Uri).(string)
	// EvenLog Api
	conf.Services.EventLog.Uri = configloader.GetVal(envConfig.Services.EventLog.Uri, conf.Services.EventLog.Uri).(string)
	conf.Services.EventLog.ApiKey = configloader.GetVal(envConfig.Services.EventLog.ApiKey, conf.Services.EventLog.ApiKey).(string)
	conf.Services.EventLog.ApiKeyWsp = configloader.GetVal(envConfig.Services.EventLog.ApiKeyWsp, conf.Services.EventLog.ApiKeyWsp).(string)
	// Sinister Api
	conf.Services.Sinister.Path = configloader.GetVal(envConfig.Services.Sinister.Path, conf.Services.Sinister.Path).(string)
	conf.Services.Sinister.Save = configloader.GetVal(envConfig.Services.Sinister.Save, conf.Services.Sinister.Save).(string)
	conf.Services.Sinister.Get = configloader.GetVal(envConfig.Services.Sinister.Get, conf.Services.Sinister.Get).(string)
	conf.Services.Sinister.Document = configloader.GetVal(envConfig.Services.Sinister.Document, conf.Services.Sinister.Document).(string)
	conf.Services.Sinister.Delete = configloader.GetVal(envConfig.Services.Sinister.Delete, conf.Services.Sinister.Delete).(string)
	conf.Services.Sinister.UrlGenAI = configloader.GetVal(envConfig.Services.Sinister.UrlGenAI, conf.Services.Sinister.UrlGenAI).(string)
	conf.Services.Sinister.TokenGenAI = configloader.GetVal(envConfig.Services.Sinister.TokenGenAI, conf.Services.Sinister.TokenGenAI).(string)
	conf.Services.Sinister.IdPromtDP = configloader.GetVal(envConfig.Services.Sinister.IdPromtDP, conf.Services.Sinister.IdPromtDP).(string)
	conf.Services.Sinister.IdPromtDM = configloader.GetVal(envConfig.Services.Sinister.IdPromtDM, conf.Services.Sinister.IdPromtDM).(string)
	conf.Services.Sinister.IdPromtDictamen = configloader.GetVal(envConfig.Services.Sinister.IdPromtDictamen, conf.Services.Sinister.IdPromtDictamen).(string)
	conf.Services.Sinister.ApiDiagnostic = configloader.GetVal(envConfig.Services.Sinister.ApiDiagnostic, conf.Services.Sinister.ApiDiagnostic).(string)
	conf.Services.Sinister.UpdateObservationIA = configloader.GetVal(envConfig.Services.Sinister.UpdateObservationIA, conf.Services.Sinister.UpdateObservationIA).(string)
	// Storage Api
	conf.Services.Storage.Path = configloader.GetVal(envConfig.Services.Storage.Path, conf.Services.Storage.Path).(string)
	conf.Services.Storage.Endpoint = configloader.GetVal(envConfig.Services.Storage.Endpoint, conf.Services.Storage.Endpoint).(string)
	// Notifications Api
	conf.Services.Notifications.Path = configloader.GetVal(envConfig.Services.Notifications.Path, conf.Services.Notifications.Path).(string)
	conf.Services.Notifications.Endpoint = configloader.GetVal(envConfig.Services.Notifications.Endpoint, conf.Services.Notifications.Endpoint).(string)
	// login qualitat
	conf.Services.Qualitat.QualitatUser = configloader.GetVal(envConfig.Services.Qualitat.QualitatUser, conf.Services.Qualitat.QualitatUser).(string)
	conf.Services.Qualitat.QualitatPass = configloader.GetVal(envConfig.Services.Qualitat.QualitatPass, conf.Services.Qualitat.QualitatPass).(string)
	return conf
}
