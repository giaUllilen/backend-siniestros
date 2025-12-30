package colletions

type SoatReturn struct {
	Poliza               string `bson:"poliza"`
	InicioVigencia       string `bson:"inicio_vigencia"`
	FinVigencia          string `bson:"fin_vigencia"`
	FechaEmisionAcsele   string `bson:"fecha_emision_acsele"`
	Estado               string `bson:"estado"`
	Canal                string `bson:"canal"`
	Uso                  string `bson:"uso"`
	Placa                string `bson:"placa"`
	TipoDocumento        string `bson:"tipo_documento"`
	NroDocumento         string `bson:"nro_documento"`
	NombreContr          string `bson:"nombre_contr"`
	FechaVenta           string `bson:"fecha_venta"`
	Prima                string `bson:"prima"`
	Primadevuelta        string `bson:"primadevuelta"`
	Primadevueltapagadas string `bson:"primadevueltapagadas"`
	Correo               string `bson:"correo"`
}