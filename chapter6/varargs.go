package chapter6

import "fmt"

// 写一个函数，该函数接受一个变长参数并对每个元素进行换行打印。

func Varargs(s ...string) {
	for _, v := range s{
		fmt.Println(v)
	}
}
