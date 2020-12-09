/*
 * @Author: Ralph
 * @Date: 2020-08-28 10:27:43
 * @LastEditTime: 2020-08-28 11:10:20
 * @LastEditors: Please set LastEditors
 * @Description: 数组的传参
 * @FilePath: \newgo\Ralph.org\element\Array\Array_parameter.go
 */
package main

import "fmt"

func getPrint(n [5]int) {
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
}

// Modifying the original array in the function will not affect the value
func GetP(n [5]int) {
	//for i:=0;i<len(n);i++{
	//	fmt.Println(n[i])
	//}
	n[2] = 10
}

func main() {
	var Numbers [5]int = [5]int{1, 2, 3, 4, 5}
	getPrint(Numbers)
	GetP(Numbers)
	for i := 0; i < len(Numbers); i++ {
		fmt.Println(Numbers[i])
	}
}
