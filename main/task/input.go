package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputURL(mirrorURLs *[]string) {
	// 读取用户批量输入的镜像站URL，并append到mirrorURLs切片中
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		// 移除末尾的换行符
		input = strings.TrimSuffix(input, "\n")
		if input == "" {
			break
		}
		*mirrorURLs = append(*mirrorURLs, input)
	}
}

func SelectMirror(mirrorURLs *[]string) {
	// 读取用户选择的镜像站序号
	var index string
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入错误:", err)
		return
	}

	// 移除换行符
	index = strings.TrimSuffix(line, "\n")

	// 以空格分割输入字符串
	data := strings.Fields(index) //data: [1 3 4]

	for _, d := range data {
		// 将字符串转换为整数
		i, err := strconv.Atoi(d)
		if err != nil {
			fmt.Println("无效的输入，请重新输入")
			continue
		}
		// 检查索引的有效性
		if i > 0 && i <= len(*mirrorURLs) {
			*mirrorURLs = append(*mirrorURLs, (*mirrorURLs)[i-1])
		} else {
			fmt.Println("存在无效的选择，默认测试Tsinghua")
		}
	}
}
