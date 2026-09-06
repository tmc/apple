package jaccl

import "unsafe"

func sameSlice(a, b []byte) bool {
	return len(a) == len(b) && (len(a) == 0 || unsafe.SliceData(a) == unsafe.SliceData(b))
}

func slicesOverlap(a, b []byte) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	aStart := uintptr(unsafe.Pointer(unsafe.SliceData(a)))
	bStart := uintptr(unsafe.Pointer(unsafe.SliceData(b)))
	aEnd := aStart + uintptr(len(a))
	bEnd := bStart + uintptr(len(b))
	return aStart < bEnd && bStart < aEnd
}
