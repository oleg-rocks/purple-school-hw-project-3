package storage

import (
	"fmt"
	"os"
	"purple_school/hw-project-3/bins"
	"purple_school/hw-project-3/files"
)

func WriteIntoFile(bin bins.Bin, name string) error {
	if !files.IsJson(name) {
		err := fmt.Errorf("Файл должен быть с расширением JSON")
		return err
	}

	file, err := os.Create(name)
	if err != nil {
		errStr := fmt.Errorf("Ошибка создания файла: %w", err)
		return errStr
	}

	defer file.Close()

	data, err := bin.ToBytes()
	if err != nil {
		errStr := fmt.Errorf("Ошибка конвертации данных: %w", err)
		return errStr
	}

	_, err = file.Write(data)
	if err != nil {
		errStr := fmt.Errorf("Ошибка записи в файл: %w", err)
		return errStr
	}

	return nil
}

func ReadFromFile(name string) ([]byte, error) {
	if !files.IsJson(name) {
		errStr := fmt.Errorf("Файл должен быть c расширением JSON")
		return nil, errStr
	}

	data, err := files.ReadFile(name)
	if err != nil {
		errStr := fmt.Errorf("Ошибка чтения файла: %w", err)
		return nil, errStr
	}

	return data, nil
}
