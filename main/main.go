package main

import (
	"bufio"
	"fmt"
	"mirrorsTestTools/main/cfginput"
	"mirrorsTestTools/main/tester"
	"os"
	"strings"
)

func main() {
	Mirrors := cfginput.ReadFile("urls.json")
	mirrorURLs := []string{}
	for _, url := range Mirrors {
		mirrorURLs = append(mirrorURLs, url)
		// fmt.Println(name, url)
	}

	// 用户开始
	fmt.Println("欢迎使用镜像测速工具")
	fmt.Println("1. 批量自定义镜像站URL. 2. 批量选择库中镜像站. 3. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("请输入要测试的镜像站的 URL后回车: (Default:Tsinghua&Aliyun)\n")
		// 读取用户批量输入的镜像站URL，并append到mirrorURLs切片中
		for {
			input, _ := reader.ReadString('\n')
			// 移除末尾的换行符
			input = strings.TrimSpace(input)
			if input == "" {
				break
			}
			mirrorURLs = append(mirrorURLs, input)
		}

	case 2:
		// coming soon...
		// mirrorURL =
		return

	case 3:
		return
	}

	// 执行测试
	for i := range mirrorURLs {
		downloadSpeed, err := tester.TestMirrorSpeed(mirrorURLs[i])

		// 处理测试结果
		if err != nil {
			fmt.Printf("测试失败：%s\n", err)
		} else {
			fmt.Printf("从 %s 下载 Debian Release 文件的速度：%f MB/s\n", mirrorURLs[i], downloadSpeed)
		}
	}
}
