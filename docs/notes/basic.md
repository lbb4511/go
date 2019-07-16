# 语句

## if

```go
package main

import "fmt"

func main() {
  var number int = 5
  if number += 4; 10 > number {
    var number int
    number += 3
    fmt.Print(number)
  } else if 10 < number {
    number -= 2
    fmt.Print(number)
  }
  fmt.Println(number)
}
```

## switch

```go
package main

import (
    "fmt"
  "math/rand"
)

func main() {
  ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
  switch v := ia[rand.Intn(4) %2 ]; interface{}(v).(type) {
  case int32 :
    fmt.Printf("Case A.")
  case byte :
    fmt.Printf("Case B.")
  default:
    fmt.Println("Unknown!")
  }
}
```

## for

```go
package main

import (
    "fmt"
)

func main() {
  map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
    for _, v := range []int{1, 2, 3, 4} {
        fmt.Printf("%d:%s\n",v,map1[v])
    }
}
```

## select

```go
package main

import "fmt"

func main() {
    ch4 := make(chan int, 1)
  for i := 0; i < 4; i++ {
    select {
    case e, ok := <-ch4:
      if !ok {
        fmt.Println("End.")
        return
      }
      fmt.Println(e)
      close(ch4)
    default:
      fmt.Println("No Data!")
      ch4 <- 1
    }
  }
}
```

## defer

```go
package main

import (
    "fmt"
)

func main() {
  for i := 0; i < 10; i++ {
    fmt.Printf("%d ", fibonacci(i))
    defer fmt.Printf("%d ", fibonacci(i))
  }
}

func fibonacci(num int) int {
  if num == 0 {
    return 0
  }
  if num < 2 {
    return 1
  }
  return fibonacci(num-1) + fibonacci(num-2)
}
```

## go

```go
package main

import (
  "fmt"
)

func main() {
  ch1 := make(chan int, 1)
  ch2 := make(chan int, 1)
  ch3 := make(chan int, 3)
  go func() {
    fmt.Println("1")
    ch1 <- 1
  }()
  go func() {
    <-ch1
    fmt.Println("2")
    ch2 <- 2
  }()
  go func() {
    <-ch2
    fmt.Println("3")
    ch3 <- 3
  }()
  <-ch3
}

```
