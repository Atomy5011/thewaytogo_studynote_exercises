package chapter7

import "fmt"

// 为练习 7.3 写一个新的版本，
// 主函数调用一个使用序列个数作为参数的函数，该函数返回一个大小为序列个数的 Fibonacci 切片。

func FibonacciFuncArray(num int) {
	// 其实7.3写了7.4的东西。。。。。。
	res := make([]int, num)
	res[0] = 1
	res[1] = 1

	for i := 2; i < num; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	fmt.Println(res)
}