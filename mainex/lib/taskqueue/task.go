/*
任务队列 By 邹慧刚 952750120@qq.com
*/

package taskqueue

import (
	"log"
	"sync"
)

type TaskQueue struct{}

var Task = TaskQueue{}

var wg sync.WaitGroup //创建一个sync.WaitGroup
var ch chan int

func (this TaskQueue) NewTask(TCount int) {
	ch = make(chan int)
	//产生任务
	go func() {
		for i := 0; i < TCount; i++ {
			ch <- i
		}
		close(ch)
	}()
}

//派遣多少个小兵去执行任务
func (this TaskQueue) Soldiers(Scount int, HandleFunc interface{}, args ...interface{}) {
	for i := 0; i < Scount; i++ {
		wg.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go this.ExecuteTask(i, HandleFunc, args)
	}
	wg.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	log.Println("success,报告大王,全部任务执行完毕。")
}

//我的小兵去执行任务
func (this TaskQueue) ExecuteTask(i int, HandleFunc interface{}, args ...interface{}) {
	defer func() {
		log.Printf("士兵 %v 结束。\r\n", i)
		wg.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
	}()
	//处理任务队列函数
	log.Printf("士兵 %v 开始行动...\r\n", i)
	for task := range ch {
		func() {

			defer func() {
				err := recover()
				if err != nil {
					log.Printf("任务失败：士兵编号=%v, task=%v, err=%v\r\n", i, task, err)
				}
			}()
			//处理任务队列函数
			HandleFunc.(func(...interface{}))(task, args)
			//fmt.Printf("任务结果=%v ，士兵编号=%v, task=%v\r\n", task*task, i, task)
		}()
	}

	//fmt.Println(i)
}
