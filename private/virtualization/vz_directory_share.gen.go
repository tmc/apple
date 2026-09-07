// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZDirectoryShare._init]
//   - [VZDirectoryShare._share]
//   - [VZDirectoryShare.DebugDescription]
//   - [VZDirectoryShare.Description]
//   - [VZDirectoryShare.Hash]
//   - [VZDirectoryShare.Superclass]
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
type IVZDirectoryShare interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_share() unsafe.Pointer
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZDirectoryShare) Init() VZDirectoryShare {
	rv := objc.SendIfResponds[VZDirectoryShare](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDirectoryShare) Autorelease() VZDirectoryShare {
	rv := objc.SendIfResponds[VZDirectoryShare](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectoryShare creates a new VZDirectoryShare instance.
func NewVZDirectoryShare() VZDirectoryShare {
	class := getVZDirectoryShareClass()
	rv := objc.SendIfResponds[VZDirectoryShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZDirectoryShare) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZDirectoryShare) _share() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_share"))
	return rv
}

// CanShare reports whether the receiver responds to the private selector _share.
func (v VZDirectoryShare) CanShare() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_share"))
}

// Share is an exported wrapper for the private property _share.
func (v VZDirectoryShare) Share() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_share")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_share"}
	}
	return v._share(), nil
}
func (v VZDirectoryShare) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZDirectoryShare) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZDirectoryShare) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZDirectoryShare) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
