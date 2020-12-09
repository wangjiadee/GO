package main

import "fmt"

// for 循环语句的语法结构
func For_Struction_Eg() {
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Println("love bb!")
	}
	fmt.Println(i)
}

// for 循环案例1 求1-100的所有数字
func For_example() {
	var i int
	var sum int
	for i = 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}

// for 循环案例2 求1-100的偶数和
func For_example2() {
	var sum int
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

// break 方法终止循环
func Break_example() {
	var i int
	for i = 0; i < 10; i++ {
		if i == 3 {
			break
		}
	}
	fmt.Println(i)
}

// 用户登录系统模拟
func User_Login() {
	var username string
	var password string
	var count int
	for {
		fmt.Println("Pelase input username:")
		fmt.Scan(&username)
		fmt.Println("Pelase input password:")
		fmt.Scan(&password)
		if username == "mdzz" && password == "sb" {
			fmt.Println("loading success!")
			break
		} else {
			count++
			if count >= 3 {
				fmt.Println("Error more than three")
				break
			}
			fmt.Println("pelase agian!")
		}
	}
}

// continue 退出本
func Continue_example() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

// 用continue实现：计算1到100(含)之间的除了能被7整除之外所有整数的和。
func Continue_example2() {
	var sum int
	for i := 1; i <= 100; i++ {
		if i%7 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}

func For_example3() {
	// 三角形
	for j := 1; j <= 5; j++ {
		for i := 1; i <= j; i++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

}

func For_example4() {
	// 2006年营业额80000元，每年增长25%，请问按此增长速度，到哪一年营业额将达到20万元

	var year int = 2006
	var money float64
	for money = 80000; money <= 200000; money = money * 1.25 {
		year += 1
	}
	fmt.Println(year)
}

func For_example5() {
	// 找出100-999间的水仙花数
	// 水仙花：是指一个三位数，它的每位数字的立方和等于其本身
	// 153 =1+125+27
	// 1: 构建循环条件
	var h int // 百位
	var t int // 十位
	var b int // 个位
	for i := 100; i <= 999; i++ {
		// ctrl+alt+l //快速排版的快捷键
		// 2: 进行计算
		h = i / 100
		t = i % 100 / 10
		b = i % 10
		if h*h*h+t*t*t+b*b*b == i {
			fmt.Println(i)
		}
	}

	// 3: 打印输出结果
}

func For_example6() {
	// 1: 考虑一行展示
	for j := 1; j <= 9; j++ { //行
		for i := 1; i <= j; i++ {
			// 1*1=1
			fmt.Printf("%d*%d=%d\t", j, i, j*i) // \t --tab
		}
		fmt.Print("\n")
	}

	// 2； 考虑多行展示
}

func For_example7() {
	i := 0
	num := 5
	for i <= num {
		fmt.Println(i)
		i++
	}
}

// #############start main function

func main() {
	// For_Struction_Eg()
	// For_example()
	// For_example2()
	// Break_example()
	// User_Login()
	// Continue_example()
	// Continue_example2()
	// For_example6()
	For_example7()
}
