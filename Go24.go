package main

import "fmt"

/**
 * 通道缓冲
 * 默认通道是无缓冲的,意味着只有对应的接收(<-chan)通道准备好接收时才允许发送(chan<-)
 * 可缓存通道允许在没有对应接收方的情况下缓存限定数量的值
 */
func main() {
	//make一个通道，最多缓存2个值
	messages := make(chan string, 2)

	//该通道有缓冲区，即使没一个对用的并发接收方依旧能发送这些值
	messages <- "buffered"
	messages <- "channel"

	//接收上面两个值
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
