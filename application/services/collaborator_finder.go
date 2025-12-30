package services

import (
	"encoding/base64"

	"is-public-api/application/models"
	"is-public-api/application/storage"
)

type collaboratorFinder struct {
	collaboratorRepository storage.ICollaboratorRepository
}

func NewCollaboratorFinder(repo storage.ICollaboratorRepository) *collaboratorFinder {
	return &collaboratorFinder{collaboratorRepository: repo}
}

func (service *collaboratorFinder) Find(txContext *models.TxContext, code string) (*models.Collaborator, error) {
	base, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return nil, err
	}
	data, err := service.collaboratorRepository.Find(txContext, string(base))
	if err != nil {
		return nil, err
	}

	collaborator := &models.Collaborator{
		IdPersona: data.IdPersona,
		Nombres: data.Nombres,
		ApellidoPaterno: data.ApellidoPaterno,
		ApellidoMaterno: data.ApellidoMaterno,
		Apellidos: data.Apellidos,
		Puesto: data.Puesto,
		DescripcionPuesto: data.DescripcionPuesto,
		CodigoArea: data.CodigoArea,
		DescripcionArea: data.DescripcionArea,
		TipoDocumento: data.TipoDocumento,
		CodigoTipoDocumento: data.CodigoTipoDocumento,
		DocumentoIdentidad: data.DocumentoIdentidad,
		CodigoIS: data.CodigoIS,
		FechaIngreso: data.FechaIngreso,
		FechaCese: data.FechaCese,
		Estado: data.Estado,
		CodigoVicepresidencia: data.CodigoVicepresidencia,
		Vicepresidencia: data.Vicepresidencia,
	}

	return collaborator, nil
}
