package main

import "C"
import "fmt"

func myPrint(s *C.char) {
  fmt.Printf("myPrint: %s\n", C.GoString(s));
}