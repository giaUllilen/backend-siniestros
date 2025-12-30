package resources

import "github.com/francoispqt/gojay"

type DocumentDownload struct {
	URL string `json:"url"`
}

func (p *DocumentDownload) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "url":
		return dec.String(&p.URL)
	}
	return nil
}

func (p *DocumentDownload) NKeys() int {
	return 1
}
