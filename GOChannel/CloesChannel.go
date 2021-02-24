package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	
	go func() {
		for i:= 0;i<5 ;i++  {
			 ch <- i
		}
		close(ch)
	}()

	for  {
		if num,flag := <- ch;flag == true {
			fmt.Println("Channel Read data :" ,num)
		}else {
			break
		}
	}
}