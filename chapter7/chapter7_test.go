package chapter7

import "testing"

func TestArrayValue(t *testing.T) {
	ArrayValue()
}

func TestForArray(t *testing.T) {
	ForArray(15)
}

func TestFibonacciArray(t *testing.T) {
	FibonacciArray(50)
}

func TestAppend(t *testing.T) {
	var d []byte = []byte{'a', 'b', 'c'}
	var s []byte = make([]byte, len(d))
	Append(s, d)
}

func TestQuestion7_5(t *testing.T) {
	ValueDoubleV1()
	t.Log("\n**** double Items value ****")
	ValueDoubleV2()
}

func TestSumArray(t *testing.T) {
	// cannot use test01 (variable of type [3]float32) as []float32 value in argument to Sum
	// var test01 = [3]float32{1.1, 2.2, 3.3} 
	var test02 = []float32{1.1, 2.2, 3.3}
	t.Log(Sum(test02))
	var test03 = []int{2, 5, 4}
	sum, avg := SumAndAverage(test03)
	t.Logf("sum: %v, avg: %.3v", sum, avg)
}

func TestQuestion7_7(t *testing.T){
	var a = []int{0, 1, 2, 3, 4, 5}
	LenSlice(a, 1)
}
