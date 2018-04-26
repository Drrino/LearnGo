package main

import "fmt"

/**
 * 结构体
 */
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})	//创建新的结构体元素

	fmt.Println(person{name: "Alice", age: 30})	//初始化指定字段

	fmt.Println(person{name: "Fred"})	//省略字段将被初始化零值

	fmt.Println(&person{name: "Ann", age: 40})	//&前缀生成一个结构体指针

	s := person{"Sean", 50}		//使用点来访问结构体字段
	fmt.Println(s.name)

	sp := &s	//可以对结构体指针使用，指针会被自动解引用
	fmt.Println(sp.age)

	sp.age = 51	//结构体参数值可变
	fmt.Println(sp.age)
}
