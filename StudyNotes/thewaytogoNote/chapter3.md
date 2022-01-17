# 构建并运行Go程序

- `gofmt`
  - 每次构建程序之前都会自动调用源码格式化工具`gofmt`并将格式化后的源码保存。
  - 如果构建成功则不会出现任何信息，构建失败会出现对应的错误，如 `a declared and not used `
- `go build`
  - 编译并安装自己身包和依赖包
- `go install`
  - 安装自身包和依赖包
- `go doc`
  - 生成代码文档，会从 Go 程序和包文件中提取顶级声明的首行注释以及每个对象的相关注释，并生成相关文档。
    - `go doc package`: 获取包的文档注释,
    - `go doc package/sibpackage`: 获取子包的文档注释，
    - `go doc package function`: 获取某个函数在某个包中的文档注释

**总结与补充：**

## go命令

```go
$ go
Go is a tool for managing Go source code.

Usage:

    go command [arguments]

The commands are:

    build       compile packages and dependencies
    clean       remove object files
    doc         show documentation for package or symbol
    env         print Go environment information
    bug         start a bug report
    fix         run go tool fix on packages
    fmt         run gofmt on package sources
    generate    generate Go files by processing source
    get         download and install packages and dependencies
    install     compile and install packages and dependencies
    list        list packages
    run         compile and run Go program
    test        test packages
    tool        run specified go tool
    version     print Go version
    vet         run go tool vet on packages

Use "go help [command]" for more information about a command.

Additional help topics:

    c           calling between Go and C
    buildmode   description of build modes
    filetype    file types
    gopath      GOPATH environment variable
    environment environment variables
    importpath  import path syntax
    packages    description of package lists
    testflag    description of testing flags
    testfunc    description of testing functions

Use "go help [topic]" for more information about that topic.
```

- go env用于打印Go语言的环境信息。
- go run命令可以编译并运行命令源码文件。
- go get可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装。
- go build命令用于编译我们指定的源码文件或代码包以及它们的依赖包。
- go install用于编译并安装指定的代码包及它们的依赖包。
- go clean命令会删除掉执行其它命令时产生的一些文件和目录。
- go doc命令可以打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。
- go test命令用于对Go语言编写的程序进行测试。
- go list命令的作用是列出指定的代码包的信息。
- go fix会把指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码。
- go vet是一个用于检查Go语言源码中静态错误的简单工具。
- go tool pprof命令来交互式的访问概要文件的内容。
