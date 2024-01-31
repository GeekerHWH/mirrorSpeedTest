package task

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Test(mirrorNames []string, mirrorURLs []string) {
	if len(mirrorURLs) == 4 {
		downloadSpeed, err := TestMirrorSpeed(mirrorURLs[1])

		// 处理测试结果
		if err != nil {
			fmt.Printf("测试失败：%s\n", err)
		} else {
			fmt.Printf("从 %s 下载 Debian Release 文件的速度：%f MB/s\n", mirrorNames[1], downloadSpeed)
		}
	} else {
		for i := len(mirrorNames); i < len(mirrorURLs); i++ {
			downloadSpeed, err := TestMirrorSpeed(mirrorURLs[i])

			// 处理测试结果
			if err != nil {
				fmt.Printf("测试失败：%s\n", err)
			} else {
				fmt.Printf("从 %s 下载 Debian Release 文件的速度：%f MB/s\n", mirrorURLs[i], downloadSpeed)
			}
		}
	}
}

// 测试函数，返回从指定镜像站下载 Debian Release 文件的下载速度（MB/s）
func TestMirrorSpeed(url string) (float64, error) {
	startTime := time.Now()

	// 发起 HTTP 请求，下载 Debian ChangeLog 文件
	resp, err := http.Get(fmt.Sprintf("%s/dists/stable/ChangeLog", url))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HTTP请求失败，状态码：%d", resp.StatusCode)
	}

	// 文件大小为3MB，使用3MB的缓冲区
	bufferSize := 3 * 1024 * 1024
	buffer := make([]byte, bufferSize)

	// 使用缓冲区，用downloadedBytes累积已下载的字节数n
	var downloadedBytes int64
	for {
		n, err := resp.Body.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, fmt.Errorf("读取响应体失败：%v", err)
		}
		//累积操作
		downloadedBytes += int64(n)
	}

	// 计算下载时间
	downloadTime := time.Since(startTime)

	// 计算下载速度（MB/s）
	downloadSpeed := float64(downloadedBytes) / (1024 * 1024) / downloadTime.Seconds()

	return downloadSpeed, nil
}
