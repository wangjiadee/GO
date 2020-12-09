package main

import "fmt"

// example for switch
func SwitchEg1() {
	var data int
	fmt.Println("pelase input data:")
	fmt.Scan(&data)
	switch data {
	case 1:
		fmt.Println("Mon")
	case 2:
		fmt.Println("Tue")
	case 3:
		fmt.Println("Wed")
	case 4:
		fmt.Println("Thu")
	case 5:
		fmt.Println("Fri")
	case 6:
		fmt.Println("Sat")
	case 7:
		fmt.Println("Sun")
	default:
		fmt.Println("NUMBER Error!")
	}
}

func SwitchEg2() {
	var score float64
	fmt.Println(">>>Pelase input score:")
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("c")
	case score >= 60:
		fmt.Println("d")
	default:
		fmt.Println("PASS")
	}
}

/*  switch 案例2：
李四的年终工作评定,如果定为A级,则工资涨500元,如果定为B级,则工资涨200元,
如果定为C级,工资不变,如果定为D级工资降200元,如果定为E级工资降500元.
设李四的原工资为5000,请用户输入李四的评级,然后显示李四来年的工资.
*/

func SwitchEg3() {
	var level string
	var salary int = 5500
	var b bool = true
	fmt.Println(">>>Pelase input u level:")
	fmt.Scan(&level)
	switch level {
	case "A":
		salary += 500
	case "B":
		salary += 200
	case "C":
	case "D":
		salary -= 200
	case "E":
		salary -= 500
	default:
		b = false
		fmt.Println("ERROR LEVEL!")
	}
	if b {
		fmt.Println("||| u salary :", salary)
	}
}

// 综合switch 和if
// 请用户输年份,再输入月份,输出该月的天数.(需要考虑闰年)
/*
	闰年的判定(符合下面两个条件之一)
	年份能够被400整除.(2000)
	年份能够被4整除但不能被100整除.(2008)
*/

/*
Go里面switch默认相当于每个case最后带有break，
匹配成功后不会自动向下执行其他case，而是跳出整个switch,
但是可以使用fallthrough强制执行后面的case代码
*/

func SwitchEg4() {
	var year int
	var mouth int
	var day int
	fmt.Println(">>>Pelase input years:")
	fmt.Scan(&year)
	fmt.Println(">>>Pelase input mouths:")
	fmt.Scan(&mouth)
	if mouth >= 1 && mouth <= 12 {
		switch mouth {
		case 1:
			fallthrough
		case 2:
			if year%400 == 0 || year%4 == 0 && year%100 != 0 {
				day = 29
			} else {
				day = 28
			}
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 7:
			fallthrough
		case 8:
			fallthrough
		case 10:
			fallthrough
		case 12:
			day = 31
		default:
			day = 30
		}
		fmt.Println("> day:", day)
	} else {
		fmt.Println("Mouth Error!")
	}
}

// ###############start main fuction

func main() {
	// SwitchEg1()
	SwitchEg2()
	SwitchEg3()
	SwitchEg4()
}
