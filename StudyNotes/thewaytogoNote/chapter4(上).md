# Go 程序的基本结构和要素

## 关键字与标识符

**关键字**

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

**预定义标识符**

| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

- Go 语言也是区分大小写的
- 有效的标识符必须以字符（可以使用任何 UTF-8 编码的字符或 _ ）开头，然后紧跟着 0 个或 多个字符或 Unicode 数字，如X56、group1、_x23、i、өԑ12
  - 以下是无效的标识符
    - 1ab（以数字开头） 
    - case（Go 语言的关键字） 
    - a+b（运算符是不允许的）
- `_`本身就是一个特殊的标识符，被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值（任何类型都 可以赋值给它），但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个 标识符作为变量对其它变量进行赋值或运算。

## Go 程序的基本结构和要素

如`hello_world.go`

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

- 包

  - 每个程序都由包（通常简称为 pkg）的概念组成，可以使用自身的包或者从其它包中导入内容。
  - 每个 Go 文件都属于且仅属于一个包，一个包可以由许多以 `.go` 为扩展名的源文件组成，因此文件名和包名一般来说都是不相同的。
  - 你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：`package main`。
    - `package main`表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 `main` 的包。
  - 所有的包名都应该使用小写字母

- 标准库

  - 在 Go 的安装文件里包含了一些可以直接使用的包，即标准库。
  - 也可以创建自己的包

- 编译顺序

  - 包的依赖关系决定了其构建顺序。

  - 属于同一个包的源文件必须全部被一起编译，一个包即是编译时的一个单元，因此根据惯例，每个目录都只包含一个包。

  - 如果对一个包进行更改或重新编译，所有引用了这个包的客户端程序都必须全部重新编译。

  - 每一段代码只会被编译一次。

  - Go 中的包模型采用了显式依赖关系的机制来达到快速编译的目的，编译器会从后缀名为 `.o` 的对象文件（需要且只需要这个文件）中提取传递依赖类型的信息。

    > 如果 `A.go` 依赖 `B.go`，而 `B.go` 又依赖 `C.go`：
    >
    > - 编译 `C.go`, `B.go`, 然后是 `A.go`.
    > - 为了编译 `A.go`, 编译器读取的是 `B.o` 而不是 `C.o`.

- 导入包

  - 通过`import`关键字

    ```go
    import "fmt"
    import "os"
    ```

    或

    ```go
    import(
    	"fmt"
    	"os"
    )
    ```

    > 推荐第二种

  - 可以对包进行重命名

    ```go
    import fm "fmt"
    ```

    这样做的好处就是之后调用这个包时可以用你自定义的名

    ```go
    package main
    
    import fm "fmt"
    
    func main() {
    	fm.Println("hello world")
    }
    ```

  - 如果你导入了一个包却没有使用它，则会在构建程序时引发错误，因为这遵循了 Go 的格言：“没有不必要的代码！“，即没用到的没必要出现在代码中。

- 可见性原则
  - 当标识符以大写字母开头，说明这种形式的标识符可以被外部包的代码所使用
    - 如Group1
    - 包括常量、变量、类型、函数名、结构字段等等
    - 类似与java中的public
  - 如果标识符以小写字母开头，则对外部包不可见
    - 类似与java中的private

## 函数

定义一个函数：

```go
func functionName(parameter_list) (return_value_list) {
   …
}
```

其中：

- parameter_list 的形式为 (param1 type1, param2 type2, …)，可以为0个或多个

- return_value_list 的形式为 (ret1 type1, ret2 type2, …)， 可以为0个或多个

- 左大括号 `{` 必须与方法的声明放在同一行，这是编译器的强制规定

- 右大括号 `}` 需要被放在紧接着函数体的下一行。如果你的函数非常简短，你也可以将它们放在同一行

  ```go
  func Sum(a, b int) int { return a + b }
  ```

所以函数最简形式如：

```go
func functionName()
```

