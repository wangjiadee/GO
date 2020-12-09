/*
 * @Author: your name
 * @Date: 2020-08-17 15:01:00
 * @LastEditTime: 2020-08-17 15:03:35
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go\src\github.com\Ralph.org\listen1\mul_return\main.go
 */
package main

import "fmt"

func Add(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	n, m := Add(5, 2)
	fmt.Println(n, m)

}
