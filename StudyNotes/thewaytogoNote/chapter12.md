# 读取用户的输入

## Scan 和 Sscan 

从键盘和标准输入 `os.Stdin` 读取输入，最简单的办法是使用 `fmt` 包提供的 Scan 和 Sscan 开头的函数。

```go
// 从控制台读取输入:
package main
import "fmt"

var (
   firstName, lastName, s string
   i int
   f float32
   input = "56.12 / 5212 / Go"
   format = "%f / %d / %s"
)

func main() {
   fmt.Println("Please enter your full name: ")
   fmt.Scanln(&firstName, &lastName)
   // fmt.Scanf("%s %s", &firstName, &lastName)
   fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
   fmt.Sscanf(input, format, &f, &i, &s)
   fmt.Println("From the string we read: ", f, i, s)
    // 输出结果: From the string we read: 56.12 5212 Go
}
```

- `Scanln` 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
- `Scanf` 与其类似，除了 `Scanf` 的第一个参数用作格式字符串，用来决定如何读取。
- `Sscan` 和以 `Sscan` 开头的函数则是从字符串读取，除此之外，与 `Scanf` 相同。

也可以使用 `bufio` 包提供的缓冲读取（buffered reader）来读取数据

```go
package main
import (
    "fmt"
    "bufio"
    "os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
    // inputReader := bufio.NewReader(os.Stdin)
    // input, err := inputReader.ReadString('\n')
    inputReader = bufio.NewReader(os.Stdin)
    fmt.Println("Please enter some input: ")
    input, err = inputReader.ReadString('\n')
    if err == nil {
        fmt.Printf("The input was: %s\n", input)
    }
}
```

- `inputReader` 是一个指向 `bufio.Reader` 的指针。`inputReader := bufio.NewReader(os.Stdin)` 这行代码，将会创建一个读取器，并将其与标准输入绑定。
- `bufio.NewReader()` 构造函数的签名为：`func NewReader(rd io.Reader) *Reader`

## 使用 `switch` 语句

```go
package main
import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("Please enter your name:")
    input, err := inputReader.ReadString('\n')

    if err != nil {
        fmt.Println("There were errors reading, exiting program.")
        return
    }

    fmt.Printf("Your name is %s", input)
    // For Unix: test with delimiter "\n", for Windows: test with "\r\n"
    switch input {
    case "Philip\r\n":  fmt.Println("Welcome Philip!")
    case "Chris\r\n":   fmt.Println("Welcome Chris!")
    case "Ivo\r\n":     fmt.Println("Welcome Ivo!")
    default: fmt.Printf("You are not welcome here! Goodbye!")
    }

    // version 2:   
    switch input {
    case "Philip\r\n":  fallthrough
    case "Ivo\r\n":     fallthrough
    case "Chris\r\n":   fmt.Printf("Welcome %s\n", input)
    default: fmt.Printf("You are not welcome here! Goodbye!\n")
    }

    // version 3:
    switch input {
    case "Philip\r\n", "Ivo\r\n":   fmt.Printf("Welcome %s\n", input)
    default: fmt.Printf("You are not welcome here! Goodbye!\n")
    }
}
```

# 文件读写

## 读文件

- 文件使用指向 `os.File` 类型的指针来表示的，也叫做文件句柄
  - 在前面章节使用到过标准输入 `os.Stdin` 和标准输出 `os.Stdout`，他们的类型都是 `*os.File`

