package dl

/*
#include <stdlib.h>

typedef int (*func_t)(char *buf);

int trampoline(func_t cb, char* buf) {
	return cb(buf);
}
*/
import "C"
import "unsafe"

type SO struct {
	handle unsafe.Pointer
}

func Open(path string) (*SO, error) {
	h, err := loadLibrary(path)
	if err != nil {
		return nil, err
	}
	return &SO{h}, nil
}

func (s *SO) Release() error {
	return unloadLibrary(s.handle)
}

func (s *SO) Func(name string) (unsafe.Pointer, error) {
	return loadFunc(s.handle, name)
}

func testFunc(ptr unsafe.Pointer, buf []byte) int {
	cb := (C.func_t)(ptr)
	return int(C.trampoline(cb, (*C.char)(unsafe.Pointer(&buf[0]))))
}
