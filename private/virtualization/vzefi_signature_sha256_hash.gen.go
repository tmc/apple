// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZEFISignatureSHA256Hash] class.
var (
	_VZEFISignatureSHA256HashClass     VZEFISignatureSHA256HashClass
	_VZEFISignatureSHA256HashClassOnce sync.Once
)

func getVZEFISignatureSHA256HashClass() VZEFISignatureSHA256HashClass {
	_VZEFISignatureSHA256HashClassOnce.Do(func() {
		_VZEFISignatureSHA256HashClass = VZEFISignatureSHA256HashClass{class: objc.GetClass("VZEFISignatureSHA256Hash")}
	})
	return _VZEFISignatureSHA256HashClass
}

// GetVZEFISignatureSHA256HashClass returns the class object for VZEFISignatureSHA256Hash.
func GetVZEFISignatureSHA256HashClass() VZEFISignatureSHA256HashClass {
	return getVZEFISignatureSHA256HashClass()
}

type VZEFISignatureSHA256HashClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZEFISignatureSHA256HashClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEFISignatureSHA256HashClass) Alloc() VZEFISignatureSHA256Hash {
	rv := objc.SendIfResponds[VZEFISignatureSHA256Hash](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZEFISignatureSHA256Hash.Data]
//   - [VZEFISignatureSHA256Hash.InitWithData]
type VZEFISignatureSHA256Hash struct {
	VZEFISignature
}

// VZEFISignatureSHA256HashFromID constructs a [VZEFISignatureSHA256Hash] from an objc.ID.
func VZEFISignatureSHA256HashFromID(id objc.ID) VZEFISignatureSHA256Hash {
	return VZEFISignatureSHA256Hash{VZEFISignature: VZEFISignatureFromID(id)}
}

// Ensure VZEFISignatureSHA256Hash implements IVZEFISignatureSHA256Hash.
var _ IVZEFISignatureSHA256Hash = VZEFISignatureSHA256Hash{}

// An interface definition for the [VZEFISignatureSHA256Hash] class.
//
// # Methods
//
//   - [IVZEFISignatureSHA256Hash.Data]
//   - [IVZEFISignatureSHA256Hash.InitWithData]
type IVZEFISignatureSHA256Hash interface {
	IVZEFISignature

	// Topic: Methods

	Data() foundation.NSData
	InitWithData(data objectivec.IObject) VZEFISignatureSHA256Hash
}

// Init initializes the instance.
func (v VZEFISignatureSHA256Hash) Init() VZEFISignatureSHA256Hash {
	rv := objc.SendIfResponds[VZEFISignatureSHA256Hash](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZEFISignatureSHA256Hash) Autorelease() VZEFISignatureSHA256Hash {
	rv := objc.SendIfResponds[VZEFISignatureSHA256Hash](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFISignatureSHA256Hash creates a new VZEFISignatureSHA256Hash instance.
func NewVZEFISignatureSHA256Hash() VZEFISignatureSHA256Hash {
	class := getVZEFISignatureSHA256HashClass()
	rv := objc.SendIfResponds[VZEFISignatureSHA256Hash](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZEFISignatureSHA256HashWithData(data objectivec.IObject) VZEFISignatureSHA256Hash {
	instance := getVZEFISignatureSHA256HashClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return VZEFISignatureSHA256HashFromID(rv)
}

func (v VZEFISignatureSHA256Hash) InitWithData(data objectivec.IObject) VZEFISignatureSHA256Hash {
	rv := objc.SendIfResponds[VZEFISignatureSHA256Hash](v.ID, objc.Sel("initWithData:"), data)
	return rv
}

func (v VZEFISignatureSHA256Hash) Data() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](v.ID, objc.Sel("data"))
	return foundation.NSData(rv)
}
