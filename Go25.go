package main

import (
	"fmt"
	"time"
)

/**
 * 通道同步
 * 使用通道来同步go协程键的执行状态
 * 这是一个使用阻塞的接收方式来等一个go协程的运行结束
 */
func main() {
	done := make(chan bool, 1)

	go worker(done)

	//程序将在接收到通道中worker发出的通知前一直阻塞
	<-done
}

//在go协程中运行的函数。done通道将用于通知其他go协程该函数已完毕
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//发送一个值通知结束
	done <- true
}
