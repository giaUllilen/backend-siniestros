package mappers

import (
	"is-public-api/application/models"
	"is-public-api/application/resources"
)

func ModelToSinisterPaymentResponse(model *models.SinisterPayment, res *resources.MapResponse) {
	*res = make(resources.MapResponse)
	(*res)["numero_documento"]=model.NumeroDocumento
	(*res)["beneficiario"]=model.Beneficiario
	(*res)["producto"]=model.Producto
	(*res)["moneda"]=model.Moneda
	(*res)["monto_pagado"]=model.MontoPagado
	(*res)["certificado"]=model.Certificado
	(*res)["num_siniestro"]=model.NumSiniestro
	(*res)["fecha_siniestro"]=model.FechaSiniestro
	(*res)["fecha_pago"]=model.FechaPago
	(*res)["cobertura"]=model.Cobertura
}