package colletions

import "time"

type Soat struct {
	Nombre                string    `bson:"nombre"`
	ApellidoPaterno       string    `bson:"apellidoPaterno"`
	ApellidoMaterno       string    `bson:"apellidoMaterno"`
	RazonSocial           string    `bson:"razonSocial"`
	TipoDocumento         string    `bson:"tipoDocumento"`
	NumeroDocumento       string    `bson:"numeroDocumento"`
	FechaNacimiento       time.Time `bson:"fechaNacimiento"`
	FechaInicio           time.Time `bson:"fechaInicio"`
	FechaFin              time.Time `bson:"fechaFin"`
	Moneda                string    `bson:"moneda"`
	Estado                string    `bson:"estado"`
	NumeroPoliza          string    `bson:"numeroPoliza"`
	NumeroCertificadoSoat string    `bson:"numeroCertificadoSoat"`
	Coberturas            []string  `bson:"coberturas"`
	PrimaNeta             string    `bson:"primaNeta"`
	Igv                   string    `bson:"igv"`
	PrimaTotal            string    `bson:"primaTotal"`
	FrecuenciaPago        string    `bson:"frecuenciaPago"`
	Placa                 string    `bson:"placa"`
	Plan                  string    `bson:"plan"`
	Uso                   string    `bson:"uso"`
	Clase                 string    `bson:"clase"`
	Marca                 string    `bson:"marca"`
	Modelo                string    `bson:"modelo"`
	FechaFabricacion      string    `bson:"fechaFabricacion"`
	NumeroAsientos        string    `bson:"numeroAsientos"`
	Correo                string    `bson:"correo"`
	Telefono              string    `bson:"telefono"`
	Departamento          string    `bson:"departamento"`
	Provincia             string    `bson:"provincia"`
	Distrito              string    `bson:"distrito"`
	Direccion             string    `bson:"direccion"`
	TtlName               string    `bson:"ttlName"`
	Descripcion           string    `bson:"descripcion"`
}
