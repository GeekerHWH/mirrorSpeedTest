package tester

import (
	"fmt"
	"net/http"
	"time"
)

// 测试函数，返回从指定镜像站下载 Debian Release 文件的下载速度（MB/s）
func TestMirrorSpeed(url string) (float64, error) {
	startTime := time.Now()

	// 发起 HTTP 请求，下载 Debian ChangeLog 文件
	resp, err := http.Get(fmt.Sprintf("%s/dists/stable/ChangeLog", url))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// 计算下载时间
	downloadTime := time.Since(startTime)

	// 计算下载速度（MB/s）
	downloadSpeed := float64(resp.ContentLength) / (1024 * 1024) / downloadTime.Seconds()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HTTP请求失败，状态码：%d", resp.StatusCode)
	}

	return downloadSpeed, nil
}
