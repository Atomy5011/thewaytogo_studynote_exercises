package main

import "fmt"

func CharLoopForOnce(c string, times int) {
	res := c
	for i := 1; i <= times; i++ {
		fmt.Println(res)
		res += c
	}
}

func CharLoopForTwice(c string, times int) {
	for i := 1; i <= times; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(c)
		}
		fmt.Println()
	}
}

// func main() {
// 	CharLoopForOnce("G", 25)
// 	fmt.Println("****************")
// 	CharLoopForTwice("G", 25)
// }
