//go:build !windows
// +build !windows

package dl

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func unloadLibrary(h unsafe.Pointer) error {
	C.dlclose(h)
	return nil
}

func loadLibrary(dlname string) (unsafe.Pointer, error) {
	s := C.CString(dlname)
	defer C.free(unsafe.Pointer(s))
	h := C.dlopen(s, C.RTLD_NOW)
	if h == nil {
		return nil, fmt.Errorf("failed to load dynamic library")
	}
	return h, nil
}

func loadFunc(h unsafe.Pointer, funcname string) (unsafe.Pointer, error) {
	n := C.CString(funcname)
	defer C.free(unsafe.Pointer(n))
	fp := C.dlsym(h, n)
	if fp == nil {
		return nil, fmt.Errorf("dynamic library does not have symbol %v", funcname)
	}

	return fp, nil
}
