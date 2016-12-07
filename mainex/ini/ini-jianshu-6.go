package main

import (
	"../structPack"
	"fmt"
	"github.com/go-ini/ini"
)

func main() {
	Write()
}

//读取值
func Read() {
	cfg, _ := ini.Load("config/CrawlerRule-jianshu.ini")

	headerSet := new(structPack.HeaderSet)
	cfg.Section("HeaderSet").MapTo(headerSet)
	fmt.Println("list done", headerSet.Host)
}

func Write() {
	cfg := ini.Empty()
	//cfg.Section("dianping").NewKey("name", "value")

	//网站信息
	siteInfo := &structPack.SiteInfo{}
	siteInfo.Name = "简书"
	siteInfo.Url = "www.jianshu.com"
	siteInfo.CateId = 6

	//头部请求信息
	headerSet := &structPack.HeaderSet{}
	headerSet.Host = "www.jianshu.com"
	headerSet.Accept = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	headerSet.Connection = "keep-alive"
	headerSet.Referer = ""
	headerSet.AcceptLanguage = "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3"
	headerSet.AcceptEncoding = "gzip, deflate"
	headerSet.UpgradeInsecureRequests = "1"
	headerSet.CacheControl = "max-age=0"
	headerSet.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome "
	headerSet.Cookie = "CNZZDATA1258679142=899942560-1478565121-https%253A%252F%252Fwww.baidu.com%252F%7C1480751075;" +
		".2.1661956874.1478570111; read_mode=day; signin_redirect=http%3A%2F%2Fwww.jianshu.com%2F; default_font" +
		"=font2; _session_id=QWx6RWNQK2ZHcHE3TjVuSHZObVdaTG1QTUdVQTNGUDBYMHcwaHFLQ3FyQTl0ODJqRTZKRkNPY3pRSWJa" +
		"Mklya0J0RjFmUXJ4RXlzT2tMNnVEL0VNbGpJSEJVU09iVFlrcUxWYTB3N0FaeUxYRDVFeFNFWm91bXNIbERwRXJ0czJaTFdETmpL" +
		"MU9oM1FLVUZZSHZwMFZZZyt3dkZUQmJ1TW1rYm16R0QwZ3JCMmNVeHlRTVFhNnBOT1ZRaU1oN1hYQjZwYkg4SVFUUG1XSWMwUDNk" +
		"bWUrbGZLalBFdnZzSSt4K0UzbVNlbjNjNFVQYmdjN2MvcWRPUml4bEZhMkhMY1ZwOHZYVnRUV2ZxaDhjUGxaRFQrd0ZuQnhqaE1y" +
		"UlVPbzhZN3l0ZWhILzRmTHQ2bTFLM0l1bWhzMjVpV21JQWxZelNCQkRNZmVNUTVPclVRSXJMbTRBPT0tLWUyTlptZFljemVOT0NWYndXY1Q5SFE9PQ" +
		"%3D%3D--35cd60d9eafde576e18e29f19dbc2533d07cb747; _gat=1"
	//抓取规则
	crawlerRule := &structPack.CrawlerRule{}
	crawlerRule.Url = "http://www.jianshu.com/collections/4314/notes?order_by=added_at&page=%d"
	crawlerRule.StartPage = 0
	crawlerRule.Task = 10
	crawlerRule.Solid = 5
	crawlerRule.ListP = "#list-container ul.article-list li"
	crawlerRule.ListA = ".title a"
	crawlerRule.ViewHost = "http://www.jianshu.com"

	//抓取内容进DB,这里面的字段跟数据库要对应,大写的,第一个为表名字全称
	crawlerContentToDb := &structPack.CrawlerContentToDb{}
	crawlerContentToDb.Title = ".article .preview  h1.title"
	crawlerContentToDb.Content = ".article .preview .show-content"

	//主配置
	mainConfig := &structPack.MainConfig{}

	mainConfig.SiteInfo = siteInfo
	mainConfig.HeaderSet = headerSet
	mainConfig.CrawlerRule = crawlerRule
	mainConfig.CrawlerContentToDb = crawlerContentToDb
	ini.ReflectFrom(cfg, mainConfig)
	err := cfg.SaveTo("config/CrawlerRule-jianshu-6.ini")
	fmt.Println(err)
	//cfg.SaveToIndent("my.ini", "\t")
}
