###自动抓取大众点评店铺数据

1.test目录存的是测试例子，各函数测试。

2.主要算法：

     a.布隆过滤算法过滤大数据中重复的url。速度快，有一定重复概率。
     b.模仿ip代理，防止网站屏蔽ip地址。
     c.调用phantomjs抓取网页，解析动态js生成的html。
     d.任务调度，防止开太多进程导致cpu100%，导致电脑不能玩。
     e.gzip解析网页，防止html抓取过来乱码。
     f.行块密度识别正文

该案例中，只用到了d,e两种算法。abc等几个算法都已实现在函数中。

### 打开本地目录

    cd D:\mnt\crawler\src

### 构建测试例子

    go build  mainex/test/g-dianping.go

    go run  mainex/test/g-dianping.go
	
	#测试数据库连接
    go run mainex/test/conn-test.go


###生成ini爬虫规则

    go run mainex/server/ini-xxx.go

### 定时运行抓取数据代码

    go run mainex/g.go

