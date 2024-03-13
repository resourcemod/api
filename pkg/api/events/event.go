package events

import "C"
import "fmt"

//export say
func say(text *C.char) {
	fmt.Println(C.GoString(text))
}
