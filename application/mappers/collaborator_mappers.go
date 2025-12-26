package mappers

import (
	"strings"

	"is-public-api/application/models"
	"is-public-api/application/resources"
)

func ModelToCollaboratorResponse(model *models.Collaborator, res *resources.MapResponse) {
	*res = make(resources.MapResponse)
	(*res)["estado"]=model.Estado
	if strings.ToLower(model.Estado) == "activo" {
		(*res)["id_persona"] = model.IdPersona
		(*res)["nombres"] = model.Nombres
		(*res)["apellido_paterno"] = model.ApellidoPaterno
		(*res)["apellido_materno"] = model.ApellidoMaterno
		(*res)["apellidos"] = model.Apellidos
		(*res)["puesto"] = model.Puesto
		(*res)["descripcion_puesto"] = model.DescripcionPuesto
		(*res)["codigo_area"] = model.CodigoArea
		(*res)["descripcion_area"] = model.DescripcionArea
		if model.CodigoTipoDocumento == "02" {
			(*res)["tipo_documento"] = "CE"
		} else {
			(*res)["tipo_documento"] = "DNI"
		}
		(*res)["numero_documento"] = model.DocumentoIdentidad
		(*res)["codigo_is"] = model.CodigoIS
		(*res)["fecha_ingreso"] = model.FechaIngreso
		(*res)["fecha_cese"] = model.FechaIngreso
		(*res)["codigo_vicepresidencia"] = model.CodigoVicepresidencia
		(*res)["vicepresidencia"] = model.Vicepresidencia
	}
}