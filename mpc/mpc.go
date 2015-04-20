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

func Cleanup(p1, p2, p3, p4 Parser) {
	C.mpc_cleanup_if(p1, p2, p3, p4)
}

func MpcaLang(language string, parsers ...Parser) {
	cLanguage := C.CString(language)
	defer C.free(unsafe.Pointer(cLanguage))
	C.mpca_lang_if(C.MPCA_LANG_DEFAULT, cLanguage,
		parsers[0], parsers[1], parsers[2], parsers[3])
}
