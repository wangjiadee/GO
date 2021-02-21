/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-18 20:04:41
 * @LastEditTime: 2021-02-18 21:33:37
 * @FilePath: \src\github.com\Ralph.org\GoPointer\PointerTest.go
 * @Effect: DO
 */
package main

import "fmt"

// 指针就是地址  指针变量就是存储地址的变量
//
// func main()  {
// 	var a int = 10
// 	var p *int = &a
// 	a =100
// 	fmt.Println("a = ",a)

// 	*p = 250   //借助a变量的地址来操作a
// 	fmt.Println("a = ",a)
// 	fmt.Println("*p = ",*p)

// 	a = 1000
// 	fmt.Println("*p = ",*p)
// }

// func main(){
// 	var p *int
// 	p = new(int)
// 	*p = 1000
// 	fmt.Println(*p)
// }

func swap(a, b int) {
	a, b = b, a
	fmt.Println("SWAP2", a, b)
}

func swap2(x, y *int) {
	*x, *y = *y, *x

}
func main() {
	a, b := 20, 40
	swap(a, b)
	fmt.Println("MAIN", a, b)
	swap2(&a, &b) //传递地址
	fmt.Println("SWAP2", a, b)
	fmt.Println("MAIN", a, b)
}
