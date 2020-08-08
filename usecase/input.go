package usecase

import (
	"mime/multipart"
	"net/http"
)

type FormInput interface {
	GetFormValue(http.ResponseWriter, *http.Request) (string, error)
}

type ParamInput interface {
	ReadParam() (string, error)
}

type FileInput interface {
	GetFile(http.ResponseWriter, *http.Request) (multipart.File, *multipart.FileHeader, error)
}
