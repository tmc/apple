// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"fmt"
)

type DispatchWalltime uint

const (
	// DISPATCH_WALLTIME_NOW: The current time.
	DISPATCH_WALLTIME_NOW DispatchWalltime = 0
)

func (e DispatchWalltime) String() string {
	switch e {
	case DISPATCH_WALLTIME_NOW:
		return "DISPATCH_WALLTIME_NOW"
	default:
		return fmt.Sprintf("DispatchWalltime(%d)", e)
	}
}
