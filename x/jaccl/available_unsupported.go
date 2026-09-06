//go:build !darwin || !arm64

package jaccl

// Available reports whether the native JACCL transport is available.
func Available() bool {
	return false
}
