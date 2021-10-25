package main

import (
	"fmt"
	"unsafe"
)

const IntSize = unsafe.Sizeof(0)
const PtrSize = unsafe.Sizeof(uintptr(0))

func main() {
	fmt.Println(strlen("hello"))
	fmt.Println(getY(Point{x: 1, y: 2}))
	fmt.Println(sum([]int{1, 2, 3, 4}))
}

func strlen(s string) int {
	// Get len field of string
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + PtrSize))
}

type Point struct{ x, y int }

func getY(p Point) int {
	// Increment p pointer by size of int
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + IntSize))
}

func sum(arr []int) int {
	// Get pointer to underlying array
	arrPtr := unsafe.Pointer(*(**int)(unsafe.Pointer(&arr)))
	// Get len field of slice
	length := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + PtrSize))

	acc := 0

	for i := 0; i < length; i++ {
		acc += *(*int)(unsafe.Pointer((uintptr(arrPtr)) + uintptr(i*int(IntSize))))
	}

	return acc
}
