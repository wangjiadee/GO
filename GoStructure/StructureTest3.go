package main

import "fmt"

type Creature struct {
	name string
	age  int
	sex  string
}

func main() {
	var p1 Creature = Creature{
		name: "A",
		age:  10,
		sex:  "M",
	}
	var p2 Creature = Creature{
		name: "B",
		age:  20,
		sex:  "W",
	}
	var p3 Creature = Creature{
		name: "A",
		age:  30,
		sex:  "M",
	}

	fmt.Println(p1 == p2)
	fmt.Println(p1 == p3)

}
