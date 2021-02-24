package main

import "fmt"

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