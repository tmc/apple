// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsASIFXPC] class.
var (
	_DiskImageParamsASIFXPCClass     DiskImageParamsASIFXPCClass
	_DiskImageParamsASIFXPCClassOnce sync.Once
)

func getDiskImageParamsASIFXPCClass() DiskImageParamsASIFXPCClass {
	_DiskImageParamsASIFXPCClassOnce.Do(func() {
		_DiskImageParamsASIFXPCClass = DiskImageParamsASIFXPCClass{class: objc.GetClass("DiskImageParamsASIF_XPC")}
	})
	return _DiskImageParamsASIFXPCClass
}

// GetDiskImageParamsASIFXPCClass returns the class object for DiskImageParamsASIF_XPC.
func GetDiskImageParamsASIFXPCClass() DiskImageParamsASIFXPCClass {
	return getDiskImageParamsASIFXPCClass()
}

type DiskImageParamsASIFXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsASIFXPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsASIFXPCClass) Alloc() DiskImageParamsASIFXPC {
	rv := objc.Send[DiskImageParamsASIFXPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DiskImageParamsASIFXPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsASIFXPCFromID constructs a [DiskImageParamsASIFXPC] from an objc.ID.
func DiskImageParamsASIFXPCFromID(id objc.ID) DiskImageParamsASIFXPC {
	return DiskImageParamsASIFXPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}

// DiskImageParamsASIF_XPCFromID is an alias for [DiskImageParamsASIFXPCFromID] for cross-framework compatibility.
func DiskImageParamsASIF_XPCFromID(id objc.ID) DiskImageParamsASIFXPC {
	return DiskImageParamsASIFXPCFromID(id)
}

// Ensure DiskImageParamsASIFXPC implements IDiskImageParamsASIFXPC.
var _ IDiskImageParamsASIFXPC = DiskImageParamsASIFXPC{}

// An interface definition for the [DiskImageParamsASIFXPC] class.
type IDiskImageParamsASIFXPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsASIFXPC) Init() DiskImageParamsASIFXPC {
	rv := objc.Send[DiskImageParamsASIFXPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsASIFXPC) Autorelease() DiskImageParamsASIFXPC {
	rv := objc.Send[DiskImageParamsASIFXPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsASIFXPC creates a new DiskImageParamsASIFXPC instance.
func NewDiskImageParamsASIFXPC() DiskImageParamsASIFXPC {
	class := getDiskImageParamsASIFXPCClass()
	rv := objc.Send[DiskImageParamsASIFXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDiskImageParamsASIF_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsASIFXPC {
	instance := getDiskImageParamsASIFXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsASIFXPCFromID(rv)
}

func NewDiskImageParamsASIF_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsASIFXPC {
	instance := getDiskImageParamsASIFXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsASIFXPCFromID(rv)
}

func NewDiskImageParamsASIF_XPCWithBackendXPCHeader(xpc objectivec.IObject, header unsafe.Pointer) DiskImageParamsASIFXPC {
	instance := getDiskImageParamsASIFXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:header:"), xpc, header)
	return DiskImageParamsASIFXPCFromID(rv)
}

func NewDiskImageParamsASIF_XPCWithCoder(coder objectivec.IObject) DiskImageParamsASIFXPC {
	instance := getDiskImageParamsASIFXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsASIFXPCFromID(rv)
}
