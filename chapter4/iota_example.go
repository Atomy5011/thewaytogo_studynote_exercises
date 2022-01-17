package main

import "fmt"

const (
	a = iota + 1
	b
	c
)

func PrintConst1() {
	fmt.Printf("a = %v, b = %v, c = %v \n", a, b, c)
}
