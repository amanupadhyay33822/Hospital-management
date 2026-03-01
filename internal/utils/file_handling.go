package utils

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	var file multipart.File
	var info *multipart.FileHeader
	var err error
	file, info, err = r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	var dst *os.File
	dst, err = os.Create("./uploads/" + info.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	///copy uploaded file to destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendResponse(w, true, "Image Uploaded Sucessfully", nil)
}
