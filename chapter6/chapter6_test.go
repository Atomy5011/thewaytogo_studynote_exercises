package chapter6

import (
	"testing"
)

func TestMultReturnval(t *testing.T) {
	r1, r2, r3 := ComputeUnName(5, 3)
	t.Logf("sum: %d, mult: %d, sub: %d.", r1, r2, r3)
	r4, r5, r6 := ComputeWithName(10, 10)
	t.Logf("sum: %d, mult: %d, sub: %d.", r4, r5, r6)
}

func TestErrorReturnval(t *testing.T) {
	normal_val := 36.00
	minus_val := -36.00
	sqrtNum1, err1 := MySqrtUnName(normal_val)
	sqrtNum2, err2 := MySqrtWithName(minus_val)
	t.Logf("sqrtNum1: %f", sqrtNum1)
	t.Logf("sqrtNum2: %f", sqrtNum2)
	t.Error(err1)
	t.Error(err2)
}

func TestVarargs(t *testing.T) {
	Varargs("0", "1", "2", "3", "4")
	Varargs("Hello", "world")
}

func TestFibonacci(t *testing.T) {
	res, num := Fibonacci(10)
	t.Logf("res: %d", res)
	t.Logf("num: %d", num)
}

func TestPringNumRecursion(t *testing.T) {
	PrintNumRecursion(1)
}

func TestFactorial(t *testing.T) {
	for i := uint64(1); i <= 30; i++ {
		t.Logf("Factorial of %d is %d\t", i, Factorial(i))
	}
}

func TestReplaceUnAscii(t *testing.T) {
	ReplaceUnAscii("test测test试test")
}

func TestClosure(t *testing.T) {
	Closure()
}