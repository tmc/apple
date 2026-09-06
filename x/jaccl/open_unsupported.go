//go:build !darwin || !arm64

package jaccl

import (
	"context"
	"fmt"
)

func open(context.Context, Config) (*Group, error) {
	return nil, fmt.Errorf("darwin/arm64 RDMA backend: %w", ErrUnsupported)
}
