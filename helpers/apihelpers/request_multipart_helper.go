package apihelpers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

type File struct {
	Key         string
	Name        string
	File        *bytes.Buffer
	ContentType string
	Size        int64
}

func MakeMultipart(form map[string]string, files ...File) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()

	for key, val := range form {
		mp.WriteField(key, val)
	}
	for _, f := range files {
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; Name="%s"; filename="%s"`, f.Key, f.Name))
		h.Set("Content-Type", f.ContentType)
		part, err := mp.CreatePart(h)
		if err != nil {
			return "", nil, err
		}
		if _, err = io.Copy(part, f.File); err != nil {
			return "", nil, err
		}
		// part.Write(f.file.Bytes())
	}
	return mp.FormDataContentType(), body, nil
}

func HttpMultipartRequest(url string, method string, body io.Reader, contentType string) (int, []byte, error) {
	req, _ := http.NewRequest(method, url, body)

	if contentType == "" {
		contentType = "application/json"
	}
	req.Header.Add("Content-Type", contentType)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		var statusCode int
		if res != nil {
			statusCode = res.StatusCode
		}
		return statusCode, nil, err
	}
	defer res.Body.Close()

	response, _ := ioutil.ReadAll(res.Body)

	return res.StatusCode, response, nil
}
