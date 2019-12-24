package main

import "fmt"
/**
 数组为值类型, 调用函数会拷贝数组
 go中一般不使用数组
 */
func printArray(arr [5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func printPointArray(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr3)
	printPointArray(&arr3)
	fmt.Println(arr3)
}
