# 基本类型和运算符

## 布尔类型 bool

布尔型的值只可以是常量 true 或者 false。

- 两个类型相同的值可以使用相等 `==` 或者不等 `!=` 运算符来进行比较并获得一个布尔型的值
- 布尔型的常量和变量也可以通过和逻辑运算符（非 `!`、与 `&&`、或 `||`）结合来产生另外一个布尔值
  - 在 Go 语言中，&& 和 || 是具有快捷性质的运算符，当运算符左边表达式的值已经能够决定整个表达式的值的时候（&& 左边的值为 false，|| 左边的值为 true），运算符右边的表达式将不会被执行。利用这个性质，如果你有多个条件判断，应当将计算过程较为复杂的表达式放在运算符的右侧以减少不必要的运算。
- 在格式化输出时，你可以使用 `%t` 来表示你要输出的值为布尔型。

## 数字类型

Go 语言中没有 float 类型。（Go语言中只有 float32 和 float64）没有double类型。

与操作系统架构无关的类型都有固定的大小，并在类型的名称中就可以看出来：

> 整数：
>
> - int8（-128 -> 127）
> - int16（-32768 -> 32767）
> - int32（-2,147,483,648 -> 2,147,483,647）
> - int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
>
> 无符号整数：
>
> - uint8（0 -> 255）
> - uint16（0 -> 65,535）
> - uint32（0 -> 4,294,967,295）
> - uint64（0 -> 18,446,744,073,709,551,615）
>
> 浮点型（IEEE-754 标准）：
>
> - float32（+- 1e-45 -> +- 3.4 * 1e38）
> - float64（+- 5 * 1e-324 -> 107 * 1e308）

- int 型是计算最快的一种类型。`%d` 用于格式化整数
- 整型的零值为 0，浮点型的零值为 0.0。
- float32 精确到小数点后 7 位，float64 精确到小数点后 15 位，`%g` 用于格式化浮点型，`%f` 输出浮点数
  - 由于精确度的缘故，你在使用 `==` 或者 `!=` 来比较浮点数时应当非常小心
  - 你应该尽可能地使用 float64，因为 `math` 包中所有有关数学运算的函数都会要求接收这个类型。
- 可以通过增加前缀 0 来表示 8 进制数（如：077）
- 增加前缀 0x 来表示 16 进制数（如：0xFF），`%x` 和 `%X` 用于格式化 16 进制表示的数字
- 使用 e 来表示 10 的连乘（如： 1e3 = 1000，或者 6.022e23 = 6.022 x 1e23），`%e` 输出科学计数表示法

- `%0nd` 用于规定输出长度为 n 的整数，其中开头的数字 0 是必须的。
- `%n.mg` 用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 或者 f，例如：使用格式化字符串 `%5.2e` 来输出 3.4 的结果为 `3.40e+00`。

## 复数

```go
complex64 (32 位实数和虚数)
complex128 (64 位实数和虚数)

示例：
var c1 complex64 = 5 + 10i
fmt.Printf("The value is: %v", c1)
// 输出： 5 + 10i
```

## 位运算

`%b` 是用于表示位的格式化标识符。

### 二元运算符

按位与 `&`、按位或 `|`、按位异或 `^`、位清除 `&^`(将指定位置上的值设置为 0。)

### 一元运算符

按位补足 `^`、位左移 `<<`、位右移 `>>`

- 按位补足 `^`：该运算符与异或运算符一同使用，即 `m^x`，对于无符号 x 使用“全部位设置为 1”，对于有符号 x 时使用 `m=-1`。例如：

  ```
    ^10 = -01 ^ 10 = -11
  ```

- 位左移 `<<`：用法：`bitP << n`。`bitP` 的位向左移动 n 位，右侧空白部分使用 0 填充
- 位右移 `>>`：用法：`bitP >> n`。`bitP` 的位向右移动 n 位，左侧空白部分使用 0 填充

**位左移常见实现存储单位的用例**

```go
type ByteSize float64
const (
	_ = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
```

## 运算符

- 逻辑运算符：`==`、`!=`、`<`、`<=`、`>`、`>=`。
  - 它们之所以被称为逻辑运算符是因为它们的运算结果总是为布尔值 `bool`。

