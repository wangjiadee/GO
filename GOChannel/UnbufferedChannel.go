package main

import "fmt"

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