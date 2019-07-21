package upload

import (
	"context"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	serviceModel "github.com/phungvandat/life-cafe-backend/model/service"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
)

// pgService implmenter for User serivce in postgres
type pgService struct {
}

// NewPGService new pg service
func NewPGService() Service {
	return &pgService{}
}

func (s *pgService) UploadImages(ctx context.Context, inputFiles []serviceModel.UploadFileInput) ([]string, error) {
	var arrImageName []string

	for _, image := range inputFiles {
		defer (*image.File).Close()
		name := strings.Split(image.Header.Filename, ".")
		fileType := name[len(name)-1]

		fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + name[0] + "." + fileType

		out, err := os.Create("public/images/" + fileName)
		defer out.Close()

		if err != nil {
			return []string{}, errors.CreateFileError
		}

		_, err = io.Copy(out, *image.File)

		if err != nil {
			return []string{}, errors.CopyFileError
		}

		arrImageName = append(arrImageName, "public/images/"+fileName)
	}
	return arrImageName, nil
}

func (s *pgService) GetImageFile(ctx context.Context, path string) (*os.File, error) {
	file, err := os.Open("public/images/" + path)

	if err != nil {
		return nil, err
	}
	return file, nil
}
