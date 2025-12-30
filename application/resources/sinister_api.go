package resources

type ResponseSinister struct {
	Code    string `json:"codigo"`
	Message string `json:"message"`
	Results struct {
		PersonDocID string `json:"id_persona_documento"`
		CaseValueID int    `json:"id_caso_valor"`
		CaseID      string `json:"id_num_caso"`
	} `json:"results"`
}
type ResponseStorage struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Stack      string `json:"stack"`
}

type ResponseApi struct {
	Codes []int `json:"codes"`
}

func NewResponseApi() *ResponseApi {
	return &ResponseApi{
		Codes: make([]int, 0),
	}
}
