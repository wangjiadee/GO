/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2020-09-01 20:57:42
 * @LastEditTime: 2020-09-01 20:58:44
 * @FilePath: \GO\src\github.com\Ralph.org\rc\run\run.go
 * @Effect: DO
 */
package main

import (
	"fmt"

	"github.com/Ralph.org/rc/test"

)

func main() {
	var s1 int = 100
	var s2 int = 200
	sum := test.Add(s1, s2)
	fmt.Println(sum)
	// Lower variables are not valid
	// fmt.Println(test.a)
}