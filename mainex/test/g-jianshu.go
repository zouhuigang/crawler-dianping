/*
SELECT page,count(goodsid) FROM `goods_auto_2` where 1=1 GROUP BY page;

SELECT * FROM `goods_auto_2` where page=0;
*/
package main

import (
	"../lib/change"
	//"../lib/taskqueue"
	"compress/gzip"
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zouhuigang/bloomfilter"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type MainConfig struct {
	*SiteInfo        `ini:"SiteInfo"`
	*HeaderSet       `ini:"HeaderSet"`
	*CrawlerRule     `ini:"CrawlerRule"`
	*DataBaseSetting `ini:"DataBaseSetting"`
}

type CrawlerRule struct {
	Url string
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

var cfg *ini.File

func init() {
	cfg, _ = ini.Load("config/CrawlerRule-jianshu.ini")

}

func main() {
	task := cfg.Section("CrawlerRule").Key("Task").MustInt()
	solid := cfg.Section("CrawlerRule").Key("Solid").MustInt()
	fmt.Println(task, solid)

	//taskqueue.Task.NewTask(task)
	//taskqueue.Task.Soldiers(solid, listFunction)

	listUrl(5)
}
func listFunction(args ...interface{}) {
	fmt.Println("list done", args[0], args[1])
	var page int = 0
	switch v := args[0].(type) {
	case int:
		//fmt.Println("整型", v)
		page = v
		break
	case string:
		//fmt.Println("字符串", v)
		break

	}
	fmt.Println("正在抓取第", page)
	listUrl(page)

}

//读取列表信息
func listUrl(page int) {
	url := cfg.Section("CrawlerRule").Key("Url").Value()
	url = fmt.Sprintf(url, page)

	html := GetHtml(url)
	//save(html, "28.html")
	Parsedocument(string(html), page)
}

//保存在文件中save(html, "28.html")
func save(body []byte, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	//写入图信息
	_, err = f.Write(body)

	defer f.Close()
}

/*解析文档结构树*/
func Parsedocument(html string, page int) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	//doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		fmt.Println("解析文档错误")
	}

	ListP := cfg.Section("CrawlerRule").Key("ListP").Value()
	ListA := cfg.Section("CrawlerRule").Key("ListA").Value()
	doc.Find(ListP).Each(func(i int, s *goquery.Selection) {
		ListAUrl, _ := s.Find(ListA).Attr("href")
		fmt.Println("获取列表成功", ListAUrl)
		ViewInfo(ListAUrl)
	})
}

func GetHtml(url string) []byte {
	proxy := "http://104.224.15.64:8080/"
	//url := "http://www.baidu.com/"
	resp, _ := GetByProxy(url, proxy)
	//fmt.Println(resp)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var body []byte

		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(resp.Body)
			body, _ = ioutil.ReadAll(reader) //dumpGZIP
		default:
			bodyByte, _ := ioutil.ReadAll(resp.Body)
			body = bodyByte
		}
		return body
	}
	//body, _ := ioutil.ReadAll(resp.Body)
	return nil
	//fmt.Println(string(body))
}

// http get by proxy 模仿ip代理
func GetByProxy(url_addr, proxy_addr string) (*http.Response, error) {
	url, _ := url.Parse(url_addr)
	request, _ := http.NewRequest("GET", url.String(), nil)

	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		return nil, err
	}

	headerSet := new(HeaderSet)
	cfg.Section("HeaderSet").MapTo(headerSet)
	map_ := change.Struct.ToMapAddr(headerSet)
	for k, v := range map_ {
		if v != "" {
			request.Header.Set(k, v.(string))
		}
	}

	client := &http.Client{
		Transport: &http.Transport{
		//Proxy: http.ProxyURL(proxy), //ip代理已注释掉
		},
	}
	fmt.Println(proxy)
	return client.Do(request)
}

//保存进数据库
func savelist(goodsname, location, price, url, imageurl string, page int) {
	//db, err := sql.Open("mysql", "用户名:密码@tcp(ip地址:端口号)/数据库名称?charset=utf8")
	UserName := cfg.Section("DataBaseSetting").Key("UserName").Value()
	PassWord := cfg.Section("DataBaseSetting").Key("PassWord").Value()
	HostIp := cfg.Section("DataBaseSetting").Key("HostIp").Value()
	Port := cfg.Section("DataBaseSetting").Key("Port").Value()
	DataBaseName := cfg.Section("DataBaseSetting").Key("DataBaseName").Value()

	con := "%s:%s@tcp(%s:%s)/%s?charset=utf8"
	con = fmt.Sprintf(con, UserName, PassWord, HostIp, Port, DataBaseName)
	db, err := sql.Open("mysql", con)
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into goods_auto_2(goodsname,location,price,url,imageurl,page)values(?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(goodsname, location, price, url, imageurl, page)
	if err != nil {
		log.Println(err)
	}
	//我们可以获得插入的id
	//id, err := rs.LastInsertId()
	//可以获得影响行数
	//affect, err := rs.RowsAffected()

}

type CrawlerContentToDb struct {
	DBTable string
	Title   string
	Content string
}

//解析列表
func ViewInfo(Url string) {
	ViewHost := cfg.Section("CrawlerRule").Key("ViewHost").Value()
	var ViewUrl = ""
	if ViewHost != "" {
		ViewUrl = ViewHost + Url
	} else {
		ViewUrl = Url
	}
	doc1, err := goquery.NewDocument(ViewUrl)
	if err != nil {
		log.Fatal(err)
	}

	crawlerContentToDb := new(CrawlerContentToDb)
	cfg.Section("CrawlerContentToDb").MapTo(crawlerContentToDb)
	map_ := change.Struct.ToMapAddr(crawlerContentToDb)
	for k, v := range map_ {
		if v != "" {
			doc.Find(rule.Author)
		}
	}

	fmt.Println(map_)

}

func bloomfilterUrl(url string) bool {
	set := bloomfilter.NewBloomFilter(1000000, 8192) // elements, false positive rate
	bloomTF1 := set.Check([]byte(url))               // => true

	if bloomTF1 {
		//fmt.Println("2存在")
		return true
	} else {
		//fmt.Println("2不存在")
		set.Add([]byte(url))
		return false
	}

	return false
}
