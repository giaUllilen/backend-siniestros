package services

import (
	"is-public-api/application/models"
	"is-public-api/application/storage"
)

type soatFinder struct {
	soatRepository storage.ISoatRepository
}

func NewSoatFinder(repo storage.ISoatRepository) *soatFinder {
	return &soatFinder{soatRepository: repo}
}

func (service *soatFinder) FindByPlate(txContext *models.TxContext, plate string) (*models.Soat, error) {
	data, err := service.soatRepository.FindByPlate(txContext, plate)
	if err != nil {
		return nil, err
	}

	soat := &models.Soat{
		Nombre:                data.Nombre,
		ApellidoPaterno:       data.ApellidoPaterno,
		ApellidoMaterno:       data.ApellidoMaterno,
		RazonSocial:           data.RazonSocial,
		TipoDocumento:         data.TipoDocumento,
		NumeroDocumento:       data.NumeroDocumento,
		FechaNacimiento:       data.FechaNacimiento,
		FechaInicio:           data.FechaInicio,
		FechaFin:              data.FechaFin,
		Moneda:                data.Moneda,
		Estado:                data.Estado,
		NumeroPoliza:          data.NumeroPoliza,
		NumeroCertificadoSoat: data.NumeroCertificadoSoat,
		Coberturas:            data.Coberturas,
		PrimaNeta:             data.PrimaNeta,
		Igv:                   data.Igv,
		PrimaTotal:            data.PrimaTotal,
		FrecuenciaPago:        data.FrecuenciaPago,
		Placa:                 data.Placa,
		Plan:                  data.Plan,
		Uso:                   data.Uso,
		Clase:                 data.Clase,
		Marca:                 data.Marca,
		Modelo:                data.Modelo,
		FechaFabricacion:      data.FechaFabricacion,
		NumeroAsientos:        data.NumeroAsientos,
		Correo:                data.Correo,
		Telefono:              data.Telefono,
		Departamento:          data.Departamento,
		Provincia:             data.Provincia,
		Distrito:              data.Distrito,
		Direccion:             data.Direccion,
		TtlName:               data.TtlName,
		Descripcion:           data.Descripcion,
	}

	return soat, nil
}

func (service *soatFinder) FindByPlateHistory(txContext *models.TxContext, plate string, dateOccurrence string) (*models.Soat, error) {
	data, err := service.soatRepository.FindByPlateHistory(txContext, plate, dateOccurrence)
	if err != nil {
		return nil, err
	}

	soat := &models.Soat{
		Nombre:                data.Nombre,
		ApellidoPaterno:       data.ApellidoPaterno,
		ApellidoMaterno:       data.ApellidoMaterno,
		RazonSocial:           data.RazonSocial,
		TipoDocumento:         data.TipoDocumento,
		NumeroDocumento:       data.NumeroDocumento,
		FechaNacimiento:       data.FechaNacimiento,
		FechaInicio:           data.FechaInicio,
		FechaFin:              data.FechaFin,
		Moneda:                data.Moneda,
		Estado:                data.Estado,
		NumeroPoliza:          data.NumeroPoliza,
		NumeroCertificadoSoat: data.NumeroCertificadoSoat,
		Coberturas:            data.Coberturas,
		PrimaNeta:             data.PrimaNeta,
		Igv:                   data.Igv,
		PrimaTotal:            data.PrimaTotal,
		FrecuenciaPago:        data.FrecuenciaPago,
		Placa:                 data.Placa,
		Plan:                  data.Plan,
		Uso:                   data.Uso,
		Clase:                 data.Clase,
		Marca:                 data.Marca,
		Modelo:                data.Modelo,
		FechaFabricacion:      data.FechaFabricacion,
		NumeroAsientos:        data.NumeroAsientos,
		Correo:                data.Correo,
		Telefono:              data.Telefono,
		Departamento:          data.Departamento,
		Provincia:             data.Provincia,
		Distrito:              data.Distrito,
		Direccion:             data.Direccion,
		TtlName:               data.TtlName,
		Descripcion:           data.Descripcion,
	}

	return soat, nil
}
