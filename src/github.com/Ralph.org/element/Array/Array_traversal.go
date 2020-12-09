/*
 * @Author: Go
 * @Date: 2020-08-28 10:15:17
 * @LastEditTime: 2020-08-28 10:22:53
 * @LastEditors: Please set LastEditors
 * @Description: 数组的遍历
 * @FilePath: \newgo\Ralph.org\element\Array\Array_traversal.go
 */
package main

import "fmt"

func Traversal_for() {
	var Numbers [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(Numbers); i++ {
		fmt.Println(Numbers[i])
	}
}

func Traversal_for_range() {
	for i, v := range Numbers {
		fmt.Println("index：", i)
		fmt.Println("value：", v)
	}
}

// Use Anonymous variable
func Traversal_Anonymous() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	for _, v := range numbers {
		//fmt.Println("biao",i)
		fmt.Println("v:", v)
	}
}
func main() {
	Traversal_for()
	Traversal_for_range()
	Traversal_Anonymous()
}
