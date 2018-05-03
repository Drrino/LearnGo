package main

import "fmt"

/**
 * 通道
 * 连接多个协程的管道。可从一个go协程将值发送到通道在其他go协程中接收
 */
func main() {
	//使用make创建新通道
	messages := make(chan string)

	//使用 channel<- 语法发送一个新值到通道中，这里在一个新的go协程中发送ping到messages通道中
	go func() { messages <- "ping" }()

	//使用 <-channel 语法从通道中接收一个值
	msg := <-messages
	fmt.Println(msg)
}
