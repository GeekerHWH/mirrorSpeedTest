package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Name []string
	URL  []string
}

func ReadFile(path string) ([]string, []string) {
	// 打开json配置文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败:", err)
	}
	defer file.Close()

	// 解码json配置文件到Mirrors中
	var Mirrors Config

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Mirrors)
	if err != nil {
		fmt.Println("解码失败:", err)
	}

	return Mirrors.Name, Mirrors.URL
}
