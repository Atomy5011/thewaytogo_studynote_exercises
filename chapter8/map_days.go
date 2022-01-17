package chapter8

import "fmt"

// 创建一个 map 来保存每周 7 天的名字，将它们打印出来并且测试是否存在 Tuesday 和 Hollyday。
// (这里应该是要判断有没有Tuesday，有没有Hollyday, 而不是存不存在Tuesday && Hollyday)
// (翻译错误？Hollyday？Holiday？)

var Days = map[int]string {
	0: "Sunday",
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
}

func DasyIsIn(days map[int]string) (bool){
	fmt.Printf("days map: %v", days)

	flag := false
	for _, v := range days {
		if v == "Tuesday" || v == "Hollyday" {
			flag = true
		}
	}

	return flag
}
