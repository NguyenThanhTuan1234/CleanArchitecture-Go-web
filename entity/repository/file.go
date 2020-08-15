package repository

import (
	"mime/multipart"
)

type FileRepository interface {
	UploadImage(multipart.File, *multipart.FileHeader) (string, error)
	MapUidToImage(int, string) (string, error)
}
