package main

import "fmt"

/**
 * 变参函数
 */
func main() {
	sum(1, 2)

	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)	//如果slice已有多个值，想作为变参使用，需func(slice...)调用
}

func sum(nums ...int) {		//使用任意数目的int作为参数
	fmt.Print(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
