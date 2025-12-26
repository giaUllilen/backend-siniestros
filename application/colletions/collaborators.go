package colletions

import "time"

type Colaborador struct {
	IdPersona             string    `bson:"idPersona"`
	Nombres               string    `bson:"nombres"`
	ApellidoPaterno       string    `bson:"apellidoPaterno"`
	ApellidoMaterno       string    `bson:"apellidoMaterno"`
	Apellidos             string    `bson:"apellidos"`
	Puesto                string    `bson:"puesto"`
	DescripcionPuesto     string    `bson:"descripcionPuesto"`
	CodigoArea            string    `bson:"codigoArea"`
	DescripcionArea       string    `bson:"descripcionArea"`
	TipoDocumento         string    `bson:"tipoDocumento"`
	CodigoTipoDocumento   string    `bson:"codigoTipoDocumento"`
	DocumentoIdentidad    string    `bson:"documentoIdentidad"`
	CodigoIS              string    `bson:"codigoIS"`
	FechaIngreso          time.Time `bson:"fechaIngreso"`
	FechaCese             time.Time `bson:"fechaCese"`
	Estado                string    `bson:"estado"`
	CodigoVicepresidencia string    `bson:"codigoVicepresidencia"`
	Vicepresidencia       string    `bson:"viceprecidencia"`
}
