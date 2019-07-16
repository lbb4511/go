package main

// #include "./hello.h"
import "C"

func main() {
	C.myPrint(C.CString("你好，欢迎！This is Cgo study."))
}