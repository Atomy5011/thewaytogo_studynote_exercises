package main

import "fmt"

const (
	FIZZ = 3
	BUZZ = 5
	FIZZBUZZ = 15
)

func ForIF() {
	for i:= 1; i <= 100; i++ {
		if i % FIZZBUZZ == 0 {
			fmt.Printf("FizzBuzz\t")
		}else if i % FIZZ == 0 {
			fmt.Printf("Fizz\t")
		}else if i % BUZZ == 0 {
			fmt.Printf("Buzz\t")
		}else {
			fmt.Printf("%d\t", i)
		}
	}
}

// func main(){
// 	for i:= 1; i <= 100; i++ {
// 		switch{
// 		case i % FIZZBUZZ == 0:
// 			fmt.Printf("FizzBuzz\t")
// 		case i % FIZZ == 0:
// 			fmt.Printf("Fizz\t")
// 		case i%BUZZ == 0:
// 			fmt.Printf("Buzz\t")
// 		default:
// 			fmt.Printf("%d\t", i)
// 		}
// 	}
// }