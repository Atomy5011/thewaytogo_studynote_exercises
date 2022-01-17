# 接口

接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

接口是一种契约，实现类型必须满足它，它描述了类型的行为，规定类型可以做什么。接口彻底将类型能做什么，以及如何做分离开来，使得相同接口的变量在不同的时刻表现出不同的行为，这就是多态的本质。

通过如下格式定义接口：

```go
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
```

上面的 `Namer` 是一个 **接口类型**。

- （按照约定，只包含一个方法的）接口的名字由方法名加 `er` 后缀组成，例如 `Printer`、`Reader`、`Writer`、`Logger`、`Converter` 等等。
- 还有一些不常用的方式（当后缀 `er` 不合适时），比如 `Recoverable`，此时接口名以 `able` 结尾，或者以 `I` 开头

示例：

在调用 `shapes[n].Area() `这个时，只知道 `shapes[n]` 是一个 `Shaper` 对象，最后它摇身一变成为了一个 `Square` 或 `Rectangle` 对象，并且表现出了相对应的行为。

```go
package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
```

```
Looping through shapes for area ...
Shape details:  {5 3}
Area of this shape is:  15
Shape details:  &{5}
Area of this shape is:  25
```

## 类型断言：如何检测和转换接口变量的类型

我们可以使用 **类型断言** 来测试在某个时刻 `varI` 是否包含类型 `T` 的值：

```go
if v, ok := varI.(T); ok {  // checked type assertion
    Process(v)
    return
}
// varI is not of type T
```

特例：假定 `v` 是一个值，然后我们想测试它是否实现了 `Stringer` 接口

```go
type Stringer interface {
    String() string
}

if sv, ok := v.(Stringer); ok {
    fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
}
```

## 类型判断：type-switch

```go
switch t := areaIntf.(type) {
case *Square:
	fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
	fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
	fmt.Printf("nil value: nothing to check?\n")
default:
	fmt.Printf("Unexpected type %T\n", t)
}
```

- 变量 `t` 得到了 `areaIntf` 的值和类型，所有 `case` 语句中列举的类型（`nil` 除外）都必须实现对应的接口
- 可以用 `type-switch` 进行运行时类型分析，但是在 `type-switch` 不允许有 `fallthrough` 。

如果仅仅是测试变量的类型，不用它的值，那么就可以不需要赋值语句

```go
switch areaIntf.(type) {
case *Square:
	// TODO
case *Circle:
	// TODO
...
default:
	// TODO
}
```

示例：一个类型分类函数，它有一个可变长度参数，可以是任意类型的数组，它会根据数组元素的实际类型执行不同的动作

```go
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
```

可以这样调用此方法：`classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)` 。

## 使用方法集与接口

**总结**

在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以根据具体类型 `P` 直接辨识的：

- 指针方法可以通过指针调用
- 值方法可以通过值调用
- 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
- 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。

Go 语言规范定义了接口方法集的调用规则：

- 类型 T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
- 类型 *T 的可调用方法集包含接受者为 *T 的所有方法
- 类型 *T 的可调用方法集不包含接受者为 T 的方法

示例：

```go
package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	// CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID: Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) // VALID: Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
```

在 `lst` 上调用 `CountInto` 时会导致一个编译器错误，因为 `CountInto` 需要一个 `Appender`，而它的方法 `Append` 只定义在指针上。 在 `lst` 上调用 `LongEnough` 是可以的，因为 `Len` 定义在值上。

在 `plst` 上调用 `CountInto` 是可以的，因为 `CountInto` 需要一个 `Appender`，并且它的方法 `Append` 定义在指针上。 在 `plst` 上调用 `LongEnough` 也是可以的，因为指针会被自动解引用。

## 空接口

```go
type Any interface{}
```

- `any` 或 `Any` 是空接口一个很好的别名或缩写
- 类似 Java 中的 Object 
- 作用：Go语言暂时没有泛型，如果一个函数需要接收任意类型的参数，就可以使用空接口

> 示例：
>
> 1. 给空接口定一个别名类型 `Element`：`type Element interface{}`
>
> 2. 定义一个容器类型的结构体 `Vector`，它包含一个 `Element` 类型元素的切片：
>
>    - `Vector` 里能放任何类型的变量，因为任何类型都实现了空接口，实际上 `Vector` 里放的每个元素可以是不同类型的变量。
>
>    ```go
>    type Vector struct {
>    	a []Element
>    }
>    ```
>
> 3. 定义一个 `At()` 方法用于返回第 `i` 个元素：
>
>    ```go
>    func (p *Vector) At(i int) Element {
>    	return p.a[i]
>    }
>    ```
>
>    再定一个 `Set()` 方法用于设置第 `i` 个元素的值：
>
>    ```go
>    func (p *Vector) Set(i int, e Element) {
>    	p.a[i] = e
>    }
>    ```

