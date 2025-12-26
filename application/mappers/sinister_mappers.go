package mappers

import (
	"is-public-api/application/models"
	"is-public-api/application/resources"
)

func RequestToSinister(p *resources.SinisterRequest) *models.Sinister {
	return &models.Sinister{
		TipoPoliza: p.TipoPoliza,
		Asegurado: models.AseguradoSinister{
			ApellidoPaterno: p.Asegurado.ApellidoPaterno,
			ApellidoMaterno: p.Asegurado.ApellidoMaterno,
			Nombres:         p.Asegurado.Nombres,
			NroDocumento:    p.Asegurado.NroDocumento,
		},
		Solicitante: models.SolicitanteSinister{
			ApellidoPaterno: p.Solicitante.ApellidoPaterno,
			ApellidoMaterno: p.Solicitante.ApellidoMaterno,
			Nombres:         p.Solicitante.Nombres,
			NroDocumento:    p.Solicitante.NroDocumento,
			Correo:          p.Solicitante.Correo,
			Celular:         p.Solicitante.Celular,
		},
		Narracion:       p.Narracion,
		FechaOcurrencia: p.FechaOcurrencia,
		MontoSolicitado: p.MontoSolicitado,
		Pagador:         p.Pagador,
		FechasIncapacidad: func(fis []resources.FechasIncapacidad) []models.FechasIncapacidadSinister {
			res := make([]models.FechasIncapacidadSinister, len(fis))
			for i, fi := range fis {
				res[i] = models.FechasIncapacidadSinister{
					FechaInicioIncapacidad: fi.FechaInicioIncapacidad,
					FechaFinIncapacidad:    fi.FechaFinIncapacidad,
				}
			}
			return res
		}(p.FechasIncapacidad),
		DeclaracionJurada: p.DeclaracionJurada,
		Beneficiarios: func(b []resources.Beneficiario) []models.BeneficiarioSinister {
			res := make([]models.BeneficiarioSinister, len(b))
			for i, ben := range b {
				res[i] = models.BeneficiarioSinister{
					MetodoPago:   ben.MetodoPago,
					NroCuenta:    ben.NroCuenta,
					Titular:      ben.Titular,
					NroDocumento: ben.NroDocumento,
					Banco:        ben.Banco,
					Moneda:       ben.Moneda,
					TipoCuenta:   ben.TipoCuenta,
				}
			}
			return res
		}(p.Beneficiarios),
		Documentos: func(ds []resources.DocumentSection) []models.DocumentSectionSinister {
			res := make([]models.DocumentSectionSinister, len(ds))
			for i, d := range ds {
				res[i] = models.DocumentSectionSinister{
					Coverage: d.Coverage,
					Documents: func(docs []resources.Document) []models.DocumentSinister {
						dr := make([]models.DocumentSinister, len(docs))
						for j, doc := range docs {
							dr[j] = models.DocumentSinister{
								Name:     doc.Name,
								Filename: doc.Filename,
								FileURL:  doc.FileURL,
							}
						}
						return dr
					}(d.Documents),
					Additional: mapperDocuments(d.Additional),
				}
			}
			return res
		}(p.Documentos),
	}
}

func mapperDocuments(add []resources.Document) []models.DocumentSinister {
	ar := make([]models.DocumentSinister, len(add))
	for j, ad := range add {
		ar[j] = models.DocumentSinister{
			Name:     ad.Name,
			Filename: ad.Filename,
			FileURL:  ad.FileURL,
		}
	}
	return ar
}

func RequestToSinisterDoc(p *resources.SinisterDocRequest) *models.SinisterDoc {
	return &models.SinisterDoc{
		TipoPoliza: p.TipoPoliza,
		Asegurado: models.AseguradoSinister{
			ApellidoPaterno: p.Asegurado.ApellidoPaterno,
			ApellidoMaterno: p.Asegurado.ApellidoMaterno,
			Nombres:         p.Asegurado.Nombres,
			NroDocumento:    p.Asegurado.NroDocumento,
		},
		Solicitante: models.SolicitanteSinister{
			ApellidoPaterno: p.Solicitante.ApellidoPaterno,
			ApellidoMaterno: p.Solicitante.ApellidoMaterno,
			Nombres:         p.Solicitante.Nombres,
			NroDocumento:    p.Solicitante.NroDocumento,
			Correo:          p.Solicitante.Correo,
			Celular:         p.Solicitante.Celular,
		},
		Narracion:       p.Narracion,
		FechaOcurrencia: p.FechaOcurrencia,
		MontoSolicitado: p.MontoSolicitado,
		Pagador:         p.Pagador,
		FechasIncapacidad: func(fis []resources.FechasIncapacidad) []models.FechasIncapacidadSinister {
			res := make([]models.FechasIncapacidadSinister, len(fis))
			for i, fi := range fis {
				res[i] = models.FechasIncapacidadSinister{
					FechaInicioIncapacidad: fi.FechaInicioIncapacidad,
					FechaFinIncapacidad:    fi.FechaFinIncapacidad,
				}
			}
			return res
		}(p.FechasIncapacidad),
		DeclaracionJurada: p.DeclaracionJurada,
		Beneficiarios: func(b []resources.Beneficiario) []models.BeneficiarioSinister {
			res := make([]models.BeneficiarioSinister, len(b))
			for i, ben := range b {
				res[i] = models.BeneficiarioSinister{
					MetodoPago:   ben.MetodoPago,
					NroCuenta:    ben.NroCuenta,
					Titular:      ben.Titular,
					NroDocumento: ben.NroDocumento,
					Banco:        ben.Banco,
					Moneda:       ben.Moneda,
					TipoCuenta:   ben.TipoCuenta,
				}
			}
			return res
		}(p.Beneficiarios),
		Coverage: p.Coverage,
		Documento: models.DocumentSinister{
			Name:     p.Documento.Name,
			Filename: p.Documento.Filename,
			FileURL:  p.Documento.FileURL,
		},
	}
}
