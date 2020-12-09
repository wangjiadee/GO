package main

import "fmt"

func funcA() int {
	x := 8
	defer func() {
		x += 1
	}()
	return x
}

func funcB() (x int) {
	defer func() {
		x += 1
	}()
	return 5
}

func funcC() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func funcD() (x int) {
	defer func(x int) {
		x += 1
	}(x)
	return 5
}

func Defer_Test() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("running")
	fmt.Println("return")
}

func Defer_Test1() {
	var i int = 0
	// defer 可以当作栈 i=0 已经存在了栈里面 后面对i的修改是无效的
	defer fmt.Printf("i = %d", i)
	i = 520
	fmt.Println(i)
}
func main() {
	fmt.Println(funcA())
	fmt.Println(funcB())
	fmt.Println(funcC())
	fmt.Println(funcD())
	Defer_Test()
	Defer_Test1()
}
