package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	//fmt.Println("Now Time:" ,time.Now())
	//myTimer := time.NewTimer(time.Second * 2)
	////nowTime := <- myTimer.C
	//NowTime := <-time.After(time.Second *2)
	//fmt.Println(NowTime)
	quit := make(chan bool)
	fmt.Println("STime:",time.Now())
	MyTicker := time.NewTicker(time.Second)

	I:=0
	go func() {
		for  {
			NowTime := <- MyTicker.C
			I++
			fmt.Println("MTime:",NowTime)
			if I == 3{
				quit <- true
				runtime.Goexit()
			}
		}
	}()

	<-quit

}

