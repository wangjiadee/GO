/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2021-02-18 22:36:42
 * @LastEditTime: 2021-02-18 23:48:26
 * @FilePath: \src\github.com\Ralph.org\GOSlice\SliceTest.go
 * @Effect: DO
 */

package main

import "fmt"

func DeleteNull(data []string) []string {
	out := data[:0]
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}
	return out

}

func DeleteNull2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "" {
			data[i] = str
			i++
		}
	}
	return data[:i]

}

func main() {
	data := []string{"Fuck", "", "Fucking", "", "", "umama", "blue"}
	data2 := DeleteNull(data)
	data3 := DeleteNull2(data)
	fmt.Println(data2)
	fmt.Println(data3)

}
