package upload

import (
	"context"
	"net/http"

	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	serviceModel "github.com/phungvandat/life-cafe-backend/model/service"
)

// UploadImagesRequest func
func UploadImagesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestModel.UploadImagesRequest

	err := r.ParseMultipartForm(200000)
	if err != nil {
		return nil, err
	}

	files := r.MultipartForm.File["images"]

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			file.Close()
			return nil, err
		}
		req = append(req, serviceModel.UploadFileInput{
			File:   &file,
			Header: fileHeader,
		})
	}
	return req, nil
}
