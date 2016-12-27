###修改ini配置文件

    cd D:\mnt\crawler\src\

然后再编辑ini.bat中添加:

    go run mainex/ini/ini-jianshu-23.go	

运行ini.bat生成爬虫配置文件

或者自己生成ini文件

    go run mainex/ini/ini-jianshu-24.go	


生成完成之后，进入config配置CrawlerRule-all-auto.ini中自定义你需要抓取的文件


最简洁的配置：

直接修改config中的ini文件，然后在CrawlerRule-all-auto.ini把文件添加进去