示例：

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
    inputFile, inputError := os.Open("input.dat")
    if inputError != nil {
        fmt.Printf("An error occurred on opening the inputfile\n" +
            "Does the file exist?\n" +
            "Have you got access to it?\n")
        return // exit the function on error
    }
    defer inputFile.Close()

    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerError := inputReader.ReadString('\n')
        fmt.Printf("The input was: %s", inputString)
        if readerError == io.EOF {
            return
        }      
    }
}
```

- 变量 `inputFile` 是 `*os.File` 类型的
- 使用 `os` 包里的 `Open` 函数来打开一个文件
- 使用 `defer inputFile.Close()` 语句确保在程序退出前关闭该文件
- 使用 `bufio.NewReader` 来获得一个读取器变量
- 使用 `ReadString('\n')` 或 `ReadBytes('\n')` 将文件的内容逐行（行结束符 '\n'）读取出来
  - 注意：**Unix 和 Linux 的行结束符是 \n，而 Windows 的行结束符是 \r\n**
  - 在使用 `ReadString` 和 `ReadBytes` 方法的时候，我们不需要关心操作系统的类型，直接使用 \n 就可以了。
  - 也可以使用 `ReadLine()` 方法来实现相同的功能
- 一旦读取到文件末尾，变量 `readerError` 的值将变成非空（事实上，其值为常量 `io.EOF`）

### 其他读文件函数

1. **将整个文件的内容读到一个字符串里**

可以使用 `io/ioutil` 包里的 `ioutil.ReadFile()` 方法，该方法第一个返回值的类型是 `[]byte`，第二个返回值是错误

```go
func main() {
    inputFile := "products.txt"
    outputFile := "products_copy.txt"
    buf, err := ioutil.ReadFile(inputFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
        // panic(err.Error())
    }
    fmt.Printf("%s\n", string(buf))
    err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
    if err != nil {
        panic(err.Error())
    }
}
```

2. **带缓冲的读取**

在很多情况下，文件的内容是不按行划分的，或者干脆就是一个二进制文件。使用 `bufio.Reader` 的 `Read()`，它只接收一个参数：

```go
buf := make([]byte, 1024)
...
n, err := inputReader.Read(buf)
if (n == 0) { break}
```

变量 n 的值表示读取到的字节数.

3. **按列读取文件中的数据**

如果数据是按列排列并用空格分隔的，你可以使用 `fmt` 包提供的以 FScan 开头的一系列函数来读取他们

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("products2.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var col1, col2, col3 []string
    for {
        var v1, v2, v3 string
        _, err := fmt.Fscanln(file, &v1, &v2, &v3)
        // scans until newline
        if err != nil {
            break
        }
        col1 = append(col1, v1)
        col2 = append(col2, v2)
        col3 = append(col3, v3)
    }

    fmt.Println(col1)
    fmt.Println(col2)
    fmt.Println(col3)
}
```

输出结果：

```go
[ABC FUNC GO]
[40 56 45]
[150 280 356]
```

4. **关于文件路径**

`path` 包里包含一个子包叫 `filepath`，这个子包提供了跨平台的函数，用于处理文件名和路径。例如 Base() 函数用于获得路径中的最后一个元素（不包含后面的分隔符）：

```
import "path/filepath"
filename := filepath.Base(path)
```

### `compress`包：读取压缩文件

`compress`包提供了读取压缩文件的功能，支持的压缩文件格式为：bzip2、flate、gzip、lzw 和 zlib。

示例：读取一个 gzip 文件

```go
package main

import (
	"fmt"
	"bufio"
	"os"
	"compress/gzip"
)

func main() {
	fName := "MyFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	defer fi.Close()
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
```

## 写文件

示例：

```go
package main

import (
	"os"
	"bufio"
	"fmt"
)

func main () {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return  
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i:=0; i<10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
```

- `OpenFile` 函数有三个参数：文件名、一个或多个标志（使用逻辑运算符“|”连接），使用的文件权限

  - `os.O_RDONLY`：只读
  - `os.O_WRONLY`：只写
  - `os.O_CREATE`：创建：如果指定文件不存在，就创建该文件。
  - `os.O_TRUNC`：截断：如果指定文件已存在，就将该文件的长度截为 0 。

  - 在读文件的时候，文件的权限是被忽略的，所以在使用 `OpenFile` 时传入的第三个参数可以用 0 。而在写文件时，不管是 Unix 还是 Windows，都需要使用 0666。

