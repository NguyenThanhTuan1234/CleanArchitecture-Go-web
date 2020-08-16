package gateway

import (
	"crypto/sha1"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func (r *fileRepository) UploadImage(mf multipart.File, fh *multipart.FileHeader) (string, error) {
	// defer mf.Close()
	// Create sha for file name
	fmt.Println(fh.Filename)
	ext := strings.Split(fh.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, mf)
	fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
	// Create new file
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "resources", "pics", fname)
	nf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
	mf.Seek(0, 0)
	io.Copy(nf, mf)
	fmt.Println(fname)
	return fname, nil
}
