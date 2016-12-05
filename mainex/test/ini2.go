package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

func main() {
	Write()
}

//读取值
func Read() {
	cfg, _ := ini.Load("config/CrawlerRuleConf.ini")

	headerSet := new(HeaderSet)
	cfg.Section("HeaderSet").MapTo(headerSet)
	fmt.Println("list done", headerSet.Host)
}

type MainConfig struct {
	*SiteInfo        `ini:"SiteInfo"`
	*HeaderSet       `ini:"HeaderSet"`
	*CrawlerRule     `ini:"CrawlerRule"`
	*DataBaseSetting `ini:"DataBaseSetting"`
}

type CrawlerRule struct {
	Url   string
	Task  int
	Solid int
}

type DataBaseSetting struct {
	UserName     string
	PassWord     string
	HostIp       string
	Port         int
	DataBaseName string
}

type SiteInfo struct {
	Name string
	Url  string
}

type HeaderSet struct {
	Host                    string `ini:"Host"`
	Accept                  string `ini:"Accept"`
	Connection              string `ini:"Connection"`
	Referer                 string `ini:"Referer"`
	AcceptLanguage          string `ini:"Accept-Language"`
	AcceptEncoding          string `ini:"Accept-Encoding"`
	UpgradeInsecureRequests string `ini:"Upgrade-InsecureRequests"`
	CacheControl            string `ini:"Cache-Control"`
	UserAgent               string `ini:"User-Agent"`
	Cookie                  string `ini:"Cookie"`
}

//request.Header.Set("Host", "www.dianping.com")
//request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
//request.Header.Set("Accept-Encoding", "gzip, deflate")
//request.Header.Set("Upgrade-Insecure-Requests", "1")
//	request.Header.Set("Cache-Control", "max-age=0")
//	request.Header.Set("Connection", "keep-alive")
//	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.59 QQ/8.5.18600.201 Safari/537.36")
//request.Header.Set("Referer", url.String())
//	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")

//	request.Header.Set("Cookie", "cy=1; cye=shanghai; _hc.v=3fe71672-ce3d-8ec0-0336-1c0b76ca1a52.1480148770; __utma=1.1733053444.1480148770.1480148770.1480148770.1; __utmz=1.1480148770.1.1.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; s_ViewType=10; aburl=1; JSESSIONID=E1285F551A896574A17C3BE5C0633C71; PHOENIX_ID=0a01070e-158adc79509-9d92de; cityid=1; pvhistory=\"5ZWG5oi3Pjo8L3NlYXJjaC9rZXl3b3JkLzEvMF8lRTUlODUlODklRTUlQTQlQTclRTQlQkMlOUElRTUlQjElOTUlRTQlQjglQUQlRTUlQkYlODMvcDI1Pjo8MTQ4MDM5MTk0ODQ0MF1fW+i/lOWbnj46PC9zdWdnZXN0L2dldEpzb25EYXRhP189MTQ4MDM5MTk1OTgzNyZjYWxsYmFjaz1qc29ucDE0ODAzOTE5NTk4NzA+OjwxNDgwMzkxOTQ5MzY1XV9b\"; m_flash2=1; selfAB=b; testName=test2")

func Write() {
	cfg := ini.Empty()
	//cfg.Section("dianping").NewKey("name", "value")

	//网站信息
	siteInfo := &SiteInfo{}
	siteInfo.Name = "大众点评"
	siteInfo.Url = "www.dianping.com"

	//头部请求信息
	headerSet := &HeaderSet{}
	headerSet.Host = "www.dianping.com"
	headerSet.Accept = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	headerSet.Connection = "keep-alive"
	headerSet.Referer = ""
	headerSet.AcceptLanguage = "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3"
	headerSet.AcceptEncoding = "gzip, deflate"
	headerSet.UpgradeInsecureRequests = "1"
	headerSet.CacheControl = "max-age=0"
	headerSet.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome"
	headerSet.Cookie = "cy=1; cye=shanghai; _hc.v=3fe71672-ce3d-8ec0-0336-1c0b76ca1a52.1480148770; __utma=1.1733053444.1480148770.1480148770.1480148770.1; __utmz=1.1480148770.1.1.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; s_ViewType=10; aburl=1; JSESSIONID=E1285F551A896574A17C3BE5C0633C71; PHOENIX_ID=0a01070e-158adc79509-9d92de; cityid=1; pvhistory='5ZWG5oi3Pjo8L3NlYXJjaC9rZXl3b3JkLzEvMF8lRTUlODUlODklRTUlQTQlQTclRTQlQkMlOUElRTUlQjElOTUlRTQlQjglQUQlRTUlQkYlODMvcDI1Pjo8MTQ4MDM5MTk0ODQ0MF1fW+i/lOWbnj46PC9zdWdnZXN0L2dldEpzb25EYXRhP189MTQ4MDM5MTk1OTgzNyZjYWxsYmFjaz1qc29ucDE0ODAzOTE5NTk4NzA+OjwxNDgwMzkxOTQ5MzY1XV9b'; m_flash2=1; selfAB=b; testName=test2"

	//抓取规则
	crawlerRule := &CrawlerRule{}
	crawlerRule.Url = "http://www.dianping.com/search/keyword/1/10_光大会展中心/p%d"
	crawlerRule.Task = 10
	crawlerRule.Solid = 5

	//数据库连接信息
	dataBaseSetting := &DataBaseSetting{}
	dataBaseSetting.UserName = "root"
	dataBaseSetting.PassWord = "TYwy2016720"
	dataBaseSetting.HostIp = "139.196.16.67"
	dataBaseSetting.Port = 3306
	dataBaseSetting.DataBaseName = "whateat"

	//主配置
	mainConfig := &MainConfig{}

	mainConfig.SiteInfo = siteInfo
	mainConfig.HeaderSet = headerSet
	mainConfig.CrawlerRule = crawlerRule
	mainConfig.DataBaseSetting = dataBaseSetting
	ini.ReflectFrom(cfg, mainConfig)
	err := cfg.SaveTo("config/CrawlerRuleConf.ini")
	fmt.Println(err)
	//cfg.SaveToIndent("my.ini", "\t")
}
