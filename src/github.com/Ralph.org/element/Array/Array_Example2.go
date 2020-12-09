/*
 * @Author: your name
 * @Date: 2020-08-28 11:25:21
 * @LastEditTime: 2020-08-28 11:32:23
 * @LastEditors: Please set LastEditors
 * @Description: 从一个数组中取出最大的数，最小的 求和 和平均值
 * @FilePath: \newgo\Ralph.org\element\Array\Array_Example2.go
 */
package main

import "fmt"

func main() {
	var (
		number [5]int = [5]int{1, 2, 3, 4, 5}
		max    int    = number[0]
		min    int    = number[0]
		sum    int
	)
	for i := 0; i < len(number); i++ {
		if number[i] > max {
			max = number[i]
		}
		if number[i] < min {
			min = number[i]
		}
		sum = sum + number[i]
	}
	fmt.Println("max", max)
	fmt.Println("min", min)
	fmt.Println("sum", sum)
	fmt.Printf("av: %.2f", float64(sum)/5)
}
