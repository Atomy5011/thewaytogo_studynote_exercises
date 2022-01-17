package chapter6

import (
	"fmt"
	"strings"
)

//练习 6.7
//包 strings 中的 Map 函数和 strings.IndexFunc() 一样都是非常好的使用例子。
//请学习它的源代码并基于该函数书写一个程序，要求将指定文本内的所有非 ASCII 字符替换成 ? 或空格。

// strings.Map example
// func main() {
// 	rot13 := func(r rune) rune {
// 		switch {
// 		case r >= 'A' && r <= 'Z':
// 			return 'A' + (r-'A'+13)%26
// 		case r >= 'a' && r <= 'z':
// 			return 'a' + (r-'a'+13)%26
// 		}
// 		return r
// 	}
// 	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
// }

// string.IndexFunc example
// func main() {
// 	f := func(c rune) bool {
// 		return unicode.Is(unicode.Han, c)
// 	}
// 	fmt.Println(strings.IndexFunc("Hello, 世界", f))
// 	fmt.Println(strings.IndexFunc("Hello, world", f))
// }

func ReplaceUnAscii(s string) {
	res := func(c rune) rune {
		if c > 127 {
			return ' '
		}
		return c
	}
	fmt.Println(strings.Map(res, s))
}
