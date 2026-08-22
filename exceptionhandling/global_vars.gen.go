// Code generated from Apple documentation. DO NOT EDIT.

package exceptionhandling

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// StackTraceKey is the key for fetching the stack trace (an [NSString] object) in the [userInfo] dictionary of the [NSException] object passed into one of the delegate methods described in [NSExceptionHandlerDelegate].
	//
	// See: https://developer.apple.com/documentation/ExceptionHandling/NSStackTraceKey
	StackTraceKey string
	// UncaughtRuntimeErrorException is identifies an Objective-C runtime error.
	//
	// See: https://developer.apple.com/documentation/ExceptionHandling/NSUncaughtRuntimeErrorException
	UncaughtRuntimeErrorException string
	// UncaughtSystemExceptionException is identifies an uncaught system exception.
	//
	// See: https://developer.apple.com/documentation/ExceptionHandling/NSUncaughtSystemExceptionException
	UncaughtSystemExceptionException string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStackTraceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StackTraceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUncaughtRuntimeErrorException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UncaughtRuntimeErrorException = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUncaughtSystemExceptionException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UncaughtSystemExceptionException = objc.GoString(cstr)
			}
		}
	}

}
