package os

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

type OsHelper interface {
	FileExist(path string) bool
	DirExist(path string) bool
	CreateDir(path string) error
	GetFileAsByteArray(path string) ([]byte, error)
	GetFileAsString(path string) (string, error)
	CreateTemporaryTextFile(content string) (string, error)
	GetEnvironmentVariable(key string) string
	CreateBinaryFile(path string, content []byte) error
	CopyFile(sourcePath string, destinationPath string) error
}

type osHelper struct {
}

func NewOsHelper() (OsHelper, error) {
	return &osHelper{},
		nil
}

func (s *osHelper) FileExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		return false
	}

	return true
}

func (s *osHelper) DirExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	if !fileInfo.IsDir() {
		return false
	}

	return true
}

func (s *osHelper) CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (s *osHelper) GetFileAsByteArray(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (s *osHelper) GetFileAsString(path string) (string, error) {
	buf, err := s.GetFileAsByteArray(path)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (s *osHelper) CreateTemporaryTextFile(content string) (string, error) {
	file, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		return "", err
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func (s *osHelper) GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

func (s *osHelper) CreateBinaryFile(path string, content []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.Write(content)
	if err != nil {
		return err
	}

	return writer.Flush()
}

func (s *osHelper) CopyFile(sourcePath string, destinationPath string) error {
	sourceFileStat, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", sourcePath)
	}

	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}

	defer source.Close()

	if err = s.CreateDir(path.Dir(destinationPath)); err != nil {
		return err
	}

	destination, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	defer destination.Close()

	_, err = io.Copy(destination, source)

	return err
}
