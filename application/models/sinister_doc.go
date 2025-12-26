package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SinisterDoc struct {
	TipoPoliza        string                      `bson:"tipo_poliza"`
	Asegurado         AseguradoSinister           `bson:"asegurado"`
	Solicitante       SolicitanteSinister         `bson:"solicitante"`
	Narracion         string                      `bson:"narracion"`
	FechaOcurrencia   string                      `bson:"fecha_ocurrencia"`
	MontoSolicitado   string                      `bson:"monto_solicitado"`
	Pagador           string                      `bson:"pagador"`
	FechasIncapacidad []FechasIncapacidadSinister `bson:"fechas_incapacidad"`
	DeclaracionJurada bool                        `bson:"declaracion_jurada"`
	Beneficiarios     []BeneficiarioSinister      `bson:"beneficiarios"`
	Coverage          string                      `bson:"coverage"`
	Documento         DocumentSinister            `bson:"documento"`
}

type Block struct {
	Confidence float64 `json:"confidence" bson:"confidence"`
	Text       string  `json:"text" bson:"text"`
}

type Page struct {
	Blocks []Block `json:"blocks" bson:"blocks"`
}

type DocumentContent struct {
	Pages []Page `json:"pages" bson:"pages"`
}

type Lesionado struct {
	ApellidoMaterno string       `json:"apellidoMaterno" bson:"apellido_materno"`
	ApellidoPaterno string       `json:"apellidoPaterno" bson:"apellido_paterno"`
	Conductor       bool         `json:"conductor" bson:"conductor"`
	Nombres         string       `json:"nombres" bson:"nombres"`
	ListaDocumentos []Documentos `json:"lista_documentos" bson:"lista_documentos"`
	Edad            int32        `json:"edad" bson:"edad"`
	Estado          string       `json:"estado" bson:"estado"`
}

type Documentos struct {
	Numero string `json:"numero" bson:"numero"`
	Tipo   string `json:"tipo" bson:"tipo"`
}

// Estructuras para updateObservationIA
type ObservationIARequest struct {
	NumeroCaso      string          `json:"numero_caso"`
	ObservacionesIA ObservacionesIA `json:"observaciones_ia"`
}

type ObservacionesIA struct {
	Confidence   float64         `json:"confidence"`
	Dictamen     DictamenIA      `json:"dictamen"`
	Observations []ObservationIA `json:"observations"`
}

type DictamenIA struct {
	Justificacion string `json:"justificacion"`
	Status        string `json:"status"`
}

type ObservationIA struct {
	Classification string `json:"classification"`
	Text           string `json:"text"`
}

type Vehiculo struct {
	Lesionados []Lesionado `json:"lesionados" bson:"lesionados"`
	Placa      string      `json:"placa" bson:"placa"`
}

type Validacion struct {
	FDO           string `json:"fdo"`
	FirmasSellos  bool   `json:"firmas_sellos"`
	Interviniente string `json:"interviniente"`
}

type Answer struct {
	FechaOcurrencia   string     `json:"fechaOcurrencia" bson:"fecha_ocurrencia"`
	Summary           string     `json:"summary" bson:"summary"`
	Validacion        Validacion `json:"validacion"`
	Vehiculos         []Vehiculo `json:"vehiculos" bson:"vehiculos"`
	DiasTranscurridos int32      `json:"dias_transcurridos" bson:"dias_transcurridos"`
}

type DataDenunciaPolicialResponseGenIA struct {
	Caso             string            `json:"caso,omitempty" bson:"caso,omitempty"`
	Answer           Answer            `json:"answer" bson:"answer"`
	AnswerID         string            `json:"answerId" bson:"answer_id"`
	ObserveDocuments *ObserveDocuments `json:"observe_documents" bson:"observe_documents"`
}

type ObserveDocuments struct {
	InformeMedico *ObserveDocument `json:"informe_medico" bson:"informe_medico"`
	DocumentMenor *ObserveDocument `json:"document_menor" bson:"document_menor"`
	DocumentPadre *ObserveDocument `json:"document_padre" bson:"document_padre"`
	DocumentMadre *ObserveDocument `json:"document_madre" bson:"document_madre"`
	DP            *ObserveDocument `json:"denuncia_policial" bson:"denuncia_policial"`
	DM            *ObserveDocument `json:"descanso_medico" bson:"descanso_medico"`
}

