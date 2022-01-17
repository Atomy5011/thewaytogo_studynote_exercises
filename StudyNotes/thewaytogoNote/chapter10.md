# 结构体

结构体定义的一般方式如下：

```go
type identifier struct {
    field1 type1
    field2 type2
    ...
}
```

- `type T struct {a, b int}` 也是合法的语法，它更适用于简单的结构体。
- 如果定义的字段从来不会被用到，可以命名为`_`
- 结构体的字段可以是任何类型，甚至是结构体本身
- 可以使用点号符给字段赋值：`structname.fieldname = value`; 同样的，使用点号符可以获取结构体字段的值：`structname.fieldname`

## 递归结构体

**链表**

![10-1](./img/10-1.jpg)

```go
type Node struct {
    data    float64
    su      *Node
}
```

**双向链表**（一个前趋节点 `pr` 和一个后继节点 `su`）

```go
type Node struct {
    pr      *Node
    data    float64
    su      *Node
}
```

**二叉树**

![img](./img/10-2.jpg)

```go
type Tree struct {
    le      *Tree
    data    float64
    ri      *Tree
}
```

## 使用工厂模式创建结构体实例

Go 语言不支持面向对象编程语言中那样的构造子方法，但是可以很容易的在 Go 中实现 “构造子工厂”方法。为了方便通常会为类型定义一个工厂，按惯例，工厂的名字以 new 或 New 开头。假设定义了如下的 File 结构体类型：

```go
type File struct {
    fd      int     // 文件描述符
    name    string  // 文件名
}
```

下面是这个结构体类型对应的工厂方法，它返回一个指向结构体实例的指针：

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }

    return &File{fd, name}
}
```

然后这样调用它：

```go
f := NewFile(10, "./test.txt")
```

- 在 Go 语言中常常像上面这样在工厂方法里使用初始化来简便的实现构造函数

- 如果 `File` 是一个结构体类型，那么表达式 `new(File)` 和 `&File{}` 是等价的

- 如果想知道结构体类型 T 的一个实例占用了多少内存，可以使用：`size := unsafe.Sizeof(T{})`

- 使用工厂方法，从而使类型变成私有的

  ```go
  type matrix struct {
      ...
  }
  
  func NewMatrix(params) *matrix {
      m := new(matrix) // 初始化 m
      return m
  }
  ```

  在其他包里使用工厂方法：

  ```go
  package main
  import "matrix"
  ...
  wrong := new(matrix.matrix)     // 编译失败（matrix 是私有的）
  right := matrix.NewMatrix(...)  // 实例化 matrix 的唯一方式
  ```

## 带标签的结构体

```go
type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}
```

类型后的字符串就是标签，它是一个附属于字段的字符串，可以是文档或其他的重要标记。标签的内容不可以在一般的编程中使用，只有包 `reflect` 能获取它。

## 匿名字段和内嵌结构体

结构体可以包含一个或多个 **匿名（或内嵌）字段**，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即 **结构体可以包含内嵌结构体**。

```go
package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b    int
	c    float32
	int  // anonymous field
	innerS //anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)			// outer.b is: 6
	fmt.Printf("outer.c is: %f\n", outer.c)			// outer.c is: 7.500000
	fmt.Printf("outer.int is: %d\n", outer.int)		// outer.int is: 60
	fmt.Printf("outer.in1 is: %d\n", outer.in1)		// outer.in1 is: 5
	fmt.Printf("outer.in2 is: %d\n", outer.in2)		// outer.in2 is: 10

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)				// outer2 is:{6 7.5 60 {5 10}}
}
```

**命名冲突**

当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？

1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。

例子：

```
type A struct {a int}
type B struct {a, b int}

type C struct {A; B}
var c C
```

规则 2：使用 `c.a` 是错误的，到底是 `c.A.a` 还是 `c.B.a` 呢？会导致编译器错误：**ambiguous DOT reference c.a disambiguate with either c.A.a or c.B.a**。

```
type D struct {B; b float32}
var d D
```

规则1：使用 `d.b` 是没问题的：它是 float32，而不是 `B` 的 `b`。如果想要内层的 `b` 可以通过 `d.B.b` 得到。

# 方法

**方法是一种特殊类型的函数。**

定义方法的一般格式如下：

```go
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
```

不好意思:joy:，我没看懂这一节说啥。。。但是看例子还是大概了解了方法的作用，可以理解为receiver类似于`java`的`this`或`python`的self下面个人归纳一下从其他书本的学习记录：

--------

节选自《go语言学习笔记》第6章 方法

1. 方法是与对象实例绑定的特殊函数，用于维护和展示对象的自身状态
2. 方法和函数定义语法区别在于前者有前置实例接收参数（receiver），编译器以此明确方法所属类型。
3. 方法不支持重载
4. receiver参数名没有限制，按惯例会用简短有意义的名称，但不推荐使用this、self。如果方法内部没有引用实例，可以省略参数名，仅保留类型。
5. receiver的类型可以是基础类型或指针类型，关系到调用时对象实例是否被复制。
6. 不能用多级指针调用方法

示例：

```
type N int

// 对应第4点
func (N) test() {
	println("hi!")
}

// 对应第5点
func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("p: %p, %v\n", n, *n)
}

func main() {
	var a N = 25
	p := &a
	
	a.value()
	a.pointer()
	
	p.value()
	p.pointer()
	
	p2 := &p		// 多级指针
	
	// 对应第6点
	// p2.value()	// 错误
	// p2.pointer()	// 错误
}
```

如何选择方法的receiver类型（基础类型还是指针类型）， 建议：

- 要修改实例状态：*T
- 无须修改状态的小对象或固定值：T
- 大对象：*T， 以减少复制成本
- 引用对象、字符串、函数等指针包装对象，直接用T
- 若包含了`Mutex`等同步字段，用*T，避免因复制造成锁操作无效
- 其他无法确定的情况都用*T

------------------

节选自《go语言圣经》

- 在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。
- 在Go语言中，我们并不会像其它语言那样用this或者self作为接收器；我们可以任意的选择接收器的名字。

## 类型的 String() 方法和格式化描述符

```go
package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)		// two1 is: (12/10)
	fmt.Println("two1 is:", two1)			// two1 is: (12/10)
	fmt.Printf("two1 is: %T\n", two1)		// two1 is: *main.TwoInts0
	fmt.Printf("two1 is: %#v\n", two1)		// two1 is: &main.TwoInts{a:12, b:10}
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
```