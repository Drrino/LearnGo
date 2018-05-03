package main

import "fmt"

/**
 * 协程
 * 在执行上为轻量级的线程
 */
func main() {
	f("direct")

	//在一个go协程充调用f函数，这个新go协程将会并行执行这个函数
	go f("goroutine")

	//为匿名函数启动一个go协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//等两个go协程在独立的go协程中异步运行结束，按下任意键结束程序
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from," : ",i)
	}
}
