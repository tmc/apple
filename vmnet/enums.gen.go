// Code generated from Apple documentation for vmnet. DO NOT EDIT.

package vmnet

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/vmnet/interface_event_t
type Interface_event_t int

const (
	// VMNET_INTERFACE_PACKETS_AVAILABLE: # Discussion
	VMNET_INTERFACE_PACKETS_AVAILABLE Interface_event_t = 1
)

func (e Interface_event_t) String() string {
	switch e {
	case VMNET_INTERFACE_PACKETS_AVAILABLE:
		return "VMNET_INTERFACE_PACKETS_AVAILABLE"
	default:
		return fmt.Sprintf("Interface_event_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/vmnet/operating_modes_t
type Operating_modes_t int

const (
)

// See: https://developer.apple.com/documentation/vmnet/vmnet_return_t
type Vmnet_return_t int

const (
)

