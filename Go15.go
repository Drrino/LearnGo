package main

import "fmt"

/**
 * 闭包
 * 通过闭包来使用匿名函数
 */
func main() {
	//调用intSeq函数将返回值赋给nextInt。这个函数包含了自己的值i，每次调用nextInt都会更新i值
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
}

//intSeq函数返回另一个在intSeq函数体内定义的函数。返回函数采用闭包方式隐藏变量i
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