特殊的：main 函数是每一个可执行程序所必须包含的，它**既没有参数，也没有返回类型**

关于函数的命名：只有当某个函数需要被外部包调用的时候才使用大写字母开头，并遵循 Pascal 命名法；否则就遵循骆驼命名法，即第一个单词的首字母小写，其余单词的首字母大写。

## 注释

**包注释：**

```
// Package superman implements methods for saving the world.
//
// Experience has shown that a small number of procedures can prove
// helpful when attempting to save the world.
package superman
```

**函数注释：**

- 出现在函数之前则要以函数名开头写注释

```
// enterOrbit causes Superman to fly into low Earth orbit, a position
// that presents several possibilities for planet salvation.
func enterOrbit() error {
   ...
}
```

## 类型

变量（或常量）包含数据，这些数据可以有不同的数据类型，简称类型。

- 使用 var 声明的变量的值会自动初始化为该类型的零值。

- 使用`type`可以自定义类型

  ```go
  type IZ int
  var a IZ = 5	
  
  // 如果你有多个类型需要定义，可以使用因式分解关键字的方式
  type (
     IZ int
     FZ float64
     STR string
  )
  ```

## Go 程序的一般结构

- 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。
- 如果存在 init 函数的话，则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）。
- 如果当前包是 main 包，则定义 main 函数。
- 然后定义其余的函数，首先是类型的方法，接着是按照 main 函数中先后调用的顺序来定义相关函数，如果有很多函数，则可以按照字母顺序来进行排序。

```go
package main

import (
   "fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
}

func main() {
   var a int
   Func1()
   // ...
   fmt.Println(a)
}

func (t T) Method1() {
   //...
}

func Func1() { // exported function Func1
   //...
}
```

## 类型转换

Go 语言不存在隐式类型转换，因此所有的转换都必须显式说明，即**类型 B 的值 = 类型 B(类型 A 的值)**

```go
a := 5.0
b := int(a)
```

- 只能在定义正确的情况下转换成功，例如从一个取值范围较小的类型转换到一个取值范围较大的类型（int16 -> int32）
- 当从一个取值范围较大的转换到取值范围较小的类型时会发生精度丢失（截断）的情况（int32 -> int16）
- 非法的类型转换时会引发编译时错误

具有相同底层类型的变量之间可以相互转换

```go
var a IZ = 5
c := int(a)
d := IZ(c)
```

# 常量

常量使用关键字 `const` 定义，用于存储不会改变的数据。**存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。**

常量的定义格式：`const identifier [type] = value`，例如：

```
const Pi = 3.14159
```

在 Go 语言中，你可以省略类型说明符 `[type]`，因为编译器可以根据变量的值来推断其类型。

- 显式类型定义： `const b string = "abc"`
- 隐式类型定义： `const b = "abc"`

未定义类型的常量会在必要时刻根据上下文来获得相关类型。

```
var n int
f(n + 5) // 无类型的数字型常量 “5” 它的类型在这里变成了 int
```

常量的值必须是能够在编译时就能够确定的

- 正确的做法：`const c1 = 2/3`
- 错误的做法：`const c2 = getNumber()` // 引发构建错误: `getNumber() used as value`

- **但内置函数可以使用，如：len()**

数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出

```go
const Ln2 = 0.693147180559945309417232121458\
			176568075500134360255254120680009
const Log2E = 1/Ln2 // this is a precise reciprocal
const Billion = 1e9 // float constant
const hardEight = (1 << 100) >> 97
```

