package main

import "fmt"

/**
 * var 声明 1 个或者多个变量
 * go自动推断已初始化的变量类型
 * 声明变量未给出对应初始值，变量将初始化为零值
 * :=语句是声明并初始化变量的简写
 */
func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}
