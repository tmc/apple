// Code generated from Apple documentation for vmnet. DO NOT EDIT.

package vmnet

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/vmnet/interface_event_t
type Interface_event_t uint32

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
type Operating_modes_t uint32

const (
	VMNET_BRIDGED_MODE Operating_modes_t = 1002
	// VMNET_HOST_MODE: # Discussion
	VMNET_HOST_MODE Operating_modes_t = 1000
	// VMNET_SHARED_MODE: # Discussion
	VMNET_SHARED_MODE Operating_modes_t = 1001
)

func (e Operating_modes_t) String() string {
	switch e {
	case VMNET_BRIDGED_MODE:
		return "VMNET_BRIDGED_MODE"
	case VMNET_HOST_MODE:
		return "VMNET_HOST_MODE"
	case VMNET_SHARED_MODE:
		return "VMNET_SHARED_MODE"
	default:
		return fmt.Sprintf("Operating_modes_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/vmnet/vmnet_return_t
type Vmnet_return_t uint32

const (
	// VMNET_BUFFER_EXHAUSTED: # Discussion
	VMNET_BUFFER_EXHAUSTED Vmnet_return_t = 1007
	// VMNET_FAILURE: # Discussion
	VMNET_FAILURE Vmnet_return_t = 1001
	// VMNET_INVALID_ACCESS: # Discussion
	VMNET_INVALID_ACCESS Vmnet_return_t = 1005
	// VMNET_INVALID_ARGUMENT: # Discussion
	VMNET_INVALID_ARGUMENT Vmnet_return_t = 1003
	// VMNET_MEM_FAILURE: # Discussion
	VMNET_MEM_FAILURE    Vmnet_return_t = 1002
	VMNET_NOT_AUTHORIZED Vmnet_return_t = 1010
	// VMNET_PACKET_TOO_BIG: # Discussion
	VMNET_PACKET_TOO_BIG Vmnet_return_t = 1006
	// VMNET_SETUP_INCOMPLETE: # Discussion
	VMNET_SETUP_INCOMPLETE     Vmnet_return_t = 1004
	VMNET_SHARING_SERVICE_BUSY Vmnet_return_t = 1009
	// VMNET_SUCCESS: # Discussion
	VMNET_SUCCESS Vmnet_return_t = 1000
	// VMNET_TOO_MANY_PACKETS: # Discussion
	VMNET_TOO_MANY_PACKETS Vmnet_return_t = 1008
)

func (e Vmnet_return_t) String() string {
	switch e {
	case VMNET_BUFFER_EXHAUSTED:
		return "VMNET_BUFFER_EXHAUSTED"
	case VMNET_FAILURE:
		return "VMNET_FAILURE"
	case VMNET_INVALID_ACCESS:
		return "VMNET_INVALID_ACCESS"
	case VMNET_INVALID_ARGUMENT:
		return "VMNET_INVALID_ARGUMENT"
	case VMNET_MEM_FAILURE:
		return "VMNET_MEM_FAILURE"
	case VMNET_NOT_AUTHORIZED:
		return "VMNET_NOT_AUTHORIZED"
	case VMNET_PACKET_TOO_BIG:
		return "VMNET_PACKET_TOO_BIG"
	case VMNET_SETUP_INCOMPLETE:
		return "VMNET_SETUP_INCOMPLETE"
	case VMNET_SHARING_SERVICE_BUSY:
		return "VMNET_SHARING_SERVICE_BUSY"
	case VMNET_SUCCESS:
		return "VMNET_SUCCESS"
	case VMNET_TOO_MANY_PACKETS:
		return "VMNET_TOO_MANY_PACKETS"
	default:
		return fmt.Sprintf("Vmnet_return_t(%d)", e)
	}
}