- 常见二元运算符有 `+`、`-`、`*` 、`/`、`%`、`++`、`--`。
  - `/` 对于整数运算而言，结果依旧为整数，例如：`9 / 4 -> 2`。

  - 取余运算符只能作用于整数：`9 % 4 -> 1`。

  - 可以将语句 `b = b + a` 简写为 `b += a`，同样的写法也可用于 `-=`、`*=`、`/=`、`%=`。

  - 一元运算符 `++`（递增）和 `--`（递减），只能用于后缀：

    ```
    i++ -> i += 1 -> i = i + 1
    i-- -> i -= 1 -> i = i - 1
    ```

  - 带有 `++` 和 `--` 的只能作为语句，而非表达式，因此 `n = i++` 这种写法是无效的，

- 运算符与优先级，由上至下代表优先级由高到低：

  ```
  优先级 	运算符
   7 		^ !
   6 		* / % << >> & &^
   5 		+ - | ^
   4 		== != < <= >= >
   3 		<-
   2 		&&
   1 		||
  ```

## 随机数

`rand` 包实现了伪随机数的生成

例：生成 10 个非负随机数

```go
package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}
```

- 函数 `rand.Float32` 和 `rand.Float64` 返回介于 [0.0, 1.0) 之间的伪随机数，其中包括 0.0 但不包括 1.0。
- 函数 `rand.Intn` 返回介于 [0, n) 之间的伪随机数。
- 你可以使用 `rand.Seed(value)` 函数来提供伪随机数的生成种子，一般情况下都会使用当前时间的纳秒级数字

## 类型别名

在使用某个类型时，你可以给它起另一个名字，然后你就可以在你的代码中使用新的名字（用于简化名称或解决名称冲突）

```
type TZ int
```

- TZ 就是 int 类型的新名称（用于表示程序中的时区），然后就可以使用 TZ 来操作 int 类型的数据

- 实际上，类型别名得到的新类型并非和原类型完全相同，新类型不会拥有原类型所附带的方法

## 字符类型

严格来说，这并不是 Go 语言的一个类型，字符只是整数的特殊用例。`byte` 类型是 `uint8` 的别名

在 ASCII 码表中，A 的值是 65，而使用 16 进制表示则为 41，所以下面的写法是等效的:

```
var ch byte = 65 或 var ch byte = '\x41'
```

包 `unicode` 包含了一些针对测试字符的非常有用的函数（其中 `ch` 代表字符）：

- 判断是否为字母：`unicode.IsLetter(ch)`
- 判断是否为数字：`unicode.IsDigit(ch)`
- 判断是否为空白符号：`unicode.IsSpace(ch)`

# 字符串

- `string` 类型的零值为长度为零的字符串，即空字符串 `""`。

- 通过函数 `len()` 来获取字符串所占的字节长度，例如：`len(str)`。

- 字符串的内容（纯字节）可以通过标准索引法来获取，在中括号 `[]` 内写入索引，索引从 0 开始计数：

  - 字符串 str 的第 1 个字节：`str[0]`
  - 第 i 个字节：`str[i - 1]`
  - 最后 1 个字节：`str[len(str)-1]`

- 字符串拼接符 `+`

  - 由于编译器行尾自动补全分号的缘故，加号 `+` 必须放在第一行。

    ```
    str := "Beginning of the string " +
    	"second part of the string
    ```

# strings 和 strconv 包

## 前缀和后缀

- `HasPrefix` 判断字符串 `s` 是否以 `prefix` 开头：

  - ```go
    strings.HasPrefix(s, prefix string) bool
    ```

- `HasSuffix` 判断字符串 `s` 是否以 `suffix` 结尾：

  - ```go
    strings.HasSuffix(s, suffix string) bool
    ```

## 字符串包含关系

- `Contains` 判断字符串 `s` 是否包含 `substr`：

  - ```
    strings.Contains(s, substr string) bool
    ```

## 判断子字符串或字符在父字符串中出现的位置（索引）

- `Index` 返回字符串 `str` 在字符串 `s` 中的索引（`str` 的第一个字符的索引），-1 表示字符串 `s` 不包含字符串 `str`：

  - ```
    strings.Index(s, str string) int
    ```

- `LastIndex` 返回字符串 `str` 在字符串 `s` 中最后出现位置的索引（`str` 的第一个字符的索引），-1 表示字符串 `s` 不包含字符串 `str`：

  - ```
    strings.LastIndex(s, str string) int
    ```

- 如果需要查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位：

  - ```
    strings.IndexRune(s string, r rune) int
    ```

## 字符串替换

- `Replace` 用于将字符串 `str` 中的前 `n` 个字符串 `old` 替换为字符串 `new`，并返回一个新的字符串，如果 `n = -1` 则替换所有字符串 `old` 为字符串 `new`：

  - ```
    strings.Replace(str, old, new, n) string
    ```

## 统计字符串出现次数

