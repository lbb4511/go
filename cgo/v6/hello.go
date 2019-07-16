package main

// extern void myPrint(_GoString_ s);
import "C"
import "fmt"

func main() {
	C.myPrint("你好，欢迎！This is Cgo study.")
}

// export myPrint
func myPrint(s string) {
  fmt.Printf("myPrint: %s\n", s);
}