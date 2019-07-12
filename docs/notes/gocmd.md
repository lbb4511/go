# Go 命令

| 命令     | 解释                                           |
| -------- | ---------------------------------------------- |
| build    | compile packages and dependencies              |
| clean    | remove object files                            |
| doc      | show documentation for package or symbol       |
| env      | print Go environment information               |
| fix      | run go tool fix on packages                    |
| fmt      | run gofmt on package sources                   |
| generate | generate Go files by processing source         |
| get      | download and install packages and dependencies |
| install  | compile and install packages and dependencies  |
| list     | list packages                                  |
| run      | compile and run Go program                     |
| test     | test packages                                  |
| tool     | run specified go tool                          |
| version  | print Go version                               |
| vet      | run go tool vet on packages                    |

1. **_`go build`_** 编译指定的源码文件或代码包以及它们的依赖包。`go build [-o output] [-i] [build flags] [packages]`
   - `-a`强行对所有涉及到的代码包（包含标准库中的代码包）进行重新构建，即使它们已经是最新的了。
   - `-n`打印编译期间所用到的其它命令，但是并不真正执行它们。
   - `-p n`指定编译过程中执行各任务的并行数量（确切地说应该是并发数量）。在默认情况下，该数量等于 CPU 的逻辑核数。但是在 darwin/arm 平台（即 iPhone 和 iPad 所用的平台）下，该数量默认是 1。_n 为正整数_
   - `-race`开启竞态条件的检测。不过此标记目前仅在 linux/amd64、freebsd/amd64、darwin/amd64 和 windows/amd64 平台下受到支持。
   - `-v`打印出那些被编译的代码包的名字。
   - `-work`打印出编译时生成的临时工作目录的路径，并在编译结束时保留它。在默认情况下，编译结束时会删除该目录。
   - `-x`打印编译期间所用到的其它命令。注意它与-n 标记的区别。
   - `-asmflags`此标记可以后跟另外一些标记，如`-D`、`-I`、`-S`等。这些后跟的标记用于控制 Go 语言编译器编译汇编语言文件时的行为。
   - `-buildmode`此标记用于指定编译模式，使用方式如-buildmode=default（这等同于默认情况下的设置）。此标记支持的编译模式目前有 6 种。借此，我们可以控制编译器在编译完成后生成静态链接库（即.a 文件，也就是我们之前说的归档文件）、动态链接库（即.so 文件）或/和可执行文件（在 Windows 下是.exe 文件）。
   - `-gcflags`此标记用于指定需要传递给 go tool compile 命令的标记的列表。`go build -gcflags "-N -l" -o test test.go`
2. **_`go clean`_** 删除掉执行其它命令时产生的一些文件和目录，包括：
   - 使用`go build`命令时在当前代码包下生成的与包名同名或者与 Go 源码文件同名的可执行文件。
   - 执行`go test`命令并加入`-c`标记时在当前代码包下生成的以包名加“.test”后缀为名的文件。
   - 还有一些目录和文件是在编译 Go 或 C 源码文件时留在相应目录中的。包括：“\_obj”和“\_test”目录，名称为“\_testmain.go”、“test.out”、“build.out”或“a.out”的文件，名称以“.5”、“.6”、“.8”、“.a”、“.o”或“.so”为后缀的文件。这些目录和文件是在执行 go build 命令时生成在临时目录中的。`go clean [-i] [-r] [-n] [-x] [build flags] [packages]`
   - `-r`，包括当前代码包的所有依赖包的上述目录和文件。
   - `-i`，同时删除安装当前代码包时所产生的结果文件。如果当前代码包中只包含库源码文件，则结果文件指的就是在工作区的 pkg 目录的相应目录下的归档文件。如果当前代码包中只包含一个命令源码文件，则结果文件指的就是在工作区的 bin 目录下的可执行文件。
