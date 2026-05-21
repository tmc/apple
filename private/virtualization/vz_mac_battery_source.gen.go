// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

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
//   - [VZMacBatterySource.RegisterObserver]
//   - [VZMacBatterySource.UnregisterObserver]
//   - [VZMacBatterySource.DebugDescription]
//   - [VZMacBatterySource.Description]
//   - [VZMacBatterySource.Hash]
//   - [VZMacBatterySource.Superclass]
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
//   - [IVZMacBatterySource.RegisterObserver]
//   - [IVZMacBatterySource.UnregisterObserver]
//   - [IVZMacBatterySource.DebugDescription]
//   - [IVZMacBatterySource.Description]
//   - [IVZMacBatterySource.Hash]
//   - [IVZMacBatterySource.Superclass]
type IVZMacBatterySource interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_source() unsafe.Pointer
	RegisterObserver(observer objectivec.IObject)
	UnregisterObserver(observer objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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

func (v VZMacBatterySource) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZMacBatterySource) RegisterObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("registerObserver:"), observer)
}
func (v VZMacBatterySource) UnregisterObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("unregisterObserver:"), observer)
}

func (v VZMacBatterySource) _source() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("_source"))
	return rv
}

// CanSource reports whether the receiver responds to the private selector _source.
func (v VZMacBatterySource) CanSource() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_source"))
}

// Source is an exported wrapper for the private property _source.
func (v VZMacBatterySource) Source() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_source")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_source"}
	}
	return v._source(), nil
}
func (v VZMacBatterySource) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMacBatterySource) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMacBatterySource) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZMacBatterySource) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
