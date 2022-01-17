# 函数

```go
func g() {
}
```

**函数也可以以申明的方式被使用，作为一个函数类型**，就像：

```
type binOp func(int, int) int
```

- 除了main()、init()函数外，其它所有类型的函数都可以有参数与返回值

- 鉴于可读性的需求，最好把 `main()` 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）
- 编写多个函数的主要目的是将一个需要很多行代码的复杂问题分解为一系列简单的任务（那就是函数）来解决

- Go 里面有三种类型的函数：
  - 普通的带有名字的函数
  - 匿名函数或者lambda函数
  - 方法（Methods，在第十章）

# 值传递和引用传递





# 传递变长参数

采用 `...type` 的形式

```
func myFunc(a, b, arg ...int) {}
```

- 如果参数被存储在一个 slice 类型的变量 `slice` 中，则可以通过 `slice...` 的形式来传递参数，调用变参函数。

  - ```
    slice := []int{7,9,3,5,1}
    x = min(slice...)
    ```

- 如果变长参数的类型并不是都相同
  - 定义一个结构类型，用以存储所有可能的参数
  - 使用空接口

# defer 和追踪

defer 允许我们推迟到函数返回之前（或任意位置执行 `return` 语句之后）一刻才执行某个语句或函数

关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 `finally` 语句块，它一般用于释放某些已分配的资源。

## defer的例子

1. 使用 defer 的语句同样可以接受参数，**下面这个例子就会在执行 defer 语句时打印 `0`**：

```
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
```

2. 当有多个 defer 行为被注册时，它们会**以逆序执行**（类似栈，即后进先出）：

```
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
// 结果： 4, 3, 2, 1, 0
```

3. 关键字 defer 允许我们进行一些函数执行完成后的收尾工作

   - 关闭文件流 : `defer file.Close()`

   - 解锁一个加锁的资源:

     ```
     mu.Lock()  
     defer mu.Unlock() 
     ```

   - 打印最终报告:

     ```
     printHeader()  
     defer printFooter()
     ```

   - 关闭数据库链接 : `defer disconnectFromDB()`

# 内置函数

Go 语言拥有一些不需要进行导入操作就可以使用的内置函数。

| 名称               | 说明                                                         |
| ------------------ | ------------------------------------------------------------ |
| close              | 用于管道通信                                                 |
| len、cap           | len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道，不能用于 map） |
| new、make          | new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：`v := new(int)`。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）**new() 是一个函数，不要忘记它的括号** |
| copy、append       | 用于复制和连接切片                                           |
| panic、recover     | 两者均用于错误处理机制                                       |
| print、println     | 底层打印函数，在部署环境中建议使用 fmt 包                    |
| complex、real imag | 用于创建和操作复数                                           |

# 递归函数

当一个函数在其函数体内调用自身，则称之为递归。

# 将函数作为参数

函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调。

将函数作为参数的最好的例子是函数 `strings.IndexFunc()`：该函数的签名是 `func IndexFunc(s string, f func(c rune) bool) int`，它的返回值是字符串s中第一个使函数`f(c)`返回`true`的Unicode字符的索引值。如果找不到，则返回-1。

# 闭包

当我们不希望给函数起名字的时候，可以使用匿名函数，例如：`func(x, y int) int { return x + y }`。这样的一个函数不能够独立存在，但可以被赋值于某个变量，即保存函数的地址到变量中，然后通过变量名对函数进行调用.

```go
package main

import "fmt"

func main() {
	f()
}
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
```

输出：

```
0 - g is of type func(int) and has value 0x681a80
1 - g is of type func(int) and has value 0x681b00
2 - g is of type func(int) and has value 0x681ac0
3 - g is of type func(int) and has value 0x681400
```

我们可以看到变量 `g` 代表的是 `func(int)`，变量的值是一个内存地址。

所以我们实际上拥有的是一个函数值：**匿名函数可以被赋值给变量并作为值使用**。

**关键字 `defer`经常配合匿名函数使用，它可以用于改变函数的命名返回值。**

#  计算函数执行时间

可以使用 `time` 包中的 `Now()` 和 `Sub` 函数

```go
start := time.Now()
longCalculation()
end := time.Now()
delta := end.Sub(start)
fmt.Printf("longCalculation took this amount of time: %s\n", delta)
```

# 通过内存缓存来提升性能

内存缓存的技术在使用计算成本相对昂贵的函数时非常有用（不仅限于例子中的递归），譬如大量进行相同参数的运算。这种技术还可以应用于纯函数中，即相同输入必定获得相同输出的函数。

```go
package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
```

