package main

import (
	"fmt"
	"strings"
)

// 字符串翻转(首位交换)
func main() {
	str := "012345678"
	arr := strings.Split(str, "")
	fmt.Println(arr)
	first := 0
	last := len(arr) - 1
	for {
		if first < last {
			arr[first], arr[last] = arr[last], arr[first]
			first++
			last--
		} else {
			break
		}
	}
	str = strings.Join(arr, "")
	fmt.Println(str)
}
