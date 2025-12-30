package models

type SinisterPayment struct {
	NumeroDocumento string 
	Beneficiario    string 
	Producto        string 
	Moneda          string 
	MontoPagado     string 
	Certificado     string 
	NumSiniestro    string 
	FechaSiniestro  string 
	FechaPago       string 
	Cobertura       string 
}