- 创建一个写入器`outputWriter := bufio.NewWriter(outputFile)`，将字符串写入缓冲区`outputWriter.WriteString(outputString)`
- 缓冲区的内容紧接着被完全写入文件：`outputWriter.Flush()`

如果写入的东西很简单，我们可以使用 `fmt.Fprintf(outputFile, "Some test data.\n")` 直接将内容写入文件。`fmt` 包里的 F 开头的 Print 函数可以直接写入任何 `io.Writer`，包括文件

使用 `os.Stdout.WriteString("hello, world\n")`，我们可以输出到屏幕。

# 文件拷贝

```go
// filecopy.go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("target.txt", "source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
```

**注意 `defer` 的使用：当打开 dst 文件时发生了错误，那么 `defer` 仍然能够确保 `src.Close()` 执行。如果不这么做，src 文件会一直保持打开状态并占用资源。**

# 从命令行读取参数

## 通过os包

os 包中有一个 string 类型的切片变量 `os.Args`，用来处理一些基本的命令行参数，它在程序启动后读取命令行输入的参数。

- 类似于python的sys.args

示例：

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
```

可以在命令行加入参数，像这样：`os_args John Bill Marc Luke`，将得到这样的输出：`Good Morning Alice John Bill Marc Luke`

## 通过flag包

在 flag 包中有一个 Flag 被定义成一个含有如下字段的结构体：

```go
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}
```

```go
package main

