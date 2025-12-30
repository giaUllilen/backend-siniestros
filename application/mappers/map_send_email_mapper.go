package mappers

import (
	"fmt"
	"strings"
)

type SendEmail struct {
}

func SendEmailToMapRequest(subDomain string, tipoPoliza string, caseID string, now string, solicitante, asegurado map[string]interface{}) map[string]interface{} {

	insuredName := "Asegurado"
	days := "8"
	producto := strings.ToLower(tipoPoliza)
	if producto == "soat" {
		insuredName = "Lesionado"
	}
	if producto == "protección tarjeta" || producto == "protección financiera" || producto == "renta hospitalaria" {
		days = "2"
	}
	data := map[string]interface{}{
		"name":         strings.Split(solicitante["nombres"].(string), " ")[0],
		"insured_name": insuredName,
		"insured":      fmt.Sprint(strings.Split(asegurado["nombres"].(string), " ")[0], " ", asegurado["apellidoPaterno"]),
		"author":       fmt.Sprint(strings.Split(solicitante["nombres"].(string), " ")[0], " ", solicitante["apellidoPaterno"]),
		"case":         fmt.Sprint(caseID),
		"date":         now,
		"subdomain":    subDomain,
		"days":         days,
		"file":         "https://" + subDomain + ".interseguro.pe/siniestros",
	}
	return data
}
