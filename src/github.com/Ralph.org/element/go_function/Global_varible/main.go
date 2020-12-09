/*
 * @Author: your name
 * @Date: 2020-08-27 15:13:29
 * @LastEditTime: 2020-08-27 15:13:53
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\element\go_function\Local_var\main.go
 */
package main

import "fmt"

var b int

func main() {
	b = 20
	TestB()
	fmt.Println("main:", b)
}
func TestB() {
	var b int
	b = 30
	fmt.Println("TestB:", b)

}
