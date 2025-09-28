//go:build windows
// +build windows

package dl

import (
	"syscall"
)

func unloadLibrary(handle uintptr) error {
	panic("unloading library")
	return nil
}

func loadLibrary(dlname string) (uintptr, error) {
	h, err := syscall.LoadLibrary(dlname)
	return uintptr(h), err
}

func loadFunc(h uintptr, funcname string) (uintptr, error) {
	addr, err := syscall.GetProcAddress(syscall.Handle(h), funcname)

	return uintptr(addr), err
}
