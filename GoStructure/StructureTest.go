/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-19 21:51:06
 * @LastEditTime: 2021-02-19 22:10:22
 * @FilePath: \GO\src\github.com\Ralph.org\GoStructure\StructureTest.go
 * @Effect: DO
 */
package main

import "fmt"

type Person struct {
	name string
	sex string
	age int	
}

type Student struct{
	p []Person
	id []int
	score [1000]int
}



func main(){
	// 1. 顺序初始化, 必须全部初始化完整
	var man1 Person = Person{"Jerry","M",19}
	fmt.Println(man1)
	// 2.部分初始化
	man2 := Person{name:"MM",age:20}
	fmt.Println(man2)

	// Index member variable "."
	fmt.Printf("man2.name = %q\n", man2.name)   //man2.name = "MM"

	var man3 Person
	man3.name = "mike"
	man3.sex = "M"
	man3.age = 99
	fmt.Println("man3:", man3)
	man3.age = 1973
	fmt.Println("man3:", man3)


	
}