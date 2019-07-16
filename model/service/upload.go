package service

import (
	"mime/multipart"
)

type UploadFileInput struct {
	Header *multipart.FileHeader
	File   *multipart.File
}
