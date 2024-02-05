package file_rhash

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Compare_hash() {
	root := "E:\\test"
	currentHashes := make(map[string][2]string)

	// 计算当前文件的哈希值
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			md5Hash, sha1Hash, err := calculateHash(path)
			if err != nil {
				fmt.Printf("Error calculating hash for file %s: %v\n", path, err)
			} else {
				currentHashes[path] = [2]string{md5Hash, sha1Hash}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

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

	// 读取之前的哈希表
	previousHashes := make(map[string][2]string)
	Filehash_old := config.Filehash_old
	previousHashesFile, err := os.Open(Filehash_old)
	if err == nil {
		defer previousHashesFile.Close()
		jsonParser := json.NewDecoder(previousHashesFile)
		if err = jsonParser.Decode(&previousHashes); err != nil {
			fmt.Println("Error decoding previous hashes:", err)
		}
	}

	// 比较当前哈希值和之前的哈希值
	for file, currentHash := range currentHashes {
		previousHash, ok := previousHashes[file]
		if !ok || previousHash != currentHash {
			fmt.Printf("File: %s, MD5: %s, SHA1: %s (changed)\n", file, currentHash[0], currentHash[1])
		}
	}
}
