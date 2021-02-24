package main

import (
	"fmt"
	"time"
)

var Mychannel = make(chan int)

func Printer(s string)  {
	for _,str := range s {
		fmt.Printf("%c",str)
		time.Sleep(500*time.Millisecond)
	}
}

func P1()  {
	Printer("Fuck")
	Mychannel <- 888
}

func P2()  {

	Printer("You")
}

func main() {
	go P1()
	<- Mychannel
	go P2()
	for {
		;
	}
}