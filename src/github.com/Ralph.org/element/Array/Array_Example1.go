/*
 * @Author: your name
 * @Date: 2020-08-28 11:12:12
 * @LastEditTime: 2020-08-28 11:24:11
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\element\Array\Array_Example1.go
 */
package main

import "fmt"

func main() {
	var num1 [5]int = [5]int{1, 2, 3, 4, 5}
	var num2 [5]int = [5]int{1, 2, 3, 4, 5}
	b := complexV(num1, num2)
	if b {
		fmt.Println("success!")
	} else {
		fmt.Println("no !")
	}
	// fmt.Println(num1 == num2)
	//	比较的类型长度是一直的才行
}
func complexV(n1 [5]int, n2 [5]int) bool {
	var b bool = true
	if len(n1) == len(n2) {
		for i := 0; i < len(n1); i++ {
			if n1[i] == n2[i] {
				continue
			} else {
				b = false
				break
			}
		}
	} else {
		b = false
	}
	return b
}
