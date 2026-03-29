// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZBinaryBootLoaderSegment] class.
var (
	_VZBinaryBootLoaderSegmentClass     VZBinaryBootLoaderSegmentClass
	_VZBinaryBootLoaderSegmentClassOnce sync.Once
)

func getVZBinaryBootLoaderSegmentClass() VZBinaryBootLoaderSegmentClass {
	_VZBinaryBootLoaderSegmentClassOnce.Do(func() {
		_VZBinaryBootLoaderSegmentClass = VZBinaryBootLoaderSegmentClass{class: objc.GetClass("_VZBinaryBootLoaderSegment")}
	})
	return _VZBinaryBootLoaderSegmentClass
}

// GetVZBinaryBootLoaderSegmentClass returns the class object for _VZBinaryBootLoaderSegment.
func GetVZBinaryBootLoaderSegmentClass() VZBinaryBootLoaderSegmentClass {
	return getVZBinaryBootLoaderSegmentClass()
}

type VZBinaryBootLoaderSegmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBinaryBootLoaderSegmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBinaryBootLoaderSegmentClass) Alloc() VZBinaryBootLoaderSegment {
	rv := objc.Send[VZBinaryBootLoaderSegment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZBinaryBootLoaderSegment.BinaryURL]
//   - [VZBinaryBootLoaderSegment.LoadAddress]
//   - [VZBinaryBootLoaderSegment.InitWithBinaryURLLoadAddress]
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoaderSegment
type VZBinaryBootLoaderSegment struct {
	objectivec.Object
}

// VZBinaryBootLoaderSegmentFromID constructs a [VZBinaryBootLoaderSegment] from an objc.ID.
func VZBinaryBootLoaderSegmentFromID(id objc.ID) VZBinaryBootLoaderSegment {
	return VZBinaryBootLoaderSegment{objectivec.Object{ID: id}}
}
// Ensure VZBinaryBootLoaderSegment implements IVZBinaryBootLoaderSegment.
var _ IVZBinaryBootLoaderSegment = VZBinaryBootLoaderSegment{}

// An interface definition for the [VZBinaryBootLoaderSegment] class.
//
// # Methods
//
//   - [IVZBinaryBootLoaderSegment.BinaryURL]
//   - [IVZBinaryBootLoaderSegment.LoadAddress]
//   - [IVZBinaryBootLoaderSegment.InitWithBinaryURLLoadAddress]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoaderSegment
type IVZBinaryBootLoaderSegment interface {
	objectivec.IObject

	// Topic: Methods

	BinaryURL() foundation.INSURL
	LoadAddress() uint64
	InitWithBinaryURLLoadAddress(url foundation.INSURL, address uint64) VZBinaryBootLoaderSegment
}

// Init initializes the instance.
func (v VZBinaryBootLoaderSegment) Init() VZBinaryBootLoaderSegment {
	rv := objc.Send[VZBinaryBootLoaderSegment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZBinaryBootLoaderSegment) Autorelease() VZBinaryBootLoaderSegment {
	rv := objc.Send[VZBinaryBootLoaderSegment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBinaryBootLoaderSegment creates a new VZBinaryBootLoaderSegment instance.
func NewVZBinaryBootLoaderSegment() VZBinaryBootLoaderSegment {
	class := getVZBinaryBootLoaderSegmentClass()
	rv := objc.Send[VZBinaryBootLoaderSegment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoaderSegment/initWithBinaryURL:loadAddress:
func NewVZBinaryBootLoaderSegmentWithBinaryURLLoadAddress(url foundation.INSURL, address uint64) VZBinaryBootLoaderSegment {
	instance := getVZBinaryBootLoaderSegmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBinaryURL:loadAddress:"), url, address)
	return VZBinaryBootLoaderSegmentFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoaderSegment/initWithBinaryURL:loadAddress:
func (v VZBinaryBootLoaderSegment) InitWithBinaryURLLoadAddress(url foundation.INSURL, address uint64) VZBinaryBootLoaderSegment {
	rv := objc.Send[VZBinaryBootLoaderSegment](v.ID, objc.Sel("initWithBinaryURL:loadAddress:"), url, address)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoaderSegment/binaryURL
func (v VZBinaryBootLoaderSegment) BinaryURL() foundation.INSURL {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("binaryURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoaderSegment/loadAddress
func (v VZBinaryBootLoaderSegment) LoadAddress() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("loadAddress"))
	return rv
}

