package chapter6

import "fmt"

//练习 6.5
//使用递归函数从 10 打印到 1。

func PrintNumRecursion(n int){
	if n > 10{
		return
	}
	PrintNumRecursion(n + 1)
	fmt.Println(n)
}
       
