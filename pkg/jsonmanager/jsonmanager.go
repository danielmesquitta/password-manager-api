package jsonmanager

import (
	"io"
	"os"

	"github.com/goccy/go-json"
)

func Open(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
}

func Decode(file *os.File, data any) error {
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func Encode(file *os.File, data any) error {
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
