package main

import (
	"fmt"
	"mirrorsTestTools/main/task"
)

func main() {

	if task.CheckOS() == 0 {
		fmt.Println("您的系统不是Debian 12 or Ubuntu 22.04, 请待支持")
		return
	}

	// Use cfginput.ReadFile() to read the configuration file and convert it to a slice.
	// 调用cfginput.ReadFile()函数读取配置文件，并将其转换为切片
	mirrorNames, mirrorURLs := task.ReadFile("urls.json")
	// fmt.Println(mirrorNames, mirrorURLs)

	fmt.Println("欢迎使用镜像测速工具")
	fmt.Println("1. 批量自定义镜像站URL. 2. 批量选择库中镜像站. 3. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		fmt.Print("请输入要测试的镜像站的 URL: (Default:Tsinghua)\n")

		task.InputURL(&mirrorURLs)

		task.Test(mirrorNames, mirrorURLs)

	case 2:
		fmt.Println("请选择要测试的镜像站(中间用空格隔开)，0 表示全选:")
		for i, name := range mirrorNames {
			fmt.Printf("%d. %s\n", i+1, name)
		}

		task.SelectMirror(&mirrorNames, &mirrorURLs)

		task.Test(mirrorNames, mirrorURLs)

	case 3:
		return
	}

}
