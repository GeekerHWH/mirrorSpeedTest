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
	// 调用cfginput.ReadFile()函数读取配置文件，并将其转换为切片
	Mirrors := cfginput.ReadFile("urls.json")
	mirrorNames := []string{}
	mirrorURLs := []string{}
	for name, url := range Mirrors {
		mirrorNames = append(mirrorNames, name)
		mirrorURLs = append(mirrorURLs, url)
		// fmt.Println(name, url)
	}

	fmt.Println("欢迎使用镜像测速工具")
	fmt.Println("1. 批量自定义镜像站URL. 2. 批量选择库中镜像站. 3. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("请输入要测试的镜像站的 URL: (Default:Tsinghua)\n")

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
	if len(mirrorURLs) == 4 {
		downloadSpeed, err := tester.TestMirrorSpeed(mirrorURLs[1])

		// 处理测试结果
		if err != nil {
			fmt.Printf("测试失败：%s\n", err)
		} else {
			fmt.Printf("从 %s 下载 Debian Release 文件的速度：%f MB/s\n", mirrorNames[1], downloadSpeed)
		}
	} else {
		for i := len(mirrorNames); i < len(mirrorURLs); i++ {
			downloadSpeed, err := tester.TestMirrorSpeed(mirrorURLs[i])

			// 处理测试结果
			if err != nil {
				fmt.Printf("测试失败：%s\n", err)
			} else {
				fmt.Printf("从 %s 下载 Debian Release 文件的速度：%f MB/s\n", mirrorURLs[i], downloadSpeed)
			}
		}
	}

}
