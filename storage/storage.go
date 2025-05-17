package storage

import (
	"fmt"
	"os"
	"purple_school/hw-project-3/bins"
	"purple_school/hw-project-3/files"
)

func WriteIntoFile(bin bins.Bin, name string) {
	if !files.IsJson(name) {
		fmt.Println("Файл должен быть с расширением JSON")
		return
	}

	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}

	defer file.Close()

	data, err := bin.ToBytes()
	if err != nil {
		fmt.Println("Ошибка конвертации данных:", err)
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	}
}

func ReadFromFile(name string) ([]byte, error) {
	if !files.IsJson(name) {
		errStr := fmt.Errorf("Файл должен быть c расширением JSON")
		return nil, errStr
	}

	data, err := files.ReadFile(name)
	if err != nil {
		errStr := fmt.Errorf("Ошибка чтения файла:", err)
		return nil, errStr
	}

	return data, nil
}
