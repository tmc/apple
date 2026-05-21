// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsPluginXPC] class.
var (
	_DiskImageParamsPluginXPCClass     DiskImageParamsPluginXPCClass
	_DiskImageParamsPluginXPCClassOnce sync.Once
)

func getDiskImageParamsPluginXPCClass() DiskImageParamsPluginXPCClass {
	_DiskImageParamsPluginXPCClassOnce.Do(func() {
		_DiskImageParamsPluginXPCClass = DiskImageParamsPluginXPCClass{class: objc.GetClass("DiskImageParamsPlugin_XPC")}
	})
	return _DiskImageParamsPluginXPCClass
}

// GetDiskImageParamsPluginXPCClass returns the class object for DiskImageParamsPlugin_XPC.
func GetDiskImageParamsPluginXPCClass() DiskImageParamsPluginXPCClass {
	return getDiskImageParamsPluginXPCClass()
}

type DiskImageParamsPluginXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsPluginXPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsPluginXPCClass) Alloc() DiskImageParamsPluginXPC {
	rv := objc.Send[DiskImageParamsPluginXPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DiskImageParamsPluginXPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsPluginXPCFromID constructs a [DiskImageParamsPluginXPC] from an objc.ID.
func DiskImageParamsPluginXPCFromID(id objc.ID) DiskImageParamsPluginXPC {
	return DiskImageParamsPluginXPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}

// DiskImageParamsPlugin_XPCFromID is an alias for [DiskImageParamsPluginXPCFromID] for cross-framework compatibility.
func DiskImageParamsPlugin_XPCFromID(id objc.ID) DiskImageParamsPluginXPC {
	return DiskImageParamsPluginXPCFromID(id)
}

// Ensure DiskImageParamsPluginXPC implements IDiskImageParamsPluginXPC.
var _ IDiskImageParamsPluginXPC = DiskImageParamsPluginXPC{}

// An interface definition for the [DiskImageParamsPluginXPC] class.
type IDiskImageParamsPluginXPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsPluginXPC) Init() DiskImageParamsPluginXPC {
	rv := objc.Send[DiskImageParamsPluginXPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsPluginXPC) Autorelease() DiskImageParamsPluginXPC {
	rv := objc.Send[DiskImageParamsPluginXPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsPluginXPC creates a new DiskImageParamsPluginXPC instance.
func NewDiskImageParamsPluginXPC() DiskImageParamsPluginXPC {
	class := getDiskImageParamsPluginXPCClass()
	rv := objc.Send[DiskImageParamsPluginXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDiskImageParamsPlugin_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsPluginXPC {
	instance := getDiskImageParamsPluginXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsPluginXPCFromID(rv)
}

func NewDiskImageParamsPlugin_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsPluginXPC {
	instance := getDiskImageParamsPluginXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsPluginXPCFromID(rv)
}

func NewDiskImageParamsPlugin_XPCWithCoder(coder objectivec.IObject) DiskImageParamsPluginXPC {
	instance := getDiskImageParamsPluginXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsPluginXPCFromID(rv)
}
