package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SinisterCoverageIA representa la estructura de la colección siniestro_coberturas_ia
type SinisterCoverageIA struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Producto        string             `bson:"producto" json:"producto"`
	Cobertura       string             `bson:"cobertura" json:"cobertura"`
	UsuarioCreacion string             `bson:"usuario_creacion" json:"usuario_creacion"`
	FechaCreacion   string             `bson:"fecha_creacion" json:"fecha_creacion"`
	Activo          bool               `bson:"activo" json:"activo"`
}

// SinisterCoverageIAFilter representa los filtros para consultas
type SinisterCoverageIAFilter struct {
	Producto  *string `bson:"producto,omitempty" json:"producto,omitempty"`
	Cobertura *string `bson:"cobertura,omitempty" json:"cobertura,omitempty"`
	Activo    *bool   `bson:"activo,omitempty" json:"activo,omitempty"`
}
