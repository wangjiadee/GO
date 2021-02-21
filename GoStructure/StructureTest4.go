package main

import (
	"fmt"
	"unsafe"
)

type Animal struct {
	name string
	sex string
	age int
}

func test(a *Animal)  {
	fmt.Println("test",unsafe.Sizeof(a))
	a.name = "CAT"
	a.age =  3
	a.sex = "G"
}

func main()  {
	var a1 *Animal = &Animal{
		name: "AA",
		sex:  "G",
		age:  10,
	}
	fmt.Println(a1)  //&{AA G 10}

	var a2 Animal = Animal{
		name: "AA",
		sex:  "G",
		age:  10,
	}
	var a3 *Animal
	a3 = &a2
	fmt.Println(a3) //&{AA G 10}

	a4 := new(Animal)
	a4.name = "BB"
	a4.sex = "M"
	a4.age = 20
	fmt.Printf("a4, type= %T\n", a4) //type= *main.Animal
	fmt.Println(a4)
	fmt.Println("a4:",unsafe.Sizeof(a4))

}