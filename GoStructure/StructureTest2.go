/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-19 22:10:36
 * @LastEditTime: 2021-02-19 22:12:13
 * @FilePath: \GO\src\github.com\Ralph.org\GoStructure\StructureTest2.go
 * @Effect: DO
 */
package main

import "fmt"

type Student struct {
	id int
	name string
	sex string
	age int
	addr string
}

	
func main() {
	var s1 *Student = &Student{
		id:   3,
		name: "Q",
		sex:  "M",
		age:  29,
		addr: "hahaha",
	}

	s2 := &Student{
		id:   34,
		name: "W",
		sex:  "M",
		age:  30,
		addr: "lalala",
	}
	fmt.Println(s1,s2)
}