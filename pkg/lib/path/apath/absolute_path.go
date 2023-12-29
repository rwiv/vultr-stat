package apath

import (
	"fmt"
	"path/filepath"
	"regexp"

	"vultr-stat/pkg/lib/path/pwd"
)

const (
	projectName = "vultr-stat"
)

func GetProjectAbsPath() (string, error) {
	ex, err := filepath.Abs("")
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return GetProjectAbsPathBy(exPath, string(filepath.Separator))
}

func GetProjectAbsPathBy(path string, separator string) (string, error) {
	if separator == `\` {
		separator = `\\`
	}
	expr := fmt.Sprintf(".*?%s%s", separator, projectName)
	reg, err := regexp.Compile(expr)
	if err != nil {
		return "", err
	}
	m := reg.FindString(path)

	return m, nil
}

func GetAbsPath(path string) (string, error) {
	if path[0] == '/' {
		return path, nil
	}

	pwdPath, err := pwd.GetPwd()
	if err != nil {
		return "", err
	}

	return pwdPath + "/" + path, nil
}

func IsAbsPath(path string) bool {
	return path[0] == '/'
}
