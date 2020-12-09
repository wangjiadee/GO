/*
 * @Author: your name
 * @Date: 2020-08-28 14:17:11
 * @LastEditTime: 2020-08-28 14:33:42
 * @LastEditors: Please set LastEditors
 * @Description: 通过数组的方式计算输入的数字的大小/通过数组的下标计算元素的和
 * @FilePath: \newgo\Ralph.org\element\Array\Array_Example4.go
 */
package main

import "fmt"

func String_sum() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 8}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}

func TwoSum(a []int, target int) {
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		for j := 0; j < len(a); j++ {
			if a[j] == other {
				fmt.Println(i, j)
			}
		}
	}
}
func TestArr() {
	b := [...]int{1, 2, 3, 4, 5, 6, 7, 9}
	TwoSum(b, 10)
}

func main() {
	String_sum()
	TestArr()
}