type ObserveDocument struct {
	Nombre     string `bson:"nombre"`
	Path       string `bson:"path,omitempty"`
	Requerido  bool   `bson:"requirido"`
	Completado bool   `bson:"completado"`
}

type ReportDM struct {
	FechaInicioIncapacidad string      `json:"fecha_inicio_incapacidad" bson:"fecha_inicio_incapacidad"`
	FechaFinIncapacidad    string      `json:"fecha_fin_incapacidad" bson:"fecha_fin_incapacidad"`
	DiasReposo             string      `json:"dias_reposo" bson:"dias_reposo"`
	Lesionado              LesionadoDM `json:"lesionado" bson:"lesionado"`
	Medico                 MedicoDM    `json:"medico" bson:"medico"`
	Summary                string      `json:"summary" bson:"summary"`
}

type LesionadoDM struct {
	Nombres           string `json:"nombres" bson:"nombres"`
	ApellidoPaterno   string `json:"apellidoPaterno" bson:"apellido_paterno"`
	ApellidoMaterno   string `json:"apellidoMaterno" bson:"apellido_materno"`
	NumeroDocumento   string `json:"numeroDocumento" bson:"numero_documento"`
	DiagnosticoMedico string `json:"diagnostico_medico" bson:"diagnostico_medico"`
}

type MedicoDM struct {
	Nombre       string `json:"nombre" bson:"nombre"`
	CMP          string `json:"cmp" bson:"cmp"`
	Especialidad string `json:"especialidad" bson:"especialidad"`
	Firma        bool   `json:"firma" bson:"firma"`
	Sello        bool   `json:"sello" bson:"sello"`
}

type ReportsDM struct {
	Reports []ReportDM `json:"reports" bson:"reports"`
}

type DataDescansoMedicoResponseGenIA struct {
	Answer           ReportsDM         `json:"answer" bson:"answer"`
	AnswerID         string            `json:"answerId" bson:"answer_id"`
	ObserveDocuments *ObserveDocuments `json:"observe_documents" bson:"observe_documents"`
}

type SinisterPreviousCases struct {
	Document string `json:"document" bson:"document"`
}

type Documento struct {
	Numero string `json:"numero" bson:"numero"`
	Tipo   string `json:"tipo" bson:"tipo"`
}

type Denuncia struct {
	Caso              string      `json:"caso" bson:"caso"`
	Answer            Answer      `json:"answer" bson:"answer"`
	AnswerID          string      `json:"answerid" bson:"answerid"`
	DocumentOptionals interface{} `json:"documentoptionals" bson:"documentoptionals"`
}

type DenunciaDocumento struct {
	Name     string `json:"name" bson:"name"`
	FileName string `json:"filename" bson:"filename"`
	FileURL  string `json:"fileurl" bson:"fileurl"`
}

type SoatCasoSiniestro struct {
	NumeroPoliza string      `json:"numero_poliza" bson:"numero_poliza"`
	FechaInicio  time.Time   `json:"fecha_inicio" bson:"fecha_inicio"`
	FechaFin     time.Time   `json:"fecha_fin" bson:"fecha_fin"`
	Estado       string      `json:"estado" bson:"estado"`
	Coberturas   interface{} `json:"coberturas" bson:"coberturas"`
}

type CasoSiniestro struct {
	ID                primitive.ObjectID  `json:"_id" bson:"_id"`
	CreatedDate       time.Time           `json:"created_date" bson:"created_date"`
	Denuncia          Denuncia            `json:"denuncia" bson:"denuncia"`
	DenunciaDocumento DenunciaDocumento   `json:"denuncia_documento" bson:"denuncia_documento"`
	Lesionado         Lesionado           `json:"lesionado" bson:"lesionado"`
	Placa             string              `json:"placa" bson:"placa"`
	NumeroDocumento   string              `json:"numero_documento" bson:"numero_documento"`
	DataQualitat      DataQualitat        `json:"qualitat" bson:"qualitat"`
	Soat              SoatCasoSiniestro   `json:"soat" bson:"soat"`
	Observation       *ObservationWrapper `json:"observaciones" bson:"observaciones"`
	ObserveDocuments  *ObserveDocuments   `json:"observe_documents" bson:"observe_documents"`
}

