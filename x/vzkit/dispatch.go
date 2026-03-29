package vzkit

import (
	"github.com/tmc/apple/dispatch"
	vmruntime "github.com/tmc/apple/x/vzkit/vm"
)

// Queue wraps a dispatch queue for VM operations.
type Queue = vmruntime.Queue

// NewQueue creates a serial dispatch queue with the given label.
func NewQueue(label string) *Queue { return vmruntime.NewQueue(label) }

// WrapQueue creates a Queue from an existing dispatch.Queue.
func WrapQueue(q dispatch.Queue) *Queue { return vmruntime.WrapQueue(q) }
