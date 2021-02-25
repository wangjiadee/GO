/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-25 20:30:05
 * @LastEditTime: 2021-02-25 20:50:03
 * @FilePath: \GO\src\ci.Ralph.org\GOSelect\SelectTimeOutTest.go
 * @Effect: DO
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	Quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("Num: ", num)
			//Set to wait for 5s and then the program exits
			case <-time.After(5 * time.Second):
				fmt.Println("[Warning :]Time out!")
				Quit <- true
				break
			}
		}
	}()

	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}

	<-Quit
	fmt.Println("Finished!")
}
