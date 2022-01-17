package chapter7


func Sum(f []float32)(sum float32){
	for _, v := range f{
		sum += v
	}
	return
}

func SumAndAverage(s []int)(int, float32){
	sum := 0
	for _, v := range s{
		sum += v
	}
	return sum, float32(sum) / float32(len(s)) 
	// 注：参考答案是float32(sum / len(a)), 返回的不是正确的平均值，因为sum / len(a)返回了int类型的数据
}