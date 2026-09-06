package jaccl

import "errors"

var (
	// ErrUnsupported reports that the local platform or RDMA provider is not supported.
	ErrUnsupported = errors.New("jaccl: unsupported")

	// ErrInvalidConfig reports an invalid group configuration.
	ErrInvalidConfig = errors.New("jaccl: invalid config")

	// ErrProtocol reports an invalid or mismatched peer protocol message.
	ErrProtocol = errors.New("jaccl: protocol error")

	// ErrClosed reports use of a closing or closed group.
	ErrClosed = errors.New("jaccl: group closed")

	// ErrPoisoned reports a group made unsafe by a transport failure.
	ErrPoisoned = errors.New("jaccl: group poisoned")
)
