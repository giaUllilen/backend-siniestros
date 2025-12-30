package apihelpers

type EndpointConfig struct {
	CloudFunction    Endpoint
	EventLog         EventLog
	Sinister         Sinister
	Storage          Service
	Notification     Service
	QualitatSinister QualitatSinister
	Qualitat         Qualitat
}

type Endpoint struct {
	Uri string
}

type Service struct {
	Path     string
	Endpoint string
}

type Qualitat struct {
	QualitatUser string
	QualitatPass string
}

type EventLog struct {
	Uri       string
	ApiKey    string
	ApiKeyWsp string
}

type Sinister struct {
	Path                string
	Save                string
	Get                 string
	Document            string
	History             string
	Delete              string
	UrlGenAI            string
	TokenGenAI          string
	IdPromtDP           string
	IdPromtDM           string
	IdPromtDictamen     string
	ApiDiagnostic       string
	UpdateObservationIA string
}

type QualitatSinister struct {
	Document string
}

type EndpointHelper struct {
	Conf *EndpointConfig
}

func NewEndpointHelper(conf *EndpointConfig) *EndpointHelper {
	return &EndpointHelper{Conf: conf}
}
