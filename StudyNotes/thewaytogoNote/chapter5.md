# if-else 结构

## 非常基础的三种结构

1. if

   ```
   if condition {
   	// do something	
   }
   ```

2. if-else

   ```
   if condition {
   	// do something	
   } else {
   	// do something	
   }
   ```

3. if-else if-else

   ```
   if condition1 {
   	// do something	
   } else if condition2 {
   	// do something else	
   } else {
   	// catch-all or default
   }
   ```

- else-if 分支的数量是没有限制的，但是为了代码的可读性，还是不要在 if 后面加入太多的 else-if 结构。如果你必须使用这种形式，则把尽可能先满足的条件放在前面。

## 有用的例子

1. 判断一个字符串是否为空：

   - `if str == "" { ... }`
   - `if len(str) == 0 {...}`

2. 判断运行 Go 程序的操作系统类型，这可以通过常量 `runtime.GOOS` 来判断(第 2.2 节)。

   ```
    if runtime.GOOS == "windows"	 {
    	.	..
    } else { // Unix-like
    	.	..
    }
   ```

   这段代码一般被放在 init() 函数中执行。这儿还有一段示例来演示如何根据操作系统来决定输入结束的提示：

   ```
    var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."
    
    func init() {
    	if runtime.GOOS == "windows" {
    		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")		
    	} else { //Unix-like
    		prompt = fmt.Sprintf(prompt, "Ctrl+D")
    	}
    }
   ```

3. 函数 `Abs()` 用于返回一个整型数字的绝对值:

   ```
    func Abs(x int) int {
    	if x < 0 {
    		return -x
    	}
    	return x	
    }
   ```

4. `isGreater` 用于比较两个整型数字的大小:

   ```
    func isGreater(x, y int) bool {
    	if x > y {
    		return true	
    	}
    	return false
    }
   ```

## if可以包含一个初始化语句

```
if initialization; condition {
	// do something
}
```

- 这种写法具有固定的格式（在初始化语句后方必须加上分号）

例如：

例如:

```
val := 10
if val > max {
	// do something
}
```

你也可以这样写:

```
if val := 10; val > max {
	// do something
}
```

## 测试多返回值函数的错误

**习惯用法**

```
value, err := pack1.Function1(param1)
if err != nil {
	fmt.Printf("An error occured in pack1.Function1 with parameter %v", param1)
	return err
}
```

如：

```
f, err := os.Open(name)
if err != nil {
	return err
}
doSomething(f) // 当没有错误发生时，文件对象被传入到某个函数中
doSomething
```

也可以将错误的获取放置在 if 语句的初始化部分：

```
if err := file.Chmod(0664); err != nil {
	fmt.Println(err)
	return err
}
```

或者将 ok-pattern 的获取放置在 if 语句的初始化部分，然后进行判断：

```
if value, ok := readData(); ok {
…
}
```

# switch结构

**表达式**

```
switch var1 {
	case val1:
		...
	case val2:
		...
	default:
		...
}
```

- 变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。
  - 类型不被局限于常量或整数，但必须是相同的类型
  - 或者最终结果为相同类型的表达式
- 可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：`case val1, val2, val3`
- 每一个 `case` 分支都是**唯一**的，**从上至下**逐一测试，**直到匹配为止**。

- 还希望继续执行后续分支的代码，可以使用 `fallthrough` 关键字来达到目的。

- 可选的 `default` 分支可以出现在任何顺序，但最好将它放在最后。
  - 它的作用类似与 `if-else` 语句中的 `else`，表示不符合任何已给出条件时，执行相关语句。

## 问题5.1

```go
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
```

输出：

```
was <= 6
was <= 7    
was <= 8    
default case
```

- fallthrough会让这段代码走到default分支

  - fallthrough是继续执行后续分支，default在case8后面，故而会继续执行

  - 如果我们移动default到别的地方，case8的fallthrough就会报错了，因为它没有后续分支

  - 如果我们把default放在case7后面

    ```go
    	case 6:
    		fmt.Println("was <= 6")
    		fallthrough
    	case 7:
    		fmt.Println("was <= 7")
    		fallthrough
    	default:
    		fmt.Println("default case")
    	case 8:
    		fmt.Println("was <= 8")
    ```

    此时只会返回

    ```
    was <= 6
    was <= 7       
    default case
    ```

## 练习5.2

写一个 Season 函数，要求接受一个代表月份的数字，然后返回所代表月份所在季节的名称（不用考虑月份的日期）。

```go
package main

import "fmt"

func Season(month int) string {
	switch month{
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	case 12, 1, 2:
		return "Winter"
	}
	return "Isn't month"
}

func main(){
	fmt.Println(Season(11))
}
```

