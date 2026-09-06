//go:build darwin && arm64

package jaccl

import "github.com/tmc/apple/rdma"

// Available reports whether the Apple RDMA provider is present. It does not
// establish that a device has an active route suitable for opening a group.
func Available() bool {
	return rdma.Available()
}
