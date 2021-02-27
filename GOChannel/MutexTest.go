/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-25 21:11:39
 * @LastEditTime: 2021-02-25 21:51:38
 * @FilePath: \GO\src\go.ralph.org\GOChannel\MutexTest.go
 * @Effect: DO
 */

package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

/*
共享数据：打印机函数
*/
func Printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
}

func person1() { // 先
	Printer("hello")
	ch <- 98
}

func person2() { // 后
	<-ch
	Printer("world")
}

func main() {
	go person1()
	go person2()
	for {

	}
}
