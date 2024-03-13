package api

import "C"
import "fmt"

//export go_register
func go_register(name *C.char) {
	fmt.Println(C.GoString(name))
}
