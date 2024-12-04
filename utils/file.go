package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type fileUtils struct{}

type FileUtils interface {
	CheckNumFolder(dir string) (int, error)
	CreateFile(name string, dirSave string, file multipart.File, header *multipart.FileHeader) (path string, ext string, err error)
}

func (u *fileUtils) CheckNumFolder(dir string) (int, error) {
	fileCount := 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Kiểm tra nếu đó là file và không phải là thư mục
		if !info.IsDir() {
			fileCount++
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return fileCount, nil
}

func (u *fileUtils) CreateFile(name string, dirSave string, file multipart.File, header *multipart.FileHeader) (path string, ext string, err error) {
	fileExtension := filepath.Ext(header.Filename)

	outputFileName := fmt.Sprintf("%s/%s%s", dirSave, name, fileExtension)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return "", "", err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)
	if err != nil {
		return "", "", err
	}

	return outputFileName, fileExtension, nil
}

func NewFileUtils() FileUtils {
	return &fileUtils{}
}
