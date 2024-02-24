package rest

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"

	"github.com/labstack/echo/v4"
)

var (
	currDir, _ = os.Getwd()

	driveDir = path.Join(currDir, "drive")
)

func ErrorResp(code int, err error) *echo.HTTPError {
	log.Println(err)

	return echo.NewHTTPError(code, err.Error())
}

func SaveFile(file *multipart.FileHeader, fileName string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	loc := path.Join(driveDir, fileName)

	dst, err := os.Create(loc)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return loc, nil
}