- `Count` 用于计算字符串 `str` 在字符串 `s` 中出现的非重叠次数：

  - ```
    strings.Count(s, str string) int
    ```

## 重复字符串

- `Repeat` 用于重复 `count` 次字符串 `s` 并返回一个新的字符串
  - ```
    strings.Repeat(s, count int) string
    ```

## 修改字符串大小写

- `ToLower` 将字符串中的 Unicode 字符全部转换为相应的小写字符：

  - ```
    strings.ToLower(s) string
    ```

- `ToUpper` 将字符串中的 Unicode 字符全部转换为相应的大写字符：
  - ```
    strings.ToUpper(s) string
    ```

## 修剪字符串

- 使用 `strings.TrimSpace(s)` 来剔除字符串开头和结尾的空白符号

- 想要剔除指定字符，则可以使用 `strings.Trim(s, subs)`，该函数的第二个参数可以包含任何字符
- 如果你只想剔除开头或者结尾的字符串，则可以使用 `TrimLeft` 或者 `TrimRight` 来实现。

## 分割字符串

- `strings.Fields(s)` 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。

- `strings.Split(s, sep)` 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。

## 拼接 slice 到字符串

- `Join` 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串：

  - ```
    strings.Join(sl []string, sep string) string
    ```

示例：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
}
```

输出

```
Splitted in slice: [The quick brown fox jumps over the lazy dog]
The - quick - brown - fox - jumps - over - the - lazy - dog -
Splitted in slice: [GO1 The ABC of Go 25]
GO1 - The ABC of Go - 25 -
sl2 joined by ;: GO1;The ABC of Go;25
```

## 从字符串中读取内容

函数 `strings.NewReader(str)` 用于生成一个 `Reader` 并读取字符串中的内容，然后返回指向该 `Reader` 的指针，从其它类型读取内容的函数还有：

- `Read()` 从 []byte 中读取内容。
- `ReadByte()` 和 `ReadRune()` 从字符串中读取下一个 byte 或者 rune。

## 字符串与其它类型的转换

- 与字符串相关的类型转换都是通过 `strconv` 包实现的。
- 任何类型 **T** 转换为字符串总是成功的。

- 针对从数字类型转换到字符串，Go 提供了以下函数：

  - `strconv.Itoa(i int) string` 返回数字 i 所表示的字符串类型的十进制数。

  - `strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string` 将 64 位浮点型的数字转换为字符串，其中 `fmt` 表示格式（其值可以是 `'b'`、`'e'`、`'f'` 或 `'g'`），`prec` 表示精度，`bitSize` 则使用 32 表示 float32，用 64 表示 float64。

- 针对从字符串类型转换为数字类型，Go 提供了以下函数：
  - `strconv.Atoi(s string) (i int, err error)` 将字符串转换为 int 型。
  - `strconv.ParseFloat(s string, bitSize int) (f float64, err error)` 将字符串转换为 float64 型。

# 时间和日期

- `time` 包为我们提供了一个数据类型 `time.Time`（作为值使用）以及显示和测量时间和日期的功能函数。
- 当前时间可以使用 `time.Now()` 获取
- 或者使用 `t.Day()`、`t.Minute()` 等等来获取时间的一部分

- Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。

- Location 类型映射某个时区的时间，UTC 表示通用协调世界时间。
- 包中的一个预定义函数 `func (t Time) Format(layout string) string` 可以根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串
  - 你可以使用一些预定义的格式，如：`time.ANSIC` 或 `time.RFC822`。

```go
package main
import (
	"fmt"
	"time"
)

var week time.Duration
func main() {
	t := time.Now()
	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011
	t = time.Now().UTC()
	fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
	// The time must be 2006-01-02 15:04:05
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
}
```

- 如果你需要在应用程序在经过一定时间或周期执行某项任务（事件处理的特例），则可以使用 `time.After` 或者 `time.Ticker`，`time.Sleep（d Duration）` 可以实现对某个进程（实质上是 goroutine）时长为 d 的暂停。这些会在后续章节出现

# 指针

- Go 语言的取地址符是 `&`，放到一个变量前使用就会返回相应变量的内存地址。
- 不能进行指针运算
- **一个指针变量可以指向任何一个值的内存地址**
- 当一个指针被定义后没有分配到任何变量时，它的值为 `nil`。
- 一个指针变量通常缩写为 `ptr`。
- 不能获取字面量（`&10`中10就是字面量）或常量的地址
- 使用指针会减少内存占用和提高效率。被指向的变量也保存在内存中，直到没有任何指针指向它们，所以从它们被创建开始就具有相互独立的生命周期。