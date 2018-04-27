# 并发编程
###### 一个Go例程(Goroutines)就是一个和其它Go例程在同一地址空间里但却独立运行的函数，就像是在shell里使用 & 标记启动一个命令。
### Goroutine 的并发安全性  
无论是 Windows 还是 Linux，基本上来说是用操作系统的线程来实现的，不过，Goroutine不是线程。它有个特性，如果当一个Go例程阻塞时，所在的线程会阻塞，但其它Go例程不受影响，多个例程可以在系统线程上做多路通信。堆栈初始很小，但随着需求会增长或收缩，Go例程不是不耗资源，但它们很轻量级的。
###### 这并不是真正的并发，如果你要真正的并发，你需要在你的 main 函数的第一行加上下面的这段代码：
```golang
import "runtime"
...
NCPU := runtime.NumCPU()
runtime.GOMAXPROCS(NCPU)
```
### 并发(Concurrency)不是并行(Parallelism)  
*--Rob Pike*
###### Go语言提供：
- 并发执行(goroutines)
    - 并发是一种将一个程序分解成小片段独立执行的程序设计方法，通信是指各个独立的执行任务间的合作。
        - `并发`将相互独立的执行过程综合到一起的编程技术。(这里是指通常意义上的执行过程，而不是Linux进程。很难定义。)
        - `并行`同时执行(通常是相关的)计算任务的编程技术。
    - 并发是指同时*处理*很多事情，而并行是指同时*能完成*很多事情，两者不同，但相关，一个重点是组合，一个重点是执行。
    - 并发提供了一种方式让我们能够设计一种方案将问题*(非必须的)*并行的解决。
- 同步和消息传输(channels)
- 多路并发控制(select)

### 通道 Channels
###### 通道是类型化的值，能够被Go例程用来做同步或交互信息。
```golang
timerChan := make(chan time.Time)
go func() {
    time.Sleep(deltaT)
    timerChan <- time.Now() // send time on timerChan
}()
// Do something else; when ready, receive.
// Receive will block until timerChan delivers.
// Value sent is other goroutine's completion time.
completedAt := <-timerChan
```
### Select
###### 这select语句很像switch，但它的判断条件是基于通信，而不是基于值的等量匹配。
```golang
select {
case v := <-ch1:
    fmt.Println("channel 1 sends", v)
case v := <-ch2:
    fmt.Println("channel 2 sends", v)
default: // optional
    fmt.Println("neither channel was ready")
}
```
### 闭包
###### 它让一些并发运算更容易表达：
```golang
func Compose(f, g func(x float) float)
                  func(x float) float {
     return func(x float) float {
        return f(g(x))
    }
}
print(Compose(sin, cos)(0.5))
```
### Demo
```golang
// 使用闭包封装一个后台操作：
go func() { // 从输入通道拷贝数据到输出通道
    for val := range input {
        output <- val
    }
}()//这个for range操作会一直执行到处理掉通道内最后一个值。

// 数据类型：
type Work struct {
    x, y, z int
}

// 一个worker的任务：
func worker(in <-chan *Work, out chan<- *Work) {
   for w := range in {
      w.z = w.x * w.y
      Sleep(w.z)
      out <- w
   }
}//必须保证当一个worker阻塞时其他worker仍能运行。

// runner:
func Run() {
   in, out := make(chan *Work), make(chan *Work)
   for i := 0; i < NumWorkers; i++ {
       go worker(in, out)
   }
   go sendLotsOfWork(in)
   receiveLotsOfResults(out)
} // 很简单的任务，但如果没有并发机制，你仍然很难这么简单的解决。
```
这个负载均衡的例子具有很明显的并行和可扩展性，Worker数可以非常巨大。Go语言的这种并发特征能的开发一个安全的、好用的、可扩展的、并行的软件变得很容易。
并发简化了同步，没有明显的需要同步的操作，程序的这种设计隐含的实现了同步。
### Demo(查询数据库)
```golang
func Query(conns []Conn, query string) Result {
    ch := make(chan Result, len(conns))  // buffered
    for _, conn := range conns {
        go func(c Conn) {
            ch <- c.DoQuery(query):
        }(conn)
    }
    return <-ch
}
```
并发和垃圾回收机制让这成为一个很小很容易解决的问题。



- [GO并发编程实战](http://ifeve.com/go-concurrent-waitgroup/)
- [Go语言并发之美](https://studygolang.com/articles/1506)