package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}

func uploadsHandler(writer http.ResponseWriter, request *http.Request) {
	// get file from request
	uploadFile, header, err := request.FormFile("upload_file")
	// handle error
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, err)
		return
	}
	// newFile file
	dirName := "./uploads"
	os.MkdirAll(dirName, 0777)
	filePath := fmt.Sprintf("%s/%s", dirName, header.Filename)
	newFile, err := os.Create(filePath)
	defer newFile.Close()

	// handle error
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, err)
		return
	}

	// write file
	io.Copy(newFile, uploadFile)
	writer.WriteHeader(http.StatusCreated)

}
