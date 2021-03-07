package main

import (
	"fmt"
	"time"
)

type MyContext struct {
	workingCh chan struct{}
	stopCh chan struct{}
}

func (m MyContext) working(){//通知工作的channel
	m.workingCh <- struct{}{}
}

func (m MyContext) end()  {//通知关闭的channel
	m.stopCh <- struct{}{}
}



func main()  {
	//初始化
	my := MyContext{make(chan struct{}),make(chan struct{})}

	go func() {
		defer fmt.Println("finish")
		for true {
			select {
			case <-my.workingCh:
				fmt.Println("working")
				//.......
			case <-my.stopCh:
				fmt.Println("end")
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		time.Sleep(1*time.Second)
		my.working()
	}
	time.Sleep(3 * time.Second)
	my.end()
	time.Sleep(50 * time.Millisecond) // 给一定的时间打印信息
}