- 空接口是反射实现基础，反射库就是将相关具体的类型转换并赋值给空接口后才去处理

- 空接口不一定是nil

  ```go
  type Intf interface {
  	Say()
  }
  
  type St struct {}
  
  func (St) Say{
  	fmt.Println("Hi!")
  }
  
  var st *St = nil
  var intf Intf = st
  
  if intf != nil {
  	intf.Say()			// Hi!
  }
  ```

  - 因为空接口有两个字段，一个是实例类型，另一个是指向绑定实例的指针，两者皆为nil，空接口才为nil

# 反射

反射可以在运行时检查类型和变量，例如它的大小、方法和 `动态` 的调用这些方法。例如两个简单的函数，`reflect.TypeOf` 和 `reflect.ValueOf`，返回被检查对象的类型和值。

> 这对于没有源代码的包尤其有用。这是一个强大的工具，除非真得有必要，否则应当避免使用或小心使用。

备注：个人觉得本书的反射降低不是很好，以下笔记来自《go语言圣经》的笔记

-----------

## 何为反射

- 能够在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作，而不需要在编译时就知道这些变量的具体类型。这种机制被称为反射。

- 为何需要反射
  - 有时候我们需要编写一个函数能够处理一类并不满足普通公共接口的类型的值，也可能是因为它们并没有确定的表示方式，或者是在我们设计该函数的时候这些类型可能还不存在。
  - 比如用switch类型分支来测试输入参数是否实现了String方法，我们可以处理有限个类型，case string、case int、case bool等，但是这些组合类型的数目基本是无穷的，如果是`[]float64`、`map[string][]string`等类型呢？我们没有办法来检查未知类型的表示方式。但是反射可以帮助我们检查这类类型，这就是我们为何需要反射的原因。

## 最常用的`reflect.Type` 和 `reflect.Value`

- 函数 `reflect.TypeOf` 接受任意的 interface{} 类型，并以 `reflect.Type` 形式返回其动态类型

  ```Go
  t := reflect.TypeOf(3)  // a reflect.Type
  fmt.Println(t.String()) // "int"
  fmt.Println(t)          // "int"
  ```

- `reflect.TypeOf` 返回的是一个动态类型的接口值，它总是返回具体的类型(因此打印 "*os.File" 而不是 "io.Writer")

  ```Go
  var w io.Writer = os.Stdout
  fmt.Println(reflect.TypeOf(w)) // "*os.File"
  ```

- reflect.Type 接口是满足 fmt.Stringer 接口， fmt.Printf 提供了一个缩写 %T 参数，内部使用 reflect.TypeOf 来输出

  ```Go
  fmt.Printf("%T\n", 3) // "int"
  ```

- reflect 包中另一个重要的类型是 Value。一个 reflect.Value 可以装载任意类型的值。函数 reflect.ValueOf 接受任意的 interface{} 类型，并返回一个装载着其动态值的 reflect.Value。和 reflect.TypeOf 类似，reflect.ValueOf 返回的结果也是具体的类型，但是 reflect.Value 也可以持有一个接口值。

  ```Go
  v := reflect.ValueOf(3) // a reflect.Value
  fmt.Println(v)          // "3"
  fmt.Printf("%v\n", v)   // "3"
  fmt.Println(v.String()) // NOTE: "<int Value>"
  ```

  - 和 reflect.Type 类似，reflect.Value 也满足 fmt.Stringer 接口，但是除非 Value 持有的是字符串，否则 String 方法只返回其类型。

- 对 Value 调用 Type 方法将返回具体类型所对应的 reflect.Type

  ```Go
  t := v.Type()           // a reflect.Type
  fmt.Println(t.String()) // "int"
  ```

- reflect.ValueOf 的逆操作是 reflect.Value.Interface 方法。它返回一个 interface{} 类型，装载着与 reflect.Value 相同的具体值：

  ```Go
  v := reflect.ValueOf(3) // a reflect.Value
  x := v.Interface()      // an interface{}
  i := x.(int)            // an int
  fmt.Printf("%d\n", i)   // "3"
  ```

