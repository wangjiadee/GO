/*
 * @Author: your name
 * @Date: 2020-08-19 15:02:31
 * @LastEditTime: 2020-08-19 16:01:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\listen1\string_strut\main.go
 */
/*
Notes：

Contains :  判断字符串s是否包含了某个子字符
			存在返回T，不存在返回F
			func Contains(s, substr string) bool

join：		将一系列字符串连接为一个字符串，之间用sep来分隔。

index:		在某个字符串中查找某个字符串的位置
			步长从0开始

Repeat ：   某个字符串重复多少次，返回的是重复后的字符串

Replace：   在s字符串中，把old字符串替换为new字符串，
			n表示替换的次数，小于0表示全部替换

Split  ：   把s字符串按照sep分割，返回slice（切片）


@@LINK@@:	https://studygolang.com/pkgdoc
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// 以某一个固定的字符去切割字符串
	var str1 []string = []string{"172.171.171.1", "172.17.171.2", "172.17.171.3"}
	ret := strings.Join(str1, "_")
	fmt.Println(ret)

	// 指定某一个字符串来划分
	ips := "17.28.29.20,172.29.28.27"
	ret1 := strings.Split(ips, ",")
	fmt.Println(ret1)

	// Contains 判断是包含某一个字符串 并换回波尔类型
	ret2 := strings.Contains(ips, "17.28.29.20")
	fmt.Println(ret2)

	// 判断字符串以什么开头 以什么结尾
	str2 := "https://www.pornhub.com"
	if strings.HasPrefix(str2, "https") {
		fmt.Println("This is save url!")
	} else {
		fmt.Println("This is unsave url")
	}
	if strings.HasSuffix(str2, "pornhub.com") {
		fmt.Println("This is X url")
	} else {
		fmt.Println("This is health url")
	}

	Index := strings.Index(str2, "porn")
	fmt.Println(Index)

	str := "oppo find x"
	str3 := strings.Replace(str, "x", "x2", 1)
	fmt.Println(str3)

	s := []string{"Oppo", "vivo", "huawei"}
	str := strings.Join(s, "--")
	fmt.Println(str)
}
