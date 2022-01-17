package chapter7

import "fmt"

// 原题目中的方式
func ValueDoubleV1() {
	items := [...]int{10, 20, 30, 40, 50}

	for _, item := range items {
		item *= 2
	}

	fmt.Printf("Items value: %v", items)
}

// 让值double
func ValueDoubleV2() {
	items := [...]int{10, 20, 30, 40, 50}
	
	for idx := range items {
		items[idx] *= 2
	}

	fmt.Printf("Items value: %v", items)
}