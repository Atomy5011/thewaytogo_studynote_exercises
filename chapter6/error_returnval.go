package chapter6

import (
	"errors"
	"math"
)

//编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，
//如果参数是一个负数的话将返回一个错误。编写两个版本，一个是非命名返回值，一个是命名返回值。

func MySqrtUnName(num float64) (float64, error) {
	if num < 0 {
		return float64(math.NaN()), errors.New("please pass in either 0 or a positive number")
	}
	return math.Sqrt(num), nil
}

func MySqrtWithName(num float64) (ret float64, err error) {
	if num < 0 {
		ret = float64(math.NaN())
		err = errors.New("please pass in either 0 or a positive number")
	} else {
		ret = math.Sqrt(num)
	}
	return
}