- 我们使用 reflect.Value 的 Kind 方法来替代之前的类型 switch。虽然还是有无穷多的类型，但是它们的 kinds 类型却是有限的

  - Bool、String 和 所有数字类型的基础类型；Array 和 Struct 对应的聚合类型；Chan、Func、Ptr、Slice 和 Map 对应的引用类型；interface 类型；还有表示空值的 Invalid 类型。

    ```Go
    func formatAtom(v reflect.Value) string {
        switch v.Kind() {
        case reflect.Invalid:
            return "invalid"
        case reflect.Int, reflect.Int8, reflect.Int16,
            reflect.Int32, reflect.Int64:
            return strconv.FormatInt(v.Int(), 10)
        case reflect.Uint, reflect.Uint8, reflect.Uint16,
            reflect.Uint32, reflect.Uint64, reflect.Uintptr:
            return strconv.FormatUint(v.Uint(), 10)
        // ...floating-point and complex cases omitted for brevity...
        case reflect.Bool:
            return strconv.FormatBool(v.Bool())
        case reflect.String:
            return strconv.Quote(v.String())
        case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
            return v.Type().String() + " 0x" +
                strconv.FormatUint(uint64(v.Pointer()), 16)
        default: // reflect.Array, reflect.Struct, reflect.Interface
            return v.Type().String() + " value"
        }
    }
    ```

- 有一些reflect.Values是可取地址的；其它一些则不可以。考虑以下的声明语句：

  ```Go
  x := 2                   // value   type    variable?	 备注
  a := reflect.ValueOf(2)  // 2       int     no			其中a对应的变量不可取地址。因为a中的值仅仅是整数2的拷贝副本。
  b := reflect.ValueOf(x)  // 2       int     no			b中的值也同样不可取地址
  c := reflect.ValueOf(&x) // &x      *int    no			c中的值还是不可取地址，它只是一个指针&x的拷贝
  d := c.Elem()            // 2       int     yes (x)		对于d，它是c的解引用方式生成的，指向另一个变量，因此是可取地址的。
  ```

  - 实际上，所有通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的。

  - 我们可以通过调用reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的Value。

  - 我们可以通过调用reflect.Value的CanAddr方法来判断其是否可以被取地址：

    ```Go
    fmt.Println(a.CanAddr()) // "false"
    fmt.Println(b.CanAddr()) // "false"
    fmt.Println(c.CanAddr()) // "false"
    fmt.Println(d.CanAddr()) // "true"
    ```

  - 通过调用可取地址的reflect.Value的reflect.Value.Set方法来更新对应的值：`d.Set(reflect.ValueOf(4))`

    - Set方法将在运行时执行和编译时进行类似的可赋值性约束的检查。以上代码，变量和值都是int类型，但是如果变量是int64类型，那么程序将抛出一个panic异常，所以关键问题是要确保改类型的变量可以接受对应的值：`d.Set(reflect.ValueOf(int64(5))) // panic: int64 is not assignable to int`
    - 同样，对一个不可取地址的reflect.Value调用Set方法也会导致panic异常
    - 有很多用于基本数据类型的Set方法：SetInt、SetUint、SetString和SetFloat等
      - 从某种程度上说，这些Set方法总是尽可能地完成任务。以SetInt为例，只要变量是某种类型的有符号整数就可以工作，即使是一些命名的类型、甚至只要底层数据类型是有符号整数就可以，而且如果对于变量类型值太大的话会被自动截断。但需要谨慎的是：对于一个引用interface{}类型的reflect.Value调用SetInt会导致panic异常，即使那个interface{}变量对于整数类型也不行。

## 获取结构体的字段标签

```go
package main

import (
    "fmt"
    "reflect"
)

type resume struct {
    Name string `json:"name" doc:"我的名字"`
}

func findDoc(stru interface{}) map[string]string {
    t := reflect.TypeOf(stru).Elem()
    doc := make(map[string]string)

    for i := 0; i < t.NumField(); i++ {
        doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
    }

    return doc

}

func main() {
    var stru resume
    doc := findDoc(&stru)
    fmt.Printf("name字段为：%s\n", doc["name"])
}
```

或

```go
package main
 
import (
    "encoding/json"
    "fmt"
    "reflect"
)
 
func main() {
    type User struct {
        UserId   int    `json:"user_id" bson:"user_id"`
        UserName string `json:"user_name" bson:"user_name"`
    }
    // 输出json格式
    u := &User{UserId: 1, UserName: "tony"}
    j, _ := json.Marshal(u)
    fmt.Println(string(j))
    // 输出内容：{"user_id":1,"user_name":"tony"}
 
    // 获取tag中的内容
    t := reflect.TypeOf(u)
    field := t.Elem().Field(0)
    fmt.Println(field.Tag.Get("json"))
    // 输出：user_id
    fmt.Println(field.Tag.Get("bson"))
    // 输出：user_id
}
```
