/*
 * @Author: your name
 * @Date: 2020-08-28 11:33:12
 * @LastEditTime: 2020-08-28 14:17:22
 * @LastEditors: Please set LastEditors
 * @Description: 二维数组
 * @FilePath: \newgo\Ralph.org\element\Array\Secondary_array.go
 */
package main

import "fmt"

func Initialize_all() {
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {5, 6, 7}}
	fmt.Println(arr)
}

func Initialize_Partial() {
	var arr [2][3]int = [2][3]int{{1, 2}, {6}}
	fmt.Println(arr)
}

func Specify_init() {
	var arr [2][3]int = [2][3]int{0: {1: 6}}
	fmt.Println(arr)
}

func Custom_init() {
	// 行的下标可以用"..."来代替，但是列的下标不能用"..."来代替。
	arr := [...][3]int{{1, 2, 3}, {5, 6}}
	fmt.Println(arr)
}

func For_Sec_arr() {
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {5, 6, 7}}
	for i := 0; i < len(arr); i++ { // 遍历的行
		for j := 0; j < len(arr[0]); j++ { // 遍历的是列。
			fmt.Println(arr[i][j])
		}
	}

	// for range eg:
	for _, v := range arr {
		for j, data := range v {
			fmt.Println("j:", j)
			fmt.Println("data:", data)
		}
	}
}

// ##############start main function
func main() {
	Initialize_all()
	Initialize_Partial()
	Specify_init()
	Custom_init()
	For_Sec_arr()
}
