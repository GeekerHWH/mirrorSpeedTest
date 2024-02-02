package task

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

func IsDebian12() bool {
	// 检查操作系统是否为Linux
	if runtime.GOOS != "linux" {
		return false
	}

	// 读取/etc/os-release文件以获取操作系统信息
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return false
	}
	defer file.Close()

	// 逐行读取文件
	var ID, VERSION string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.Trim(strings.TrimSpace(parts[1]), "\"")
			switch key {
			case "ID":
				ID = value
			case "VERSION_ID":
				VERSION = value
			}
		}
	}

	// 检查操作系统是否为Debian 12
	return ID == "debian" && VERSION == "12"
}

func IsUbuntu2204() bool {
	// 检查操作系统是否为Linux
	if runtime.GOOS != "linux" {
		return false
	}

	// 读取/etc/os-release文件以获取操作系统信息
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return false
	}
	defer file.Close()

	// 逐行读取文件
	var ID, VERSION string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.Trim(strings.TrimSpace(parts[1]), "\"")
			switch key {
			case "ID":
				ID = value
			case "VERSION_ID":
				VERSION = value
			}
		}
	}

	// 检查操作系统是否为Ubuntu 22.04
	return ID == "ubuntu" && VERSION == "22.04"
}
