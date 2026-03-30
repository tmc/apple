// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMMIORegion] class.
var (
	_VZMMIORegionClass     VZMMIORegionClass
	_VZMMIORegionClassOnce sync.Once
)

func getVZMMIORegionClass() VZMMIORegionClass {
	_VZMMIORegionClassOnce.Do(func() {
		_VZMMIORegionClass = VZMMIORegionClass{class: objc.GetClass("_VZMMIORegion")}
	})
	return _VZMMIORegionClass
}

// GetVZMMIORegionClass returns the class object for _VZMMIORegion.
func GetVZMMIORegionClass() VZMMIORegionClass {
	return getVZMMIORegionClass()
}

type VZMMIORegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMMIORegionClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMMIORegionClass) Alloc() VZMMIORegion {
	rv := objc.Send[VZMMIORegion](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMMIORegion.BaseAddress]
//   - [VZMMIORegion.EncodeWithEncoder]
//   - [VZMMIORegion.Length]
//   - [VZMMIORegion.WriteSynchronously]
//   - [VZMMIORegion.InitWithBaseAddressLength]
//   - [VZMMIORegion.InitWithBaseAddressLengthWriteSynchronously]
//   - [VZMMIORegion.DebugDescription]
//   - [VZMMIORegion.Description]
//   - [VZMMIORegion.Hash]
//   - [VZMMIORegion.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion
type VZMMIORegion struct {
	objectivec.Object
}

// VZMMIORegionFromID constructs a [VZMMIORegion] from an objc.ID.
func VZMMIORegionFromID(id objc.ID) VZMMIORegion {
	return VZMMIORegion{objectivec.Object{ID: id}}
}

// Ensure VZMMIORegion implements IVZMMIORegion.
var _ IVZMMIORegion = VZMMIORegion{}

// An interface definition for the [VZMMIORegion] class.
//
// # Methods
//
//   - [IVZMMIORegion.BaseAddress]
//   - [IVZMMIORegion.EncodeWithEncoder]
//   - [IVZMMIORegion.Length]
//   - [IVZMMIORegion.WriteSynchronously]
//   - [IVZMMIORegion.InitWithBaseAddressLength]
//   - [IVZMMIORegion.InitWithBaseAddressLengthWriteSynchronously]
//   - [IVZMMIORegion.DebugDescription]
//   - [IVZMMIORegion.Description]
//   - [IVZMMIORegion.Hash]
//   - [IVZMMIORegion.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion
type IVZMMIORegion interface {
	objectivec.IObject

	// Topic: Methods

	BaseAddress() uint64
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	Length() uint64
	WriteSynchronously() bool
	InitWithBaseAddressLength(address uint64, length uint64) VZMMIORegion
	InitWithBaseAddressLengthWriteSynchronously(address uint64, length uint64, synchronously bool) VZMMIORegion
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZMMIORegion) Init() VZMMIORegion {
	rv := objc.Send[VZMMIORegion](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMMIORegion) Autorelease() VZMMIORegion {
	rv := objc.Send[VZMMIORegion](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMMIORegion creates a new VZMMIORegion instance.
func NewVZMMIORegion() VZMMIORegion {
	class := getVZMMIORegionClass()
	rv := objc.Send[VZMMIORegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/initWithBaseAddress:length:
func NewVZMMIORegionWithBaseAddressLength(address uint64, length uint64) VZMMIORegion {
	instance := getVZMMIORegionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBaseAddress:length:"), address, length)
	return VZMMIORegionFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/initWithBaseAddress:length:writeSynchronously:
func NewVZMMIORegionWithBaseAddressLengthWriteSynchronously(address uint64, length uint64, synchronously bool) VZMMIORegion {
	instance := getVZMMIORegionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBaseAddress:length:writeSynchronously:"), address, length, synchronously)
	return VZMMIORegionFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/encodeWithEncoder:
func (v VZMMIORegion) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/initWithBaseAddress:length:
func (v VZMMIORegion) InitWithBaseAddressLength(address uint64, length uint64) VZMMIORegion {
	rv := objc.Send[VZMMIORegion](v.ID, objc.Sel("initWithBaseAddress:length:"), address, length)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/initWithBaseAddress:length:writeSynchronously:
func (v VZMMIORegion) InitWithBaseAddressLengthWriteSynchronously(address uint64, length uint64, synchronously bool) VZMMIORegion {
	rv := objc.Send[VZMMIORegion](v.ID, objc.Sel("initWithBaseAddress:length:writeSynchronously:"), address, length, synchronously)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/baseAddress
func (v VZMMIORegion) BaseAddress() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("baseAddress"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/debugDescription
func (v VZMMIORegion) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/description
func (v VZMMIORegion) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/hash
func (v VZMMIORegion) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/length
func (v VZMMIORegion) Length() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("length"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/superclass
func (v VZMMIORegion) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMMIORegion/writeSynchronously
func (v VZMMIORegion) WriteSynchronously() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("writeSynchronously"))
	return rv
}
