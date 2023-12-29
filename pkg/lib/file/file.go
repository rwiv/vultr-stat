package file

import (
	"os"
	"path/filepath"
	"strings"
)

func GetFilename(path string) string {
	s := string(filepath.Separator)
	chunks := strings.Split(path, s)
	return chunks[len(chunks)-1]
}

func GetPwd() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}
