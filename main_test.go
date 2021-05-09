package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func Test_uploadHandler(t *testing.T) {
	assert := assert.New(t)
	originalPath := "/home/potato/Studyspace/go/web/study02/easywalk.png"
	file, _ := os.Open(originalPath)
	defer file.Close()

	buf := bytes.Buffer{}
	writer := multipart.NewWriter(&buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(originalPath))
	assert.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", &buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadsHandler(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(originalPath)
	assert.FileExists(uploadFilePath)

	uploaded, _ := os.Open(uploadFilePath)
	original, _ := os.Open(originalPath)
	defer uploaded.Close()
	defer original.Close()

	uploadedData := []byte{}
	originalData := []byte{}
	uploaded.Read(uploadedData)
	original.Read(originalData)
	assert.Equal(uploadedData, originalData)
}
