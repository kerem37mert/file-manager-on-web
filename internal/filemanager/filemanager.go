package filemanager

import (
	"os"
)

func List(path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)

	return files, err
}
