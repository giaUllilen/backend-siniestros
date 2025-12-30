package apihelpers

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/francoispqt/gojay"
)

// HttpRequest function
func HttpRequest(url string, method string, Data []byte, headers ...map[string]string) (int, []byte, error) {
	data := strings.NewReader(string(Data))
	req, _ := http.NewRequest(method, url, data)

	req.Header.Add("Content-Type", "application/json")
	if len(headers) > 0 {
		firstHeaders := headers[0]
		for k, v := range firstHeaders {
			req.Header.Add(k, v)
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("[HttpRequest] error Request: ", err)
		return res.StatusCode, nil, err
	}
	defer res.Body.Close()

	response, _ := io.ReadAll(res.Body)

	return res.StatusCode, response, nil
}

// HttpRequestToStruct function
func HttpRequestToStruct(url string, method string, Data []byte, response interface{}) (int, error) {
	data := strings.NewReader(string(Data))
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		log.Println("error Request: ", err)
		return 500, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("[HttpRequestToStruct] error Request: ", err)
		if res != nil && res.Body != nil {
			bytes, _ := io.ReadAll(res.Body)
			gojay.Unmarshal(bytes, response)
		}
		return res.StatusCode, err
	}
	defer res.Body.Close()

	bytes, _ := io.ReadAll(res.Body)
	_ = gojay.Unmarshal(bytes, &response)
	return res.StatusCode, nil
}
