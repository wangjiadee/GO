/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-19 20:40:55
 * @LastEditTime: 2021-02-19 21:07:02
 * @FilePath: \GO\src\github.com\Ralph.org\GOmap\mapTest2.go
 * @Effect: DO
 */
package main

import "fmt"

func MapDelete(m map[int]string,key int){
	delete(m,key)
}

func main(){
	// Initialization of map
	var m1 map[int]string = map[int]string{1:"MDZZ",2:"SHA"}
	fmt.Println(m1)

	// Initialization of map two :=
	m2 := map[int]string{1:"Love",2:"Hate"}
	fmt.Println(m2)


	//map assignment
	m3 := make(map[int]string,100)

	m3[1] ="aaa"
	m3[1] ="ddd"  //modify value
	m3[2] ="ccc"  //add value

	fmt.Println(m3)

	m4 := map[int]string{1:"Love",2:"Hate",3:"MDZZ",4:"SHA"}
	for k,v := range m4{
		fmt.Printf("%d ===== %s\n",k,v)
	}

	for k := range m4{
		fmt.Println(k,m4[k]) 
	}

	v,BOOL := m4[1]
	fmt.Println(v,BOOL) //Love true

	MapDelete(m4,1)
	fmt.Println(m4)

	MapDelete(m4,100)
	fmt.Println(m4)
}