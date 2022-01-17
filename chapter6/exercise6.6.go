package chapter6

//练习 6.6
//实现一个输出前 30 个整数的阶乘的程序。

func Factorial(num uint64) (res uint64) {
	res = 1
	if num > 0 {
		res = num * Factorial(num-1)
		return
	}
	return
}
