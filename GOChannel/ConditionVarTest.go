package main

/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-27 14:00:15
 * @LastEditTime: 2021-02-27 14:32:58
 * @FilePath: \GO\src\go.ralph.org\GOChannel\ConditionVarTest.go
 * @Effect: DO
 */

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define global condition variables
var cond sync.Cond

func main() {
	//Channel used to end communication
	quit := make(chan bool)
	//Shared resource channel
	product := make(chan int, 3)
	//Set random number
	rand.Seed(time.Now().UnixNano())
	// 创建条件变量和使用的锁  互斥锁初始值0,未加锁状态
	cond.L = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go producer(product, i+1)
	}

	for i := 0; i < 5; i++ {
		go consumer(product, i+1)
	}

	// 主协成堵塞 不结束
	<-quit
}

func producer(out chan<- int, index int) {
	for {
		// 加锁：条件变量对应互斥锁加锁
		cond.L.Lock()
		for len(out) == 5 {
			cond.Wait() //挂起当前go程 等待条件变量满足 呗消费者唤醒
		}
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", index, num, len(out))
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(in <-chan int, index int) {
	for {
		cond.L.Lock()
		for len(in) == 0 {
			cond.Wait()
		}
		num := <-in
		fmt.Printf("---- %dth 消费者, 消费数据 %3d,公共区剩余%d个数据\n", index, num, len(in))
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 500)
	}
}
