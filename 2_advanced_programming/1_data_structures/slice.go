package main

import (
	"fmt"
	"unsafe"
)

func main() {
	ints := []int{1, 2, 3, 4}

	fmt.Println(sum(ints))
	fmt.Println(sum2(ints))
}

func sum(intSlice []int) int {
	const PtrSize = unsafe.Sizeof(uintptr(0))
	const IntSize = unsafe.Sizeof(0)

	// Get pointer to underlying array
	arrPtr := unsafe.Pointer(*(**int)(unsafe.Pointer(&intSlice)))
	// Get len field of slice
	length := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&intSlice)) + PtrSize))

	acc := 0

	for i := 0; i < length; i++ {
		acc += *(*int)(unsafe.Pointer((uintptr(arrPtr)) + uintptr(i*int(IntSize))))
	}

	return acc
}

func sum2(intSlice []int) int {
	const IntSize = unsafe.Sizeof(0)

	type sliceType struct {
		ptr unsafe.Pointer
		len int
		cap int
	}

	sliceStruct := *(*sliceType)(unsafe.Pointer(&intSlice))

	acc := 0

	for i := 0; i < sliceStruct.len; i++ {
		acc += *(*int)(unsafe.Pointer(uintptr(sliceStruct.ptr) + uintptr(int(IntSize)*i)))
	}

	return acc
}
