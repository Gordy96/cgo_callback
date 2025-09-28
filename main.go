package main

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdlib.h>
#include <stdio.h>
#include "plugin/plugin.h"

// forward declaration for trampoline
extern int goRead(char* port, char* buf);
extern int goWrite(char* port, char* buf, int size);

// trampoline wrapper
static void tInit(const char *port, init_func_t cb) {
    cb(port, goRead, goWrite);
}
static void tRun(const char *port, run_func_t r) {
	r(port);
}
static int test(char* port, char* buf) {
	return sprintf(buf, "hello, %s", port);
}
*/
import "C"

import (
	"cgo/dl"
	"fmt"
	"unsafe"
)

//export goRead
func goRead(port *C.char, buf *C.char) C.int {
	return C.test(port, buf)
}

//export goWrite
func goWrite(port *C.char, buf *C.char, size C.int) C.int {
	from := C.GoString(port)
	b := C.GoBytes(unsafe.Pointer(buf), size)
	s := string(b)

	fmt.Printf("%s wrote \"%s\"\n", from, s)

	return size
}

func main() {
	// Load the .so dynamically
	lib, err := dl.Open("./plugin/plugin.so")
	if err != nil {
		panic(err)
	}
	defer lib.Release()

	sym, err := lib.Func("init")
	if err != nil {
		panic(err)
	}

	port := C.CString("runner")

	// Cast symbol to proper function pointer
	initFunc := (C.init_func_t)(sym)

	C.tInit(port, initFunc)

	// Resolve the symbol
	sym, err = lib.Func("run")
	if err != nil {
		panic(err)
	}

	runFunc := (C.run_func_t)(sym)
	C.tRun(port, runFunc)
}
