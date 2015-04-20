package mpc

/*
#cgo LDFLAGS: -lm
#include "mpc_interface.h"
*/
import "C"

import (
	"unsafe"
)

type Parser *C.mpc_parser_t

func New(name string) Parser {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.mpc_new(cName)
}

func Cleanup(p Parser) {
	C.mpc_cleanup_if(p)
}
