package main


// #include "./hello.h"
import "C"

func main() {
	myPrint(C.CString("你好，欢迎！This is Cgo study."))
}