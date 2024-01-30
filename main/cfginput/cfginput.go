package cfginput

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(path string) map[string]string {
	// 打开json配置文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败:", err)
	}
	defer file.Close()

	// 解码json配置文件到Mirrors中
	Mirrors := make(map[string]string)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Mirrors)
	if err != nil {
		fmt.Println("解码失败:", err)
	}

	return Mirrors

	// // 打印配置信息
	// fmt.Println("镜像配置:")
	// for name, url := range Mirrors {
	// 	fmt.Printf("%s: %s\n", name, url)
	// }
}
