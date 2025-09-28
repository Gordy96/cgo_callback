package main

import (
	"cgo/adapter"
	"cgo/dl"
	"fmt"
)

type FakePort struct{}

func (f *FakePort) Name() string {
	return "port"
}

func (f *FakePort) Read(b []byte) (n int, err error) {
	const hw = "hello world"
	copy(b, hw)

	return len(hw), nil
}

func (f *FakePort) Write(b []byte) (n int, err error) {
	fmt.Printf("fake port received: %s\n", b)
	return len(b), nil
}

func main() {
	// Load the .so dynamically
	lib, err := dl.Open("./plugin/plugin.so")
	if err != nil {
		panic(err)
	}
	defer lib.Release()

	err = adapter.InitLib(lib)
	if err != nil {
		panic(err)
	}

	a, err := adapter.New("runner", []adapter.Port{&FakePort{}}, lib)
	if err != nil {
		panic(err)
	}
	err = a.Init()
	if err != nil {
		panic(err)
	}

	defer a.Close()

	err = a.TriggerPinInterrupt(2)
	if err != nil {
		panic(err)
	}
	err = a.TriggerPinInterrupt(2)
	if err != nil {
		panic(err)
	}
	err = a.TriggerPinInterrupt(2)
	if err != nil {
		panic(err)
	}
}