根据上面的例子我们可以看到，反斜杠 `\` 可以在常量表达式中作为多行的连接符使用。

常量也允许使用并行赋值的形式：

```go
const beef, two, c = "eat", 2, "veg"
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)
```

常量还可以用作枚举：

```go
const (
	Unknown = 0
	Female = 1
	Male = 2
)
```

## iota

> 原书中iota讲得并不是很详细，这里做了扩充

- `iota`是go语言的常量计数器，只能在常量的表式中使用。主要应用场景是需要枚举或者自增的地方
- iota在const关键字出现的时候将被重置为0，const中每新增一行，iota计数一次，因此，我们可以将iota视为一个行索引，即代表const的第几行（第一行的索引0）。

举个例子：

```go
const (
	a = iota
	b = iota
	c = iota
)
```

第一行iota为0，第二行iota为1，第三行iota为2，即a = 0， b = 1， c = 2

当让，我们可以更加简化

```go
const(
	a = iota  	//0
	b			//1
	c			//2
	d			//3
)
```

每次出现const时，iota就会被初始化，即重置为0

```
const a = iota	//a = 0
const(
	b = iota	//b = 0
	c			//c = 1
)
```

那么，下面的例子，a，b，c各是多少：

```go
const (
	a = iota + 1
	b
	c
)
```

> 答案：
>
>    
>
> a = 1, b = 2, c = 3 
>
> 实际上，上面的例子详细的为
>
> ```go
> const (
> 	a = iota + 1
> 	b = iota + 1
> 	c = iota + 1
> )
> ```

也就是说，被我们省略的地方，实际上是与上一行相同的表达式。利用这个特性，我们就可以做出很多种枚举了：

```go
//奇数常量表
const (
	v1 = 2 * iota + 1		//1
	v2						//3
	v3						//5
	v4						//7
)

// 以5递增
const (
    val1 = 5 * (iota + 1)	//5
    val2					//10
    val3					//15
)
```

**可以在iota声明中键入自己想要的值**

```go
const (
	n1 = iota		//0
	n2 = 100		//100
	n3 = iota		//2
)
```

这里可以看出，iota是和const的行数相关的

**多个iota在同一行**

```go
const (
	x1, x2 = iota + 1, iota + 2
	x3, x4
	x5, x6
)
```

同一行中是可以出现多个iota的，但是实际上iota在该行中代表的值是不会变动的（记住iota只与const行数相关）

同时，我们还知道，被省略的表达式与上面出现的表达式相同

因此:

```
x1 = 1, x2 = 2, x3 = 2, x4 = 3, x5 = 3, x6 = 4
```

**可跳过的值**

golang中可以使用`_`跳过常量列表中的某个值

```go
const (
	c1 = iota + 1		//1
	_
	c3					//3
	c4					//4
)
```

**iota与<<**

既然`_`可以跳过值，通过位运算，我们就可以实现更多日常中会出现的枚举

```go
// 定义数量级
const (
    _ = iota
    KB = 1 << (10 * iota) 	// 1 << (10 * 1) 1024
    MB          			// 1 << (10 * 2) 1048576
    GB          			// 1 << (10 * 3) 1073741824
)
```

# 变量

变量声明 `var` 关键字：`var identifier type`。

示例：

```go
var a int
var b bool
var str string

//也可以改写成这种形式
var (
	a int
	b bool
	str string
)
```

- 当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。记住，**所有的内存在 Go 中都是经过初始化的。**

- 变量的命名规则遵循骆驼命名法，即首个单词小写，每个新单词的首字母大写，例如：`numShips` 和 `startDate`。
  - 全局变量能够被外部所使用，需要将首字母也大写。

Go 可以根据变量的值来自动推断其类型, 如

```go
var a = 15
var b = false
var str = "Go says hello to the world!"

