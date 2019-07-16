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