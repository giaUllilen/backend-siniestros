package mappers

import (
	"is-public-api/application/models"
	"is-public-api/application/resources"
)

func ModelToSoatReturnResponse(model *models.SoatReturn, res *resources.MapResponse) {
	*res = make(resources.MapResponse)
	(*res)["estado"]=model.Estado
	(*res)["poliza"]=model.Poliza
	(*res)["canal"]=model.Canal
	(*res)["uso"]=model.Uso
	(*res)["placa"]=model.Placa
	(*res)["tipo_documento"]=model.TipoDocumento
	(*res)["nro_documento"]=model.NroDocumento
	(*res)["nombre_contr"]=model.NombreContr
}
