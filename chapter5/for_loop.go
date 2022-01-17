package main

import "fmt"

func ForLoopV1(){
	for i := 1; i <= 15; i++ {
		fmt.Printf("Round %d\n", i)
	}
}

func ForloopV2(){
	i := 1
	loop2:
	fmt.Printf("Round %d\n", i)
	i++
	if i <= 15{
		goto loop2
	}
}

// func main(){
// 	ForLoopV1()
// 	fmt.Println("****************")
// 	ForloopV2()
// }