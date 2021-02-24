package main

import (
	"fmt"
)

func test() {
	fmt.Println("SB")
	return
	defer fmt.Println("Cao")
}

func main() {
	
	go func() {
		fmt.Println("WDNMD")
		test()
		defer fmt.Println("MDZZ")
	}()

	for  {
		;
	}
}