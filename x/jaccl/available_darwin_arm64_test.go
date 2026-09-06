//go:build darwin && arm64

package jaccl

import (
	"testing"

	"github.com/tmc/apple/rdma"
)

func TestAvailableMatchesProvider(t *testing.T) {
	if got, want := Available(), rdma.Available(); got != want {
		t.Fatalf("Available = %t, want %t", got, want)
	}
}
