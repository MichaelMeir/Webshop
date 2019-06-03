package controllers

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"strconv"
	"strings"
	"webshop/api/models"
)

func (c Controller) GetProductImage(w http.ResponseWriter, req *http.Request) {
	var image = models.Image{}
	c.DB.Data.Where("uuid = ?", req.URL.Query().Get("id")).First(&image)
	w.Header().Set("Content-Type", "image/png")
	if strings.HasPrefix(image.ImageData, "data:image/png;base64,") {
		w.Write([]byte(image.ImageData))
	}else {
		w.Write([]byte("data:image/png;base64," + image.ImageData))
	}
}

func (c Controller) UploadImage(w http.ResponseWriter, req *http.Request) {
	var image = models.Image{}
	file, _, err := req.FormFile("image")
	serverError:
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " file upload failed: " + err.Error()))
		return
	}
	var buffer = bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		goto serverError
	}
	imageData := base64.StdEncoding.EncodeToString(buffer.Bytes())
	image.ImageData = imageData
	image.UUID = ""
}