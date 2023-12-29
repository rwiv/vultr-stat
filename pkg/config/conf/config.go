package conf

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"vultr-stat/pkg/lib/path/apath"
)

const (
	ConfigDirectory = "conf"
	ConfigFile      = "conf.json"
)

func GetConf() (*Conf, error) {
	b, err := getJsonBytes(filepath.Join(ConfigDirectory, ConfigFile))
	if err != nil {
		return nil, err
	}

	result := new(Conf)
	err = json.Unmarshal(b, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getJsonBytes(relPath string) ([]byte, error) {
	absPath, err := apath.GetProjectAbsPath()
	if err != nil {
		return nil, err
	}
	confPath := filepath.Join(absPath, relPath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(confPath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}
