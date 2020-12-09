/*
 * @Author: your name
 * @Date: 2020-08-17 14:51:11
 * @LastEditTime: 2020-08-17 14:56:57
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go\src\github.com\Ralph.org\listen1\hello\Thread.go
 */
package main

import (
	"fmt"
	"time"

)

func calc() {
	for i := 0; i < 20; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Run", i, "times")
	}
	fmt.Println("Finish!")
}

func main() {
	//  串行运行
	// calc()
	// 并行运行
	go calc()
	fmt.Println("Exited")
}
