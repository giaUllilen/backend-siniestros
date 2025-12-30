package services

import (
	"bytes"
	"html/template"

	"is-public-api/application/models"
	"is-public-api/application/templates"
)

type templateHtmlMaker struct {
}

func NewTemplateHtmlMaker() IServiceTemplateMaker {
	return &templateHtmlMaker{}
}

func (service *templateHtmlMaker) Make(txContext *models.TxContext, data map[string]interface{}) (*bytes.Buffer, error) {

	t, err := template.New("mailpage").Parse(templates.SinisterSuccess)
	if err != nil {
		// apihelpers.RenderError(ctx, err, fasthttp.StatusBadRequest)
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)

	return buf, nil
}
