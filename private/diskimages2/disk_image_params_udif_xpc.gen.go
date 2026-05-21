// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsUDIFXPC] class.
var (
	_DiskImageParamsUDIFXPCClass     DiskImageParamsUDIFXPCClass
	_DiskImageParamsUDIFXPCClassOnce sync.Once
)

func getDiskImageParamsUDIFXPCClass() DiskImageParamsUDIFXPCClass {
	_DiskImageParamsUDIFXPCClassOnce.Do(func() {
		_DiskImageParamsUDIFXPCClass = DiskImageParamsUDIFXPCClass{class: objc.GetClass("DiskImageParamsUDIF_XPC")}
	})
	return _DiskImageParamsUDIFXPCClass
}

// GetDiskImageParamsUDIFXPCClass returns the class object for DiskImageParamsUDIF_XPC.
func GetDiskImageParamsUDIFXPCClass() DiskImageParamsUDIFXPCClass {
	return getDiskImageParamsUDIFXPCClass()
}

type DiskImageParamsUDIFXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsUDIFXPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsUDIFXPCClass) Alloc() DiskImageParamsUDIFXPC {
	rv := objc.Send[DiskImageParamsUDIFXPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DiskImageParamsUDIFXPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsUDIFXPCFromID constructs a [DiskImageParamsUDIFXPC] from an objc.ID.
func DiskImageParamsUDIFXPCFromID(id objc.ID) DiskImageParamsUDIFXPC {
	return DiskImageParamsUDIFXPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}

// DiskImageParamsUDIF_XPCFromID is an alias for [DiskImageParamsUDIFXPCFromID] for cross-framework compatibility.
func DiskImageParamsUDIF_XPCFromID(id objc.ID) DiskImageParamsUDIFXPC {
	return DiskImageParamsUDIFXPCFromID(id)
}

// Ensure DiskImageParamsUDIFXPC implements IDiskImageParamsUDIFXPC.
var _ IDiskImageParamsUDIFXPC = DiskImageParamsUDIFXPC{}

// An interface definition for the [DiskImageParamsUDIFXPC] class.
type IDiskImageParamsUDIFXPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsUDIFXPC) Init() DiskImageParamsUDIFXPC {
	rv := objc.Send[DiskImageParamsUDIFXPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsUDIFXPC) Autorelease() DiskImageParamsUDIFXPC {
	rv := objc.Send[DiskImageParamsUDIFXPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsUDIFXPC creates a new DiskImageParamsUDIFXPC instance.
func NewDiskImageParamsUDIFXPC() DiskImageParamsUDIFXPC {
	class := getDiskImageParamsUDIFXPCClass()
	rv := objc.Send[DiskImageParamsUDIFXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDiskImageParamsUDIF_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsUDIFXPC {
	instance := getDiskImageParamsUDIFXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsUDIFXPCFromID(rv)
}

func NewDiskImageParamsUDIF_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsUDIFXPC {
	instance := getDiskImageParamsUDIFXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsUDIFXPCFromID(rv)
}

func NewDiskImageParamsUDIF_XPCWithBackendXPCHeader(xpc objectivec.IObject, header unsafe.Pointer) DiskImageParamsUDIFXPC {
	instance := getDiskImageParamsUDIFXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:header:"), xpc, header)
	return DiskImageParamsUDIFXPCFromID(rv)
}

func NewDiskImageParamsUDIF_XPCWithCoder(coder objectivec.IObject) DiskImageParamsUDIFXPC {
	instance := getDiskImageParamsUDIFXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsUDIFXPCFromID(rv)
}
