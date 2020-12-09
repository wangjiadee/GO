/*
 * @Author: your name
 * @Date: 2020-08-19 15:00:06
 * @LastEditTime: 2020-08-19 15:00:10
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\listen1\const\main.go
 */
package main

import "fmt"

func main() {
	// 写法一：
	// const a int = 520
	// const b = "xiejier"

	// fmt.Printf("a=%d b =%s \n", a, b)

	// 写法二:
	// const (
	//      a = 520
	//      b = "xiejier"
	// )

	// fmt.Printf("a=%d b =%s \n", a, b)

	// 写法三:
	// const (
	// a = iota
	// b
	// )
	// fmt.Printf("a=%d b =%d \n", a, b)

	// 写法四:
	const (
		a1 = 1 << iota
		a2 = 1 << iota
		a3 = 1 << iota
		a4 = 1 << iota
	)
	fmt.Printf("a1=%d a2=%d a3=%d a4=%d \n", a1, a2, a3, a4)
}
