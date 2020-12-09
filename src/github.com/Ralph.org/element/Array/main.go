/*
 * @Author: your name
 * @Date: 2020-08-27 17:54:20
 * @LastEditTime: 2020-08-28 10:09:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\element\Array\main.go
 */
package main

import "fmt"

// The array index starts from 0
func Array_index() {
	var number [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(number[3])
}

// The array can Partial assignment
// Not assigned a value of 0
func Array_assignment() {
	number := [5]int{1, 2}
	fmt.Println(number[1], number[4])
}

// Array can also specify assignment
func Array_assignment1() {
	var number [5]int = [5]int{2: 5, 3: 6}
	fmt.Println(number)
}

// *** You can customize the length of the array by initializing
func Array_custome() {
	Numbers := [...]int{7, 8, 9}
	fmt.Println(len(Numbers), Numbers[0])
}

// ###################start main function

func main() {
	Array_index()
	Array_assignment()
	Array_assignment1()
	Array_custome()
}
