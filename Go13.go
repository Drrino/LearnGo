package main

import "fmt"

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() //仅仅想返回值的一部分，可使用空白定义符_
	fmt.Println(c)
}

func vals() (int, int) { //返回2个int值
	return 3, 7
}
