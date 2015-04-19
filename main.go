package main

/*
#cgo LDFLAGS: -lm
#include "mpc.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func mpcNew(name string) *C.mpc_parser_t {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.mpc_new(cName)
}

func main() {
	var number *C.mpc_parser_t
	fmt.Println(number)
	number = mpcNew("number")
	fmt.Println(number)
}
