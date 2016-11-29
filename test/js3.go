package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/nladuo/go-phantomjs-fetcher"
	"strings"
)

func main() {
	//create a fetcher which seems to a httpClient
	fetcher, err := phantomjs.NewFetcher(2016, nil)
	defer fetcher.ShutDownPhantomJSServer()
	if err != nil {
		panic(err)
	}
	//inject the javascript you want to run in the webpage just like in chrome console.
	js_script := "function(){document.getElementById('site-nav').click();}"
	//run the injected js_script at the end of loading html
	js_run_at := phantomjs.RUN_AT_DOC_END
	//send httpGet request with injected js
	resp, err := fetcher.GetWithJS("http://www.dianping.com/search/keyword/1/0_%E5%85%89%E5%A4%A7%E4%BC%9A%E5%B1%95%E4%B8%AD%E5%BF%83/p30", js_script, js_run_at)
	if err != nil {
		panic(err)
	}

	//select search results by goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Content))
	if err != nil {
		panic(err)
	}
	fmt.Println("Results:")
	doc.Find("#shop-all-list ul li").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find(".tit a").Attr("href")
		imageurl, _ := s.Find("div.pic a img").Attr("data-src")
		goodsname := s.Find(".tit a h4").Text()
		price := s.Find(".comment a.mean-price b").Text()
		location := s.Find(".txt .tag-addr span.addr").Text()
		fmt.Println("list链接:\n", i, goodsname, price, imageurl, url, location)
		//savelist(goodsname, location, price, url, imageurl, page)
	})
}
