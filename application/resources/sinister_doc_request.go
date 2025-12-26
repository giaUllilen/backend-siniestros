package resources

import (
	"github.com/francoispqt/gojay"
)

type SinisterDocRequest struct {
	TipoPoliza        string              `json:"tipoPoliza"`
	Asegurado         Asegurado           `json:"asegurado"`
	Solicitante       Solicitante         `json:"solicitante"`
	Narracion         string              `json:"narracion"`
	FechaOcurrencia   string              `json:"fechaOcurrencia"`
	MontoSolicitado   string              `json:"montoSolicitado"`
	Pagador           string              `json:"pagador"`
	FechasIncapacidad []FechasIncapacidad `json:"fechasIncapacidad"`
	DeclaracionJurada bool                `json:"declaracionJurada"`
	Beneficiarios     []Beneficiario      `json:"beneficiarios"`
	Coverage          string              `json:"coverage"`
	Documento         Document            `json:"documento"`
}

type SinisterCase struct {
	Caso               string   `json:"caso,omitempty"`
	NumeroDocumento    string   `json:"numero_documento"`
	Documento          Document `json:"documento"`
	PreviousAnalysisID string   `json:"previous_analysis_id,omitempty"`
}

func (p *SinisterDocRequest) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "tipoPoliza":
		return dec.String(&p.TipoPoliza)
	case "asegurado":
		return dec.Object(&p.Asegurado)
	case "solicitante":
		return dec.Object(&p.Solicitante)
	case "narracion":
		return dec.String(&p.Narracion)
	case "fechaOcurrencia":
		return dec.String(&p.FechaOcurrencia)
	case "montoSolicitado":
		return dec.String(&p.MontoSolicitado)
	case "pagador":
		return dec.String(&p.Pagador)
	case "fechasIncapacidad":
		return dec.Array((*FechasIncapacidadArray)(&p.FechasIncapacidad))
	case "declaracionJurada":
		return dec.Bool(&p.DeclaracionJurada)
	case "beneficiarios":
		return dec.Array((*BeneficiarioArray)(&p.Beneficiarios))
	case "coverage":
		return dec.String(&p.Coverage)
	case "documento":
		return dec.Object(&p.Documento)
	}
	return nil
}

func (p *SinisterDocRequest) NKeys() int {
	return 11
}

func (p *SinisterCase) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "caso":
		return dec.String(&p.Caso)
	case "numero_documento":
		return dec.String(&p.NumeroDocumento)
	case "documento":
		return dec.Object(&p.Documento)
	case "previous_analysis_id":
		return dec.String(&p.PreviousAnalysisID)
	}
	return nil
}

func (p *SinisterCase) NKeys() int {
	return 4
}
