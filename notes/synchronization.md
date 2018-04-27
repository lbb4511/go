# 同步(Synchronization)

### 初始化
程序的初始化在一个独立的goroutine中执行。在初始化过程中创建的goroutine将在 第一个用于初始化goroutine执行完成后启动。</br>
如果包p导入了包q，包q的init初始化函数将在包p的初始化之前执行。</br>
程序的入口函数 main.main 则是在所有的 init 函数执行完成 之后启动。</br>
*在任意init函数中新创建的goroutines，将在所有的init 函数完成后执行。*
### goroutine的创建
###### 用于启动goroutine的go语句在goroutine之前运行。
```
var a string

func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	go f()
}
```
>调用hello函数，会在某个时刻打印“hello, world”（有可能是在hello函数返回之后）。
### Channel communication 管道通信
用管道通信是两个goroutines之间同步的主要方法。在管道上执行的发送操作会关联到该管道的 接收操作，这通常对应goroutines。</br>
管道上的发送操作发生在管道的接收完成之前（happens before）。
```
var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
```
可以确保会输出"hello, world"。因为，a的赋值发生在向管道 c发送数据之前，而管道的发送操作在管道接收完成之前发生。 因此，在print 的时候，a已经被赋值。</br>
从一个unbuffered管道接收数据在向管道发送数据完成之前发送。
```
package main

var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}

func main() {
	go f()
	c <- 0
	print(a)
}
```
同样可以确保输出“hello, world”。因为，a的赋值在从管道接收数据 前发生，而从管道接收数据操作在向unbuffered 管道发送完成之前发生。所以，在print 的时候，a已经被赋值。</br>
如果用的是缓冲管道（如 c = make(chan int, 1) ），将不能保证输出 “hello, world”结果（可能会是空字符串， 但肯定不会是他未知的字符串， 或导致程序崩溃）。
### 锁
包sync实现了两种类型的锁： sync.Mutex 和 sync.RWMutex。</br>
对于任意 sync.Mutex 或 sync.RWMutex 变量l。 如果 n < m ，那么第n次 l.Unlock() 调用在第 m次 l.Lock() 调用返回前发生。
```
var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}
```
可以确保输出“hello, world”结果。因为，第一次 l.Unlock() 调用（在f函数中）在第二次 l.Lock() 调用 （在main 函数中）返回之前发生，也就是在 print 函数调用之前发生。</br>
对于任何呼叫到一个sync.RWMutex变数l l.RLock，有一个n使得l.RLock第n个呼叫l.Unlock后发生（返回）和n +1'th之前的匹配l.RUnlock发生调用l.Lock。</br>
包sync提供了一个在多个goroutines中进行初始化的方法。多个goroutines可以 通过 once.Do(f) 方式调用f函数。 但是，f函数 只会被执行一次，其他的调用将被阻塞直到唯一执行的f()返回。
once.Do(f) 中唯一执行的f()发生在所有的 once.Do(f) 返回之前。
```
var once sync.Once
var a string

func setup() {
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```
> 调用twoprint会输出“hello, world”两次。第一次twoprint 函数会运行setup唯一一次。


## ***错误的同步方式***
注意：变量读操作虽然可以侦测到变量的写操作，但是并不能保证对变量的读操作就一定发生在写操作之后。
```
package main

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}
```
函数g可能输出2，也可能输出0。

这种情形使得我们必须回避一些看似合理的用法。

这里用重复检测的方法来代替同步。在例子中，twoprint函数可能得到错误的值：
```
package main

import (
	"sync"
	"time"
)

var once sync.Once

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	twoprint()
	time.Sleep(8000)
}
```
在doprint函数中，写done暗示已经给a赋值了。 但是没有办法给出保证，函数可能输出空的值（在2个goroutines中同时执行到测试语句）。

另一个错误陷阱是忙等待：
```
package main

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a)
}
```
我们没有办法保证在main中看到了done值被修改的同时也 能看到a被修改，因此程序可能输出空字符串。 更坏的结果是，main 函数可能永远不知道done被修改，因为在两个线程之间没有同步操作，这样main 函数永远不能返回。

下面的用法本质上也是同样的问题.
```
package main

type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func main() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}
```
即使main观察到了 g != nil 条件并且退出了循环，但是任何然 不能保证它看到了g.msg的初始化之后的结果。

在这些例子中，只有一种解决方法：用显示的同步。
