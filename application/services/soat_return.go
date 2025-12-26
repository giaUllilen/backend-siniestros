package services

import (
	"is-public-api/application/models"
	"is-public-api/application/storage"
)

type soatReturnFinder struct {
	soatReturnRepository storage.ISoatReturnRepository
}

func NewSoatReturnFinder(repo storage.ISoatReturnRepository) *soatReturnFinder {
	return &soatReturnFinder{soatReturnRepository: repo}
}

func (service *soatReturnFinder) FindByDocument(txContext *models.TxContext, documentNumber string) (*models.SoatReturn, error) {
	data, err := service.soatReturnRepository.FindByDocument(txContext, string(documentNumber))
	if err != nil {
		return nil, err
	}

	soatReturn := &models.SoatReturn{
		Poliza:               data.Poliza,
		InicioVigencia:       data.InicioVigencia,
		FinVigencia:          data.FinVigencia,
		FechaEmisionAcsele:   data.FechaEmisionAcsele,
		Estado:               data.Estado,
		Canal:                data.Canal,
		Uso:                  data.Uso,
		Placa:                data.Placa,
		TipoDocumento:        data.TipoDocumento,
		NroDocumento:         data.NroDocumento,
		NombreContr:          data.NombreContr,
		FechaVenta:           data.FechaVenta,
		Prima:                data.Prima,
		Primadevuelta:        data.Primadevuelta,
		Primadevueltapagadas: data.Primadevueltapagadas,
		Correo:               data.Correo,
	}

	return soatReturn, nil
}
