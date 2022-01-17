package chapter7

import "fmt"

// 在第 6.6 节我们看到了一个递归计算 Fibonacci 数值的方法。
// 但是通过数组我们可以更快的计算出 Fibonacci 数。完成该方法并打印出前 50 个 Fibonacci 数字。

func FibonacciArray(num int) {
	res := make([]int, num)
	res[0] = 1
	res[1] = 1

	for i := 2; i < num; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	fmt.Println(res)
}