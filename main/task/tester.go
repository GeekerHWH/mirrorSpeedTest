package task

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

func Test(mirrorNames []string, mirrorURLs []string) {
	if len(mirrorURLs) == 5 {
		TestMirrorSpeed(mirrorURLs[1])

	} else {
		var waitGroup sync.WaitGroup
		waitGroup.Add(len(mirrorURLs) - len(mirrorNames))

		// multi-threads speed testing
		// 多线程测速
		for i := len(mirrorNames); i < len(mirrorURLs); i++ {
			go func(url string) {
				TestMirrorSpeed(url)
				waitGroup.Done()
			}(mirrorURLs[i])
		}
		waitGroup.Wait()
	}
}

// Will print the speed of downloading the Debian Release file from the specified mirror.
// 测试函数，返回从指定镜像站下载 Debian Release 文件的下载速度（MB/s）
func TestMirrorSpeed(url string) {
	// Start clocking
	// 开始计时
	startTime := time.Now()

	// 发起 HTTP 请求，下载 Debian ChangeLog 文件
	resp, err := http.Get(fmt.Sprintf("%s/dists/stable/ChangeLog", url))
	if err != nil {
		fmt.Printf("HTTP请求失败：%s\n", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP请求失败，状态码：%d", resp.StatusCode)
		return
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
			fmt.Printf("读取响应体失败：%v", err)
			return
		}
		//累积操作
		downloadedBytes += int64(n)
	}

	// 计算下载时间
	downloadTime := time.Since(startTime)

	// 计算下载速度（MB/s）
	downloadSpeed := float64(downloadedBytes) / (1024 * 1024) / downloadTime.Seconds()

	// 处理测试结果
	if err != nil {
		fmt.Printf("测试失败：%s\n", err)
	} else {
		fmt.Printf("从 %s 下载 Debian Release 文件的速度：%f MB/s\n", url, downloadSpeed)
	}
}

func TCPPing(url string) {
	start := time.Now()

	conn, err := net.Dial("tcp", url+":80")
	if err != nil {
		fmt.Printf("TCP Ping %s: %s\n", url, err)
	}
	defer conn.Close()

	delay := time.Since(start)
	fmt.Printf("TCP Ping %s: %s\n", url, delay)
}
