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

/*CREATE TABLE `crawl_crontab` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `c_name` varchar(50) NOT NULL COMMENT '爬虫名称',
  `start_time` int(6) NOT NULL DEFAULT '0' COMMENT '爬虫生效时间',
  `end_time` int(6) NOT NULL DEFAULT '0' COMMENT '爬虫生效时间',
  `seconds` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，秒',
  `minutes` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，分',
  `hours` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，小时',
  `day` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，天，1-31号',
  `month` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，月，1-12',
  `week` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，星期，0-6',
  `is_open` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1开启状态，0关闭状态',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ctime` (`ctime`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='网站抓取规则表';

*/
//所有定时的爬虫
type Crawl_crontab struct {
	Id         int
	C_name     string
	Start_time string
	End_time   string
	Seconds    string
	Minutes    string
	Hours      string
	Day        string
	Month      string
	Week       string
	Status     int
	Entryid    int
	Is_showlog int
	Is_open    int
}
