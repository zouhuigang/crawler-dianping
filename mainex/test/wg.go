package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"sync"
)

var waitgroup sync.WaitGroup

func Afunction(shownum int) {
	fmt.Println(shownum)
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}

func main() {
	go auto()
	select {}
}

func auto() {

	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		go doCrawl(i)

	})
	c.Start()
}

//开始抓取
func doCrawl(i int) {
	log.Println("正在开始新的任务队列:", i)
	main2()
}

func main2() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go Afunction(i)
	}
	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	log.Println("所有任务队列都执行完成。")
}
