// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZEFISignature] class.
var (
	_VZEFISignatureClass     VZEFISignatureClass
	_VZEFISignatureClassOnce sync.Once
)

func getVZEFISignatureClass() VZEFISignatureClass {
	_VZEFISignatureClassOnce.Do(func() {
		_VZEFISignatureClass = VZEFISignatureClass{class: objc.GetClass("VZEFISignature")}
	})
	return _VZEFISignatureClass
}

// GetVZEFISignatureClass returns the class object for VZEFISignature.
func GetVZEFISignatureClass() VZEFISignatureClass {
	return getVZEFISignatureClass()
}

type VZEFISignatureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZEFISignatureClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEFISignatureClass) Alloc() VZEFISignature {
	rv := objc.SendIfResponds[VZEFISignature](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZEFISignature._init]
type VZEFISignature struct {
	objectivec.Object
}

// VZEFISignatureFromID constructs a [VZEFISignature] from an objc.ID.
func VZEFISignatureFromID(id objc.ID) VZEFISignature {
	return VZEFISignature{objectivec.Object{ID: id}}
}

// Ensure VZEFISignature implements IVZEFISignature.
var _ IVZEFISignature = VZEFISignature{}

// An interface definition for the [VZEFISignature] class.
//
// # Methods
//
//   - [IVZEFISignature._init]
type IVZEFISignature interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
}

// Init initializes the instance.
func (v VZEFISignature) Init() VZEFISignature {
	rv := objc.SendIfResponds[VZEFISignature](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZEFISignature) Autorelease() VZEFISignature {
	rv := objc.SendIfResponds[VZEFISignature](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFISignature creates a new VZEFISignature instance.
func NewVZEFISignature() VZEFISignature {
	class := getVZEFISignatureClass()
	rv := objc.SendIfResponds[VZEFISignature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZEFISignature) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
