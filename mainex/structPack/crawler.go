package structPack

type MainConfig struct {
	*SiteInfo           `ini:"SiteInfo"`
	*HeaderSet          `ini:"HeaderSet"`
	*CrawlerRule        `ini:"CrawlerRule"`
	*CrawlerContentToDb `ini:"CrawlerContentToDb"`
}

type CrawlerRule struct {
	Url       string
	StartPage int //开始爬取页面
	Task      int
	Solid     int
	ListP     string //容纳列表的元素
	ListA     string //在容器列表中a的位置
	ViewHost  string //防止得到简短的地址
}

type CrawlerContentToDb struct {
	Title   string
	Content string
}

type SiteInfo struct {
	Name   string
	Url    string
	CateId int //分类的id，写死
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
