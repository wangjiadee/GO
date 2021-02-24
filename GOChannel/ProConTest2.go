package main

import (
	"fmt"
	//"time"
)

type OrderInfo struct {
	id int   //共享资源 订单嗯呢
}

func Produce(Make chan<- OrderInfo){
	for i:=1;i<11 ;i++  {
		fmt.Println("美团外卖为您接单-单号: " ,i)
		info := OrderInfo{id: i }
		Make <- info  //写入缓冲区
	}

	close(Make)
}

func Consumer(Pay <- chan OrderInfo){
	for things := range Pay{
		fmt.Println("您已完成订单-订单号 : ",things.id)
		//time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan OrderInfo)
	go Produce(ch)
	Consumer(ch)
}
