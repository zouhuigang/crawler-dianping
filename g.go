/*
SELECT page,count(goodsid) FROM `goods_auto_2` where 1=1 GROUP BY page;

SELECT * FROM `goods_auto_2` where page=0;
*/
package main

import (
	"compress/gzip"
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zouhuigang/bloomfilter"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"taskqueue"
)

func main() {
	taskqueue.Task.NewTask(50)
	taskqueue.Task.Soldiers(10, list)
	//listUrl(0)
}
func list(args ...interface{}) {
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
	//page = page + 28
	pagestr := strconv.Itoa(page)
	//url := "http://www.dianping.com/search/keyword/1/0_%E5%85%89%E5%A4%A7%E4%BC%9A%E5%B1%95%E4%B8%AD%E5%BF%83/p" + pagestr
	url := "http://www.dianping.com/search/keyword/1/0_%E5%85%89%E5%A4%A7%E4%BC%9A%E5%B1%95%E4%B8%AD%E5%BF%83/p" + pagestr
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
	//html, _ := doc.Html()
	//fmt.Println("doc", html)
	doc.Find("#shop-all-list ul li").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find(".tit a").Attr("href")
		imageurl, _ := s.Find("div.pic a img").Attr("data-src")
		goodsname := s.Find(".tit a h4").Text()
		price := s.Find(".comment a.mean-price b").Text()
		location := s.Find(".txt .tag-addr span.addr").Text()
		fmt.Println("list链接:\n", i, goodsname, price, imageurl, url, location)
		savelist(goodsname, location, price, url, imageurl, page)
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
	request.Header.Set("Host", "www.dianping.com")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Encoding", "gzip, deflate")
	request.Header.Set("Upgrade-Insecure-Requests", "1")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.59 QQ/8.5.18600.201 Safari/537.36")
	//request.Header.Set("Referer", url.String())
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")

	request.Header.Set("Cookie", "cy=1; cye=shanghai; _hc.v=3fe71672-ce3d-8ec0-0336-1c0b76ca1a52.1480148770; __utma=1.1733053444.1480148770.1480148770.1480148770.1; __utmz=1.1480148770.1.1.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; s_ViewType=10; aburl=1; JSESSIONID=E1285F551A896574A17C3BE5C0633C71; PHOENIX_ID=0a01070e-158adc79509-9d92de; cityid=1; pvhistory=\"5ZWG5oi3Pjo8L3NlYXJjaC9rZXl3b3JkLzEvMF8lRTUlODUlODklRTUlQTQlQTclRTQlQkMlOUElRTUlQjElOTUlRTQlQjglQUQlRTUlQkYlODMvcDI1Pjo8MTQ4MDM5MTk0ODQ0MF1fW+i/lOWbnj46PC9zdWdnZXN0L2dldEpzb25EYXRhP189MTQ4MDM5MTk1OTgzNyZjYWxsYmFjaz1qc29ucDE0ODAzOTE5NTk4NzA+OjwxNDgwMzkxOTQ5MzY1XV9b\"; m_flash2=1; selfAB=b; testName=test2")

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
	db, err := sql.Open("mysql", "用户名:密码@tcp(ip地址:端口号)/数据库名称?charset=utf8")
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

//解析列表
func viewInfo(Url string) {
	doc1, err := goquery.NewDocument("http://www.dianping.com/" + Url)
	if err != nil {
		log.Fatal(err)
	}
	doc1.Find("#basic-info").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band, _ := s.Html()
		//title := s.Find(".tit a h4").Text()
		fmt.Println("list链接:\n", band)
	})

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
