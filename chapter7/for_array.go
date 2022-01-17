package chapter7

import "fmt"

//写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。

func ForArray(num int) {
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}
