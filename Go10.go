package main

import "fmt"

/**
 * map 为go内置关联数据类型（其他语言中成为哈希或字典）
 */
func main() {
	m := make(map[string]int) //使用内建make(map[key-type]val-type)创建空map

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2") //delete从map中移除键值对
	fmt.Println("map:", m)

	_, prs := m["k2"] //从map中取值，可选的第二返回值指示这个键是否存在于map。可以用来消除键不存在和键含有零值，如0或""产生歧义
	fmt.Println("prs", prs)

	n:=map[string]int{"foo":1,"bar":2}
	fmt.Println("map:",n)
}
