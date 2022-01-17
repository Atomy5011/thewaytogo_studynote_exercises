package chapter7

import "fmt"

func LenSlice(a []int, n int) {
	fmt.Println(len(a[n:n]))
	fmt.Println(len(a[n:n+1]))
}