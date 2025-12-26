package resources

import (
	"github.com/francoispqt/gojay"
)

type Asegurado struct {
	ApellidoPaterno string `json:"apellidoPaterno"`
	ApellidoMaterno string `json:"apellidoMaterno"`
	Nombres         string `json:"nombres"`
	NroDocumento    string `json:"nroDocumento"`
}

func (a *Asegurado) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "apellidoPaterno":
		return dec.String(&a.ApellidoPaterno)
	case "apellidoMaterno":
		return dec.String(&a.ApellidoMaterno)
	case "nombres":
		return dec.String(&a.Nombres)
	case "nroDocumento":
		return dec.String(&a.NroDocumento)
	}
	return nil
}

func (a *Asegurado) NKeys() int {
	return 4
}

type Solicitante struct {
	ApellidoPaterno string `json:"apellidoPaterno"`
	ApellidoMaterno string `json:"apellidoMaterno"`
	Nombres         string `json:"nombres"`
	NroDocumento    string `json:"nroDocumento"`
	Correo          string `json:"correo"`
	Celular         string `json:"celular"`
}

func (s *Solicitante) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "apellidoPaterno":
		return dec.String(&s.ApellidoPaterno)
	case "apellidoMaterno":
		return dec.String(&s.ApellidoMaterno)
	case "nombres":
		return dec.String(&s.Nombres)
	case "nroDocumento":
		return dec.String(&s.NroDocumento)
	case "correo":
		return dec.String(&s.Correo)
	case "celular":
		return dec.String(&s.Celular)
	}
	return nil
}

func (s *Solicitante) NKeys() int {
	return 6
}

type FechasIncapacidad struct {
	FechaInicioIncapacidad string `json:"fechaInicioIncapacidad"`
	FechaFinIncapacidad    string `json:"fechaFinIncapacidad"`
}

func (fi *FechasIncapacidad) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "fechaInicioIncapacidad":
		return dec.String(&fi.FechaInicioIncapacidad)
	case "fechaFinIncapacidad":
		return dec.String(&fi.FechaFinIncapacidad)
	}
	return nil
}

func (fi *FechasIncapacidad) NKeys() int {
	return 2
}

type FechasIncapacidadArray []FechasIncapacidad

func (arr *FechasIncapacidadArray) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var fi FechasIncapacidad
	if err := dec.Object(&fi); err != nil {
		return err
	}
	*arr = append(*arr, fi)
	return nil
}

type Beneficiario struct {
	MetodoPago   string `json:"metodoPago"`
	NroCuenta    string `json:"nroCuenta"`
	Titular      string `json:"titular"`
	NroDocumento string `json:"nroDocumento"`
	Banco        string `json:"banco"`
	Moneda       string `json:"moneda"`
	TipoCuenta   string `json:"tipoCuenta"`
}

func (b *Beneficiario) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "metodoPago":
		return dec.String(&b.MetodoPago)
	case "nroCuenta":
		return dec.String(&b.NroCuenta)
	case "titular":
		return dec.String(&b.Titular)
	case "nroDocumento":
		return dec.String(&b.NroDocumento)
	case "banco":
		return dec.String(&b.Banco)
	case "moneda":
		return dec.String(&b.Moneda)
	case "tipoCuenta":
		return dec.String(&b.TipoCuenta)
	}
	return nil
}

func (b *Beneficiario) NKeys() int {
	return 7
}

type BeneficiarioArray []Beneficiario

func (arr *BeneficiarioArray) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var b Beneficiario
	if err := dec.Object(&b); err != nil {
		return err
	}
	*arr = append(*arr, b)
	return nil
}

type Document struct {
	Name     string `json:"name"`
	Filename string `json:"filename"`
	FileURL  string `json:"file_url"`
}

func (d *Document) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "name":
		return dec.String(&d.Name)
	case "filename":
		return dec.String(&d.Name)
	case "file_url":
		return dec.String(&d.FileURL)
	}
	return nil
}

func (d *Document) NKeys() int {
	return 3
}

type DocumentArray []Document

func (arr *DocumentArray) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var d Document
	if err := dec.Object(&d); err != nil {
		return err
	}
	*arr = append(*arr, d)
	return nil
}

type DocumentSection struct {
	Coverage   string     `json:"coverage"`
	Documents  []Document `json:"documents"`
	Additional []Document `json:"additional"`
}

func (ds *DocumentSection) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "coverage":
		return dec.String(&ds.Coverage)
	case "documents":
		return dec.Array((*DocumentArray)(&ds.Documents))
	case "additional":
		return dec.Array((*DocumentArray)(&ds.Additional))
	}
	return nil
}

func (ds *DocumentSection) NKeys() int {
	return 3
}

type DocumentSectionArray []DocumentSection

func (arr *DocumentSectionArray) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var ds DocumentSection
	if err := dec.Object(&ds); err != nil {
		return err
	}
	*arr = append(*arr, ds)
	return nil
}

type SinisterRequest struct {
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
	Documentos        []DocumentSection   `json:"documentos"`
}

func (p *SinisterRequest) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
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
	case "documentos":
		return dec.Array((*DocumentSectionArray)(&p.Documentos))
	}
	return nil
}

func (p *SinisterRequest) NKeys() int {
	return 11
}
