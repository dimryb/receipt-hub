package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func GetProjectRoot() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		_, err = os.Stat(filepath.Join(path, "go.mod"))
		if errors.Is(err, os.ErrNotExist) {
			// Если текущая директория уже корневая, выходим из цикла
			if path == filepath.Dir(path) {
				panic("could not find go.mod file in any parent directory")
			}
			path = filepath.Dir(path)
			continue
		}
		if err != nil {
			panic(err)
		}

		break
	}

	return path
}
