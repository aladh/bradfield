package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(strlen("hello"))
	fmt.Println(strlen2("hello"))
}

func strlen(s string) int {
	const PtrSize = unsafe.Sizeof(uintptr(0))

	// Get len field of string
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + PtrSize))
}

func strlen2(s string) int {
	type stringType struct {
		ptr unsafe.Pointer
		len int
	}

	stringStruct := *(*stringType)(unsafe.Pointer(&s))
	return stringStruct.len
}
