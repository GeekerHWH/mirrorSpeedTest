package task

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sort"
	"sync"
	"text/tabwriter"
	"time"
)

const initNumURLs = 10

type Mirror struct {
	Name  string
	URL   string
	Speed float64
	Ping  time.Duration
}

func Test(mirrorNames []string, mirrorURLs []string) {
	if len(mirrorURLs) == initNumURLs {
		TestMirrorSpeed(mirrorURLs[1])

	} else {
		var waitGroup sync.WaitGroup
		waitGroup.Add(len(mirrorURLs) - initNumURLs)

		// multi-threads speed testing
		// 多线程测速
		var Mirrors []Mirror
		var mu sync.Mutex // 用于保护Mirrors切片的互斥锁
		for i := initNumURLs; i < len(mirrorURLs); i++ {
			go func(index int, url string) {
				defer waitGroup.Done()

				speed := TestMirrorSpeed(url)
				ping := TCPPing(url)

				mu.Lock()
				Mirrors = append(Mirrors, Mirror{Name: mirrorNames[index], URL: url, Speed: speed, Ping: ping})
				mu.Unlock()
			}(i, mirrorURLs[i])
		}
		waitGroup.Wait()

		// 按速度带宽从大到小排序
		sort.Slice(Mirrors, func(i, j int) bool {
			return Mirrors[i].Speed > Mirrors[j].Speed
		})

		// 创建一个新的 tabwriter.Writer
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)

		// 打印表头
		fmt.Fprintln(w, "Name\tSpeed\tDelay")

		// 打印数据
		for i := range Mirrors {
			// if i == 0 {
			// 	fmt.Fprintf(w, "\x1b[32m%s\t%.2fMB/s\t%v\x1b[0m\n", Mirrors[i].Name, Mirrors[i].Speed, Mirrors[i].Ping)
			// 	continue
			// }
			fmt.Fprintf(w, "%s\t%.2fMB/s\t%v\n", Mirrors[i].Name, Mirrors[i].Speed, Mirrors[i].Ping)
		}
		// 刷新并关闭 tabwriter.Writer
		w.Flush()
	}
}

// Will print the speed of downloading the Debian Release file from the specified mirror.
// 测试函数，返回从指定镜像站下载 Debian Release 文件的下载速度（MB/s）
func TestMirrorSpeed(url string) float64 {
	// Start clocking
	// 开始计时
	startTime := time.Now()

	// 发起 HTTP 请求，下载 Debian ChangeLog 文件
	resp, err := http.Get(fmt.Sprintf("http://%s/debian/dists/stable/ChangeLog", url))
	if err != nil {
		fmt.Printf("HTTP请求失败：%s\n", err)
		return 0
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s HTTP请求失败，状态码：%d\n", url, resp.StatusCode)
		return 0
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
			return 0
		}
		//累积操作
		downloadedBytes += int64(n)
	}

	// 计算下载时间
	downloadTime := time.Since(startTime)

	// 计算下载速度（MB/s）
	downloadSpeed := float64(downloadedBytes) / (1024 * 1024) / downloadTime.Seconds()

	return downloadSpeed
	// 处理测试结果
	// fmt.Printf("从 %s 下载 Debian Release 文件的速度：%f MB/s\n", url, downloadSpeed)
}

func TCPPing(url string) time.Duration {
	start := time.Now()

	conn, err := net.Dial("tcp", url+":80")
	if err != nil {
		fmt.Printf("TCP Ping %s: %s\n", url, err)
	}
	defer conn.Close()

	delay := time.Since(start)
	// fmt.Printf("TCP Ping %s: %s\n", url, delay)
	return delay
}