输出：

```
Autumn
```

# for结构

## 基于计数器的迭代

```go
for 初始化语句; 条件语句; 修饰语句 {}
```

示例：

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}
```

- 可以在循环中同时使用多个计数器`for i, j := 0, N; i < j; i, j = i+1, j-1 {}`

### 练习5.4

1. 使用 for 结构创建一个简单的循环。要求循环 15 次然后使用 fmt 包来打印计数器的值。
2. 使用 goto 语句重写循环，要求不能使用 for 关键字。

```go
package main

import "fmt"

func ForLoopV1(){
	for i := 1; i <= 15; i++ {
		fmt.Printf("Round %d\n", i)
	}
}

func ForloopV2(){
	i := 1
	loop2:
	fmt.Printf("Round %d\n", i)
	i++
	if i <= 15{
		goto loop2
	}
}

func main(){
	ForLoopV1()
	fmt.Println("****************")
	ForloopV2()
}
```

### 练习5.5

创建一个程序，要求能够打印类似下面的结果（尾行达 25 个字符为止）：

```
G
GG
GGG
GGGG
GGGGG
GGGGGG
```

1. 使用 2 层嵌套 for 循环。
2. 仅用 1 层 for 循环以及字符串连接。

```go
package main

import "fmt"

func CharLoopForOnce(c string, times int) {
	res := c
	for i := 1; i <= times; i++ {
		fmt.Println(res)
		res += c
	}
}

func CharLoopForTwice(c string, times int) {
	for i := 1; i <= times; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(c)
		}
		fmt.Println()
	}
}

func main() {
	CharLoopForOnce("G", 25)
	fmt.Println("****************")
	CharLoopForTwice("G", 25)
}
```

### 练习5.6

使用按位补码从 0 到 10，使用位表达式 `%b` 来格式化输出

```go
package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("The complement of %b is: %b\n", i, ^i)
	}
}
```

### **练习 5.7** Fizz-Buzz

写一个从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。遇到 5 的倍数时，打印 `Buzz` 而不是相应的数字。对于同时为 3 和 5 的倍数的数，打印 `FizzBuzz`（提示：使用 switch 语句）。

```go
package main

import "fmt"

const (
	FIZZ = 3
	BUZZ = 5
	FIZZBUZZ = 15
)

//if也可实现
func ForIF() {
	for i:= 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("FizzBuzz\t")
		}else if i % 3 == 0 {
			fmt.Printf("Fizz\t")
		}else if i % 5 == 0 {
			fmt.Printf("Buzz\t")
		}else {
			fmt.Printf("%d\t", i)
		}
	}
}

func main(){
	for i:= 1; i <= 100; i++ {
		switch{
		case i % FIZZBUZZ == 0:
			fmt.Printf("FizzBuzz\t")
		case i % FIZZ == 0:
			fmt.Printf("Fizz\t")
		case i%BUZZ == 0:
			fmt.Printf("Buzz\t")
		default:
			fmt.Printf("%d\t", i)
		}
	}
}
```

### **练习 5.8**

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
```

## 基于条件判断的迭代

基本形式为：`for 条件语句 {}`，类似与while

```go
for i >= 0 {
	i = i - 1
	fmt.Printf("The variable i is now: %d\n", i)
}
```

## 无限循环

一般情况写为`for { }`

- 如果 for 循环的头部没有条件语句，那么就会认为条件永远为 true，因此**循环体内必须有相关的条件判断以确保会在某个时刻退出循环。**
  - 可以使用 break 语句或 return 语句直接返回

## for-range 结构

一般形式为：`for ix, val := range coll { }`。它可以迭代任何一个集合，简单来说就是可以迭代k-v结构。

- `val` 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值
- **如果 `val` 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值**

```go
for pos, char := range str {
...
}
```

# Break 与 continue

- `break`和`continue`都是用来控制循环结构的，主要作用是停止循环。
- `break`跳出一个循环体或者完全结束一个循环
- `continue`跳过本次循环体中剩下尚未执行的语句，立即进行下一次的循环条件判定
  - 关键字 continue 只能被用于 for 循环中

# 标签与 goto

> 注：goto与标签并不推荐

某一行第一个以冒号（`:`）结尾的单词即为标签

- for、switch 或 select 语句都可以配合标签（label）形式的标识符使用

goto会跳转到标签位置

- 可以使用 goto 语句和标签配合使用来模拟循环

```go
package main

func main() {
	i:=0
	HERE:
		print(i)
		i++
		if i==5 {
			return
		}
		goto HERE
}
```

