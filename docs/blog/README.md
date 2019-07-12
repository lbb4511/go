# 博客

## GO 开发者对 GO 初学者的建议

以促进 India 的 go 编程作为 GopherConIndia 承诺的一部分。我们采访了  [40 位 Gophers](http://list.ly/list/Pak-gopher-interviews)（一个 Gopher 代表一个 GO 项目或是任何地方的 GO 程序员），得到了他们关于 GO 的意见。从 2014 年的八月到十一月，我们将每个星期发表两篇采访稿。
如果你正好刚刚开始 go 编程，他们对于我们一些问题的答案可能会对你有非常有用。看看这些。

### 应该做

- 通读  [the Go standard library](http://golang.org/pkg/)  和  [Effective Go](http://golang.org/doc/effective_go.html)，为了学习 GO 的规范，Effective Go 是被高度推荐的，尤其是如果你有其他语言的背景。
- 在  [Go tour](http://tour.golang.org/#1)  上做练习
- 看完[语言参考](https://golang.org/ref/spec)
- 练习  [Go by Example](https://gobyexample.com/)，而不仅仅是复制粘贴！
- 坚持编写 GO 代码，在几周内你将会在这门语言上变得高效
- 理解接口的功能，他们是 GO 最大的礼物之一，可能比 channels 和 goroutines 还重要。这个关于接口的文章  [article on interfaces](http://mwholt.blogspot.in/2014/08/maximizing-use-of-interfaces-in-go.html)  和 Andrew Gerrand 在 GopherCon 2014 上的 keynote [接口的描述](http://talks.golang.org/2014/go4gophers.slide#5)  会对你非常有帮助。
- 抛弃你的 OO 的思想包袱，如果你来自于其他语言，比如动态语言 Python 或是 Ruby，或者是一个编译型语言如 Java 或 C#。GO 是一个面向对象的语言，但是它不是一个基于 class 的语言和不支持继承。
- 了解继承从 GO 语言中移除了。实践组合的用法而不是继承的机会显现了，并且纠结于继承只会导致你沮丧
- 不要以其他语言的风格编写 GO
- 寻找更加有经验的 Gophers，他们能帮助你 review 代码片段和给你反馈。在 GO 社区能得到真正的支持和帮助
- 用 GO 实现你想法中的一个项目或是找到一个项目来工作。然后随着你学习的更多，不断重构你的应用。利用[邮件列表](https://groups.google.com/forum/#!forum/golang-nuts)和参加  [Gopher Academy Slack group](https://gophers.slack.com/)  来见其他的 Gophers 来得到帮助。[Dave Cheney](http://dave.cheney.net/)  的博客和  [GoingGo](http://www.goinggo.net/)  的博客也是一个非常好的开始
- 不要等待泛型和函数式被添加进语言；屏住呼吸并学习爱上我们在今天拥有的这门语言

> 注：私人添加，可以订阅 Newspaper.io 的 Golang Daily，以及  [@ASTA 谢](http://weibo.com/533452688)  的  [《Go Web 编程》](https://github.com/astaxie/build-web-application-with-golang) 【作者也出了实体书，大家可以购买】和 订阅 Golang Ask News，社区  [http://golanghome.com/](http://golanghome.com/)，[@无闻 Unknown](http://weibo.com/Obahua)  的  [《Go 编程基础》](https://github.com/Unknwon/go-fundamental-programming)，[《Go Web 基础》](https://github.com/Unknwon/go-web-foundation)  和  [《Go 名库讲解》](https://github.com/Unknwon/go-rock-libraries-showcases)

## 给 go 初学者分享的一些问题

- 对于任何人来说学习一门新语言可能都是令人挫折的。GO 社区是不可置信的活跃，你不是孤单的。利用所有的文档，博客，本地的 Meetups 和用户组，比如 Slack。不要害怕问问题和参与
- 如果你对 GO 感兴趣，使用它的一侧涉足，或是专业的使用它，如果本地有 Go meetup，考虑参与。如果你有货，考虑去分享它
- 如果你有计划旅行，并且有能力，努力去访问 GO 社区目的地
- 来访的用户群体是个证明这个社区有众多的用户，支持者和雇员的途径
- 不要浪费时间去和其他语言比较，如果你喜欢 GO，就爱上他并且去使用它
- 接受 Go 的文化和 GO 做事情的方式。你的代码会感谢你，如果你这样做了，你会得到很多
- 不要冲动的引入依赖
- 简单是 GO 最重要的特征。避免过分设计，使用简单的代码片段而不是单一的庞大的代码库
- 从其他语言移植库到 GO 是一个很好的做法，它允许你剥离他人的代码并且以符合 GO 语言的方式粘合起来。

> 注：How do you see the market for Go Programmers in the work place? What is the future for Go 这部分不翻译，请读者自己看

## GO 程序员进化史

- 来自贝尔实验室的两位大牛罗布·派克,肯·汤普逊与 Google 的罗伯特·格瑞史莫,2007 年开始设计的一种编译型，可平行化，并具有垃圾回收功能的编程语言。
  - 罗布·派克（[Rob Pike](http://baike.baidu.com/view/3463850.htm)）1980 年奥运会射箭的银牌得主(超级厉害了)
  - C 语言之父肯·汤普逊([Kenneth Lane Thompson](http://baike.baidu.com/view/6667836.htm))
  - 罗伯特·格瑞史莫（[Robert Griesemer](https://en.wikipedia.org/wiki/Robert_Griesemer)）曾在 Java 编译器方面的工作
  - Ian Lance Taylor[Github](https://github.com/ianlancetaylor)
  - [Russ Cox](http://blog.jobbole.com/82902/)
- Go 语言是基于 Inferno 操作系统所开发的。
- Go 语言于 2009 年 11 月有 Google 正式宣布推出，成为开放源代码项目，并在 Linux 及 Mac OS X 平台上进行了实现，后追加 Windows 系统下的实现。
- Go 语言的特点在[wiki](https://zh.wikipedia.org/wiki/Go)里有介绍
- Go 程序可以从 3 行到数百万行，它被写入到一个或多个文本文件扩展名“.go”;例如，hello.go。你可以用“vi”，“vim”或任何其他文本编辑器来编写 Go 程序到到文件中。
- go build hello.go 可以编译生成一个二进制的 hello linux 没有后缀
- go run hello.go 可以直接执行
- [Francesc](https://twitter.com/francesc) 是 Go 核心团队的一员, 是提倡 Google Cloud 平台的开发者. 他是一个编程语言的爱好者, Google 的技术指导大师, Go tour 的创造者之一. 这个讨论的灵感来自于另一个 [Raquel Vélez](https://twitter.com/rockbot) 在 JSConf. Slides 的讨论。
- [Sourcegraph](https://sourcegraph.com/) 是下一代编程协作工具, 用于搜索, 探索, 和审查代码. 我们参加 GopherCon India 来分享我们是怎样使用 Go 并学习别人是怎样使用它的, 对配合 liveblog 的这次讨论我们深感荣幸。

> 作为 Go 团队的开发者之一，Francesc 可能比世界上其他人接触到的 Go 语言程序员都要多。正因为有了这样的有利条件，他把 Go 语言的学习过程划分为 5 个阶段。</br>
> 这些阶段对于其他语言的学习也是成立的。理解自己处于哪个阶段，可以帮助你找到提高自己的最有效的方法，也可以避免每个阶段学习过程中的常见陷阱。</br>
> 编者按：这篇文章对于每一个学习阶段都给出了交互式的代码片段。点击函数名你就可以跳到具体的函数定义，方便进行深入的研究。请看下文。

## 这里是 GO 程序员的五个进化阶段

- 第一阶段 (菜逼): 刚刚学习了这门语言。 已经通过一些教程或者培训班了解基本的语法，可以写短的代码片段。
- 第二阶段 (探索者): 可以写一个完整的程序，但不懂一些更高级的语言特征，比如“channels”。还没有使用 GO 写一个大项目。
- 第三阶段 (大手): 你能熟练的使用 Go, 能够用 GO 去解决，生产环境中一个具体和完整的问题。已经形成了一套自己的惯用法和常用代码库。在你的编码方案中 Go 是一个非常好用的工具。
- 第四阶段 (大神): 绝逼清楚 Go 语言的设计选择和背后的动机。能理解的简洁和可组合性哲学。
- 第五阶段 (布道师): 积极地与他人分享关于 Go 语言知识和你对 Go 语言的理解。在各种合适的场所发出自己的声音, 参与邮件列表、建立 QQ 群、做专题报告。成为一个布道者不见得是一个完全独立的阶段，这个角色可以在上述的任何一个阶段中。

### 第一阶段: 菜逼

菜鸟在这个阶段使用 Go 去创建一些小项目或者玩具项目。他们应该会利用到[Go tour](https://tour.golang.org/welcome/1), [Go playground](https://play.golang.org/), [Go 文档](http://golang.org/doc/), 和邮件列表([golang-nuts](https://groups.google.com/forum/)).

```golang
func main() {
    fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))}
```

这是 Go 语言写的简单例子，这个代码段来自[golang/example](https://sourcegraph.com/github.com/golang/example)代码库里面的 hello.go。点击就可以查看完整代码撸。
一项重要的技能，新人应该试着学习如何正确提问。很多新人在邮件列表里面这样说“嘿，这报错了”，这并没有提供足够的信息，让别人能理解并帮助他们解决问题。别人看到的是一个粘贴了几百行的代码的帖子，并没有花费精力来重点说明所遇到的问题。
所以, 应该尽量避免直接粘贴代码到论坛。而应该使用可以编辑并且可以在浏览器中直接运行的 Go playground 的“分享”按钮链接到代码片段。

### 第二阶段：探索者

探索者已经可以使用 Go 写一些小的软件，但有时仍然会有些迷茫。他们可能不完全明白怎么使用 Go 的高级特性，比如通道。虽然他们还有很多东西要学习，但已掌握的足够做一些有用的事情了！他们开始对 Go 的潜能有感觉了，并对它们能使用 Go 创建的东西感到兴奋。
在探索阶段通常会经历两个步骤。第一，膨胀的预期达到顶点，你觉得可以用 Go 做所有的事情，但还并不能明白或领悟到 Go 的真谛。你大概会用所熟悉的语言的模式和惯用语来写 Go 代码，但对于什么是地道的 Go，还没有比较强烈的感觉。你开始尝试着手干这样的事情--“迁移架构 X，从 Y 语言到 Go 语言”。
到达预期膨胀的顶点之后，你会遇到理想幻灭的低谷。你开始想念语言 Y 的特性 X，此时你还没有完全的掌握地道的 Go。你还在用其他编程语言的风格来写 Go 语言的程序，你甚至开始觉得沮丧。你可能在大量使用 reflect 和 unsafe 这两个包，但这不是地道的 Go。地道的 Go 不会使用那些魔法一样的东西。
这个探索阶段产生的项目的一个很好的例子就是[Martini Web 框架](https://sourcegraph.com/github.com/go-martini/martini)。Martini 是一个 Go 语言的早期 Web 框架，它从 Ruby 的 Web 框架当中吸收了很多思想（比如依赖注入）。最初，这个框架在社区中引起了强烈的反响，但是它逐渐在性能和可调试性上受到了一些批评。Martini 框架的作者 Jeremy Saenz 积极响应这些来自 Go 社区的反馈，写了一个更加符合 Go 语言规范的库[Negroni](https://github.com/codegangsta/negroni)

```golang
func (m *Martini) RunOnAddr(addr string) {
    // TODO: Should probably be implemented using a new instance of http.Server in place of
    // calling http.ListenAndServer directly, so that it could be stored in the martini struct for later use.
    // This would also allow to improve testing when a custom host and port are passed.

    logger := m.Injector.Get(reflect.TypeOf(m.logger)).Interface().(*log.Logger)
    logger.Printf("listening on %s (%s)\n", addr, Env)
    logger.Fatalln(http.ListenAndServe(addr, m))
}
```

来自 Martini 框架的交互式代码片段，它是不地道的 Go 的例子。注意用反射包实现的依赖注入

```golang
func TestNegroniServeHTTP(t *testing.T) {
    result := ""
    response := httptest.NewRecorder()

    n := New()
    n.Use(HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
        result += "foo"
        next(rw, r)
        result += "ban"
    }))
    n.Use(HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
        result += "bar"
        next(rw, r)
        result += "baz"
    }))
    n.Use(HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
        result += "bat"
        rw.WriteHeader(http.StatusBadRequest)
    }))

    n.ServeHTTP(response, (*http.Request)(nil))

    expect(t, result, "foobarbatbazban")
    expect(t, response.Code, http.StatusBadRequest)
}
```

来自 Negroni 库的交互式代码片段，它是地道的 Go 的例子
其他语言在提供一些核心功能，比如 HTTP 处理的时候，往往需要依赖第三方库。但是 Go 语言在这一点上很不同，它的标准库非常强大。如果你认为 Go 标准库没有强大到可以做你想做的事情，那么我说你错了。Go 语言标准库难以置信的强大，值得你花时间阅读它的代码，学习它实现的模式。

```golang
func (srv *Server) ListenAndServe() error {
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
```

Go 标准库中的 ListenAndServe 函数片段。如果你写过 Go 程序，你可能已经调用过这个函数很多次了，但是你曾经花时间看过它的实现么？去点击上面的代码片段吧。
幻灭的低谷中的幻灭感来自于这样的事实：你还在用其他语言的模式来想问题，而且你还没有完全探索过 Go 能提供给你什么。下面是一些好玩的事情，你可以做一下来打破困境，进一步探索这门语言中好玩的事。
go generate
现在来看看 go generate。go generate 是一个你可以用来自动自成 Go 代码的命令。你可以结合例如 jsonenums(一个用于为枚举类型自动生成 JSON 编组样板代码的类库)这样的元编程来使用 go generate 快速自动实现重复乏味代码的编写。在 Go 标准类库里面已经有大量可以用于解析 AST 的接口，而 AST 使得编写元编程工具更简单，更容易。在会议上，有另外两次讨论（Go 语言中的元编程实践和拥抱标准类库）谈及到了这一点。

```golang
func main() {
    flag.Parse()
    if len(*typeNames) == 0 {
        log.Fatalf("the flag -type must be set")
    }
    types := strings.Split(*typeNames, ",")

    // Only one directory at a time can be processed, and the default is ".".
    dir := "."
    if args := flag.Args(); len(args) == 1 {
        dir = args[0]
    } else if len(args) > 1 {
        log.Fatalf("only one directory at a time")
    }

    pkg, err := parser.ParsePackage(dir, *outputSuffix+".go")
    if err != nil {
        log.Fatalf("parsing package: %v", err)
    }

    var analysis = struct {
        Command        string
        PackageName    string
        TypesAndValues map[string][]string
    }{
        Command:        strings.Join(os.Args[1:], " "),
        PackageName:    pkg.Name,
        TypesAndValues: make(map[string][]string),
    }

    // Run generate for each type.
    for _, typeName := range types {
        values, err := pkg.ValuesOfType(typeName)
        if err != nil {
            log.Fatalf("finding values for type %v: %v", typeName, err)
        }
        analysis.TypesAndValues[typeName] = values

        var buf bytes.Buffer
        if err := generatedTmpl.Execute(&buf, analysis); err != nil {
            log.Fatalf("generating code: %v", err)
        }

        src, err := format.Source(buf.Bytes())
        if err != nil {
            // Should never happen, but can arise when developing this code.
            // The user can compile the output to see the error.
            log.Printf("warning: internal error: invalid Go generated: %s", err)
            log.Printf("warning: compile the package to analyze the error")
            src = buf.Bytes()
        }

        output := strings.ToLower(typeName + *outputSuffix + ".go")
        outputPath := filepath.Join(dir, output)
        if err := ioutil.WriteFile(outputPath, src, 0644); err != nil {
            log.Fatalf("writing output: %s", err)
        }
    }
}
```

一段互动的片段演示了如何编写 jsonenums 命令。
OpenGL
许多人使用 Go 作 web 服务，但是你知道你也可以用 Go 写出很 cool 的图形应用吗？查看 Go 在[OpenGL](https://sourcegraph.com/github.com/go-gl/gl)中的捆绑。

```golang
func main() {
    glfw.SetErrorCallback(errorCallback)

    if !glfw.Init() {
        panic("Can't init glfw!")
    }
    defer glfw.Terminate()

    window, err := glfw.CreateWindow(Width, Height, Title, nil, nil)
    if err != nil {
        panic(err)
    }

    window.MakeContextCurrent()

    glfw.SwapInterval(1)

    gl.Init()

    if err := initScene(); err != nil {
        fmt.Fprintf(os.Stderr, "init: %s\n", err)
        return
    }
    defer destroyScene()

    for !window.ShouldClose() {
        drawScene()
        window.SwapBuffers()
        glfw.PollEvents()
    }
}
```

交互式的片段正说明 Go 的 OpenGL 捆绑能制作 Gopher cube。点击函数或方法名去探索。
黑客马拉松和挑战
你也可以观看挑战和黑客马拉松，类似[Gopher Gala](http://gophergala.com/)和[Go Challenge](http://ww2.golang-challenge.com/)。在过去，来自世界各地的程序员一起挑战一些真实的酷项目，你可以从中获取灵感。

### 第三阶段: 老手

作为一个老手，这意味着你可以解决很多 Go 语言中你关心的问题。新的需要解决的问题会带来新的疑问，经过试错，你学会了在这门语言中什么是可以做的，什么是不能做的。此时，你已经对这门语言的习惯和模式有了一个坚实的理解。你可以非常高效地工作，写出可读，文档完善，可维护的代码。
成为老手的一个很好的方法就是在大项目上工作。如果你自己有一个项目的想法，开始动手去做吧（当然你要确定它并不是已经存在了）。大多数人也许并没有一个很大的项目的想法，所以他们可以对已经存在的项目做出贡献。Go 语言已经有很多大型项目，而且它们正在被广泛使用，比如 Docker, Kubernetes 和 Go 本身。可以看看这个项目[列表](https://github.com/golang/go/wiki/Projects)

```golang
func (cli *DockerCli) CmdRestart(args ...string) error {
    cmd := cli.Subcmd("restart", "CONTAINER [CONTAINER...]", "Restart a running container", true)
    nSeconds := cmd.Int([]string{"t", "-time"}, 10, "Seconds to wait for stop before killing the container.")
    cmd.Require(flag.Min, 1)

    utils.ParseFlags(cmd, args, true)

    v := url.Values{}
    v.Set("t", strconv.Itoa(*nSeconds))

    var encounteredError error
    for _, name := range cmd.Args() {
        _, _, err := readBody(cli.call("POST", "/containers/"+name+"/restart?"+v.Encode(), nil, false))
        if err != nil {
            fmt.Fprintf(cli.err, "%s\n", err)
            encounteredError = fmt.Errorf("Error: failed to restart one or more containers")
        } else {
            fmt.Fprintf(cli.out, "%s\n", name)
        }
    }
    return encounteredError
}
```

Docker 项目的交互式代码片段。点击函数名，开始探索之旅吧。
老手应该对 Go 生态系统的工具有一个很强的掌握，因为这些工具真的提高生产效率。你应该了解 go generate，go vet，go test-race， 和 gofmt/goimports/goreturns。你应该使用 go fmt，因为它会自动把你的代码按照 Go 社区的风格标准来格式化。goimports 可以做同样的事情，而且还会添加丢失的 imports。goretures 不光做了前面所说的事情，还可以在返回表达式添加丢失的错误，这是大家都讨厌的地方。
在老手阶段，你一定要开始做 code review。code review 的意义并不是要修改或者找到错误（那是测试人员做的事情）。code review 可以帮助维持统一的编程风格，提高软件的总体质量，还可以在别人的反馈中提高你自己的编程技术。几乎所有的大型开源项目都对每一个提交做 code review。
下面是一个从人类的反馈当中学习的例子：Google 的 Go 团队以前都在 main 函数的外面声明命令行标记。在去年的 GopherCon 会议上，Francesc 遇到了 SoundCloud 公司的 Peter Bourgon（@peterbourgon）。Peter Bourgon 说在 SoundCloud，他们都在 main 函数内部声明标记，这样他们不会错误地在外部使用标记。Francesc 现在认为这是最佳实践。

### 第四阶段：专家

作为一个专家，你很好地了解了语言的哲学思想。对于 Go 语言的特性，你知道何时应该使用，何时不应该使用。例如，Jeremy Saenz 在 dotGo 风暴讨论中谈论到了何时不该使用接口。

```golang
func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call {
    call := new(Call)
    call.ServiceMethod = serviceMethod
    call.Args = args
    call.Reply = reply
    if done == nil {
        done = make(chan *Call, 10) // buffered.
    } else {
        // If caller passes done != nil, it must arrange that
        // done has enough buffer for the number of simultaneous
        // RPCs that will be using that channel.  If the channel
        // is totally unbuffered, it's best not to run at all.
        if cap(done) == 0 {
            log.Panic("rpc: done channel is unbuffered")
        }
    }
    call.Done = done
    client.send(call)
    return call
}
```

来自标准类库的一小块交互代码片段使用了频道。理解标准类库里面的模式背后的决策原因是成为一个专家必经之路。
但是不要成为只局限于单一语言的专家。跟其他任何语言一样，Go 仅仅只是一个工具。你还应该去探索其他语言，并且学习他们的模式和风格。Francesc 从他使用 Go 的经验中找到了编写 JavaScript 的启发。他还喜欢重点关注于不可变性和致力于避免易变性的 Haskell 语言，并从中获得了如何编写 Go 代码的灵感。

### 第五阶段：布道者

作为一个布道者，你分享自己的知识，传授你学会的和你提出的最佳实践。你可以分享自己对 Go 喜欢或者不喜欢的地方。全世界各地都有 Go 会议，[找到离你最近的](http://www.meetup.com/topics/golang/)。
你可以在任何一个阶段成为布道者，不要等到你成为这个领域的专家的时候才发出自己的声音。在你学习 Go 的任何一个阶段，提出问题，结合你的经验给出反馈，不要羞于提出自己不喜欢的地方。你提出的反馈可以帮助社区改善做事情的方法，也可能改变你自己对编程的看法。

```golang
func main() {
    httpAddr := flag.String("http", "127.0.0.1:3999", "HTTP service address (e.g., '127.0.0.1:3999')")
    originHost := flag.String("orighost", "", "host component of web origin URL (e.g., 'localhost')")
    flag.StringVar(&basePath, "base", "", "base path for slide template and static resources")
    flag.BoolVar(&present.PlayEnabled, "play", true, "enable playground (permit execution of arbitrary user code)")
    nativeClient := flag.Bool("nacl", false, "use Native Client environment playground (prevents non-Go code execution)")
    flag.Parse()

    if basePath == "" {
        p, err := build.Default.Import(basePkg, "", build.FindOnly)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Couldn't find gopresent files: %v\n", err)
            fmt.Fprintf(os.Stderr, basePathMessage, basePkg)
            os.Exit(1)
        }
        basePath = p.Dir
    }
    err := initTemplates(basePath)
    if err != nil {
        log.Fatalf("Failed to parse templates: %v", err)
    }

    ln, err := net.Listen("tcp", *httpAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer ln.Close()

    _, port, err := net.SplitHostPort(ln.Addr().String())
    if err != nil {
        log.Fatal(err)
    }
    origin := &url.URL{Scheme: "http"}
    if *originHost != "" {
        origin.Host = net.JoinHostPort(*originHost, port)
    } else if ln.Addr().(*net.TCPAddr).IP.IsUnspecified() {
        name, _ := os.Hostname()
        origin.Host = net.JoinHostPort(name, port)
    } else {
        reqHost, reqPort, err := net.SplitHostPort(*httpAddr)
        if err != nil {
            log.Fatal(err)
        }
        if reqPort == "0" {
            origin.Host = net.JoinHostPort(reqHost, port)
        } else {
            origin.Host = *httpAddr
        }
    }

    if present.PlayEnabled {
        if *nativeClient {
            socket.RunScripts = false
            socket.Environ = func() []string {
                if runtime.GOARCH == "amd64" {
                    return environ("GOOS=nacl", "GOARCH=amd64p32")
                }
                return environ("GOOS=nacl")
            }
        }
        playScript(basePath, "SocketTransport")
        http.Handle("/socket", socket.NewHandler(origin))
    }
    http.Handle("/static/", http.FileServer(http.Dir(basePath)))

    if !ln.Addr().(*net.TCPAddr).IP.IsLoopback() &&
        present.PlayEnabled && !*nativeClient {
        log.Print(localhostWarning)
    }

    log.Printf("Open your web browser and visit %s", origin.String())
    log.Fatal(http.Serve(ln, nil))
}
```

流行的 present 命令的 main 函数，很多 Go 的用户使用它来制作幻灯片。许多演讲者修改了这个模块来满足自己的需要。

### Q&A

> - Q: 在 GO 语言中，我所怀念的一项功能是一个好的调试器。

- A: 我们正在做了，不只是调试器，我们还会提供一个更好的总体监视工具可以让你在程序运行时更好地洞察程序在干什么(显示出所有正在运行的 goroutine 的状态)。在 GO 1.5 中探索它吧。

## go 语言学习历程

接触 go 是 2012 年的时候，真正开始系统的学习和开发系统是 2014 年了，go 语言的学习也算自己 2014 年的重要工作之一，对 go 语言学习的总结，也算是年底总结之一

### 学习 go 的原因和动机

1. 先前做过 2 年 Unix c 开发经验, 对于 C 系的语言有特殊的感情，go 特别适合我胃口，用过后爱不释手；
2. go 语言团队太过耀眼和强大：Thompson 图灵奖获得者，unix 和 C 的共同发明人；Pike PLAN9 操作系统的主要开发者、UTF-8 发明者；Robert Griesemer 参与 java 的 HotSpot, js v8 引擎开发者；
3. 国内传道者的极力推荐：许式伟兄，谢孟军兄等强力推荐及相关书籍问世；

### 学习资料

书籍是：老许的《go 语言编程》、 老谢的《go web 编程》、 雨痕的《go 语言学习笔记》、 golang.org 上面的《Effective Go》、《The Go Programming Language Specification》、go 标准库和 github 上众多开源库

当然无闻的视频教程也是非常适合初学者的，跟着他一起把代码敲一遍事半功倍

### go 学习体会

go 语言基础知识非常简单，简单到几天就可以学完，并能够上手开发；但是要做到精通，没有一定的代码量和几年的经验很难达到，这是学习任何一门编程语言都必须要经历的，你唯一能做的是：不停的写代码，不停的思考，不停的总结，不停的读别人代码，向高手请教；

### go 学习难点

将我在学习中遇到的难点，以及相关参考资料索引出来，这些知识点对新入门的学习者有点难，但是对于想全面掌握 go 技能的开发人员来说，我认为是非常有价值的，这些知识点都是个人一步步学习趟过坑之后去发掘的：

1. go map slice string array interface 底层数据模型，其中 array 和 slice 是引起混乱的根源；参见：Russ Cox 非常经典文章

   [Go Data Structures](https://research.swtch.com/godata)

   [Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices)

2. go defer panic recover 是 go 特有的，go 的错误就是错误，异常就是异常，没有混为一谈，这个设计表面上看代码烦乱了，实际上更清晰；try catch 表面上是清晰，实际上隐藏着太多的问题。（两种错误处理机制：异常抛出机制-错误码处理机制）

   [使用 Defer 的几个场景](./)

3. go interface 接口的底层实现机制（能深入到源码）（深入才能理解：接口赋值，接口转换，接口断言及 go 的动态性）；go 就是一门面向接口、面向组合编程的语言，对于 go 语言来说接口是灵魂一点都不为过；

   参见：老许《Go 语言编程》 第 9 章 9.5 节 [接口机理](./)

   参加：国外一位大佬写的：[How to use interfaces in Go](http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)

   参见：[interfaces_and_types](https://golang.org/doc/effective_go.html#interfaces_and_types)

   参见：[Learning Go - Interfaces & Reflections](http://www.laktek.com/2012/02/13/learning-go-interfaces-reflections/)

4. go 类型系统：这个对掌握一个语言非常重要：static type（语言层面就是静态类型语言）, dynamic type（动态类型针对接口而言） ,underly type(底层类型针对强制转换和赋值）；go 的类型系统比较简练，干净，都是值类型，即使像 slice,map,channel,func 等看起来是引用类型，其实也是值类型，只是其内部数据结构封装了指针罢了；不像 JAVA 有两套类型：基本的是值类型，对象是引用类型，中间又有装箱，拆箱动作，更有奇葩的 string；

   参见：[The Go Programming Language Specification-type](https://golang.org/ref/spec#Types)

   参见：[Learning Go - Types](http://www.laktek.com/2012/01/27/learning-go-types/)

5. go function :多值返回；带命名的返回参数用法；闭包；函数是一等公民；高阶函数；函数也是一种基础类型，可以 type xxx func 为函数新定义一种类型；并发 go 以 h 函数为载体；对象是附有行为的数据，而闭包是附有数据的行为

   参见国外大牛：[Function Types in Go (golang)](http://jordanorelli.com/post/42369331748/function-types-in-go-golang)

   go 闭包：[函数编程之闭包漫谈(Closure)](http://www.cnblogs.com/Jifangliang/archive/2008/08/05/1260602.html) [Go 语言(Golang) - 闭包](http://blog.sina.com.cn/s/blog_487109d101018fcx.html)

   go 函数式编程：[go 函数式编程实践](http://blog.sina.com.cn/s/blog_6e1bd8350101k4r3.html)

6. go 参数传递：函数参数全部是传值：即使传递的是指针，传递的也是指针的拷贝；闭包引用外部变量是引用

   所谓引用是指使用的不是指针，但是却有指针的效果,引用：a 做为参数传递函数内部，函数内部修改 a 却改变了外部 a 的值

   指针：*a 作为参数传递到函数内部，函数内部修改了*a 的值，外部 a 指向的值也发生改变；

   参见：[Go 语言的传参和传引用](./) 这篇文章分析的非常到位时难得好文章

7. go error 处理机制,error 与 nil 关系，参见[Go 中 error 类型的 nil 值和 nil](./)

   国外这两篇文章写得也比较好，教你如何自定义 error 以返回更具体的错误；

   [Error Handling In Go, Part I](https://www.ardanlabs.com/blog/2014/10/error-handling-in-go-part-i.html)

   [Error Handling In Go, Part II](https://www.ardanlabs.com/blog/2014/11/error-handling-in-go-part-ii.html)

8. go nil 也一直是困扰 go 学习者的一个绊脚石

   上面有篇 error nil 相关的文章

   这里有陈一回的 Interface nil 的文章 [golang: 详解 interface 和 nil](https://my.oschina.net/goal/blog/194233) 陈兄关于 go 的几篇文章都非常接地气，建议都看了

9. go package、全局 const(常量）、var（变量）加载顺序，及 package 引用机制：

   参见：老许翻译的那本《Go 语言程序设计》

10. go reflect ：反射是一个强大的武器，是一个新手成为老手的必须涉猎的东西，也是元编程的一种方式，其性能相对来说不太高（灵活性也带来性能问题）

    参见：官方版 [laws-of-reflection](https://blog.golang.org/laws-of-reflection)

    翻译版的：[The-laws-of-reflection](./) 我认为这是翻译的最好的一篇文章，融入了作者的思考和感悟

    还有我的简写版的，更多的是个 API 的指引：[The laws of reflection](https://blog.csdn.net/hittata/article/details/26159959)

    反射小试身手：参考这篇在 [GOLANG 中用名字调用函数](https://mikespook.com/2012/07/%E5%9C%A8-golang-%E4%B8%AD%E7%94%A8%E5%90%8D%E5%AD%97%E8%B0%83%E7%94%A8%E5%87%BD%E6%95%B0/) Mikespook 翻译的其他文章也非常棒

    martini 框架使用的经典 DI 库：[inject](https://github.com/codegangsta/inject) 教科书办实现注入

    inject 库代码非常晦涩，可以参考陈兄的这边经典文章 golang: [Martini 之 inject 源码分析](https://my.oschina.net/goal/blog/195036) 这篇文章深入浅出，写的非常好

    反射与接口、go 类型系统关系非常密切，refelct.TypeOf 返回的是接口的 dynamic type , func (v Value) Type() Type 返回 value 的 type, reflect.Type 表述的是 underlying type

11. go channel 和 gorutone 使用方法；并发编程模式；实现原理. 如果形容 go 语言是一个皇冠，goroutine 和 channel 则是皇冠上的明珠，go 是 CSB 并发编程理论的实践者和改良派；go 将并发编程的复杂性降低了一个数量级，从此并发编程不再是某些自称“技术淫人”的专利；世界本应该如此简单，语言和编译器能做好的事情就不要让程序员瞎忙活，这是 go 的哲学；

    相关参考：

    - 参见《go 并发编程》(说实在的，这本书写的没有达到我期望的水准，写的都是基础，语言有点啰嗦，挖的够深，但是拔的不高，所谓拔的不高就是没有系统的介绍并发设计模式）

    - 参见[goroutine 背后的系统知识](https://my.oschina.net/chai2010/blog/119867)

    - go 内存模型 [英文版](https://golang.org/ref/mem) 中文版 英文需要翻墙，有点晦涩难懂，中文的这一篇写的非常好，融入译者自己的思考，推荐阅读；

    - [go 语言并发之美](https://studygolang.com/articles/1506) 这篇文章写得非常好，对常用的并发模式写的深入浅出，这点《go 并发编程》一本很厚的书居然没有这些内容，实在让人失望，有点徒有虚名

    - Google IO 大会上大牛的几篇文章（有墙）
      - [Concurrency is not Parallelism](https://talks.golang.org/2012/waza.slide#1) 这里有篇翻译的并发不是并行
      - [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#1)
      - [Advanced Go Concurrency Patterns](https://talks.golang.org/2013/advconc.slide#1)
      - [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines) 这里有篇翻译：[Go 并发模式：管道和显式取消](http://blog.jobbole.com/65552/)

12. go 调度器模型，go 内存管理，GC，go 调试与性能分析，跨平台；这些都是高级命题，当程序遇到性能问题时，你可能需要去了解 go 并发的实现机制，找到问题点，规避或者改进或者找到替代方案；除非你有强烈的好奇心，一般的程序员不会去读整个运行时的实现。

    参见[雨痕学习笔记](https://github.com/qyuhen/book)

    [The Go scheduler](http://morsmachine.dk/go-scheduler) 这里有篇翻译 go 调度器

    阿里 skoo 的几篇文章 对调度器过程写的非常生动

    《go 并发编程》对 go 调度也做了论述

    国外的这篇 PPT 写的也不错 ，还有一篇

    github 上有本[电子书](https://github.com/tiancaiamao/go-internals/) 写的非常深入，对 go 的几个关键点实现进行深入剖析，非常好的文章

### 框架学习

学习了谢大的[beego](https://beego.me/)框架，beego 非常容易入门，模块化设计，并且模块非常齐全；谢大人比较热情，QQ 群较活跃,，我的两个小系统都是基于 beego 开发的；[小黑的这篇导读](https://my.oschina.net/fuxiaohei/blog/228999#OSC_h3_3)对于想看框架源码的人来说是个福音；

Martini 只是看了 inject 那部分，2015 年希望有时间细看一下 martini 和 revel。

### 期待

期待有个牛人能出一本专门介绍 go 如何设计大系统的书，go 语言设计模式和面向对象设计模式有很大差别；老许有一篇 PPT 里面介绍 go 的面向连接和组合的语言，以七牛系统的规模，应该可以抽象出一套模式出来，有人做吗？并发的相关设计模式，网上有多文章，但是还不是很系统；希望 2015 有人能站出来做这件事情，我们好站在巨人的肩膀上继续前行。

其实不是为了学习而学习： 我很大一部分时间还是边开发，边了解标准库，边学习；遇到比较大的通用的模块到 github 上找有无已经实现的，如有借鉴过来吧，如果你认为自己改写的比原作者好，可以 pull request. 当然在开发中遇到自己知识的盲点，就需要有股专研的精神，把它搞明白，技术也就自然得到提高，个人薄见，谨慎参考。

---

> 注：原文地址为:

- [http://www.jb51.net/article/62728.htm](http://www.jb51.net/article/62728.htm)
- [https://segmentfault.com/a/1190000000654351](https://segmentfault.com/a/1190000000654351)
- [Advise from Go developers to Go programming newbies](http://www.gophercon.in/blog/2014/08/23/adviceforgonewbies/)
