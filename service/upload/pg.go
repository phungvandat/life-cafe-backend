package upload

import (
	"context"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	serviceModel "github.com/phungvandat/life-cafe-backend/model/service"
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
			return []string{}, CreateFileError
		}

		_, err = io.Copy(out, *image.File)

		if err != nil {
			return []string{}, CopyFileError
		}

		arrImageName = append(arrImageName, fileName)
	}
	return arrImageName, nil
}
