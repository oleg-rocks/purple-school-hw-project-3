package files

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return nil, err
	}
	return data, nil
}

func IsJson(path string) bool {
	return strings.HasSuffix(path, ".json")
}
