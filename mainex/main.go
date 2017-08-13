package main

import (
	"math/rand"
	//"path/filepath"
	//"github.com/robfig/cron"
	cron "gopkg.in/robfig/cron.v2"
	"log"
	"mainex/structPack"
	//"mainex/server"
	. "db"
	"fmt"
	"time"
)

var (
	c      *cron.Cron                  //定义定时任务
	mainid cron.EntryID                //主定时任务，不能移除掉
	clist  []*structPack.Crawl_crontab //所有的开放的定时任务
)

func init() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
	//测试数据库是否连通
	/*	err := db.TestDB()
		if err != nil {
			log.Println("数据库连接出错,请检查配置文件:", err)
			os.Exit(0)
		}*/
}

//得到所有开放的爬虫
func getCronList() ([]*structPack.Crawl_crontab, error) {
	beans := make([]*structPack.Crawl_crontab, 0)
	err := MasterDB.
		Where("is_open=1").Find(&beans)
	return beans, err
}

//根据爬虫id，得到最新的信息
func getNewestInfoById(id int) *structPack.Crawl_crontab {
	info := new(structPack.Crawl_crontab)
	for _, v := range clist {
		if id == v.Id {
			info = v
		}
	}
	return info
}

//检测当前爬虫是否已完成一轮爬取行为,一般存入redis
//func getRefreshStatus(id int) (status int) {
//	return 0
//}

//设置为0
func setRefreshStart(id int) {
	beans := &structPack.Crawl_crontab{}
	beans.Status = 0
	MasterDB.Where("id=?", id).Cols("status").Update(beans)
}

func setEntryID(id int, entryid int) {
	beans := &structPack.Crawl_crontab{}
	beans.Entryid = entryid
	MasterDB.Where("id=?", id).Cols("entryid").Update(beans)
}

func setRefreshDone(id int) {
	beans := &structPack.Crawl_crontab{}
	beans.Status = 1
	MasterDB.Where("id=?", id).Cols("status").Update(beans)
}

// 执行爬虫npc刷新
func doNpcRefresh(id int) bool {
	crinfo := getNewestInfoById(id)
	if crinfo.Status == 0 { //如果任务没有完成，则不能开始新的一轮
		fmt.Printf("爬虫%d刷新未完成，不进行新一轮刷新\n", id)
		return false
	}

	//判断之前的程序是否在运行中,问题，可能刷新配置的时间刚好和任务开始执行的时间交叉，出现Status=1之后，到这里又变成Status=0

	//start := time.Now()
	//开始爬取任务
	spec := fmt.Sprintf("%s %s %s %s %s %s", crinfo.Seconds, crinfo.Minutes, crinfo.Hours, crinfo.Day, crinfo.Month, crinfo.Week)
	// 定时增量
	i := 0
	//spec := "*/5 * * * * ?"
	entryID, _ := c.AddFunc(spec, func() {
		i++
		go doCrawl(crinfo.Id, i)

	})
	setEntryID(id, int(entryID))
	//end

	return true
}

//移掉不正常的任务
func removeNotNormalRuning() {

	for _, entry := range c.Entries() {
		if entry.ID == mainid {
			continue
		}
		var isnotNormalRuning bool = true
		for _, v := range clist {
			ecronId := cron.EntryID(v.Entryid)
			ecron := c.Entry(ecronId) //正常任务
			if entry.ID == ecronId && ecron.Valid() {
				isnotNormalRuning = false
			}
			//log.Printf("removeNotNormalRuning====:%v,%v\n", v.Entryid, entry.ID)
		}

		if isnotNormalRuning {
			c.Stop()
			c.Remove(entry.ID)
			log.Printf("正在删除定时任务:%v\n", entry.ID)
			c.Start()
		}

	}

}

func main() {

	c = cron.New()
	spec := "*/10 * * * * *"
	mainid, _ = c.AddFunc(spec, func() {
		//每秒检测下数据库配置是否更新了
		log.Printf("正在读取最新配置.....\n")
		var err error
		clist, err = getCronList()
		if err != nil {
			log.Println("读取所有的爬虫列表失败", err)
		}

		if len(clist) <= 0 {
			log.Println("等待新任务.....", err)
		}

		//fmt.Println("====================检测NPC刷新====================")
		//移掉不正常的任务
		removeNotNormalRuning()

		for _, v := range clist {
			ps := cron.EntryID(v.Entryid)
			entry := c.Entry(ps)
			//log.Printf("cron init:%v,%v\n", entry.Valid(), v.Is_open)
			if entry.Valid() { //正在运行
				continue
			}

			is_finished := doNpcRefresh(v.Id)
			if !is_finished { //未完成任务，则不刷新,不开始新的任务
				continue
			}

		}

	})

	c.Start()

	select {}

}

//开始抓取
func doCrawl(id int, i int) {
	crinfo := getNewestInfoById(id)
	setRefreshStart(id)
	//server.Ginit("config/" + v)
	if crinfo.Is_showlog == 1 {
		log.Printf("爬虫[%d-%s]正在第%d次抓取...", id, crinfo.C_name, i)
	}
	setRefreshDone(id)
}
