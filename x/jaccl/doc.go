// Package jaccl implements RDMA collective communication in Go.
//
// It uses the Apple RDMA bindings directly. It does not call libjaccl or use
// cgo. The initial implementation targets Apple Thunderbolt RDMA on darwin/
// arm64; unsupported platforms return ErrUnsupported.
package jaccl
