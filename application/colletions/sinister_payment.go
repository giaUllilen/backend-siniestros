package colletions

type SinisterPayment struct {
	NumeroDocumento string `bson:"numero_documento"`
	Beneficiario    string `bson:"beneficiario"`
	Producto        string `bson:"producto"`
	Moneda          string `bson:"moneda"`
	MontoPagado     string `bson:"monto_pagado"`
	Certificado     string `bson:"certificado"`
	NumSiniestro    string `bson:"num_siniestro"`
	FechaSiniestro  string `bson:"fecha_siniestro"`
	FechaPago       string `bson:"fecha_pago"`
	Cobertura       string `bson:"cobertura"`
}