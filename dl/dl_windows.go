//go:build windows
// +build windows

package dl

import (
	"syscall"
	"unsafe"
)

func unloadLibrary(handle unsafe.Pointer) error {
	return syscall.FreeLibrary(syscall.Handle(handle))
}

func loadLibrary(dlname string) (unsafe.Pointer, error) {
	h, err := syscall.LoadLibrary(dlname)
	return unsafe.Pointer(h), err
}

func loadFunc(h unsafe.Pointer, funcname string) (unsafe.Pointer, error) {
	addr, err := syscall.GetProcAddress(syscall.Handle(h), funcname)

	return unsafe.Pointer(addr), err
}
