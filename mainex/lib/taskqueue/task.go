/*
任务队列 By 邹慧刚 952750120@qq.com
*/

package taskqueue

import (
	"fmt"
	"sync"
)

type TaskQueue struct{}

var Task = TaskQueue{}

var wg sync.WaitGroup //创建一个sync.WaitGroup
var ch chan int = make(chan int)

func (this TaskQueue) NewTask(TCount int) {
	//产生任务
	go func() {
		for i := 0; i < TCount; i++ {
			ch <- i
		}
		close(ch)
	}()
}

//多少个士兵去执行任务
func (this TaskQueue) Soldiers(Scount int, HandleFunc interface{}, args ...interface{}) {
	//开始执行
	wg.Add(Scount)
	for i := 0; i < Scount; i++ {
		i := i
		go func() {
			defer func() { wg.Done() }()

			fmt.Printf("士兵 %v 开始行动...\r\n", i)

			for task := range ch {
				func() {

					defer func() {
						err := recover()
						if err != nil {
							fmt.Printf("任务失败：士兵编号=%v, task=%v, err=%v\r\n", i, task, err)
						}
					}()
					//处理任务队列函数
					HandleFunc.(func(...interface{}))(task, args)
					fmt.Printf("任务结果=%v ，士兵编号=%v, task=%v\r\n", task*task, i, task)
				}()
			}

			fmt.Printf("士兵 %v 结束。\r\n", i)
		}()

	}

	//等待所有任务完成
	wg.Wait()
	print("全部任务结束")
}
