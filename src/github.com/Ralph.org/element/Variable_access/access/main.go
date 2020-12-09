/*
 * @Author: your name
 * @Date: 2020-08-31 09:30:27
 * @LastEditTime: 2020-09-01 18:55:16
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GO\src\github.com\Ralph.org\element\Variable_access\access\main.go
 */
package main

import (
	"fmt"

	"github.com/Ralph.org/element/variable_access/test"

)

func main() {
	var s1 int = 100
	var s2 int = 200
	sum := test.Add(s1, s2)
	fmt.Println(sum)
	// Lower variables are not valid
	// fmt.Println(test.a)
}
