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
	cfg, _ := ini.Load("config/CrawlerRule-oschina.ini")

	headerSet := new(structPack.HeaderSet)
	cfg.Section("HeaderSet").MapTo(headerSet)
	fmt.Println("list done", headerSet.Host)
}

func Write() {
	cfg := ini.Empty()
	//cfg.Section("dianping").NewKey("name", "value")

	//网站信息
	siteInfo := &structPack.SiteInfo{}
	siteInfo.Name = "开源中国"
	siteInfo.Url = "www.oschina.net"
	siteInfo.CateId = 5

	//头部请求信息
	headerSet := &structPack.HeaderSet{}
	headerSet.Host = "www.oschina.net"
	headerSet.Accept = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	headerSet.Connection = "keep-alive"
	headerSet.Referer = "http://www.oschina.net/search?q=go&scope=blog&onlytitle=1&sort_by_time=1"
	headerSet.AcceptLanguage = "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3"
	headerSet.AcceptEncoding = "gzip, deflate"
	headerSet.UpgradeInsecureRequests = "1"
	headerSet.CacheControl = "max-age=0"
	headerSet.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0 "
	headerSet.Cookie = "_user_behavior_=f4ad7712-26f5-49f3-bb7a-2ee58cee0bd5; Hm_lvt_a411c4d1664dd70048ee98afe7b28f0b=1480927398" +
		",1481002344,1481010967,1481081978; Hm_lpvt_a411c4d1664dd70048ee98afe7b28f0b=1481081978"
	//抓取规则
	crawlerRule := &structPack.CrawlerRule{}
	crawlerRule.Url = "http://www.oschina.net/search?scope=blog&q=go&p=%d"
	crawlerRule.StartPage = 0
	crawlerRule.Task = 10
	crawlerRule.Solid = 5
	crawlerRule.ListP = "ul #results li.obj_type_3"
	crawlerRule.ListA = "h3 a"
	crawlerRule.ViewHost = ""

	//抓取内容进DB,这里面的字段跟数据库要对应,大写的,第一个为表名字全称
	crawlerContentToDb := &structPack.CrawlerContentToDb{}
	crawlerContentToDb.Title = ".blog-content .blog-heading .title"
	crawlerContentToDb.Content = ".blog-content .blogBody"

	//主配置
	mainConfig := &structPack.MainConfig{}

	mainConfig.SiteInfo = siteInfo
	mainConfig.HeaderSet = headerSet
	mainConfig.CrawlerRule = crawlerRule
	mainConfig.CrawlerContentToDb = crawlerContentToDb
	ini.ReflectFrom(cfg, mainConfig)
	err := cfg.SaveTo("config/CrawlerRule-oschina-5.ini")
	fmt.Println(err)
	//cfg.SaveToIndent("my.ini", "\t")
}
