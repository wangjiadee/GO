package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int ,5)  //存满5元素之前不会堵塞
	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	go func() {
		for i := 0;i<10;i++ {
			ch <- i
			len:= len(ch)
			cap:= cap(ch)
			fmt.Println("Sub go : ", " len: ",len, " cap: ", cap)
			//fmt.Println(i)
		}
	}()
	time.Sleep(time.Second *5)
	for i:= 0 ;i< 8 ;i++  {
		num := <-ch
		fmt.Println("Main go : " ,num)
	}
}
