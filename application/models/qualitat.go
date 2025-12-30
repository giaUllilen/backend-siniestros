package models

// QualitatResponse representa la estructura de respuesta de Qualitat
type QualitatResponse struct {
	FechaSiniestro      string `json:"fecha_siniestro"`
	NroSiniestroCliente string `json:"nro_siniestro_cliente"`
	NroSiniestro        string `json:"nro_siniestro"`
	NroPoliza           string `json:"nro_poliza"`
	Placa               string `json:"placa"`
	NroCaso             string `json:"nro_caso"`
	Nombres             string `json:"nombres"`
	ApPaterno           string `json:"ap_paterno"`
	ApMaterno           string `json:"ap_materno"`
	DNI                 string `json:"dni"`
	CentroMedico        string `json:"centro_medico"`
	TipoSiniestro       string `json:"tipo_siniestro"`
	Recepcion           string `json:"recepcion"`
	Ocupante            string `json:"ocupante"`
	Fallecido           string `json:"fallecido"`
	FechaFallecimiento  string `json:"fecha_fallecimiento"`
	EstadoSiniestro     string `json:"estado_siniestro"`
}
