package main

import "fmt"

/**
 * range迭代数据结构
 */
func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { //range统计slice元素个数，数组可采用此方式
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums { //range在数组和slice中提供每项索引和值，不需要索引则使用空值定义符_来忽略
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs { //range在map中迭代键值对
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" { //range在字符串中迭代unicode编码，第一个返回rune起始字节位置，第二个为rune自己
		fmt.Println(i, c)
	}
}
