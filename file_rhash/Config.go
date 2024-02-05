package file_rhash

import (
	"encoding/json"
	"os"
)

// 定义文件目录和hash文件名
type Config struct {
	Root         string `json:"path"`
	Filehash_old string `json:"filehash_old"`
}

func CheckConfigFile(configFilePath string) (bool, error) {
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func CreateDefaultConfig(configFilePath string) error {
	defaultConfig := Config{
		Root:         "/path/to/root",
		Filehash_old: "filehash_old.json",
	}

	configData, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, configData, 0644)
	if err != nil {
		return err
	}

	return nil
}
