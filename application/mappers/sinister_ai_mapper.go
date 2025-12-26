package mappers

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"is-public-api/application/colletions"
	"is-public-api/application/resources"
)

func SinisterToMapResponse(c colletions.SinisterCase) resources.MapResponse {
	result := resources.MapResponse{
		"case":                  c.Case,
		"placa":                 c.Placa,
		"fecha":                 c.CreatedDate,
		"lesionado":             MapperBsonToMapResponse(c.Lesionado),
		"soat":                  MapperBsonToMapResponse(c.Soat),
		"denuncia":              MapperBsonToMapResponse(c.Denuncia),
		"denuncia_documento":    MapperBsonToMapResponse(c.DenunciaDocumento),
		"certificado":           MapperBsonToMapResponse(c.Certificado),
		"certificado_documento": MapperBsonToMapResponse(c.CertificadoDocumento),
		"observations":          MapperBsonToMapResponse(c.Observations),
		"numero_documento":      c.NumeroDocumento,
	}
	return result
}

func MapperBsonToMapResponse(bsonMap map[string]interface{}) resources.MapResponse {
	result := make(resources.MapResponse)

	for key, value := range bsonMap {
		switch v := value.(type) {
		case map[string]interface{}:
			result[key] = MapperBsonToMapResponse(v)
		case bson.M:
			result[key] = MapperBsonToMapResponse(v)
		case bson.A:
			result[key] = MapperBsonToListArray(v)
		case primitive.DateTime:
			result[key] = v.Time()
		default:
			result[key] = v
		}
	}

	return result
}

func MapperBsonToListArray(bsonArray bson.A) resources.ListArray {
	result := make(resources.ListArray, len(bsonArray))

	for i, value := range bsonArray {
		switch v := value.(type) {
		case map[string]interface{}:
			result[i] = MapperBsonToMapResponse(v)
		case bson.M:
			result[i] = MapperBsonToMapResponse(v)
		case bson.A:
			result[i] = MapperBsonToListArray(v)
		default:
			result[i] = v
		}
	}

	return result
}
