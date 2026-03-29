// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineSaveOptions] class.
var (
	_VZVirtualMachineSaveOptionsClass     VZVirtualMachineSaveOptionsClass
	_VZVirtualMachineSaveOptionsClassOnce sync.Once
)

func getVZVirtualMachineSaveOptionsClass() VZVirtualMachineSaveOptionsClass {
	_VZVirtualMachineSaveOptionsClassOnce.Do(func() {
		_VZVirtualMachineSaveOptionsClass = VZVirtualMachineSaveOptionsClass{class: objc.GetClass("_VZVirtualMachineSaveOptions")}
	})
	return _VZVirtualMachineSaveOptionsClass
}

// GetVZVirtualMachineSaveOptionsClass returns the class object for _VZVirtualMachineSaveOptions.
func GetVZVirtualMachineSaveOptionsClass() VZVirtualMachineSaveOptionsClass {
	return getVZVirtualMachineSaveOptionsClass()
}

type VZVirtualMachineSaveOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineSaveOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineSaveOptionsClass) Alloc() VZVirtualMachineSaveOptions {
	rv := objc.Send[VZVirtualMachineSaveOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtualMachineSaveOptions.Compress]
//   - [VZVirtualMachineSaveOptions.SetCompress]
//   - [VZVirtualMachineSaveOptions.Encrypt]
//   - [VZVirtualMachineSaveOptions.SetEncrypt]
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineSaveOptions
type VZVirtualMachineSaveOptions struct {
	objectivec.Object
}

// VZVirtualMachineSaveOptionsFromID constructs a [VZVirtualMachineSaveOptions] from an objc.ID.
func VZVirtualMachineSaveOptionsFromID(id objc.ID) VZVirtualMachineSaveOptions {
	return VZVirtualMachineSaveOptions{objectivec.Object{ID: id}}
}
// Ensure VZVirtualMachineSaveOptions implements IVZVirtualMachineSaveOptions.
var _ IVZVirtualMachineSaveOptions = VZVirtualMachineSaveOptions{}

// An interface definition for the [VZVirtualMachineSaveOptions] class.
//
// # Methods
//
//   - [IVZVirtualMachineSaveOptions.Compress]
//   - [IVZVirtualMachineSaveOptions.SetCompress]
//   - [IVZVirtualMachineSaveOptions.Encrypt]
//   - [IVZVirtualMachineSaveOptions.SetEncrypt]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineSaveOptions
type IVZVirtualMachineSaveOptions interface {
	objectivec.IObject

	// Topic: Methods

	Compress() bool
	SetCompress(value bool)
	Encrypt() bool
	SetEncrypt(value bool)
}

// Init initializes the instance.
func (v VZVirtualMachineSaveOptions) Init() VZVirtualMachineSaveOptions {
	rv := objc.Send[VZVirtualMachineSaveOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineSaveOptions) Autorelease() VZVirtualMachineSaveOptions {
	rv := objc.Send[VZVirtualMachineSaveOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineSaveOptions creates a new VZVirtualMachineSaveOptions instance.
func NewVZVirtualMachineSaveOptions() VZVirtualMachineSaveOptions {
	class := getVZVirtualMachineSaveOptionsClass()
	rv := objc.Send[VZVirtualMachineSaveOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineSaveOptions/compress
func (v VZVirtualMachineSaveOptions) Compress() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("compress"))
	return rv
}
func (v VZVirtualMachineSaveOptions) SetCompress(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setCompress:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineSaveOptions/encrypt
func (v VZVirtualMachineSaveOptions) Encrypt() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("encrypt"))
	return rv
}
func (v VZVirtualMachineSaveOptions) SetEncrypt(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setEncrypt:"), value)
}

