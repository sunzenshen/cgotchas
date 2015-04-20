package mpc

/*
#cgo LDFLAGS: -lm
#include "mpc_interface.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

type Ast *C.mpc_ast_t
type Error *C.mpc_err_t
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

func Parse(input string, parser Parser) (C.mpc_result_t, error) {
	var r C.mpc_result_t
	cInput := C.CString(input)
	defer C.free(unsafe.Pointer(cInput))
	stdin := C.CString("<stdin>")
	defer C.free(unsafe.Pointer(stdin))
	var err error
	if C.mpc_parse(stdin, cInput, parser, &r) == C.int(0) {
		err = errors.New("mpc: failed to parse input string")
	}
	return r, err
}

func GetOutput(result *C.mpc_result_t) Ast {
	cast := (*Ast)(unsafe.Pointer(result))
	return *cast
}

func GetError(result *C.mpc_result_t) Error {
	cast := (*Error)(unsafe.Pointer(result))
	return *cast
}

func ErrDelete(result *C.mpc_result_t) {
	C.mpc_err_delete(GetError(result))
}

func ErrPrint(result *C.mpc_result_t) {
	C.mpc_err_print(GetError(result))
}

func AstDelete(result *C.mpc_result_t) {
	C.mpc_ast_delete(GetOutput(result))
}

// PrintAst parses an input string for its AST representation
func PrintAst(input string, parser Parser) {
	r, err := Parse(input, parser)
	if err != nil {
		ErrPrint(&r)
		ErrDelete(&r)
	} else {
		C.mpc_ast_print(GetOutput(&r))
		AstDelete(&r)
	}
}
