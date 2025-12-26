package apihelpers

import (
	"bufio"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

func DetectContentType(fileHeader *multipart.FileHeader) string {
	file, _ := fileHeader.Open()
	defer file.Close()
	buffer := make([]byte, 512)
	file.Read(buffer)
	return http.DetectContentType(buffer)
}

func ReadFile(fileHeader *multipart.FileHeader) (*bytes.Buffer, error) {
	var part []byte
	var count int
	var err error
	file, _ := fileHeader.Open()
	defer file.Close()
	reader := bufio.NewReader(file)
	buf := bytes.NewBuffer(make([]byte, 0))
	chunkSize := 1024
	part = make([]byte, chunkSize)
	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buf.Write(part[:count])
	}
	if err != io.EOF {
		return nil, err
	} else {
		err = nil
	}
	return buf, nil
}
