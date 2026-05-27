package rdma

import (
	"errors"

	bindings "github.com/tmc/apple/rdma"
)

var (
	ErrRTRUnsafe = errors.New("rdma rtr unsafe")
)

// ErrnoText returns a compact name for common Apple RDMA errno values.
func ErrnoText(errno int) string {
	return bindings.ErrnoText(errno)
}

// ErrnoName returns the symbolic name for common Apple RDMA errno values.
func ErrnoName(errno int) string {
	return bindings.ErrnoName(errno)
}
