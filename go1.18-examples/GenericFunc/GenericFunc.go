package main

import "fmt"

// GenericFunc 一个标准的泛型函数模板
func GenericFunc[T any](args T) {
	// logic code
}

// e.g.
func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// Go语言在泛型中引入一个叫类型集合概念
// 约束参数类型只能为数值类型
func add[T int64 | float64](a, b T) T {
	return a + b
}

// 泛型里面引入一个~内置符号
// 这个符号会限制泛型参数底层是基于某种类型实现的变体或者别名
type MyInt int8

func add2[T int64 | float64 | ~int8](a, b T) T {
	return a + b
}

func bubbleSort[T int64 | float64](sequence []T) {
	for i := 0; i < len(sequence)-1; i++ {
		for j := 0; j < len(sequence)-1-i; j++ {
			if sequence[j] > sequence[j+1] {
				sequence[j], sequence[j+1] = sequence[j+1], sequence[j]
			}
		}
	}
}

func main() {
	print([]int{1, 2, 3, 4, 5, 6})
	print([]string{"1", "2", "3", "4", "5", "6"})

	fmt.Println("1+2=", add[int64](1, 2))
	fmt.Println("1.2+2.3=", add[float64](1.2, 2.3))

	fmt.Println("MyInt 3+4=", add2[MyInt](3, 4))

	sequence := []int64{1, 4, 3, 9, 6, 5}
	bubbleSort[int64](sequence)
	fmt.Println("after bubble sort:", sequence)
	sequence2 := []float64{120.13, 2.3, 112.3, 66.5, 78.12, 1.2, 8}
	bubbleSort[float64](sequence2)
	fmt.Println("after bubble sort:", sequence2)
}
