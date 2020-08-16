package controller

import (
	"mime/multipart"
	"net/http"
)

type FileControllerPreset struct {
	_     struct{}
	Value string
}

func (c *FileControllerPreset) GetFile(w http.ResponseWriter, r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(c.Value)
}
