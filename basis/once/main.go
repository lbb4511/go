package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

// sync.Once.Do(f func())是一个挺有趣的东西,能保证once只执行一次，无论你是否更换once.Do(xx)这里的方法,这个sync.Once块只会执行一次。
// 整个程序，只会执行onces()函数一次,onced()函数是不会被执行的。
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(func() { fmt.Println("onced") })
			fmt.Println("213")
		}()
	}
	for i, v := range make([]string, 10) {
		once.Do(func() { fmt.Println("onces") })
		fmt.Println("count:", v, "---", i)
	}
	time.Sleep(10)
}
