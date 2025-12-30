package endpoints

import (
	"bytes"
	"fmt"
	"io"
	"is-public-api/application/models"
	helpers_ai "is-public-api/helpers/apihelpers"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// Constantes para URLs
const (
	baseURL          = "https://qualitatasesoria.com/siniestros"
	loginURL         = baseURL + "/Login/Login.aspx"
	validateLoginURL = baseURL + "/Login/ValidarLogin.aspx"
	saveConfigURL    = baseURL + "/General/GuardarConfiguracion.aspx"
	beneficiaryURL   = baseURL + "/General/ListadoBeneficiario.aspx"
	sinisterURL      = baseURL + "/Consultas/ListadoConsultaSiniestroCli.aspx"
)

// Headers comunes
var commonHeaders = map[string]string{
	"accept":             "*/*",
	"content-type":       "application/x-www-form-urlencoded",
	"origin":             "https://qualitatasesoria.com",
	"sec-ch-ua":          `"Chromium";v="130", "Google Chrome";v="130", "Not?A_Brand";v="99"`,
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": `"Windows"`,
	"x-requested-with":   "XMLHttpRequest",
}

var initialHeaders = map[string]string{
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"accept-language":           "es-ES,es;q=0.9,en;q=0.8",
	"cache-control":             "max-age=0",
	"priority":                  "u=0, i",
	"referer":                   baseURL + "/Login/Principal.aspx",
	"sec-ch-ua":                 `"Chromium";v="130", "Google Chrome";v="130", "Not?A_Brand";v="99"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        `"Windows"`,
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
}

type qualitatEndpoint struct {
	conf *helpers_ai.QualitatSinister
	env  *helpers_ai.EndpointHelper
}

func NewQualitatEndpoint(conf *helpers_ai.EndpointConfig, env *helpers_ai.EndpointHelper) *qualitatEndpoint {
	return &qualitatEndpoint{conf: &helpers_ai.QualitatSinister{Document: conf.QualitatSinister.Document}, env: env}
}

func (service *qualitatEndpoint) QualitatStart(Document string, InjuredCompleteName string) (*[]models.QualitatResponse, error) {
	// Crear un cliente HTTP con soporte para cookies
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("error creando cookie jar: %w", err)
	}
	client := &http.Client{Jar: jar}

	// Endpoint 1: Login inicial
	resp, err := makeRequest(client, loginURL, "GET", nil, initialHeaders)
	if err != nil {
		return nil, fmt.Errorf("error en login inicial: %w", err)
	}
	defer resp.Body.Close()

	return service.QualitatLogin(client, Document, InjuredCompleteName)
}

func (service *qualitatEndpoint) QualitatLogin(client *http.Client, Document string, InjuredCompleteName string) (*[]models.QualitatResponse, error) {
	data := url.Values{}
	data.Set("desperfil", "")
	data.Set("txtusuario", service.env.Conf.Qualitat.QualitatUser)
	data.Set("txtclave", service.env.Conf.Qualitat.QualitatPass)

	headers := mergeHeaders(commonHeaders, map[string]string{
		"referer": loginURL,
	})

	resp, err := makeRequest(client, validateLoginURL, "POST", bytes.NewBufferString(data.Encode()), headers)
	if err != nil {
		return nil, fmt.Errorf("error en validación de login: %w", err)
	}
	defer resp.Body.Close()

	return service.QualitatSaveConfig(client, Document, InjuredCompleteName)
}

func (service *qualitatEndpoint) QualitatSaveConfig(client *http.Client, Document string, InjuredCompleteName string) (*[]models.QualitatResponse, error) {
	data := url.Values{}
	data.Set("id_perfil", "2")
	data.Set("confcliente", "1")
	data.Set("ddlconfriesgo", "2")
	data.Set("ddlconframo", "4")

	headers := mergeHeaders(commonHeaders, map[string]string{
		"referer": loginURL,
	})

	resp, err := makeRequest(client, saveConfigURL, "POST", bytes.NewBufferString(data.Encode()), headers)
	if err != nil {
		return nil, fmt.Errorf("error guardando configuración: %w", err)
	}
	defer resp.Body.Close()

	// Intentar búsqueda por documento primero
	result, err := service.QualitatGetListBeneficiaries(client, Document, "")
	if err != nil {
		return nil, fmt.Errorf("error obteniendo lista de beneficiarios por documento: %w", err)
	}

	// Si no se encontraron resultados por documento y hay nombre disponible, buscar por nombre
	if result == nil && InjuredCompleteName != "" {
		result, err = service.QualitatGetListBeneficiaries(client, "", InjuredCompleteName)
		if err != nil {
			return nil, fmt.Errorf("error obteniendo lista de beneficiarios por nombre: %w", err)
		}
	}

	return result, nil
}

func (service *qualitatEndpoint) QualitatGetListBeneficiaries(client *http.Client, Document string, InjuredCompleteName string) (*[]models.QualitatResponse, error) {
	data := url.Values{}
	data.Set("idbeneficiario", "")
	data.Set("modo", "C")
	data.Set("tipo", "")
	data.Set("pagina", "1")
	data.Set("regxpag", "10")
	data.Set("paginador", "10")
	data.Set("seleccion", "")

	// Configurar búsqueda según el campo proporcionado
	if Document != "" {
		data.Set("txtdni", Document)
		data.Set("txtdatos", "")
	} else if InjuredCompleteName != "" {
		data.Set("txtdni", "")
		data.Set("txtdatos", InjuredCompleteName)
	} else {
		data.Set("txtdni", "NOT_PROVIDED")
		data.Set("txtdatos", "NOT_PROVIDED")
	}

	headers := mergeHeaders(commonHeaders, map[string]string{
		"referer": loginURL,
	})

	resp, err := makeRequest(client, beneficiaryURL, "POST", bytes.NewBufferString(data.Encode()), headers)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo lista de beneficiarios: %w", err)
	}
	defer resp.Body.Close()

	// Lee la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta de beneficiarios: %w", err)
	}

	// Analiza el HTML
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error analizando HTML de beneficiarios: %w", err)
	}

	// Extraer los parámetros del onclick
	params, err := extractOndblclickParameters(doc, "tblBeneficiario")
	if err != nil {
		return nil, fmt.Errorf("error extrayendo parámetros de beneficiarios: %w", err)
	}

	if len(params) > 0 {
		// Acceder a la primera fila
		firstRow := params[0]
		return service.QualitatGetSinister(client, firstRow[0], firstRow[2], firstRow[3], firstRow[4], Document)
	}
	return nil, nil
}

func (service *qualitatEndpoint) QualitatGetSinister(client *http.Client, ID string, lastName1 string, lastName2 string, firstName string, Document string) (*[]models.QualitatResponse, error) {
	data := url.Values{}
	data.Set("pagina", "1")
	data.Set("regxpag", "14")
	data.Set("paginador", "10")
	data.Set("txtubigeo", "")
	data.Set("tmpfecha", "")
	data.Set("idbeneficiario", ID)
	data.Set("txtsiniestro", "")
	data.Set("txtpoliza", "")
	data.Set("txtplaca", "")
	data.Set("txtbeneficiario", fmt.Sprintf("%s - %s %s %s", Document, lastName1, lastName2, firstName))
	data.Set("txtsiniestrocliente", "")

	headers := mergeHeaders(commonHeaders, map[string]string{
		"referer": loginURL,
	})

	resp, err := makeRequest(client, sinisterURL, "POST", bytes.NewBufferString(data.Encode()), headers)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo siniestros: %w", err)
	}
	defer resp.Body.Close()

	// Lee la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta de siniestros: %w", err)
	}

	// Analiza el HTML
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error analizando HTML de siniestros: %w", err)
	}

	// Extrae y procesa los datos de la tabla
	return service.processSinisterData(extractTableData(doc))
}

// processSinisterData convierte los datos de la tabla HTML en objetos QualitatResponse
func (service *qualitatEndpoint) processSinisterData(tableData [][]string) (*[]models.QualitatResponse, error) {
	if len(tableData) == 0 {
		return nil, nil
	}

	dataQualitat := make([]models.QualitatResponse, len(tableData))
	for i, row := range tableData {
		if len(row) >= 17 {
			dataQualitat[i] = models.QualitatResponse{
				FechaSiniestro:      row[0],
				NroSiniestroCliente: row[1],
				NroSiniestro:        row[2],
				NroPoliza:           row[3],
				Placa:               row[4],
				NroCaso:             row[5],
				Nombres:             row[6],
				ApPaterno:           row[7],
				ApMaterno:           row[8],
				DNI:                 row[9],
				CentroMedico:        row[10],
				TipoSiniestro:       row[11],
				Recepcion:           row[12],
				Ocupante:            row[13],
				Fallecido:           row[14],
				FechaFallecimiento:  row[15],
				EstadoSiniestro:     row[16],
			}
		}
	}
	return &dataQualitat, nil
}

// mergeHeaders combina headers base con headers adicionales
func mergeHeaders(base, additional map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range base {
		result[k] = v
	}
	for k, v := range additional {
		result[k] = v
	}
	return result
}

// makeRequest realiza una solicitud HTTP con el cliente proporcionado, método, cuerpo y encabezados.
func makeRequest(client *http.Client, url string, method string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error creando la solicitud: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error enviando la solicitud: %w", err)
	}

	return resp, nil
}

// extractTableData extrae datos de tablas HTML de manera más eficiente
func extractTableData(n *html.Node) [][]string {
	var rows [][]string

	var traverse func(*html.Node, []string) []string
	traverse = func(n *html.Node, currentRow []string) []string {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "td":
				// Extrae el texto dentro de la celda
				cellContent := extractText(n)
				return append(currentRow, cellContent)
			case "tr":
				if len(currentRow) > 0 {
					// Agrega la fila completa cuando termina un <tr>
					rows = append(rows, currentRow)
					currentRow = []string{}
				}
			}
		}

		// Procesa los nodos hijos
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			currentRow = traverse(c, currentRow)
		}
		return currentRow
	}

	traverse(n, []string{})
	return rows
}

// extractText extrae texto de nodos HTML de manera recursiva
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}

	var result strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if text := extractText(c); text != "" {
			if result.Len() > 0 {
				result.WriteString(" ")
			}
			result.WriteString(text)
		}
	}
	return result.String()
}

// extractOndblclickParameters extrae los parámetros de la función 'escogerBeneficiario' en el atributo ondblclick
func extractOndblclickParameters(doc *html.Node, tableID string) ([][]string, error) {
	tableNode := findTableByID(doc, tableID)
	if tableNode == nil {
		return nil, fmt.Errorf("table with id '%s' not found", tableID)
	}

	// Regex para capturar los parámetros de 'escogerBeneficiario(...)'
	regex := regexp.MustCompile(`escogerBeneficiario\("([^"]*)","([^"]*)","([^"]*)","([^"]*)","([^"]*)"\)`)

	var params [][]string
	var extractFromTable func(*html.Node)

	extractFromTable = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			// Buscar atributo ondblclick en el <tr>
			for _, attr := range n.Attr {
				if attr.Key == "ondblclick" {
					if matches := regex.FindStringSubmatch(attr.Val); len(matches) == 6 {
						params = append(params, matches[1:])
					}
					break
				}
			}
		}
		// Recorrer los nodos hijos
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractFromTable(c)
		}
	}

	// Procesar el <tbody> de la tabla
	for c := tableNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "tbody" {
			extractFromTable(c)
			break
		}
	}

	return params, nil
}

// findTableByID busca una tabla por su ID de manera recursiva
func findTableByID(n *html.Node, tableID string) *html.Node {
	if n.Type == html.ElementNode && n.Data == "table" {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == tableID {
				return n
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findTableByID(c, tableID); result != nil {
			return result
		}
	}

	return nil
}
