# cgo 简单笔记

## 基于 CGO 的 Hello, 世界

```go
package main

// #include <stdio.h>
import "C"

func main() {
  C.puts(C.CString("你好，欢迎！This is Cgo study."))
}
```

其中 `import "C"` 前注释中可以导入 C 语言 函数/变量/宏 等到一个虚拟的 "C" 包中。
我们可以直接使用 C.前缀访问 C 中的函数。

当前的例子中，使用 C 的 puts 函数输出"你好，欢迎！This is Cgo study."。之所以没有使用 C 语言经典的 printf 函数，
是因为 printf 的参数是可变的，而 cgo 不支持访问你可变参数的函数。

如果需要使用 printf 函数，需要再做一次包装。

```go
package main

/*
#include <stdio.h>

static void myPrint(const char* msg) {
  printf("myPrint: %s\n", msg);
}
*/
import "C"

func main() {
  C.puts(C.CString("puts: 你好，欢迎！This is Cgo study."))
  // C.printf(C.CString("你好，欢迎！This is Cgo study.")) // error
  C.myPrint(C.CString("你好，欢迎！This is Cgo study."))
}
```

我们通过 myPrint 来屏蔽 printf 的可变参数特性。在 import "C" 定义的函数一般建议定义为 static。

### C 代码模块化

模块化是一种重要的编程方法。当程序中的一行语句太长的时候，我们希望将代码拆分为多行；当代码语句多到一定层度，我们就会将代码拆分为函数；如果一个文件中的函数太多，则希望将函数拆分到多个文件中重新组织。以上这些都是采用模块化的思路来简化代码的组织。

对于前面的 myPrint 函数，我们也可以采用模块化的测试来重新组织。首选创建一个 hello.h 头文件，里面包含 myPrint 函数的声明。然后将 myPrint 函数的实现放到 hello.c 文件中。在 CGO 代码中，就可以通过 #include "hello.h" 的方式直接引用 myPrint 函数。

hello.h

```c
extern void myPrint(const char *s);
```

hello.c

```c
#include "hello.h"

#include <stdio.h>

void myPrint(const char *msg)
{
  printf("myPrint: %s\n", msg);
}
```

main.go

```go
package main

// #include "./hello.h"
import "C"

func main() {
  C.myPrint(C.CString("你好，欢迎！This is Cgo study."))
}
```

创建 hello.h 头文件是模块化编程的一个重要里程碑。对于 SayHello 函数的用户来说，我们只需要知道 SayHello 函数满足 C 语言函数的调用规约即可。至于 SayHello 函数是采用 C 语言或 C++ 语言、甚至是其它任何语言实现的，对于 SayHello 函数的用户并没有区别。因此，我们可以该 Go 语言重新实现 SayHello 函数。

hello.h

```c
extern void myPrint(/* const */ char *s);
```

hello.go

```go
package main

import "C"
import "fmt"

func myPrint(s *C.char) {
  fmt.Printf("myPrint: %s\n", C.GoString(s));
}
```

main.go

```go
package main

// #include "./hello.h"
import "C"

func main() {
  C.myPrint(C.CString("你好，欢迎！This is Cgo study."))
}
```

hello.h 头文件包含 myPrint 函数声明，但是 hello.c 变成了 hello.go，函数本身从 C 语言实现改成了用 Go 语言实现。Go 语言实现的函数和 C 语言版本的函数名字和参数类型几乎是完全一致的（Go 导出的 C 函数不支持 const 修饰），因此对于 myPrint 函数的用户来说并没有太多的差异。现在可以说我是采用 C 语言思维编程的 Go 语言码农。

### 整合

在模块化的基础上，我们采用 Go 重新实现了 C 语言规格的 myPrint 函数。现在我们可以尝试打破模块化编程的思路：删除 hello.h 头文件，将全部的 CGO 代码统一到一个 Go 源文件中：

hello.go

```go
package main

import "C"
import "fmt"

func main() {
  myPrint(C.CString("你好，欢迎！This is Cgo study."))
}

func myPrint(s *C.char) {
  fmt.Printf("myPrint: %s\n", C.GoString(s));
}
```

这时候虽然没有了头文件的函数声明，但是 myPrint 函数的声明在我们 Go 语言码农的心中。我们通过 extern 的方式在 CGO 中手工声明 myPrint 函数。然后在 main 函数中调用一个目前还不真是存在的 myPrint 函数进行字符串输出。这个例子其实 90%以上是 Go 语言代码，但是编程的思维是 C 语言。

在导出 myPrint 函数时，依然采用了 C 语言的字符串格式。为此，在输出 Go 语言字符串时，需要先转换为 C 语言格式的字符串；然后在 Go 语言中输出 C 语言字符串时，有需要转回 Go 语言字符串；最后还需要释放中间临时创建的 C 语言字符串。这是心中编程思维被 C 语言字符串固化的结果。我们需要忘掉 C 语言原有的字符串结构：其实从 C 语言角度看来，Go 语言字符串也是一种特殊格式的字符串。

---

<https://blog.csdn.net/u010884123/article/details/60872980>
<https://blog.csdn.net/RA681t58CJxsgCkJ31/article/details/80504482>
