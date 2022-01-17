package chapter7

import (
	"fmt"
	"math"
)

func MinSlice(s []int) {
	min := math.MaxInt32 //max := 0

	for _, v := range s{
		if v < min{		// v > max
			min = v
		}
	}

	fmt.Printf("The min value is: %v", min)
}