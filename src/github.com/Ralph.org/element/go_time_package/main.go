/*
 * @Author: your name
 * @Date: 2020-08-26 14:53:28
 * @LastEditTime: 2020-08-26 14:58:30
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\element\go_time_package\main.go
 */
package main

import (
	"fmt"
	"time"
)

func TestTime() {
	fmt.Println(time.Now())

	// 获取各种类型的片段输出
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())

	// 格式化输出1
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	// go语言的时间戳表示
	fmt.Println(time.Now().Unix())
}

// 获取当前时间2 并格式化
func TimeFormat() {
	timeStr := time.Now().Format("2019/11/09 20:00:00")
	fmt.Printf(timeStr)
}

func TimeChange() {
	// 时间戳转换时间
	timestamp := time.Now().Unix()
	TimeObj := time.Unix(timestamp, 0)
	fmt.Println(TimeObj)
}

// 定时器
func TimeTicker() {
	ticker := time.Tick(2 * time.Second)
	for i := range ticker {
		fmt.Println(i)
		fmt.Println("process run")
	}
}

func TestCost() {
	Start := time.Now().UnixNano() // 精确到纳秒
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	End := time.Now().UnixNano()
	Cost := (End - Start) / 1000
	fmt.Println(Cost)

}

// ###################start main function
func main() {
	// TestTime()
	// TimeChange()
	// TimeTicker()
	// TimeFormat()
	TestCost()
}
