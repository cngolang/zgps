// main
package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"time"
)

var MaxQueue int //最大下载个数
//var Max_Queue_Downurl int
var Threads_Maximum int //下载线程最大数
var Down_Proxy string   //下载文件代理方式
var Site_Proxy string   //获取下载地址代理方式
func main() {
	start := time.Now()
	flag.IntVar(&MaxQueue, "q", -1, "最大下载个数")
	flag.IntVar(&Threads_Maximum, "t", 10, "下载线程最大数")
	flag.StringVar(&Down_Proxy, "dp", "", "下载文件代理方式，注意开头无socks5,只支持tcp socks5代理,如： 127.0.0.1:8888")
	flag.StringVar(&Site_Proxy, "sp", "", "获取下载地址代理方式，如： socks5://127.0.0.1:8889,http://127.0.0.1:8888")
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("请输入专辑地址")
		return
	} else {
		album_url := os.Args[len(os.Args)-1]
		_, err := url.ParseRequestURI(album_url)
		if err != nil {
			fmt.Println("请输入正确的专辑地址，如： http://shantianfang.zgpingshu.com/1040/#play")
		} else {
			var audiobooks_slice []Audiobooks
			if album_url == "" {
				album_url = "http://shantianfang.zgpingshu.com/1040/#play"
			}
			//			getList(album_url, audiobooks_slice, maxqueue, threads_maximum)
			getList(album_url, audiobooks_slice) //, maxqueue, threads_maximum)
		}
	}
	fmt.Printf("运行共消耗 : %s \n", time.Since(start))
}

func testDemo() {
	//zgps.exe  -q  1  -sp   -dp 127.0.0.1:8888 http://shantianfang.zgpingshu.com/1040/#play
	MaxQueue = 1
	Down_Proxy = "127.0.0.1:8889"
	Site_Proxy = "http://127.0.0.1:8888"

	var audiobooks_slice []Audiobooks
	album_url := "http://shantianfang.zgpingshu.com/1040/#play"
	getList(album_url, audiobooks_slice)
}
