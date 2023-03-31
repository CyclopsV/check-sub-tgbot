package pars

import (
	"encoding/json"
	"log"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	log.Printf("Извлечение данных из файла `%v`", filePath)
	content, err := os.ReadFile(filePath)
	return content, err
}

func JSON[T any](storage *T, path string) error {
	content, err := ReadFile(path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, &storage); err != nil {
		return err
	}
	return nil
}
