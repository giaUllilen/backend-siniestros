package models

import "time"

type Soat struct {
	Nombre                string
	ApellidoPaterno       string
	ApellidoMaterno       string
	RazonSocial           string
	TipoDocumento         string
	NumeroDocumento       string
	FechaNacimiento       time.Time
	FechaInicio           time.Time
	FechaFinVigencia      time.Time
	FechaFin              time.Time
	Moneda                string
	NumeroPoliza          string
	NumeroCertificadoSoat string
	Coberturas            []string
	PrimaNeta             string
	Estado                string
	Igv                   string
	PrimaTotal            string
	FrecuenciaPago        string
	Placa                 string
	Plan                  string
	Uso                   string
	Clase                 string
	Marca                 string
	Modelo                string
	FechaFabricacion      string
	NumeroAsientos        string
	Correo                string
	Telefono              string
	Departamento          string
	Provincia             string
	Distrito              string
	Direccion             string
	TtlName               string
	Descripcion           string
}
