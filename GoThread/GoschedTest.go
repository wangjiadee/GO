package main

import (
	"fmt"
)

func main() {
	go func(s string) {
		for i:=0;i<2 ;i++  {
			fmt.Println(s)
		}
	}("WDNMD")

	for i:= 0;i<2 ;i++  {
		//runtime.Gosched()
		fmt.Println("Hello")

	}
}