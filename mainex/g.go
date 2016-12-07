package main

import (
	"flag"
	"math/rand"
	//"path/filepath"
	"github.com/go-ini/ini"
	"github.com/robfig/cron"
	"log"
	"mainex/server"
	"time"
)

func init() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {

	var (
		needAll           bool
		crawlConfFilename string
		whichSite         string
	)
	flag.BoolVar(&needAll, "all", false, "是否需要全量抓取，默认否")
	flag.StringVar(&crawlConfFilename, "config", "config/CrawlerRule-all-auto.ini", "自动抓取配置文件")
	flag.StringVar(&whichSite, "site", "", "抓取配置中哪个站点（空表示所有配置站点）")
	//flag.XxxVar(a,b,c,d) 将命令行参数绑定到变量上。a变量名称,b命令行赋值时的元素如-site sss 就是将sss赋值给site+
	//c是默认值,d是备注说明
	flag.Parse()
	go autocrawl(needAll, crawlConfFilename, whichSite)

	select {}
}

func autocrawl(needAll bool, crawlConfFile string, whichSite string) {
	log.Println(crawlConfFile)
	cfg, err := ini.Load(crawlConfFile)
	if err != nil {
		log.Println("配置文件加载失败", err)
	}

	rule := cfg.Section("AllSite").KeysHash()

	// 定时增量
	i := 0
	c := cron.New()
	spec := "* * */1 * * ?"
	c.AddFunc(spec, func() {
		i++
		for k, v := range rule {
			go doCrawl(k, v, i)
		}

	})
	c.Start()
}

//开始抓取
func doCrawl(k string, v string, i int) {
	server.Ginit("config/" + v)
	log.Println("cron running:", i, k, v)
}
