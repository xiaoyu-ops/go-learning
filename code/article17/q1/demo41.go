package main

import "fmt"

func main() {
	// 示例1。
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
	fmt.Println()

	// 示例2。
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}//在go语言中，[...]表示编译器根据初始化值的个数自动推断数组的长度，会复制整个数组作为迭代源。
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
	fmt.Println()

	// 示例3。
	numbers3 := []int{1, 2, 3, 4, 5, 6}//这里会改变的原因是不是复制的切片结构体，而不是复制的底层数组。
	maxIndex3 := len(numbers2) - 1
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
}
