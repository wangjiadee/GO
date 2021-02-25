/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-25 20:50:33
 * @LastEditTime: 2021-02-25 21:11:47
 * @FilePath: \GO\src\ci.Ralph.org\GOChannel\DeadLock1.go
 * @Effect: DO
 */
package main

import "fmt"

/*
死锁方式1：go程自己死锁
*/
func DeadLock1() {
	ch := make(chan int)
	ch <- 676 //执行后堵塞
	num := <-ch
	fmt.Println(num)
}

/*
go程间channel访问顺序导致死锁
*/
func DeadLock2() {
	ch := make(chan int)
	num := <-ch //读然后堵塞
	fmt.Println("NUM:", num)

	go func() {
		ch <- 345
	}()
}

/*
多go程，多channel 交叉死锁
*/
func DeadLock3() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()

	for {
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}
}

func main() {
	DeadLock1()
	DeadLock2()
	DeadLock3()

}
