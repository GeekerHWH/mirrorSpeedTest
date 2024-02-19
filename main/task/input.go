package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SelectMirror(mirrorNames *[]string, mirrorURLs *[]string) ([]string, []string) {
	// 读取用户选择的镜像站序号
	var (
		index       string
		selectNames []string
		selectURLs  []string
	)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入错误(默认测试Tsinghua):", err)
		return []string{"Tsinghua"}, []string{"mirrors.tuna.tsinghua.edu.cn"}
	}

	// 移除换行符
	index = strings.TrimSuffix(line, "\n")

	// 以空格分割输入字符串
	data := strings.Fields(index) //data: [1 3 4]

	for _, d := range data {
		// 将字符串转换为整数i(选择)
		i, err := strconv.Atoi(d)
		if err != nil {
			fmt.Println("无效的输入，请重新输入")
			continue
		}
		// 检查索引的有效性
		if i == 0 {
			//全选
			return *mirrorNames, *mirrorURLs
		} else if i > 0 && i <= len(*mirrorURLs) {
			selectNames = append(selectNames, (*mirrorNames)[i-1])
			selectURLs = append(selectURLs, (*mirrorURLs)[i-1])
		} else {
			fmt.Println("存在无效的选择组合，默认测试Tsinghua")
			return []string{"Tsinghua"}, []string{"mirrors.tuna.tsinghua.edu.cn"}
		}
	}
	// 如果没有有效选择，则默认返回测试Tsinghua
	if len(selectNames) == 0 {
		return []string{"Tsinghua"}, []string{"mirrors.tuna.tsinghua.edu.cn"}
	}

	return selectNames, selectURLs
}

func InputURL() []string {
	var (
		URLs []string
	)

	// 读取用户批量输入的镜像站URL
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		// 移除末尾的换行符
		input = strings.TrimSuffix(input, "\n")
		if input == "" {
			break
		}
		URLs = append(URLs, input)
	}

	//如果没有输入则默认使用清华源
	if len(URLs) == 0 {
		return []string{"mirrors.tuna.tsinghua.edu.cn"}
	}

	return URLs
}
