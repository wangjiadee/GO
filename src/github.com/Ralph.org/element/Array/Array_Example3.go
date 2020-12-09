/*
 * @Author: your name
 * @Date: 2020-08-28 11:29:36
 * @LastEditTime: 2020-08-28 11:32:04
 * @LastEditors: Please set LastEditors
 * @Description: 查找最大的字符串
 * @FilePath: \newgo\Ralph.org\element\Array\Array_Example3.go
 */
package main

import "fmt"

func main() {
	names := [...]string{"马龙", "迈克尔乔丹", "雷吉米勒", "蒂姆邓肯", "科比布莱恩特"}
	var max string = names[0]
	for i := 0; i < len(names); i++ {
		if len((names[i])) > len(max) {
			max = names[i]

		}
	}
	fmt.Println(max)
}
