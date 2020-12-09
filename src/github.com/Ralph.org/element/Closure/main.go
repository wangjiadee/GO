/*
 * @Author: Ralph
 * @Date: 2020-09-01 11:43:48
 * @LastEditTime: 2020-09-01 11:44:23
 * @LastEditors: Please set LastEditors
 * @Description: 闭包
 * @FilePath: \newgo\Ralph.org\element\Closure\main.go
 */
package main

import (
	"fmt"
	"strings"
)

// function survive that x always exists
func Ader() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

// Closure can be functions
func Closurefunc(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func makeSuffixFunc_test() {
	func1 := makeSuffixFunc(".avi")
	func2 := makeSuffixFunc(".mp4")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
}

//########################start main function
func main() {
	f := Ader()
	ret := f(20)
	ret1 := f(40)
	fmt.Println(ret, ret1)
	tmp1 := Closurefunc(10)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := Closurefunc(200)
	fmt.Println(tmp2(1), tmp2(2))
	makeSuffixFunc_test()
}
