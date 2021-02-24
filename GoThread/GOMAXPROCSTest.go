package main

import (
	"fmt"
	"runtime"
)

func main() {
	n:= runtime.GOMAXPROCS(1)
	fmt.Println("n: ",n)

	for  {
		go fmt.Print("Fuck") // sub_go

		fmt.Print("What")    //main_go
	}
}