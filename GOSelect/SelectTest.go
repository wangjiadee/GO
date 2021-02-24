package main
/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-24 21:22:15
 * @LastEditTime: 2021-02-24 22:13:18
 * @FilePath: \01-Go语言高级-第04天（channel）f:\github\GO\GO\src\GOSelect\SelectTest.go
 * @Effect: DO
 */
import (
	"fmt"
	"runtime"
)

func fibonacci(ch <- chan int, Quit <- chan bool){
	for  {
		select {
		case num:= <- ch:
			fmt.Println(num)
		case <-Quit:
			runtime.Goexit()
		}
	}
}


func main() {
	ch := make(chan int)
	Quit := make(chan bool)

	go fibonacci(ch,Quit)

	x,y := 1,1
	for i:=0;i<20;i++ {
		ch <- x
		x,y = y,x+y
	}
	Quit <- true
}