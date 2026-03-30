// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacBatterySource] class.
var (
	_VZMacBatterySourceClass     VZMacBatterySourceClass
	_VZMacBatterySourceClassOnce sync.Once
)

func getVZMacBatterySourceClass() VZMacBatterySourceClass {
	_VZMacBatterySourceClassOnce.Do(func() {
		_VZMacBatterySourceClass = VZMacBatterySourceClass{class: objc.GetClass("_VZMacBatterySource")}
	})
	return _VZMacBatterySourceClass
}

// GetVZMacBatterySourceClass returns the class object for _VZMacBatterySource.
func GetVZMacBatterySourceClass() VZMacBatterySourceClass {
	return getVZMacBatterySourceClass()
}

type VZMacBatterySourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacBatterySourceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacBatterySourceClass) Alloc() VZMacBatterySource {
	rv := objc.Send[VZMacBatterySource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacBatterySource._init]
//   - [VZMacBatterySource._source]
//   - [VZMacBatterySource.EncodeWithEncoder]
//   - [VZMacBatterySource.RegisterObserver]
//   - [VZMacBatterySource.UnregisterObserver]
//   - [VZMacBatterySource.DebugDescription]
//   - [VZMacBatterySource.Description]
//   - [VZMacBatterySource.Hash]
//   - [VZMacBatterySource.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource
type VZMacBatterySource struct {
	objectivec.Object
}

// VZMacBatterySourceFromID constructs a [VZMacBatterySource] from an objc.ID.
func VZMacBatterySourceFromID(id objc.ID) VZMacBatterySource {
	return VZMacBatterySource{objectivec.Object{ID: id}}
}

// Ensure VZMacBatterySource implements IVZMacBatterySource.
var _ IVZMacBatterySource = VZMacBatterySource{}

// An interface definition for the [VZMacBatterySource] class.
//
// # Methods
//
//   - [IVZMacBatterySource._init]
//   - [IVZMacBatterySource._source]
//   - [IVZMacBatterySource.EncodeWithEncoder]
//   - [IVZMacBatterySource.RegisterObserver]
//   - [IVZMacBatterySource.UnregisterObserver]
//   - [IVZMacBatterySource.DebugDescription]
//   - [IVZMacBatterySource.Description]
//   - [IVZMacBatterySource.Hash]
//   - [IVZMacBatterySource.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource
type IVZMacBatterySource interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_source() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	RegisterObserver(observer objectivec.IObject)
	UnregisterObserver(observer objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZMacBatterySource) Init() VZMacBatterySource {
	rv := objc.Send[VZMacBatterySource](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacBatterySource) Autorelease() VZMacBatterySource {
	rv := objc.Send[VZMacBatterySource](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacBatterySource creates a new VZMacBatterySource instance.
func NewVZMacBatterySource() VZMacBatterySource {
	class := getVZMacBatterySourceClass()
	rv := objc.Send[VZMacBatterySource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource/_init
func (v VZMacBatterySource) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource/encodeWithEncoder:
func (v VZMacBatterySource) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource/registerObserver:
func (v VZMacBatterySource) RegisterObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("registerObserver:"), observer)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource/unregisterObserver:
func (v VZMacBatterySource) UnregisterObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("unregisterObserver:"), observer)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource/_source
func (v VZMacBatterySource) _source() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_source"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource/debugDescription
func (v VZMacBatterySource) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource/description
func (v VZMacBatterySource) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource/hash
func (v VZMacBatterySource) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatterySource/superclass
func (v VZMacBatterySource) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
