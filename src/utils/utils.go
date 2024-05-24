package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(dir string, fileName string) (*os.File, error) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("os.MkdirAll: %w", err)
	}
	filepath := filepath.Join(dir, fileName)
	f, err := os.Create(filepath)
	if err != nil {
		return nil, fmt.Errorf("os.Create: %w", err)
	}
	return f, nil
}
