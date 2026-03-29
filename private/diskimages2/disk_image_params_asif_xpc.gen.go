// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsASIF_XPC] class.
var (
	_DiskImageParamsASIF_XPCClass     DiskImageParamsASIF_XPCClass
	_DiskImageParamsASIF_XPCClassOnce sync.Once
)

func getDiskImageParamsASIF_XPCClass() DiskImageParamsASIF_XPCClass {
	_DiskImageParamsASIF_XPCClassOnce.Do(func() {
		_DiskImageParamsASIF_XPCClass = DiskImageParamsASIF_XPCClass{class: objc.GetClass("DiskImageParamsASIF_XPC")}
	})
	return _DiskImageParamsASIF_XPCClass
}

// GetDiskImageParamsASIF_XPCClass returns the class object for DiskImageParamsASIF_XPC.
func GetDiskImageParamsASIF_XPCClass() DiskImageParamsASIF_XPCClass {
	return getDiskImageParamsASIF_XPCClass()
}

type DiskImageParamsASIF_XPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsASIF_XPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsASIF_XPCClass) Alloc() DiskImageParamsASIF_XPC {
	rv := objc.Send[DiskImageParamsASIF_XPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DiskImageParamsASIF_XPC.SetHeader]
//   - [DiskImageParamsASIF_XPC.InitWithBackendXPCHeader]
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsASIF_XPC
type DiskImageParamsASIF_XPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsASIF_XPCFromID constructs a [DiskImageParamsASIF_XPC] from an objc.ID.
func DiskImageParamsASIF_XPCFromID(id objc.ID) DiskImageParamsASIF_XPC {
	return DiskImageParamsASIF_XPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}
// Ensure DiskImageParamsASIF_XPC implements IDiskImageParamsASIF_XPC.
var _ IDiskImageParamsASIF_XPC = DiskImageParamsASIF_XPC{}

// An interface definition for the [DiskImageParamsASIF_XPC] class.
//
// # Methods
//
//   - [IDiskImageParamsASIF_XPC.SetHeader]
//   - [IDiskImageParamsASIF_XPC.InitWithBackendXPCHeader]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsASIF_XPC
type IDiskImageParamsASIF_XPC interface {
	IDiskImageParamsXPC

	// Topic: Methods

	SetHeader(header objectivec.IObject)
	InitWithBackendXPCHeader(xpc objectivec.IObject, header objectivec.IObject) DiskImageParamsASIF_XPC
}

// Init initializes the instance.
func (d DiskImageParamsASIF_XPC) Init() DiskImageParamsASIF_XPC {
	rv := objc.Send[DiskImageParamsASIF_XPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsASIF_XPC) Autorelease() DiskImageParamsASIF_XPC {
	rv := objc.Send[DiskImageParamsASIF_XPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsASIF_XPC creates a new DiskImageParamsASIF_XPC instance.
func NewDiskImageParamsASIF_XPC() DiskImageParamsASIF_XPC {
	class := getDiskImageParamsASIF_XPCClass()
	rv := objc.Send[DiskImageParamsASIF_XPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func NewDiskImageParamsASIF_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsASIF_XPC {
	instance := getDiskImageParamsASIF_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsASIF_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func NewDiskImageParamsASIF_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsASIF_XPC {
	instance := getDiskImageParamsASIF_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsASIF_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsASIF_XPC/initWithBackendXPC:header:
func NewDiskImageParamsASIF_XPCWithBackendXPCHeader(xpc objectivec.IObject, header objectivec.IObject) DiskImageParamsASIF_XPC {
	instance := getDiskImageParamsASIF_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:header:"), xpc, header)
	return DiskImageParamsASIF_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsASIF_XPC/initWithCoder:
func NewDiskImageParamsASIF_XPCWithCoder(coder objectivec.IObject) DiskImageParamsASIF_XPC {
	instance := getDiskImageParamsASIF_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsASIF_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsASIF_XPC/setHeader:
func (d DiskImageParamsASIF_XPC) SetHeader(header objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("setHeader:"), header)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsASIF_XPC/initWithBackendXPC:header:
func (d DiskImageParamsASIF_XPC) InitWithBackendXPCHeader(xpc objectivec.IObject, header objectivec.IObject) DiskImageParamsASIF_XPC {
	rv := objc.Send[DiskImageParamsASIF_XPC](d.ID, objc.Sel("initWithBackendXPC:header:"), xpc, header)
	return rv
}

