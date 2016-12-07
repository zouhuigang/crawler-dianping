package server

import (
	"compress/gzip"
	. "db"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-ini/ini"
	"github.com/zouhuigang/bloomfilter"
	"io/ioutil"
	"log"
	"mainex/lib/change"
	"mainex/lib/public"
	"mainex/lib/taskqueue"
	"mainex/structPack"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var cfg *ini.File
var bloom *bloomfilter.BloomFilter

//func init() {
//	cfg, _ = ini.Load("config/CrawlerRule-jianshu.ini")
//	bloom = bloomfilter.NewBloomFilter(1000000, 8192)
//}

func Ginit(crawlerRuleConf string) {
	ini.DefaultHeader = true
	cfg, _ = ini.Load(crawlerRuleConf)
	bloom = bloomfilter.NewBloomFilter(1000000, 8192)

	log.Println("定时任务正在执行")

	gmain(crawlerRuleConf)
}

//修改startpage
func UpdateStartPage(crawlerRuleConf string, page int) {
	// 通过Itoa方法转换
	str1 := strconv.Itoa(page)
	cfg.Section("CrawlerRule").Key("StartPage").SetValue(str1)
	cfg.SaveTo(crawlerRuleConf)
}

func gmain(crawlerRuleConf string) {
	task := cfg.Section("CrawlerRule").Key("Task").MustInt()
	solid := cfg.Section("CrawlerRule").Key("Solid").MustInt()
	startPage := cfg.Section("CrawlerRule").Key("StartPage").MustInt(0)
	//fmt.Println(task, solid)
	var upage int = task + startPage
	//修改
	defer UpdateStartPage(crawlerRuleConf, upage)

	taskqueue.Task.NewTask(task)
	taskqueue.Task.Soldiers(solid, listFunction)

	//listUrl(5)
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
	//fmt.Println("正在抓取第", page)
	listUrl(page)

}

//读取列表信息
func listUrl(page int) {
	startPage := cfg.Section("CrawlerRule").Key("StartPage").MustInt(0)
	page = page + startPage
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
		ViewInfo(ListAUrl, page)
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

	headerSet := new(structPack.HeaderSet)
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
	fmt.Println("代理地址未使用:", proxy)
	return client.Do(request)
}

//解析列表
func ViewInfo(Url string, pages int) error {
	ViewHost := cfg.Section("CrawlerRule").Key("ViewHost").Value()
	var ViewUrl = ""
	if ViewHost != "" {
		ViewUrl = ViewHost + Url
	} else {
		ViewUrl = Url
	}
	if bloom.Check([]byte(ViewUrl)) {
		fmt.Println("该链接已爬取过" + ViewUrl)
		return nil
	} else {
		bloom.Add([]byte(ViewUrl))
	}

	doc, err := goquery.NewDocument(ViewUrl)
	if err != nil {
		log.Fatal(err)
	}

	rule := cfg.Section("CrawlerContentToDb").KeysHash()
	tmpArticle := &structPack.Anote{}

	//解析文章
	tmpArticle.Content, err = doc.Find(rule["Content"]).Html()
	if err != nil {
		log.Fatal("文章内容解析失败")
	}
	tmpArticle.Title = doc.Find(rule["Title"]).Text()
	tmpArticle.Is_auto = 1
	tmpArticle.Is_open = 1
	tmpArticle.Cateid = 5
	tmpArticle.Is_auto_page = pages
	tmpArticle.Is_auto_source = cfg.Section("SiteInfo").Key("Name").Value()
	tmpArticle.Url = ViewUrl
	//save([]byte(tmpArticle.Content), tmpArticle.Title+".html")
	//封面图
	cover := public.Publics.CoverGirl(tmpArticle.Content, 3) //[]string
	lang, _ := json.Marshal(cover)                           //转json str
	var scover string
	if len(cover) != 0 {
		scover = string(lang)
	} else {
		scover = ""
	}
	tmpArticle.Cover = scover
	tmpArticle.Newslist_tpl = len(cover)

	//解析文章end

	_, err = MasterDB.Insert(tmpArticle)
	if err != nil {
		log.Fatal("插入数据库失败")
		return nil
	}

	return nil

}

func bloomfilterUrltest(url string) bool {
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
