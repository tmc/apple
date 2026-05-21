// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsRawXPC] class.
var (
	_DiskImageParamsRawXPCClass     DiskImageParamsRawXPCClass
	_DiskImageParamsRawXPCClassOnce sync.Once
)

func getDiskImageParamsRawXPCClass() DiskImageParamsRawXPCClass {
	_DiskImageParamsRawXPCClassOnce.Do(func() {
		_DiskImageParamsRawXPCClass = DiskImageParamsRawXPCClass{class: objc.GetClass("DiskImageParamsRaw_XPC")}
	})
	return _DiskImageParamsRawXPCClass
}

// GetDiskImageParamsRawXPCClass returns the class object for DiskImageParamsRaw_XPC.
func GetDiskImageParamsRawXPCClass() DiskImageParamsRawXPCClass {
	return getDiskImageParamsRawXPCClass()
}

type DiskImageParamsRawXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsRawXPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsRawXPCClass) Alloc() DiskImageParamsRawXPC {
	rv := objc.Send[DiskImageParamsRawXPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DiskImageParamsRawXPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsRawXPCFromID constructs a [DiskImageParamsRawXPC] from an objc.ID.
func DiskImageParamsRawXPCFromID(id objc.ID) DiskImageParamsRawXPC {
	return DiskImageParamsRawXPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}

// DiskImageParamsRaw_XPCFromID is an alias for [DiskImageParamsRawXPCFromID] for cross-framework compatibility.
func DiskImageParamsRaw_XPCFromID(id objc.ID) DiskImageParamsRawXPC {
	return DiskImageParamsRawXPCFromID(id)
}

// Ensure DiskImageParamsRawXPC implements IDiskImageParamsRawXPC.
var _ IDiskImageParamsRawXPC = DiskImageParamsRawXPC{}

// An interface definition for the [DiskImageParamsRawXPC] class.
type IDiskImageParamsRawXPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsRawXPC) Init() DiskImageParamsRawXPC {
	rv := objc.Send[DiskImageParamsRawXPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsRawXPC) Autorelease() DiskImageParamsRawXPC {
	rv := objc.Send[DiskImageParamsRawXPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsRawXPC creates a new DiskImageParamsRawXPC instance.
func NewDiskImageParamsRawXPC() DiskImageParamsRawXPC {
	class := getDiskImageParamsRawXPCClass()
	rv := objc.Send[DiskImageParamsRawXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDiskImageParamsRaw_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsRawXPC {
	instance := getDiskImageParamsRawXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsRawXPCFromID(rv)
}

func NewDiskImageParamsRaw_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsRawXPC {
	instance := getDiskImageParamsRawXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsRawXPCFromID(rv)
}

func NewDiskImageParamsRaw_XPCWithCoder(coder objectivec.IObject) DiskImageParamsRawXPC {
	instance := getDiskImageParamsRawXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsRawXPCFromID(rv)
}
