// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSharedDirectory] class.
var (
	_VZSharedDirectoryClass     VZSharedDirectoryClass
	_VZSharedDirectoryClassOnce sync.Once
)

func getVZSharedDirectoryClass() VZSharedDirectoryClass {
	_VZSharedDirectoryClassOnce.Do(func() {
		_VZSharedDirectoryClass = VZSharedDirectoryClass{class: objc.GetClass("VZSharedDirectory")}
	})
	return _VZSharedDirectoryClass
}

// GetVZSharedDirectoryClass returns the class object for VZSharedDirectory.
func GetVZSharedDirectoryClass() VZSharedDirectoryClass {
	return getVZSharedDirectoryClass()
}

type VZSharedDirectoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSharedDirectoryClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSharedDirectoryClass) Alloc() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSharedDirectory.ReadOnly]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory
type VZSharedDirectory struct {
	objectivec.Object
}

// VZSharedDirectoryFromID constructs a [VZSharedDirectory] from an objc.ID.
func VZSharedDirectoryFromID(id objc.ID) VZSharedDirectory {
	return VZSharedDirectory{objectivec.Object{ID: id}}
}

// Ensure VZSharedDirectory implements IVZSharedDirectory.
var _ IVZSharedDirectory = VZSharedDirectory{}

// An interface definition for the [VZSharedDirectory] class.
//
// # Methods
//
//   - [IVZSharedDirectory.ReadOnly]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory
type IVZSharedDirectory interface {
	objectivec.IObject

	// Topic: Methods

	ReadOnly() bool
}

// Init initializes the instance.
func (s VZSharedDirectory) Init() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSharedDirectory) Autorelease() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSharedDirectory creates a new VZSharedDirectory instance.
func NewVZSharedDirectory() VZSharedDirectory {
	class := getVZSharedDirectoryClass()
	rv := objc.Send[VZSharedDirectory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/readOnly
func (s VZSharedDirectory) ReadOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("readOnly"))
	return rv
}
