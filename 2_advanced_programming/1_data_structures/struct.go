package main

import (
	"fmt"
	"unsafe"
)

type Point struct{ x, y int }

func main() {
	fmt.Println(getY(Point{x: 1, y: 2}))
}

func getY(p Point) int {
	const IntSize = unsafe.Sizeof(0)

	// Increment p pointer by size of int
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + IntSize))
}
