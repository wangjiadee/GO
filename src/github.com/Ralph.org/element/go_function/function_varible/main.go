/*
 * @Author: Ralph
 * @Date: 2020-08-31 14:57:36
 * @LastEditTime: 2020-08-31 15:25:34
 * @LastEditors: Please set LastEditors
 * @Description: 函数作为函数的参数
 * @FilePath: \newgo\Ralph.org\element\go_function\function_varible\main.go
 */

package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a + b
}

// 函数的格式类型要保持一样就可以
func calc(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func main() {
	sum := calc(100, 300, Add)
	sub := calc(520, 120, Sub)
	fmt.Println(sum, sub)
}
