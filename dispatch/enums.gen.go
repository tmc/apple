// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"fmt"
)

type DispatchAutoreleaseFrequency uintptr

const (
	// DISPATCH_AUTORELEASE_FREQUENCY_INHERIT: The queue inherits its autorelease frequency from its target queue.
	DISPATCH_AUTORELEASE_FREQUENCY_INHERIT DispatchAutoreleaseFrequency = 0
	// DISPATCH_AUTORELEASE_FREQUENCY_NEVER: The queue does not set up an autorelease pool around executed blocks.
	DISPATCH_AUTORELEASE_FREQUENCY_NEVER DispatchAutoreleaseFrequency = 2
	// DISPATCH_AUTORELEASE_FREQUENCY_WORK_ITEM: The queue configures an autorelease pool before the execution of a block and releases the objects in that pool after the block finishes executing.
	DISPATCH_AUTORELEASE_FREQUENCY_WORK_ITEM DispatchAutoreleaseFrequency = 1
)

func (e DispatchAutoreleaseFrequency) String() string {
	switch e {
	case DISPATCH_AUTORELEASE_FREQUENCY_INHERIT:
		return "DISPATCH_AUTORELEASE_FREQUENCY_INHERIT"
	case DISPATCH_AUTORELEASE_FREQUENCY_NEVER:
		return "DISPATCH_AUTORELEASE_FREQUENCY_NEVER"
	case DISPATCH_AUTORELEASE_FREQUENCY_WORK_ITEM:
		return "DISPATCH_AUTORELEASE_FREQUENCY_WORK_ITEM"
	default:
		return fmt.Sprintf("DispatchAutoreleaseFrequency(%d)", e)
	}
}

type DispatchBlock uint

const (
	// DISPATCH_BLOCK_ASSIGN_CURRENT: Set the attributes of the work item to match the attributes of the current execution context.
	DISPATCH_BLOCK_ASSIGN_CURRENT DispatchBlock = 0x4
	// DISPATCH_BLOCK_BARRIER: Cause the work item to act as a barrier block when submitted to a concurrent queue.
	DISPATCH_BLOCK_BARRIER DispatchBlock = 0x1
	// DISPATCH_BLOCK_DETACHED: Disassociate the work item’s attributes from the current execution context.
	DISPATCH_BLOCK_DETACHED DispatchBlock = 0x2
	// DISPATCH_BLOCK_ENFORCE_QOS_CLASS: Prefer the quality-of-service class associated with the block.
	DISPATCH_BLOCK_ENFORCE_QOS_CLASS DispatchBlock = 0x20
	// DISPATCH_BLOCK_INHERIT_QOS_CLASS: Prefer the quality-of-service class associated with the current execution context.
	DISPATCH_BLOCK_INHERIT_QOS_CLASS DispatchBlock = 0x10
	// DISPATCH_BLOCK_NO_QOS_CLASS: Execute the work item without assigning a quality-of-service class.
	DISPATCH_BLOCK_NO_QOS_CLASS DispatchBlock = 0x8
)

func (e DispatchBlock) String() string {
	switch e {
	case DISPATCH_BLOCK_ASSIGN_CURRENT:
		return "DISPATCH_BLOCK_ASSIGN_CURRENT"
	case DISPATCH_BLOCK_BARRIER:
		return "DISPATCH_BLOCK_BARRIER"
	case DISPATCH_BLOCK_DETACHED:
		return "DISPATCH_BLOCK_DETACHED"
	case DISPATCH_BLOCK_ENFORCE_QOS_CLASS:
		return "DISPATCH_BLOCK_ENFORCE_QOS_CLASS"
	case DISPATCH_BLOCK_INHERIT_QOS_CLASS:
		return "DISPATCH_BLOCK_INHERIT_QOS_CLASS"
	case DISPATCH_BLOCK_NO_QOS_CLASS:
		return "DISPATCH_BLOCK_NO_QOS_CLASS"
	default:
		return fmt.Sprintf("DispatchBlock(%d)", e)
	}
}

type DispatchWalltime uint

const (
	// DISPATCH_WALLTIME_NOW: The current time.
	DISPATCH_WALLTIME_NOW DispatchWalltime = 18446744073709551614
)

func (e DispatchWalltime) String() string {
	switch e {
	case DISPATCH_WALLTIME_NOW:
		return "DISPATCH_WALLTIME_NOW"
	default:
		return fmt.Sprintf("DispatchWalltime(%d)", e)
	}
}
