package models

import (
	"time"
)

type Collaborator struct {
	IdPersona             string    
	Nombres               string    
	ApellidoPaterno       string    
	ApellidoMaterno       string    
	Apellidos             string
	Puesto                string    
	DescripcionPuesto     string    
	CodigoArea            string    
	DescripcionArea       string    
	TipoDocumento         string    
	CodigoTipoDocumento   string    
	DocumentoIdentidad    string    
	CodigoIS              string    
	FechaIngreso          time.Time 
	FechaCese             time.Time 
	Estado                string    
	CodigoVicepresidencia string    
	Vicepresidencia       string    
}
