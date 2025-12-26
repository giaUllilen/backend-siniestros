package services

import (
	"is-public-api/application/models"
	"is-public-api/application/storage"
)

type sinisterPaymentFinder struct {
	sinisterPaymentRepository storage.ISinisterPaymentRepository
}

func NewSinisterPaymentFinder(repo storage.ISinisterPaymentRepository) *sinisterPaymentFinder {
	return &sinisterPaymentFinder{sinisterPaymentRepository: repo}
}

func (service *sinisterPaymentFinder) FindByDocumentNumber(txContext *models.TxContext, documentNumber string) (*models.SinisterPayment, error) {
	data, err := service.sinisterPaymentRepository.FindByDocumentNumber(txContext, string(documentNumber))
	if err != nil {
		return nil, err
	}

	payment := &models.SinisterPayment{
		NumeroDocumento: data.NumeroDocumento,
		Beneficiario:    data.Beneficiario,
		Producto:        data.Producto,
		Moneda:          data.Moneda,
		MontoPagado:     data.MontoPagado,
		Certificado:     data.Certificado,
		NumSiniestro:    data.NumSiniestro,
		FechaSiniestro:  data.FechaSiniestro,
		FechaPago:       data.FechaPago,
		Cobertura:       data.Cobertura,
	}

	return payment, nil
}
