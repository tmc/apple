// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"unsafe"
	"github.com/tmc/apple/objc"
)

// SliceToNSArray converts a Go slice to an NSArray.
// The toID function extracts the objc.ID from each element.
// Returns an empty NSArray if the slice is empty.
// Note: The returned NSArray is autoreleased. If you need to keep it beyond
// the current autorelease pool scope, call Retain() on the result.
func SliceToNSArray[T any](slice []T, toID func(T) objc.ID) objc.ID {
	if len(slice) == 0 {
		return objc.Send[objc.ID](objc.ID(objc.GetClass("NSArray")), objc.Sel("array"))
	}
	ids := make([]objc.ID, len(slice))
	for i, v := range slice {
		ids[i] = toID(v)
	}
	return objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSArray")),
		objc.Sel("arrayWithObjects:count:"),
		unsafe.Pointer(&ids[0]), uintptr(len(ids)),
	)
}

// StringSliceToNSArray converts a Go []string to an NSArray of NSString objects.
func StringSliceToNSArray(slice []string) objc.ID {
	return SliceToNSArray(slice, func(s string) objc.ID {
		return objc.String(s)
	})
}

// IObjectSliceToNSArray converts a slice of IObject-implementing types to an NSArray.
func IObjectSliceToNSArray[T IObject](slice []T) objc.ID {
	return SliceToNSArray(slice, func(obj T) objc.ID {
		return obj.GetID()
	})
}

// ClassSliceToNSArray converts a slice of objc.Class to an NSArray.
func ClassSliceToNSArray(slice []objc.Class) objc.ID {
	return SliceToNSArray(slice, func(c objc.Class) objc.ID {
		return objc.ID(c)
	})
}

// IDSliceToNSArray converts a slice of objc.ID to an NSArray.
func IDSliceToNSArray(slice []objc.ID) objc.ID {
	return SliceToNSArray(slice, func(id objc.ID) objc.ID {
		return id
	})
}


