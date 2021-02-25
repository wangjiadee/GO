/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-24 20:55:17
 * @LastEditTime: 2021-02-25 20:19:03
 * @FilePath: \GO\src\ci.Ralph.org\GOSelect\SelectTest.go
 * @Effect: DO
 */
package main

import (
	"fmt"
	"runtime"
)

func fibonacci(ch <-chan int, Quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-Quit:
			runtime.Goexit()
		}
	}
}

func main() {
	ch := make(chan int)
	Quit := make(chan bool)

	go fibonacci(ch, Quit)

	x, y := 1, 1
	for i := 0; i < 20; i++ {
		ch <- x
		x, y = y, x+y
	}
	Quit <- true
}
