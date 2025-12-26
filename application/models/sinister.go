package models

type Sinister struct {
	TipoPoliza        string
	Asegurado         AseguradoSinister
	Solicitante       SolicitanteSinister
	Narracion         string
	FechaOcurrencia   string
	MontoSolicitado   string
	Pagador           string
	FechasIncapacidad []FechasIncapacidadSinister
	DeclaracionJurada bool
	Beneficiarios     []BeneficiarioSinister
	Documentos        []DocumentSectionSinister
}

type AseguradoSinister struct {
	ApellidoPaterno string
	ApellidoMaterno string
	Nombres         string
	NroDocumento    string
}

type SolicitanteSinister struct {
	ApellidoPaterno string
	ApellidoMaterno string
	Nombres         string
	NroDocumento    string
	Correo          string
	Celular         string
}

type FechasIncapacidadSinister struct {
	FechaInicioIncapacidad string
	FechaFinIncapacidad    string
}

type BeneficiarioSinister struct {
	MetodoPago   string
	NroCuenta    string
	Titular      string
	NroDocumento string
	Banco        string
	Moneda       string
	TipoCuenta   string
}

type DocumentSinister struct {
	Name     string
	Filename string
	FileURL  string
}

type DocumentSectionSinister struct {
	Coverage   string
	Documents  []DocumentSinister
	Additional []DocumentSinister
}
