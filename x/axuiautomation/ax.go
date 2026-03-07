package axuiautomation

import (
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// Type aliases for accessibility types
type AXUIElementRef = uintptr
type AXObserverRef = uintptr
type AXValueRef = uintptr
type AXError = int32
type AXValueType = int32
type AXObserverCallback = uintptr

// AX function pointers
var (
	axUIElementCreateApplication      func(pid int32) AXUIElementRef
	axUIElementCopyAttributeValue     func(element AXUIElementRef, attribute uintptr, value *uintptr) AXError
	axUIElementSetAttributeValue      func(element AXUIElementRef, attribute uintptr, value uintptr) AXError
	axUIElementGetAttributeValueCount func(element AXUIElementRef, attribute uintptr, count *int) AXError
	axUIElementPerformAction          func(element AXUIElementRef, action uintptr) AXError
	axUIElementGetPid                 func(element AXUIElementRef, pid *int32) AXError
	axValueCreate                     func(valueType AXValueType, valuePtr unsafe.Pointer) AXValueRef
	axValueGetValue                   func(value AXValueRef, valueType AXValueType, valuePtr unsafe.Pointer) bool
	axObserverCreate                  func(pid int32, callback AXObserverCallback, observer *AXObserverRef) AXError
	axObserverGetRunLoopSource        func(observer AXObserverRef) uintptr
	axObserverAddNotification         func(observer AXObserverRef, element AXUIElementRef, notification uintptr, refcon unsafe.Pointer) AXError
	axIsProcessTrusted                func() bool
	axIsProcessTrustedWithOptions     func(options uintptr) bool

	axInitOnce sync.Once
	axLoaded   bool
)

func initAX() {
	axInitOnce.Do(func() {
		lib, err := purego.Dlopen("/System/Library/Frameworks/ApplicationServices.framework/ApplicationServices", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err != nil {
			return
		}

		purego.RegisterLibFunc(&axUIElementCreateApplication, lib, "AXUIElementCreateApplication")
		purego.RegisterLibFunc(&axUIElementCopyAttributeValue, lib, "AXUIElementCopyAttributeValue")
		purego.RegisterLibFunc(&axUIElementSetAttributeValue, lib, "AXUIElementSetAttributeValue")
		purego.RegisterLibFunc(&axUIElementGetAttributeValueCount, lib, "AXUIElementGetAttributeValueCount")
		purego.RegisterLibFunc(&axUIElementPerformAction, lib, "AXUIElementPerformAction")
		purego.RegisterLibFunc(&axUIElementGetPid, lib, "AXUIElementGetPid")
		purego.RegisterLibFunc(&axValueCreate, lib, "AXValueCreate")
		purego.RegisterLibFunc(&axValueGetValue, lib, "AXValueGetValue")
		purego.RegisterLibFunc(&axObserverCreate, lib, "AXObserverCreate")
		purego.RegisterLibFunc(&axObserverGetRunLoopSource, lib, "AXObserverGetRunLoopSource")
		purego.RegisterLibFunc(&axObserverAddNotification, lib, "AXObserverAddNotification")
		purego.RegisterLibFunc(&axIsProcessTrusted, lib, "AXIsProcessTrusted")
		purego.RegisterLibFunc(&axIsProcessTrustedWithOptions, lib, "AXIsProcessTrustedWithOptions")

		axLoaded = true
	})
}

// Wrapper functions that match the expected signatures

func AXUIElementCreateApplication(pid int32) AXUIElementRef {
	initAX()
	if !axLoaded || axUIElementCreateApplication == nil {
		return 0
	}
	return axUIElementCreateApplication(pid)
}

func AXUIElementCopyAttributeValue(element AXUIElementRef, attribute uintptr, value *uintptr) AXError {
	initAX()
	if !axLoaded || axUIElementCopyAttributeValue == nil {
		return -1
	}
	return axUIElementCopyAttributeValue(element, attribute, value)
}

// AXUIElementCopyAttributeValueCF is a version that works with CFTypeRef pointers
func AXUIElementCopyAttributeValueCF(element AXUIElementRef, attribute uintptr, value *unsafe.Pointer) AXError {
	initAX()
	if !axLoaded || axUIElementCopyAttributeValue == nil {
		return -1
	}
	var ptr uintptr
	result := axUIElementCopyAttributeValue(element, attribute, &ptr)
	if result == 0 && ptr != 0 {
		*value = unsafe.Pointer(ptr)
	}
	return result
}

func AXUIElementSetAttributeValue(element AXUIElementRef, attribute uintptr, value uintptr) AXError {
	initAX()
	if !axLoaded || axUIElementSetAttributeValue == nil {
		return -1
	}
	return axUIElementSetAttributeValue(element, attribute, value)
}

func AXUIElementGetAttributeValueCount(element AXUIElementRef, attribute uintptr, count *int) AXError {
	initAX()
	if !axLoaded || axUIElementGetAttributeValueCount == nil {
		return -1
	}
	return axUIElementGetAttributeValueCount(element, attribute, count)
}

func AXUIElementPerformAction(element AXUIElementRef, action uintptr) AXError {
	initAX()
	if !axLoaded || axUIElementPerformAction == nil {
		return -1
	}
	return axUIElementPerformAction(element, action)
}

func AXUIElementGetPid(element AXUIElementRef, pid *int32) AXError {
	initAX()
	if !axLoaded || axUIElementGetPid == nil {
		return -1
	}
	return axUIElementGetPid(element, pid)
}

func AXValueCreate(valueType AXValueType, valuePtr unsafe.Pointer) AXValueRef {
	initAX()
	if !axLoaded || axValueCreate == nil {
		return 0
	}
	return axValueCreate(valueType, valuePtr)
}

func AXValueGetValue(value AXValueRef, valueType AXValueType, valuePtr unsafe.Pointer) bool {
	initAX()
	if !axLoaded || axValueGetValue == nil {
		return false
	}
	return axValueGetValue(value, valueType, valuePtr)
}

func AXObserverCreate(pid int32, callback AXObserverCallback, observer *AXObserverRef) AXError {
	initAX()
	if !axLoaded || axObserverCreate == nil {
		return -1
	}
	return axObserverCreate(pid, callback, observer)
}

func AXObserverGetRunLoopSource(observer AXObserverRef) uintptr {
	initAX()
	if !axLoaded || axObserverGetRunLoopSource == nil {
		return 0
	}
	return axObserverGetRunLoopSource(observer)
}

func AXObserverAddNotification(observer AXObserverRef, element AXUIElementRef, notification uintptr, refcon unsafe.Pointer) AXError {
	initAX()
	if !axLoaded || axObserverAddNotification == nil {
		return -1
	}
	return axObserverAddNotification(observer, element, notification, refcon)
}

func AXIsProcessTrusted() bool {
	initAX()
	if !axLoaded || axIsProcessTrusted == nil {
		return false
	}
	return axIsProcessTrusted()
}

func AXIsProcessTrustedWithOptions(options uintptr) bool {
	initAX()
	if !axLoaded || axIsProcessTrustedWithOptions == nil {
		return false
	}
	return axIsProcessTrustedWithOptions(options)
}
