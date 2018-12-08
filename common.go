// commonfunction
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

//获取专辑中音频列表信息
//func getList(url string, audiobooks_slice []Audiobooks, downCount int, threads_maximum int) {
func getList(url string, audiobooks_slice []Audiobooks) {

	maxqueue := MaxQueue               //	最大下载个数
	downCount := MaxQueue              //	下载线程最大数
	threads_maximum := Threads_Maximum //	下载线程最大数

	c := colly.NewCollector()
	if Site_Proxy != "" { //	访问网站代理
		rp, err := proxy.RoundRobinProxySwitcher(Site_Proxy) //	("socks5://127.0.0.1:8889", "http://127.0.0.1:8888")
		if err != nil {
			log.Fatal(err)
		}
		c.SetProxyFunc(rp)
	}

	categoryName := ""
	c.OnHTML("#categoryHeader > h1", func(e *colly.HTMLElement) {
		categoryName, _ = DecodeToGBK(e.Text)
	})
	itemNum := 0
	c.OnHTML(".player > a[href]", func(e *colly.HTMLElement) {
		item_name, _ := DecodeToGBK(e.Text)
		var audiobooks Audiobooks
		audiobooks.item_id = itemNum
		audiobooks.save_name = item_name + ".mp3"
		audiobooks.category_name = categoryName
		audiobooks_slice = append(audiobooks_slice, audiobooks)
		itemNum++
	})
	downitemNum := 0
	c.OnHTML(".down > a[href]", func(e *colly.HTMLElement) {
		item_down_website := e.Attr("href")
		item_down_website, _ = DecodeToGBK(item_down_website)
		audiobooks_slice[downitemNum].item_down_website = "http:" + item_down_website
		savePath := audiobooks_slice[downitemNum].category_name + "/" + audiobooks_slice[downitemNum].save_name
		isDown := FileIsDownloaded(savePath)
		if isDown == false {
			if downCount == -1 {
				getDownUrl(downitemNum, audiobooks_slice)
			} else if downCount > 0 {
				getDownUrl(downitemNum, audiobooks_slice)
				downCount--
			}
		}
		downitemNum++
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("访问专辑页面错误: ", err)
	})

	c.OnScraped(func(r *colly.Response) {
		for _, audiobooks := range audiobooks_slice {
			if audiobooks.down_url != "" {
				if maxqueue == -1 {
					start := time.Now()
					_, err := DownAudiobookByFileload(audiobooks, threads_maximum)
					status := "ok"
					if err != nil {
						status = "failed"
					}
					fmt.Printf("下载第%d个音频： %s  Total cost: %s \n", audiobooks.item_id+1, status, time.Since(start))
				} else if maxqueue > 0 {
					start := time.Now()
					_, err := DownAudiobookByFileload(audiobooks, threads_maximum)
					status := "ok"
					if err != nil {
						status = "failed"
					}
					fmt.Printf("下载第%d个音频： %s  Total cost: %s \n", audiobooks.item_id+1, status, time.Since(start))
					maxqueue--
				}
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Referer", url)
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
		r.Headers.Set("User-Agent", GetRandomUserAgent())
	})

	c.Visit(url)
}

//获取音频下载地址
func getDownUrl(key int, audiobooks_slice []Audiobooks) {
	c := colly.NewCollector()
	if Site_Proxy != "" { //	访问网站代理
		rp, err := proxy.RoundRobinProxySwitcher(Site_Proxy) //	("socks5://127.0.0.1:8889", "http://127.0.0.1:8888")
		if err != nil {
			log.Fatal(err)
		}
		c.SetProxyFunc(rp)
	}

	c.OnHTML("#down", func(e *colly.HTMLElement) {
		down_url := e.Attr("href")
		down_url, _ = DecodeToGBK(down_url)
		audiobooks_slice[key].down_url = down_url
		fmt.Printf("获取第%d个音频下载地址： ok \n", key+1)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("获取第%d个音频下载地址： fail \n", key+1)
		log.Println("访问音频下载地址页面错误: ", err)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Referer", audiobooks_slice[key].item_down_website)
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
		r.Headers.Set("User-Agent", GetRandomUserAgent())
	})

	c.Visit(audiobooks_slice[key].item_down_website)
}
