// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"fmt"
)

type NS uint32

const (
	// NSHandleOtherExceptionMask: The exception handler handles exceptions caught by handlers lower than the top-level handler by converting them to NSException objects containing a stack trace.
	NSHandleOtherExceptionMask NS = 512
	// NSHandleTopLevelExceptionMask: The exception handler handles exceptions caught by the top-level handler by converting them to NSException objects containing a stack trace.
	NSHandleTopLevelExceptionMask NS = 128
	// NSHandleUncaughtExceptionMask: The exception handler handles uncaught exceptions by terminating the thread in which they occur.
	NSHandleUncaughtExceptionMask NS = 2
	// NSHandleUncaughtRuntimeErrorMask: The exception handler handles uncaught runtime errors by converting them to NSException objects containing a stack trace.
	NSHandleUncaughtRuntimeErrorMask NS = 32
	// NSHandleUncaughtSystemExceptionMask: The exception handler handles uncaught system exceptions by converting them to NSException objects containing a stack trace.
	NSHandleUncaughtSystemExceptionMask NS = 8
	// NSLogOtherExceptionMask: The exception handler logs exceptions caught by handlers lower than the top-level handler.
	NSLogOtherExceptionMask NS = 256
	// NSLogTopLevelExceptionMask: The exception handler logs exceptions that would be caught by the top-level handler.
	NSLogTopLevelExceptionMask NS = 64
	// NSLogUncaughtExceptionMask: The exception handler logs uncaught exceptions.
	NSLogUncaughtExceptionMask NS = 1
	// NSLogUncaughtRuntimeErrorMask: The exception handler logs uncaught runtime errors.
	NSLogUncaughtRuntimeErrorMask NS = 16
	// NSLogUncaughtSystemExceptionMask: The exception handler logs uncaught system exceptions.
	NSLogUncaughtSystemExceptionMask NS = 4
)

func (e NS) String() string {
	switch e {
	case NSHandleOtherExceptionMask:
		return "NSHandleOtherExceptionMask"
	case NSHandleTopLevelExceptionMask:
		return "NSHandleTopLevelExceptionMask"
	case NSHandleUncaughtExceptionMask:
		return "NSHandleUncaughtExceptionMask"
	case NSHandleUncaughtRuntimeErrorMask:
		return "NSHandleUncaughtRuntimeErrorMask"
	case NSHandleUncaughtSystemExceptionMask:
		return "NSHandleUncaughtSystemExceptionMask"
	case NSLogOtherExceptionMask:
		return "NSLogOtherExceptionMask"
	case NSLogTopLevelExceptionMask:
		return "NSLogTopLevelExceptionMask"
	case NSLogUncaughtExceptionMask:
		return "NSLogUncaughtExceptionMask"
	case NSLogUncaughtRuntimeErrorMask:
		return "NSLogUncaughtRuntimeErrorMask"
	case NSLogUncaughtSystemExceptionMask:
		return "NSLogUncaughtSystemExceptionMask"
	default:
		return fmt.Sprintf("NS(%d)", e)
	}
}

type NSHangOn uint32

const (
	// NSHangOnOtherExceptionMask: The exception handler suspends execution when it detects an exception that would be handled by an object other than the top-level handler.
	NSHangOnOtherExceptionMask NSHangOn = 16
	// NSHangOnTopLevelExceptionMask: The exception handler suspends execution when it detects an exception that would be handled by the top-level handler.
	NSHangOnTopLevelExceptionMask NSHangOn = 8
	// NSHangOnUncaughtExceptionMask: The exception handler suspends execution when it detects an uncaught exception (other than a system exception or runtime error).
	NSHangOnUncaughtExceptionMask NSHangOn = 1
	// NSHangOnUncaughtRuntimeErrorMask: The exception handler suspends execution when it detects an uncaught runtime error.
	NSHangOnUncaughtRuntimeErrorMask NSHangOn = 4
	// NSHangOnUncaughtSystemExceptionMask: The exception handler suspends execution when it detects an uncaught system exception.
	NSHangOnUncaughtSystemExceptionMask NSHangOn = 2
)

func (e NSHangOn) String() string {
	switch e {
	case NSHangOnOtherExceptionMask:
		return "NSHangOnOtherExceptionMask"
	case NSHangOnTopLevelExceptionMask:
		return "NSHangOnTopLevelExceptionMask"
	case NSHangOnUncaughtExceptionMask:
		return "NSHangOnUncaughtExceptionMask"
	case NSHangOnUncaughtRuntimeErrorMask:
		return "NSHangOnUncaughtRuntimeErrorMask"
	case NSHangOnUncaughtSystemExceptionMask:
		return "NSHangOnUncaughtSystemExceptionMask"
	default:
		return fmt.Sprintf("NSHangOn(%d)", e)
	}
}
