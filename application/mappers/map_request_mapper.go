package mappers

import (
	"is-public-api/application/resources"
)

func MapStringInterfaceToMapRequest(model map[string]interface{}) *resources.MapRequest {
	aux := resources.MapRequest(model)
	return &aux
}

func MapStringToMapRequest(model map[string]string) *resources.MapString {
	aux := resources.MapString(model)
	return &aux
}
