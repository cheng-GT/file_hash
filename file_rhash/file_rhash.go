package file_rhash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func calculateHash(filePath string) (string, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	md5Hash := md5.New()
	sha1Hash := sha1.New()

	if _, err := io.Copy(md5Hash, file); err != nil {
		return "", "", err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return "", "", err
	}

	if _, err := io.Copy(sha1Hash, file); err != nil {
		return "", "", err
	}

	md5HashInBytes := md5Hash.Sum(nil)
	sha1HashInBytes := sha1Hash.Sum(nil)

	md5HashStr := hex.EncodeToString(md5HashInBytes)
	sha1HashStr := hex.EncodeToString(sha1HashInBytes)

	return md5HashStr, sha1HashStr, nil
}

func Json_writ() {
	// 读取JSON配置文件
	configFile, err := os.Open("Config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer configFile.Close()

	// 解码JSON配置文件
	var config Config
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return
	}

	root := config.Root
	hashTable := make(map[string][2]string)

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			md5Hash, sha1Hash, err := calculateHash(path)
			if err != nil {
				fmt.Printf("Error calculating hash for file %s: %v\n", path, err)
			} else {
				hashTable[path] = [2]string{md5Hash, sha1Hash}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将哈希表写入JSON文件
	jsonData, err := json.MarshalIndent(hashTable, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	Filehash_old := config.Filehash_old
	jsonFile, err := os.Create(Filehash_old)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
}
