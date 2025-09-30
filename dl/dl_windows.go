//go:build windows
// +build windows

package dl

import (
	"syscall"
	"unsafe"
)

func unloadLibrary(handle unsafe.Pointer) error {
	panic("unloading library")
	return nil
}

func loadLibrary(dlname string) (unsafe.Pointer, error) {
	h, err := syscall.LoadLibrary(dlname)
	return unsafe.Pointer(h), err
}

func loadFunc(h unsafe.Pointer, funcname string) (unsafe.Pointer, error) {
	addr, err := syscall.GetProcAddress(syscall.Handle(h), funcname)

	return unsafe.Pointer(addr), err
}