import (
	"flag" // command line option parser
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
```

- `flag.Parse()` 扫描参数列表（或者常量列表）并设置 flag, `flag.Arg(i)` 表示第 i 个参数。`Parse()` 之后 `flag.Arg(i)` 全部可用，`flag.Arg(0)` 就是第一个真实的 flag，而不是像 `os.Args(0)` 放置程序的名字
- `flag.Narg()` 返回参数的数量。
- `flag.VisitAll(fn func(*Flag))` 是另一个有用的功能：按照字典顺序遍历 flag，并且对每个标签调用 fn 

# 用 buffer 读取文件

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}
```

# 用切片读写文件

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

func main() {
	flag.Parse() // Scans the arg list and sets up flags
	if flag.NArg() == 0 {
		cat(os.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if f == nil {
			fmt.Fprintf(os.Stderr, "cat: can't open %s: error %s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		cat(f)
		f.Close()
	}
}
```

# 用 defer 关闭文件

`defer` 关键字对于在函数结束时关闭打开的文件非常有用，例如下面的代码片段：

```go
func data(name string) string {
	f, _ := os.OpenFile(name, os.O_RDONLY, 0)
	defer f.Close() // idiomatic Go code!
	contents, _ := ioutil.ReadAll(f)
	return string(contents)
}
```

在函数 return 后执行了 `f.Close()`

# golang中对不同数据格式的处理

- 数据结构 --> 指定格式 = `序列化` 或 `编码`（传输之前）
- 指定格式 --> 数据结构 = `反序列化` 或 `解码`（传输之后）

序列化是在内存中把数据转换成指定格式（data -> string），反之亦然（string -> data）。

编码也是一样的，只是输出一个数据流（实现了 io.Writer 接口）；解码是从一个数据流（实现了 io.Reader）输出到一个数据结构。

## JSON

Go 语言的 json 包可以让你在程序中方便的读取和写入 JSON 数据。）`json.Marshal`

```go
// json.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
```

- 出于安全考虑，在 web 应用中最好使用 `json.MarshalforHTML()` 函数，其对数据执行 HTML 转码，所以文本可以被安全地嵌在 HTML `<script>` 标签中。

- JSON 与 Go 类型对应如下：
  - bool 对应 JSON 的 boolean
  - float64 对应 JSON 的 number
  - string 对应 JSON 的 string
  - nil 对应 JSON 的 null=

- 反序列化`json.Unmarshal()`
- json 包使用 `map[string]interface{}` 和 `[]interface{}` 储存任意的 JSON 对象和数组；其可以被反序列化为任何的 JSON blob 存储到接口值中

### 解码任意的数据：

json 包使用 `map[string]interface{}` 和 `[]interface{}` 储存任意的 JSON 对象和数组；其可以被反序列化为任何的 JSON blob 存储到接口值中。

来看这个 JSON 数据，被存储在变量 b 中：

```go
b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
```

不用理解这个数据的结构，我们可以直接使用 Unmarshal 把这个数据编码并保存在接口值中：

```go
var f interface{}
err := json.Unmarshal(b, &f)
```

f 指向的值是一个 map，key 是一个字符串，value 是自身存储作为空接口类型的值：

```go
map[string]interface{} {
	"Name": "Wednesday",
	"Age":  6,
	"Parents": []interface{} {
		"Gomez",
		"Morticia",
	},
}
```

要访问这个数据，我们可以使用类型断言

```go
m := f.(map[string]interface{})
```

我们可以通过 for range 语法和 type switch 来访问其实际类型：

```go
for k, v := range m {
	switch vv := v.(type) {
	case string:
		fmt.Println(k, "is string", vv)
	case int:
		fmt.Println(k, "is int", vv)

	case []interface{}:
		fmt.Println(k, "is an array:")
		for i, u := range vv {
			fmt.Println(i, u)
		}
	default:
		fmt.Println(k, "is of a type I don’t know how to handle")
	}
}
```

通过这种方式，你可以处理未知的 JSON 数据，同时可以确保类型安全。

### 编码和解码流

json 包提供 Decoder 和 Encoder 类型来支持常用 JSON 数据流读写。NewDecoder 和 NewEncoder 函数分别封装了 io.Reader 和 io.Writer 接口。

```
func NewDecoder(r io.Reader) *Decoder
func NewEncoder(w io.Writer) *Encoder
```

要想把 JSON 直接写入文件，可以使用 json.NewEncoder 初始化文件（或者任何实现 io.Writer 的类型），并调用 Encode()；反过来与其对应的是使用 json.NewDecoder 和 Decode() 函数：

```
func NewDecoder(r io.Reader) *Decoder
func (dec *Decoder) Decode(v interface{}) error
```

## XML

- 如同 json 包一样，也有 `xml.Marshal()` 和 `xml.Unmarshal()` 从 XML 中编码和解码数据
- encoding/xml 包实现了一个简单的 XML 解析器（SAX），用来解析 XML 数据内容

示例：

```go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func main()  {
	//读取xml文件到缓存中，返回[]byte
	XmlParam, err := ioutil.ReadFile("test.xml")
	if err != nil {
		fmt.Println("xml文件读取失败！")
		return
	}
	fmt.Println(string(XmlParam))
	//预期：<message><name>chen</name><age>26</age><sex>男</sex></message>


	/*
	  如果不想使用xml文件获取参数则模拟http发送的xml参数,如下:
	  var XmlParam=`<message><name>chen</name><age>26</age><sex>男</sex></message>`
	*/
	

    //将xmL映射到结构体中
	var Per Person
	err1:=xml.Unmarshal(XmlParam,&Per)
	if err1!=nil{
		fmt.Println("Unmarshal error")
		return
	}
	//获取其中一个标签的值
	fmt.Println(Per.Name) //预期：chen
}

//反序列化结构体定义
type Person struct {
	Name  string `xml:"name"`  //注意这里有个反引号
	Age   string `xml:"age"`
	Sex   string `xml:"sex"`
}
/*
  结构体中的反引号说明:
     `xml:"name"` 表示:将xml数据反序列化时,xml标签的数据name对应到结构体的Name成员中,需要
  使用xml中的name标签值时,从Person中的Name成员中获取即可.
*/
```

## GOB

- Gob 是 Go 自己的以二进制形式序列化和反序列化程序数据的格式
  - 类似于 Python 的 "pickle" 和 Java 的 "Serialization"
- 两个用 Go 写的服务之间的通信。这样的话服务可以被实现得更加高效和优化。