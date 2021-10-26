package main

import (
	"fmt"
	"unsafe"
)

const BucketSize = 8

type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32
	buckets   unsafe.Pointer
}

type bmap struct {
	tophash  [BucketSize]uint8
	keys     [BucketSize]int
	values   [BucketSize]int
	overflow unsafe.Pointer
}

func main() {
	fmt.Println(maxValue(buildMap(100)))
}

func buildMap(count int) map[int]int {
	m := make(map[int]int, count)

	for i := 0; i < count; i++ {
		m[i] = i
	}

	return m
}

func maxValue(m map[int]int) int {
	// See https://github.com/golang/go/blob/0d0193409492b96881be6407ad50123e3557fdfb/src/runtime/map.go#L93
	const EmptyRest = 0
	const EmptyOne = 1

	// Map variables are always pointers, so we need to deref twice here
	mapStruct := *(**hmap)(unsafe.Pointer(&m))
	numBuckets := 1 << mapStruct.B
	bucketSize := int(unsafe.Sizeof(bmap{}))

	maxVal := 0

	for bucketIndex := 0; bucketIndex < numBuckets; bucketIndex++ {
		// Get bucket at index from buckets array
		bucket := *(*bmap)(unsafe.Pointer(uintptr(mapStruct.buckets) + uintptr(bucketSize*bucketIndex)))

		for itemIndex := 0; itemIndex < BucketSize; itemIndex++ {
			if bucket.tophash[itemIndex] == EmptyRest {
				// No additional items are present in this bucket
				break
			}

			if bucket.tophash[itemIndex] == EmptyOne {
				// Item is empty
				continue
			}

			if val := bucket.values[itemIndex]; val > maxVal {
				maxVal = val
			}
		}
	}

	return maxVal
}
