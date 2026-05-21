// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsLockedXPC] class.
var (
	_DiskImageParamsLockedXPCClass     DiskImageParamsLockedXPCClass
	_DiskImageParamsLockedXPCClassOnce sync.Once
)

func getDiskImageParamsLockedXPCClass() DiskImageParamsLockedXPCClass {
	_DiskImageParamsLockedXPCClassOnce.Do(func() {
		_DiskImageParamsLockedXPCClass = DiskImageParamsLockedXPCClass{class: objc.GetClass("DiskImageParamsLocked_XPC")}
	})
	return _DiskImageParamsLockedXPCClass
}

// GetDiskImageParamsLockedXPCClass returns the class object for DiskImageParamsLocked_XPC.
func GetDiskImageParamsLockedXPCClass() DiskImageParamsLockedXPCClass {
	return getDiskImageParamsLockedXPCClass()
}

type DiskImageParamsLockedXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsLockedXPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsLockedXPCClass) Alloc() DiskImageParamsLockedXPC {
	rv := objc.Send[DiskImageParamsLockedXPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DiskImageParamsLockedXPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsLockedXPCFromID constructs a [DiskImageParamsLockedXPC] from an objc.ID.
func DiskImageParamsLockedXPCFromID(id objc.ID) DiskImageParamsLockedXPC {
	return DiskImageParamsLockedXPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}

// DiskImageParamsLocked_XPCFromID is an alias for [DiskImageParamsLockedXPCFromID] for cross-framework compatibility.
func DiskImageParamsLocked_XPCFromID(id objc.ID) DiskImageParamsLockedXPC {
	return DiskImageParamsLockedXPCFromID(id)
}

// Ensure DiskImageParamsLockedXPC implements IDiskImageParamsLockedXPC.
var _ IDiskImageParamsLockedXPC = DiskImageParamsLockedXPC{}

// An interface definition for the [DiskImageParamsLockedXPC] class.
type IDiskImageParamsLockedXPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsLockedXPC) Init() DiskImageParamsLockedXPC {
	rv := objc.Send[DiskImageParamsLockedXPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsLockedXPC) Autorelease() DiskImageParamsLockedXPC {
	rv := objc.Send[DiskImageParamsLockedXPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsLockedXPC creates a new DiskImageParamsLockedXPC instance.
func NewDiskImageParamsLockedXPC() DiskImageParamsLockedXPC {
	class := getDiskImageParamsLockedXPCClass()
	rv := objc.Send[DiskImageParamsLockedXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDiskImageParamsLocked_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsLockedXPC {
	instance := getDiskImageParamsLockedXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsLockedXPCFromID(rv)
}

func NewDiskImageParamsLocked_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsLockedXPC {
	instance := getDiskImageParamsLockedXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsLockedXPCFromID(rv)
}

func NewDiskImageParamsLocked_XPCWithCoder(coder objectivec.IObject) DiskImageParamsLockedXPC {
	instance := getDiskImageParamsLockedXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsLockedXPCFromID(rv)
}
