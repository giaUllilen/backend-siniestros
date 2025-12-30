package colletions

import "time"

// Constantes de colecciones MongoDB
const (
	CollectionSinisterCoverageIA = "siniestro_coberturas_ia"
)

type SinisterCase struct {
	Case                 string                 `bson:"case"`
	Placa                string                 `bson:"placa"`
	Lesionado            map[string]interface{} `bson:"lesionado"`
	Soat                 map[string]interface{} `bson:"soat"`
	CreatedDate          time.Time              `bson:"created_date"`
	NumeroDocumento      string                 `bson:"numero_documento"`
	Denuncia             map[string]interface{} `bson:"denuncia"`
	DenunciaDocumento    map[string]interface{} `bson:"denuncia_documento"`
	Certificado          map[string]interface{} `bson:"certificado"`
	CertificadoDocumento map[string]interface{} `bson:"certificado_documento"`
	Observations         map[string]interface{} `bson:"observations"`
}
