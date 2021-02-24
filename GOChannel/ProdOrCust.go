package main

import (
	"fmt"
	"time"
)

func Produce(Make chan<- int){
	for i:=1;i<11 ;i++  {
		fmt.Println("Produce Make : " ,i)
		Make <- i  //写入缓冲区
	}

	close(Make)
}

func Consumer(Pay <- chan int){
	for things := range Pay{
		fmt.Println("Consumer Pay: ",things)
		time.Sleep(time.Second)
	}
}


func main() {
	ch := make(chan int,4)
	go Produce(ch)
	Consumer(ch)

}