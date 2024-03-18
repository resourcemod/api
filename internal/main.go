package main

import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}

//export go_register
func go_register(name *C.char, fn unsafe.Pointer) {
	fmt.Println(C.GoString(name))
}
