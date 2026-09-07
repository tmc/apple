//go:build darwin

package lldb

import (
	"os"
	"unsafe"
)

// cstr returns a NUL-terminated C string for s. The returned pointer is only
// valid for the duration of the call it is passed to.
func cstr(s string) *byte {
	b := make([]byte, len(s)+1)
	copy(b, s)
	return &b[0]
}

// cstrv builds a NULL-terminated array of C strings from ss and returns a
// pointer to its first element, or nil for an empty slice. The array and its
// strings are heap-allocated; the caller must keep them referenced for the
// duration of the call they are passed to.
func cstrv(ss []string) **byte {
	if len(ss) == 0 {
		return nil
	}
	arr := make([]*byte, len(ss)+1)
	for i, s := range ss {
		arr[i] = cstr(s)
	}
	return &arr[0]
}

// goString copies a NUL-terminated C string at p into a Go string. A nil pointer
// yields "".
func goString(p unsafe.Pointer) string {
	if p == nil {
		return ""
	}
	var n int
	for *(*byte)(unsafe.Add(p, n)) != 0 {
		n++
	}
	return string(unsafe.Slice((*byte)(p), n))
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
