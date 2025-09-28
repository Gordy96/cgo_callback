package dl

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
