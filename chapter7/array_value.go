package chapter7

import "fmt"

//证明当数组赋值时，发生了数组内存拷贝。

func ArrayValue(){
	arr1 := [...]int{1, 2, 3}

	arr2 := arr1
	arr2[1] = 100
	
	for idx, val := range arr1{
		fmt.Printf("Array arr1 at index %d is %d\n", idx, val)
	}

	fmt.Println("*********************")

	for idx, val := range arr2{
		fmt.Printf("Array arr2 at index %d is %d\n", idx, val)
	}

}