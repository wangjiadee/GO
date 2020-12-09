/*
 * @Author: Ralph
 * @Date: 2020-08-17 16:24:23
 * @LastEditTime: 2020-08-17 17:51:59
 * @LastEditors: Please set LastEditors
 * @Description: GO
 * @FilePath: \go\src\github.com\Ralph.org\listen1\variable\main.go
 */
package main

import "fmt"

func main() {
	var a int
	var b bool
	var s string
	//  两种定义方式
	// var (
	// 	a1 int
	// 	b1 bool
	// 	s1 string
	// )

	a = 10
	b = true
	s = "xiejier"

	fmt.Println(a, b, s)
	// fmt.Println("a1 =%d b1 = %t s1 = %f\n", a1, b1, s1)
	// 	a1= 10
	// 	b1= true
	// 	s1 = 'xx'
}
