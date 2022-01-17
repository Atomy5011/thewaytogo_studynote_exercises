package chapter7

import "fmt"

// 给定切片 sl，将一个 []byte 数组追加到 sl 后面。
// 写一个函数 Append(slice, data []byte) []byte，该函数在 sl 不能存储更多数据的时候自动扩容。

func Append(slice, data []byte) {
	idx := 0
	temp := byte(0)
	for _, val := range slice {
		if val != temp {
			slice[idx] = val
			fmt.Println(slice[idx])
			idx++
		}
		temp = val
	}

	fmt.Println(slice)
}