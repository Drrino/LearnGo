package main

import "fmt"

/**
 * 指针
 * 允许程序中通过引用传递值或者数据结构
 * 解引用：返回内存地址中保存的值
 */
func main() {
	i := 1
	fmt.Println("initial:", i)

	//通过zeroval和zeroptr来比较指针和值类型的不同
	//zeroval从调用它的函数得到ival形参的拷贝
	zeroval(i)
	fmt.Println("zeroval:", i)

	//&int参数意味着它用了一个int的指针，函数体内*iptr解引用这个指针，从他内存地址获取这个地址的当前值。
	//对一个解引用的指针赋值会将改变这个指针引用的真实地址的值。
	//zeroval不能改变i的值，zeroptr可以，因为它有一个这个变量的内存地址的引用
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	//通过&i语法取得i的内存地址
	fmt.Println("pointer:", &i)
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func zeroval(ival int) {
	ival = 0
}
