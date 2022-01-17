package chapter6

// 编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。

func ComputeUnName(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

func ComputeWithName(i, j int) (sum, mult, sub int) {
	sum = i + j
	mult = i * j
	sub = i - j
	return sum, mult, sub
}
