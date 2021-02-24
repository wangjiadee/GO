package main

import "fmt"

func main() {
	ch := make(chan int)
	var WriteCh chan <- int = ch
	WriteCh <- 123

	var ReadCh <- chan int = ch
	num := <-ReadCh
	fmt.Println(num)
}