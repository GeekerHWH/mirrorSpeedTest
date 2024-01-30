package main

import (
	"fmt"
	"mirrorsTestTools/main/tester"
)

// mirrorURL 是要测试的 TUNA Debian 镜像站的 URL
// const mirrorURL = "https://mirrors.tuna.tsinghua.edu.cn/debian"

func main() {
	var mirrorURL string = "https://mirrors.tuna.tsinghua.edu.cn/debian"
	fmt.Print("请输入要测试的 TUNA Debian 镜像站的 URL: ")
	fmt.Scanln(&mirrorURL)
	// 执行测试
	downloadSpeed, err := tester.TestMirrorSpeed(mirrorURL)

	// 处理测试结果
	if err != nil {
		fmt.Printf("测试失败：%s\n", err)
	} else {
		fmt.Printf("从 %s 下载 Debian Release 文件的速度：%f MB/s\n", mirrorURL, downloadSpeed)
	}
}
