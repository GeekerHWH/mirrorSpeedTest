package mode

import (
	"flag"
	"fmt"
	"mirrorsTestTools/main/task"
)

func ParametersRun() {
	// 定义命令行参数
	var showHelp bool
	var url string
	var country string
	var interactiveCountry string

	// 设置命令行参数
	flag.BoolVar(&showHelp, "h", false, "Print program introduction and usage of its parameters")
	flag.StringVar(&url, "url", "", "Specify a single URL")
	flag.StringVar(&country, "c", "", "Specify the country")
	flag.StringVar(&interactiveCountry, "i", "", "Specify the country and Enter interactive mode")

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
		task.Test(URL, URL)
		return
	}

	// 如果有-c参数，则将其后的值赋给country变量
	if country != "" {
		// 在这里可以使用country变量进行相应的操作
		countryConfigNames, countryConfigURLs, err := task.ReadFile("urls.json", country)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		task.Test(countryConfigNames, countryConfigURLs)
		return
	}

	// 如果有-i参数，则进入交互模式
	if interactiveCountry != "" {
		Interactive(interactiveCountry)
		return
	}

	// 如果没有匹配的参数，则执行默认操作
	fmt.Println("No valid options provided. Use -h for help.")
}
