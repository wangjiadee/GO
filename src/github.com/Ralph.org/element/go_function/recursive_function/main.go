/*
 * @Author: your name
 * @Date: 2020-08-27 16:54:42
 * @LastEditTime: 2020-08-27 17:02:34
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\element\go_function\recursive_function\main.go
 */
package main

import "fmt"

var s int = 1

func Re_func(n int) int {
	// 只有第一排的人才知道自己的排数
	if n == 1 {
		return 1
	}
	// 如果不是第一排，问一下前一排的人
	r := Re_func(n - 1)
	fmt.Println("前一排的排数：", r)
	// 把前一排人的排数+1，计算出自己的排数。
	return r + 1
}

// 阶乘案例
func Re_demo(n int) {
	if n == 1 {
		return
	}
	s *= n
	Re_demo(n - 1)
}

func main() {
	c := Re_func(4)
	fmt.Println(c)
	Re_demo(4)
}
