package main

// #include <stdio.h>
import "C"

func main() {
	C.puts(C.CString("你好，欢迎！This is Cgo study."))
}