3. **_`go doc`_** 打印附于 Go 语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。
   `go doc [-u] [-c] [package|[package.]symbol[.method]]` - `-c`加入此标记后会使 go doc 命令区分参数中字母的大小写。 - `-cmd`打印出 main 包中的可导出的程序实体（其名称的首字母大写）的文档。 - `-u`打印出不可导出的程序实体（其名称的首字母小写）的文档。 - `godoc`是一个很强大的工具，同样用于展示指定代码包的文档。在 Go 语言的 1.5 版本以后，它是一个内置的标准命令。 - `godoc fmt` `godoc -http=:6060`标记-http 的值:6060 表示启动的 Web 服务器使用本机的 6060 端口。之后，我们就可以通过在网络浏览器的地址栏中输入[](http://localhost:6060)[http://localhost:6060](http://localhost:6060)来查看以网页方式展现的 Go 文档了。
4. **_`go env`_** 用于打印 Go 语言的环境信息。 `go env [var ...]` [Go 语言的 GOPATH 与工作目录详解](http://www.jb51.net/article/56779.htm)

5. **_`go fix`_** 用来修复以前老版本的代码到新版本。命令`go fix`其实是命令`go tool fix`的简单封装。这甚至比`go fmt`命令对`gofmt`命令的封装更简单。像其它的 Go 命令一样，`go fix`命令会先对作为参数的代码包导入路径进行验证，以确保它是正确有效的。
   `go fix [packages]` - `-diff`不将修正后的内容写入文件，而只打印修正前后的内容的对比信息到标准输出。 - `-r`只对目标源码文件做有限的修正操作。该标记的值即为允许的修正操作的名称。多个名称之间用英文半角逗号分隔。 - `-force`使用此标记后，即使源码文件中的代码已经与 Go 语言的最新版本相匹配了，也会强行执行指定的修正操作。该标记的值就是需要强行执行的修正操作的名称，多个名称之间用英文半角逗号分隔。
6. **_`go fmt`_** 格式化 I / O 的功能,和 C 的 printf 和 scanf 类似。[golang 中 fmt 用法](http://blog.csdn.net/chenbaoke/article/details/39932845) / [fmt 包 Printf](http://www.cnblogs.com/dasn/articles/5028811.html) 在 go 中，代码则有标准的风格 `go fmt [-n] [-x] [packages]`
   - `-n`打印编译期间所用到的其它命令，但是并不真正执行它们。
   - `-x`可以看到 go get 命令执行过程中所使用的所有命令。
7. **_`go generate`_**是一个你可以用来自动自成 Go 代码的命令，你可以结合例如 jsonenums(一个用于为枚举类型自动生成 JSON 编组样板代码的类库)这样的元编程来使用 go generate 快速自动实现重复乏味代码的编写。在 Go 标准类库里面已经有大量可以用于解析 AST 的接口，而 AST 使得编写元编程工具更简单，更容易。
8. **_`go get`_**可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装。`go get [-d] [-f] [-fix] [-insecure] [-t] [-u] [build flags] [packages]`
   - `-d`让命令程序只执行下载动作，而不执行安装动作。
   - `-f`仅在使用-u 标记时才有效。该标记会让命令程序忽略掉对已下载代码包的导入路径的检查。如果下载并安装的代码包所属的项目是你从别人那里 Fork 过来的，那么这样做就尤为重要了。
   - `-fix`让命令程序在下载代码包后先执行修正动作，而后再进行编译和安装。之前解决相关性或建设中的代码。
   - `-insecure`允许命令程序使用非安全的 scheme（如 HTTP）去下载指定的代码包。如果你用的代码仓库（如公司内部的 Gitlab）没有 HTTPS 支持，可以添加此标记。请在确定安全的情况下使用它。
   - `-t`让命令程序同时下载并安装指定的代码包中的测试源码文件中依赖的代码包。
   - `-u`让命令利用网络来更新已有代码包及其依赖包。默认情况下，该命令只会从网络上下载本地不存在的代码包，而不会更新已有的代码包。
   - `-x`可以看到 go get 命令执行过程中所使用的所有命令
   - 智能下载 命令 go get 还有一个很值得称道的功能。在使用它检出或更新代码包之后，它会寻找与本地已安装 Go 语言的版本号相对应的标签（tag）或分支（branch）。比如，本机安装 Go 语言的版本是 1.x，那么 go get 命令会在该代码包的远程仓库中寻找名为“go1”的标签或者分支。如果找到指定的标签或者分支，则将本地代码包的版本切换到此标签或者分支。如果没有找到指定的标签或者分支，则将本地代码包的版本切换到主干的最新版本。
9. **_`go install`_**用于编译并安装指定的代码包及它们的依赖包。实际上，`go install`命令只比`go build`命令多做了一件事，即：安装编译后的结果文件到指定目录。`go install [build flags] [packages]`
10. **_`go list`_**命令的作用是列出指定的代码包的信息。与其他命令相同，我们需要以代码包导入路径的方式给定代码包。被给定的代码包可以有多个。这些代码包对应的目录中必须直接保存有 Go 语言源码文件，其子目录中的文件不算在内。否则，代码包将被看做是不完整的。`go list [-e] [-f format] [-json] [build flags] [packages]`
11. **_`go run`_**命令可以编译并运行命令源码文件。由于它其中包含了编译动作，因此它也可以接受所有可用于`go build`命令的标记。</br>除了标记之外，`go run`命令只接受 Go 源码文件作为参数，而不接受代码包。</br>与`go build`命令和`go install`命令一样，`go run`命令也不允许多个命令源码文件作为参数，即使它们在同一个代码包中也是如此。而原因也是一致的，多个命令源码文件会都有 main 函数声明。`go run [build flags] [-exec xprog] gofiles... [arguments...]`</br>如果命令源码文件可以接受参数，那么在使用 go run 命令运行它的时候就可以把它的参数放在它的文件名后面，像这样：`go run showds.go -p ~/golang/goc2p`。
12. **_`go test`_** `go test [build/test flags] [packages] [build/test flags & test binary flags]`
    - `-c`生成用于运行测试的可执行文件，但不执行它。这个可执行文件会被命名为“pkg.test”，其中的“pkg”即为被测试代码包的导入路径的最后一个元素的名称。
    - `-i`安装/重新安装运行测试所需的依赖包，但不编译和运行测试代码。
    - `-o`指定用于运行测试的可执行文件的名称。追加该标记不会影响测试代码的运行，除非同时追加了标记-c 或-i。
13. **_`go tool`_** `go tool [-n] command [args...]`
    - `go tool pprof`命令来交互式的访问概要文件的内容。命令将会分析指定的概要文件，并会根据我们的要求为我们提供高可读性的输出信息。
    - `go tool cgo`这个工具可以使我们创建能够调用 C 语言代码的 Go 语言源码文件。这使得我们可以使用 Go 语言代码去封装一些 C 语言的代码库，并提供给 Go 语言代码或项目使用。
14. **_`go version`_** 查看 Go 版本
15. **_`go vet`_** 是一个用于检查 Go 语言源码中静态错误的简单工具 `go vet [-n] [-x] [build flags] [packages]`
    - `go vet`命令是`go tool vet`命令的简单封装。它会首先载入和分析指定的代码包，并把指定代码包中的所有 Go 语言源码文件和以“.s”结尾的文件的相对路径作为参数传递给`go tool vet`命令。

> - \*标记`-n`会让命令在执行过程中打印用到的系统命令，但不会真正执行它们。

- \*标记`-x`既打印命令又执行命令。
- \*标记`-r`，包括当前代码包的所有依赖包的上述目录和文件。

---

> 文档

- [Go 命令教程](http://wiki.jikexueyuan.com/project/go-command-tutorial/0.0.html)
