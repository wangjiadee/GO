/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-19 21:15:48
 * @LastEditTime: 2021-02-19 21:49:07
 * @FilePath: \GO\src\github.com\Ralph.org\GOmap\mapTest3.go
 * @Effect: DO
 */
package main
import (
	"strings"
	"fmt"
)

func WordCountFunc(str string) map[string]int{
	s := strings.Fields(str)    // 将字符串，拆分成 字符串切片s
	m := make(map[string]int)	// 创建一个用于存储 word 出现次数的 map

	for i:=0;i< len(s);i++{
		if _,ok := m[s[i]];ok {  //ok ==true 往下走
			m[s[i]]++     //m[s[i]] = m[s[i]] + 1
		} else {
			m[s[i]] = 1   //只出现一次默认为1
		}
	}
	return m
}

func main() {
	str := "I Fuck my work and I I I I Fuck Fuck Fuck my family too"
	mRet := WordCountFunc(str)

	// 遍历map ，展示每个word 出现的次数：
	for k, v := range mRet {
		fmt.Printf("%q:%d\n", k, v)
	}
}