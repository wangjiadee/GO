## GO中的channel



1. channel是一个`数据类型`，主要用来解决协程的同步问题以及协程之间数据共享（数据传递）的问题。
2. goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
3. 引⽤类型 channel可用于多个 goroutine 通讯。其内部实现了同步，确保并发安全。

![img](file:///C:\Users\ADMINI~1\AppData\Local\Temp\ksohtml1140\wps1.png)	 



###### 定义channel变量

​		和map类似，channel也一个对应make创建的底层数据结构的引用。当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。

```
make(chan Type)  //等价于make(chan Type, 0)
make(chan Type, capacity)
make(chan 在channel中传递的数据类型,容量)  容量= 0： 无缓冲channel， 容量 > 0 ：有缓冲channel
```

###### 	channel有两个端：

​		另一端： 写端（传入端）    chan <-
​		另一端： 读端（传出端）    <- chan
​		要求：读端和写端必须同时满足条件，才在shan上进行数据流动。否则，则阻塞。

```go
package main

import (
	"fmt"
	"time"
	"runtime"
)
// 全局定义channel， 用来完成数据同步
var channel = make(chan int)

// 定义一台打印机
func printer(s string)  {
	for _, ch := range s {
		fmt.Printf("%c", ch)			// 屏幕：stdout
		time.Sleep(300 * time.Millisecond)
	}
}

// 定义两个人使用打印机
func person1()  {				// person 先执行。
	printer("hello")
	channel <- 8
}
func person2()  {				// person 后执行
						// 		<- channel
	printer("wrold")
}

func main()  {
	go person1()
	<- channel
	go person2()
	runtime.Goexit()
}

```

###### channel中的len和cap

```go
func main() {
	ch:= make(chan string)
	fmt.Println("len(ch) 表示channel中剩余未读取个数： ",len(ch))
	fmt.Println("cap(ch) 表示channel中通道容量： ",cap(ch))
	go func(){
		for i :=0;i<2 ;i++  {
			fmt.Println("i : ",i)
		}
		ch <- "子go执行完成"
	}()


	str := <-ch
	fmt.Println("str",str )

}
```

###### 无缓冲Channel

​		无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。这种类型的通道要求发送goroutine和接收goroutine同时准备好，才能完成发送和接收操作。否则，通道会导致先执行发送或接收操作的 goroutine 阻塞等待。
​		这种对通道进行发送和接收的交互行为本身就是同步的。其中任意一个操作都无法离开另一个操作单独存在。

- ​	阻塞：由于某种原因数据没有到达，当前协程（线程）持续处于等待状态，直到条件满足，才接触阻塞。
- ​	同步：在两个或多个协程（线程）间，保持数据内容一致性的机制。

```go
func main() {
	ch := make(chan int)
	go func() {
		for i:=0;i<5 ;i++  {
			fmt.Println("Sub go i Write: " ,i)
			ch <- i
		}
	}()

	for i:=0;i<5 ;i++  {
		num := <- ch
		fmt.Println("Main go read i :", num)
	}
}
```



































