package main

// extern void myPrint(char *s);
import "C"
import "fmt"

func main() {
	myPrint(C.CString("你好，欢迎！This is Cgo study."))
}

func myPrint(s *C.char) {
  fmt.Printf("myPrint: %s\n", C.GoString(s));
}