/*
 * @Author: your name
 * @Date: 2020-08-26 15:42:01
 * @LastEditTime: 2020-08-27 11:26:59
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\element\go_syntax\if_syntax\main.go
 */
package main

import "fmt"

// eg about if syntax
func If_Eg() {
	fmt.Println("pelase input u age:")
	var age int
	fmt.Scan(&age)
	// if 的使用方式
	if age >= 18 {
		fmt.Println("u can buy Yellow vedio")
	}
}

// eg2 about if syntax
func If_Eg2() {
	var chinese int
	var math int
	fmt.Println("pelase input u chinese grade:")
	fmt.Scan(&chinese)
	fmt.Println("pelase input u math grade:")
	fmt.Scan(&math)
	if chinese > 70 && math == 100 {
		fmt.Println("奖励绿帽子！")
	}
}

// About if...else syntax
func Ifelse() {
	fmt.Println("pls input u age:")
	var age int
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("u can see yellow video!")
	} else {
		fmt.Println("studying hard now!")
	}
}

// example for if...else

func Ifelse2() {
	var money float64
	fmt.Println("Pelase input money:")
	fmt.Scan(&money)
	if money >= 2 {
		var count int
		fmt.Println("Pelase input count imformation:")
		fmt.Scan(&count)
		if count > 0 {
			fmt.Println("down!")
		} else {
			fmt.Println("up!")
		}
	} else {
		fmt.Println("gun !")
	}
}

// example for if...else...if
func Ifelseif() {
	var score float64
	fmt.Println("input u score:")
	fmt.Scan(&score)
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}

func Ifelseif2() {
	var username string
	var password string
	fmt.Println("Pelase input u username:")
	fmt.Scan(&username)
	fmt.Println("Pelase input u Password:")
	fmt.Scan(&password)
	if username == "mdzz" && password == "sb" {
		fmt.Println("loading success!")
	} else if username == "mdzz" {
		fmt.Println("password Error!")
	} else if password == "sd" {
		fmt.Println("Username Error!")
	} else {
		fmt.Println("password or Username Error!")
	}
}

// ###########################start main function
func main() {
	// If_Eg()
	// If_Eg2()
	// Ifelse()
	// Ifelse2()
	// Ifelseif()
	Ifelseif2()
}