// 或
var (
	a = 15
	b = false
	str = "Go says hello to the world!"
	numShips = 50
	city string
)
```

在**函数体内**声明局部变量，可以使用简短声明语法 `:=`，例如：

```go
a := 1
```

## 值类型与引用类型

-  int、float、bool 和 string 这些基本类型都属于值类型，像数组和结构这些复合类型也是值类型。
- 指针属于引用类型。其他引用类型还有

> - 值类型：所有像int、float、bool和string这些类型都属于值类型，使用这些类型的变量直接指向存在内存中的值，值类型的变量的值存储在栈中。当使用等号=将一个变量的值赋给另一个变量时，如 j = i ,实际上是在内存中将 i 的值进行了拷贝。可以通过 **&i** 获取变量 i 的内存地址
> - 引用类型：复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。一个引用类型的变量r1存储的是r1的值所在的内存地址（数字），或内存地址中第一个字所在的位置，这个内存地址被称之为指针，这个指针实际上也被存在另外的某一个字中。

## 打印

之前学过Printf，它可以在 fmt 包外部使用。该函数主要用于打印输出到控制台。

现在，我们学习了变量，除了直接输出变量，我们还可以以另一种方式将变量打印出来，即格式化

```go
"The operating system is: %s\n"
```

`%..`，其中 `..` 可以被不同类型所对应的标识符替换，如 `%s` 代表字符串标识符、`%v` 代表使用类型的默认输出格式的标识符。这些标识符所对应的值从格式化字符串后的第一个逗号开始按照相同顺序添加，如果参数超过 1 个则同样需要使用逗号分隔。使用这些占位符可以很好地控制格式化输出的文本。

## init函数

变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。

init函数不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。

- ~~每个源文件都只能包含一个init函数~~。(这里原书表述有误？实际上一个源文件可以包含多个init函数，下面会有例子示例)
  - 初始化总是以单线程执行，并且按照包的依赖关系顺序执行。
- 一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。

- init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine

  ```
  func init() {
     // setup preparations
     go backend()
  }
  ```

**总结与补充：**

init函数有下面的特性：

- init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等

- 每个包可以拥有多个init函数

- 包的每个源文件也可以拥有多个init函数

  ```go
  package main
  
  import "fmt"
  
  func init() {
     fmt.Println("init 1")
  }
  
  func init() {
     fmt.Println("init 2")
  }
  
  func main() {
     fmt.Println("main")
  }
  ```

  结果：

  ```
  init 1
  init 2
  main
  ```

- init函数没有输入参数、返回值

- 同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个执行顺序。

- 不同包的init函数按照包导入的依赖关系决定执行顺序。

  - 对同一个go文件的`init()`调用顺序是从上到下的。
  - 对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的`init()`函数。
  - 对于不同的`package`，如果不相互依赖的话，按照main包中"先`import`的后调用"的顺序调用其包中的`init()`，如果`package`存在依赖，则先调用最早被依赖的`package`中的`init()`，最后调用`main`函数。

- 如果`init`函数中使用了`println()`或者`print()`你会发现在执行过程中这两个不会按照你想象中的顺序执行。这两个函数官方只推荐在测试环境中使用，对于正式环境不要使用。

**init与main的区别**

- Go语言程序的默认入口函数(主函数)：func main()
- 相同：两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
- 不同：
  - init可以应用于任意包中，且可以重复定义多个。
  - main函数只能用于main包中，且只能定义一个。

# 4.4的课后作业

local_scope.go

```go
package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() { print(a) }

func m() {
   a := "O"
   print(a)
}
```

答案：GOG

global_scope.go:

```go
package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() {
   print(a)
}

func m() {
   a = "O"
   print(a)
}
```

答案：GOO

function_calls_function.go：

```
package main

var a string

func main() {
   a = "G"
   print(a)
   f1()
}

func f1() {
   a := "O"
   print(a)
   f2()
}

func f2() {
   print(a)
}
```

答案：GOG

解析：

- 例一中使用的是`:=`，这会创建一个局部变量a，不会影响到全局变量a的值
- 例二中使用的是`=`，这会将全局变量a的值替换为O
- 例三的道理同例一

> 在golang中“=”是赋值，“:=”是**声明变量**并赋值
>
> - =要和var 关键字一起使用。var可以在函数中使用，也可以在函数外使用。
> -  :=只能在函数中使用，所以只能定义局部变量。
> - 用var定义之后，用=赋值之后，还可以改变。但是用:=定义的变量不能改变值（也就是b:="q",在函数中不能再重新这样定义 b:="s",但是可以b="s"）