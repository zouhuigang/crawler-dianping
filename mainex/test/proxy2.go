package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// http get by proxy
func GetByProxy(url_addr, proxy_addr string) (*http.Response, error) {
	request, _ := http.NewRequest("GET", url_addr, nil)
	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	return client.Do(request)
}

func main() {
	proxy := "http://177.42.184.153:8080/"
	url := "http://www.baidu.com/"
	resp, _ := GetByProxy(url, proxy)
	fmt.Println(resp)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
