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
	cfg := ini.Empty()
	cfg.Section("author").Key("GITHUB").String() // how about continuation lines?
	cfg.Section("package").Key("FULL_NAME").String()
	cfg.Section("dianping").Key("name").String()
}

func Write() {
	cfg := ini.Empty()
	cfg.Section("dianping").NewKey("name", "value")
	cfg.Section("dianping.HeaderSet").NewKey("host", "www.dianping.com")
	err := cfg.SaveTo("config/CrawlerRuleConf.ini")
	fmt.Println(err)
	//cfg.SaveToIndent("my.ini", "\t")
}
