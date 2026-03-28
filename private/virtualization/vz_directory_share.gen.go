// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDirectoryShare] class.
var (
	_VZDirectoryShareClass     VZDirectoryShareClass
	_VZDirectoryShareClassOnce sync.Once
)

func getVZDirectoryShareClass() VZDirectoryShareClass {
	_VZDirectoryShareClassOnce.Do(func() {
		_VZDirectoryShareClass = VZDirectoryShareClass{class: objc.GetClass("VZDirectoryShare")}
	})
	return _VZDirectoryShareClass
}

// GetVZDirectoryShareClass returns the class object for VZDirectoryShare.
func GetVZDirectoryShareClass() VZDirectoryShareClass {
	return getVZDirectoryShareClass()
}

type VZDirectoryShareClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDirectoryShareClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDirectoryShareClass) Alloc() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZDirectoryShare._init]
//   - [VZDirectoryShare._share]
//   - [VZDirectoryShare.DebugDescription]
//   - [VZDirectoryShare.Description]
//   - [VZDirectoryShare.Hash]
//   - [VZDirectoryShare.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare
type VZDirectoryShare struct {
	objectivec.Object
}

// VZDirectoryShareFromID constructs a [VZDirectoryShare] from an objc.ID.
func VZDirectoryShareFromID(id objc.ID) VZDirectoryShare {
	return VZDirectoryShare{objectivec.Object{ID: id}}
}
// Ensure VZDirectoryShare implements IVZDirectoryShare.
var _ IVZDirectoryShare = VZDirectoryShare{}

// An interface definition for the [VZDirectoryShare] class.
//
// # Methods
//
//   - [IVZDirectoryShare._init]
//   - [IVZDirectoryShare._share]
//   - [IVZDirectoryShare.DebugDescription]
//   - [IVZDirectoryShare.Description]
//   - [IVZDirectoryShare.Hash]
//   - [IVZDirectoryShare.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare
type IVZDirectoryShare interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_share() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (d VZDirectoryShare) Init() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VZDirectoryShare) Autorelease() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectoryShare creates a new VZDirectoryShare instance.
func NewVZDirectoryShare() VZDirectoryShare {
	class := getVZDirectoryShareClass()
	rv := objc.Send[VZDirectoryShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare/_init
func (d VZDirectoryShare) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare/_share
func (d VZDirectoryShare) _share() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("_share"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare/debugDescription
func (d VZDirectoryShare) DebugDescription() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare/description
func (d VZDirectoryShare) Description() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare/hash
func (d VZDirectoryShare) Hash() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare/superclass
func (d VZDirectoryShare) Superclass() objc.Class {
	rv := objc.Send[objc.Class](d.ID, objc.Sel("superclass"))
	return rv
}

