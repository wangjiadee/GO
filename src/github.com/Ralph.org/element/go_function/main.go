/*
 * @Author: your name
 * @Date: 2020-08-27 14:38:16
 * @LastEditTime: 2020-08-27 15:02:32
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go\src\github.com\Ralph.org\element\go_function\main.go
 */
package main

import "fmt"

// function define
func PlayGame() {
	fmt.Println("fucking game!")
}
func Play() {
	fmt.Println("lalala")
}

// Function parameter passing     1--100之间所有数字之和
func SumAdd(num int) { //形参
	var sum int
	for i := 1; i <= num; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func Add(num1 int, num2 int) { //定义两个形参
	fmt.Println(num1 + num2)
}

// 传多个参数
func TestSum(args ...int) {
	// args参数可以认为就是一个集合，在该集合中存储了传递过来的数据，
	// 会对每个数据加上一个编号，该编号是从0开始计算的。所以可以通过该编号获取集合中存储的具体的数据。
	/*
	 fmt.Println(args[0])
	 fmt.Println(args[1])
	 fmt.Println(args[2])
	 fmt.Println(args[3])
	*/
	/*
		 for i:=0;i<len(args);i++{ // len() 可以获取参数的存储的数据的个数
			 fmt.Println(args[i])
		 }
	*/
	// 集合。 数组，切片
	// 如果只写一个变量，该变量中存储的是集合的编号
	// _:匿名变量。匿名变量不会保存具体的数据
	for _, v := range args {
		//fmt.Println("a=",a) // 存储的是集合的编号
		fmt.Println("v=", v) // 存储的是具体的值。
	}

}

// 多个参数的顺序问题---固定参数放在前面，不定参数放在后面
func TestA(num int, args ...int) {
	fm参数ln(num)
	fmt.Println("argsp[0]=", args[0])
	fmt.Println("argsp[1]=", args[1])
}

// 函数的返回值案例集合
func AddResult(num1 int, num2 int) int { // 表示指定函数返回的数据的类型。
	var sum int
	sum = num1 + num2
	return sum //将变量sum中存储的值返回。
}

// 表明：最终会返回整型变量sum中的值。
// 在函数体中没有必要在重新创建sum变量。
func AddResult1(num1 int, num2 int) (sum int) {
	// var sum int
	sum = num1 + num2
	return sum //将变量sum中存储的值返回。
}

func AddResult2(num1 int, num2 int) (sum int) {
	// var sum int
	sum = num1 + num2
	return //如果已经指定了返回的变量的名称，那么return后面可以不用在加上变量的名称。
}
func GetResult() (num1 int, num2 int) {
	//var num1 int = 10
	// var num2 int = 20
	num1 = 10
	num2 = 20
	//return num1,num2 //表明返回两个变量的值。
	return
}

// ###############start main function

func main() {
	var (
		s  int
		s1 int
		s2 int
	)
	PlayGame()
	Play()
	SumAdd(200) //实参
	Add(3, 6)   //传递两个实参, 实参与形参的个数要保持一致，类型也要保持一致。
	TestA(10, 8, 9)
	s = AddResult(3, 6) //将函数返回的结果赋值给了变量s.
	fmt.Println(s / 3)
	s1, s2 = GetResult() // 在这里，函数返回两个值，所以定义两个变量来进行接收。
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
}
