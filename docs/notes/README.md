# Go 编码规范指南

## 序言

看过很多方面的编码规范，可能每一家公司都有不同的规范，这份编码规范是写给我自己的，同时希望我们公司内部同事也能遵循这个规范来写[Go](http://lib.csdn.net/base/go)代码。
如果你的代码没有办法找到下面的规范，那么就遵循标准库的规范，多阅读标准库的源码，标准库的代码可以说是我们写代码参考的标杆。

## 格式化规范

go 默认已经有了`gofmt`工具，但是我们强烈建议使用 goimport 工具，这个在 gofmt 的基础上增加了自动删除和引入包. `go get golang.org/x/tools/cmd/goimports`</br>
不同的编辑器有不同的配置, [sublime 的配置教程](http://lbb4511.cn/software/_book/tools/sublimetext/go.html)
LiteIDE 默认已经支持了 goimports，如果你的不支持请点击属性配置->golangfmt->勾选 goimports
保存之前自动 fmt 你的代码。

## 行长约定

一行最长不超过 80 个字符，超过的请使用换行展示，尽量保持格式优雅。`go vet`</br>
vet 工具可以帮我们静态分析我们的源码存在的各种问题，例如多余的代码，提前 return 的逻辑，struct 的 tag 是否符合标准等。</br>
`go get golang.org/x/tools/cmd/vet` 使用如下：`go vet .`

## package 名字

保持 package 的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，尽量和标准库不要冲突。

## import 规范

import 在多行的情况下，goimports 会自动帮你格式化，但是我们这里还是规范一下 import 的一些规范，如果你在一个文件里面引入了一个 package，还是建议采用如下格式：
`import "fmt"`
如果你的包引入了三种类型的包，标准库包，程序内部包，第三方包，建议采用如下方式进行组织你的包：

```go
import (
    "encoding/json"
    "strings" "myproject/models"
    "myproject/controller"
    "myproject/utils"
    "github.com/astaxie/beego"
    "github.com/go-sql-driver/mysql"
)
```

> 有顺序的引入包，不同的类型采用空格分离，第一种实标准库，第二是项目包，第三是第三方包。

在项目中不要使用相对路径引入包：

```go
// 这是不好的导入
import "../net"
// 这是正确的做法
import "github.com/repo/proj/src/net"
```

## 变量申明

变量名采用驼峰标准，不要使用\_来命名变量名，多个变量申明放在一起

```go
var (
    Found bool
    count int
)
```

> 在函数外部申明必须使用 var,不要采用:=，容易踩到变量的作用域的问题。

## 自定义类型的 string 循环问题

如果自定义的类型定义了 String 方法，那么在打印的时候会产生隐藏的一些 bug

```go
type MyInt
 intfunc (m MyInt) String() string {
    return fmt.Sprint(m)//BUG:死循环
}
 func(m MyInt) String() string {
    return fmt.Sprint(int(m)) //这是安全的,因为我们内部进行了类型转换
}
```

## 避免返回命名的参数

如果你的函数很短小，少于 10 行代码，那么可以使用，不然请直接使用类型，因为如果使用命名变量很容易引起隐藏的 bug

```go
func Foo(a int, b int) (string, ok){
}
```

当然如果是有多个相同类型的参数返回，那么命名参数可能更清晰：

```go
func (f *Foo) Location() (float64, float64, error)
```

下面的代码就更清晰了：

```go
// Location returns f's latitude and longitude.
// Negative values mean south and west, respectively.
func (f *Foo) Location() (lat, long float64, err error)
```

## 错误处理

错误处理的原则就是不能丢弃任何有返回 err 的调用，不要采用\_丢弃，必须全部处理。接收到错误，要么返回 err，要么实在不行就 panic，或者使用 log 记录下来

_error 信息_ error 的信息不要采用大写字母，尽量保持你的错误简短，但是要足够表达你的错误的意思。

## 长句子打印或者调用，使用参数进行格式化分行

**_我们在调用 fmt.Sprint 或者 log.Sprint 之类的函数时，有时候会遇到很长的句子，我们需要在参数调用处进行多行分割：_**

下面是错误的方式：

```go
log.Printf(“A long format string: %s %d %d %s”, myStringParameter, len(a),
    expected.Size, defrobnicate(“Anotherlongstringparameter”,
        expected.Growth.Nanoseconds() /1e6))`
应该是如下的方式：
`log.Printf(
    “A long format string: %s %d %d %s”,
    myStringParameter,
    len(a),
    expected.Size,
    defrobnicate(
        “Anotherlongstringparameter”,
        expected.Growth.Nanoseconds()/1e6,
    ),
)
```

## 注意闭包的调用

在循环中调用函数或者 goroutine 方法，一定要采用显示的变量调用，不要再闭包函数里面调用循环的参数

```go
fori:=0;i<limit;i++{
    go func(){ DoSomething(i) }() //错误的做法
    go func(i int){ DoSomething(i) }(i)//正确的做法
}
```

[Race on loop counter](https://golang.google.cn/doc/articles/race_detector.html#Race_on_loop_counter)

## 在逻辑处理中禁用 panic

在 main 包中只有当实在不可运行的情况采用 panic，例如文件无法打开，数据库无法连接导致程序无法正常运行，但是对于其他的 package 对外的接口不能有 panic，只能在包内采用。</br>
强烈建议在 main 包中使用 log.Fatal 来记录错误，这样就可以由 log 来结束程序。

## 注释规范

注释可以帮我们很好的完成文档的工作，写得好的注释可以方便我们以后的维护。详细的如何写注释可以参考：[Commentary](http://golang.google.cn/doc/effective_go.html#commentary) ###_bug 注释_

针对代码中出现的 bug，可以采用如下教程使用特殊的注释，在 godocs 可以做到注释高亮：

```go
// BUG(astaxie):This divides by zero. var i float = 1/0
```

## struct 规范

**_struct 申明和初始化格式采用多行：_**

定义如下：

```go
type User struct{
    Username  string
    Email     string
}
```

初始化如下：

```go
u := User{
    Username: "astaxie",
    Email:    "astaxie@gmail.com",
}
```

**_recieved 是值类型还是指针类型_**

到底是采用值类型还是指针类型主要参考如下原则：

```go
func(w Win) Tally(playerPlayer)int // w不会有任何改变
func(w *Win) Tally(playerPlayer)int // w会改变数据
```

**_带 mutex 的 struct 必须是指针 receivers_**</br>
如果你定义的 struct 中带有 mutex,那么你的 receivers 必须是指针

## 参考资料

- [代码审查注释](https://code.google.com/p/go-wiki/wiki/CodeReviewComments)
- [Effective Go](http://golang.org/doc/effective_go.html)
