package mode

import (
	"flag"
	"fmt"
	"mirrorsTestTools/main/task"
)

func ParametersRun() {
	// 定义命令行参数
	var interactiveMode bool
	var showHelp bool
	var url string

	// 设置命令行参数
	flag.BoolVar(&interactiveMode, "i", false, "Enter interactive mode")
	flag.BoolVar(&showHelp, "h", false, "Print program introduction and usage of its parameters")
	flag.StringVar(&url, "url", "", "Specify a URL")

	// 定制Usage信息
	flag.Usage = func() {
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	// 如果有-h参数或没有任何参数，则打印帮助信息
	if showHelp || flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	// 如果有--url参数，则将其后的值赋给url变量
	if url != "" {
		var URL = []string{url}
		// 在这里可以使用url变量进行相应的操作
		task.Test(URL, URL, 0)
		return
	}

	// 如果有-i参数，则进入交互模式
	if interactiveMode {
		Interactive()
		return
	}

	// 如果没有匹配的参数，则执行默认操作
	fmt.Println("No valid options provided. Use -h for help.")
}
