// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"fmt"
)

type DispatchWalltimeNow uint

const (
	// DISPATCH_WALLTIME_NOW: The current time.
	DISPATCH_WALLTIME_NOW DispatchWalltimeNow = 0
)

func (e DispatchWalltimeNow) String() string {
	switch e {
	case DISPATCH_WALLTIME_NOW:
		return "DISPATCH_WALLTIME_NOW"
	default:
		return fmt.Sprintf("DispatchWalltimeNow(%d)", e)
	}
}

