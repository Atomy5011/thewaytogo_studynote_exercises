package chapter6

//练习 6.4
//重写本节中生成斐波那契数列的程序并返回两个命名返回值（详见第 6.2 节），即数列中的位置和对应的值，例如 5 与 4，89 与 10。

func Fibonacci(num int)(res, idx int){
	if num <= 1 {
		res = 1
	}else{
		temp1, _ := Fibonacci(num - 1)
		temp2, _ := Fibonacci(num - 2)
		res = temp1 + temp2
	}
	idx = num	//整这玩意有啥用，不传多少是多少嘛
	return
}