type DataQualitat struct {
	EstadoSiniestro        string `json:"estado_siniestro" bson:"estado_siniestro"`
	CentroMedico           string `json:"centro_medico" bson:"centro_medico"`
	NumeroSiniestro        string `json:"numero_siniestro" bson:"numero_siniestro"`
	NumeroSiniestroCliente string `json:"numero_siniestro_cliente" bson:"numero_siniestro_cliente"`
	NumeroPoliza           string `json:"numero_poliza" bson:"numero_poliza"`
	TipoSiniestro          string `json:"tipo_siniestro" bson:"tipo_siniestro"`
}

type SinisterHistory struct {
	NumeroSiniestro        string              `bson:"NUMERO_SINIESTRO" json:"NUMERO_SINIESTRO"`
	NumeroCaso             string              `bson:"NUMERO_CASO" json:"NUMERO_CASO"`
	FechaOcurrencia        string              `bson:"FECHA_OCURRENCIA" json:"FECHA_OCURRENCIA"`
	MontoPagado            string              `bson:"MONTO_PAGADO" json:"MONTO_PAGADO"`
	EstadoDictamen         string              `bson:"ESTADO_DICTAMEN" json:"ESTADO_DICTAMEN"`
	FechaSolicitud         string              `bson:"FECHA_SOLICITUD" json:"FECHA_SOLICITUD"`
	FechaInicioIncapacidad string              `bson:"FECHA_INICIO_INCAPACIDAD" json:"FECHA_INICIO_INCAPACIDAD"`
	FechaFinIncapacidad    string              `bson:"FECHA_FIN_INCAPACIDAD" json:"FECHA_FIN_INCAPACIDAD"`
	FechasIncapacidad      []FechasIncapacidad `bson:"FECHAS_INCAPACIDAD" json:"FECHAS_INCAPACIDAD"`
}

type FechasIncapacidad struct {
	FechaInicioIncapacidad string `bson:"FECHA_INICIO_INCAPACIDAD" json:"FECHA_INICIO_INCAPACIDAD"`
	FechaFinIncapacidad    string `bson:"FECHA_FIN_INCAPACIDAD" json:"FECHA_FIN_INCAPACIDAD"`
}

type SinisterHistoryRequest struct {
	Filtro             []string `json:"filtro"`
	CurrentPage        int      `json:"currentPage"`
	PageSize           int      `json:"pageSize"`
	EstadoAprobacion   string   `json:"ESTADO_APROBACION"`
	EstadoDictamen     string   `json:"ESTADO_DICTAMEN"`
	FechaDesde         string   `json:"FECHA_DESDE"`
	FechaHasta         string   `json:"FECHA_HASTA"`
	FechaDictamenDesde string   `json:"FECHA_DICTAMEN_DESDE"`
	FechaDictamenHasta string   `json:"FECHA_DICTAMEN_HASTA"`
	FechaIngresoDesde  string   `json:"FECHA_INGRESO_DESDE"`
	FechaIngresoHasta  string   `json:"FECHA_INGRESO_HASTA"`
	NumeroDocumento    string   `json:"NUMERO_DOCUMENTO"`
	NumeroTicket       string   `json:"NUMERO_TICKET"`
	Producto           string   `json:"PRODUCTO"`
	TipoDocumento      string   `json:"TIPO_DOCUMENTO"`
	NumeroSiniestro    string   `json:"NUMERO_SINIESTRO"`
}

// / OBSERVACIONES
type ObservationData struct {
	Classification string `json:"classification"`
	Text           string `json:"text"`
}

type ObservationDictamen struct {
	Justificacion string `json:"justificacion"`
	Status        string `json:"status"`
}

type ObservationAnswer struct {
	Dictamen     ObservationDictamen `json:"dictamen"`
	Observations []ObservationData   `json:"observations"`
}

type ObservationDataAnswer struct {
	Answer ObservationAnswer `json:"answer"`
}

type ObservationWrapper struct {
	Data ObservationDataAnswer `json:"data"`
}
