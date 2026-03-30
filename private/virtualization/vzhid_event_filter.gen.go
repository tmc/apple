// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZHIDEventFilter] class.
var (
	_VZHIDEventFilterClass     VZHIDEventFilterClass
	_VZHIDEventFilterClassOnce sync.Once
)

func getVZHIDEventFilterClass() VZHIDEventFilterClass {
	_VZHIDEventFilterClassOnce.Do(func() {
		_VZHIDEventFilterClass = VZHIDEventFilterClass{class: objc.GetClass("_VZHIDEventFilter")}
	})
	return _VZHIDEventFilterClass
}

// GetVZHIDEventFilterClass returns the class object for _VZHIDEventFilter.
func GetVZHIDEventFilterClass() VZHIDEventFilterClass {
	return getVZHIDEventFilterClass()
}

type VZHIDEventFilterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZHIDEventFilterClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZHIDEventFilterClass) Alloc() VZHIDEventFilter {
	rv := objc.Send[VZHIDEventFilter](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZHIDEventFilter.GetHIDReportsFromHIDEvent]
//   - [VZHIDEventFilter.GetHIDReportsFromNSEvent]
//   - [VZHIDEventFilter.UpdateCoordinateTransformIsFlipped]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZHIDEventFilter
type VZHIDEventFilter struct {
	objectivec.Object
}

// VZHIDEventFilterFromID constructs a [VZHIDEventFilter] from an objc.ID.
func VZHIDEventFilterFromID(id objc.ID) VZHIDEventFilter {
	return VZHIDEventFilter{objectivec.Object{ID: id}}
}

// Ensure VZHIDEventFilter implements IVZHIDEventFilter.
var _ IVZHIDEventFilter = VZHIDEventFilter{}

// An interface definition for the [VZHIDEventFilter] class.
//
// # Methods
//
//   - [IVZHIDEventFilter.GetHIDReportsFromHIDEvent]
//   - [IVZHIDEventFilter.GetHIDReportsFromNSEvent]
//   - [IVZHIDEventFilter.UpdateCoordinateTransformIsFlipped]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZHIDEventFilter
type IVZHIDEventFilter interface {
	objectivec.IObject

	// Topic: Methods

	GetHIDReportsFromHIDEvent(hIDEvent objectivec.IObject) objectivec.IObject
	GetHIDReportsFromNSEvent(nSEvent objectivec.IObject) objectivec.IObject
	UpdateCoordinateTransformIsFlipped(transform corefoundation.CGRect, flipped bool)
}

// Init initializes the instance.
func (v VZHIDEventFilter) Init() VZHIDEventFilter {
	rv := objc.Send[VZHIDEventFilter](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZHIDEventFilter) Autorelease() VZHIDEventFilter {
	rv := objc.Send[VZHIDEventFilter](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHIDEventFilter creates a new VZHIDEventFilter instance.
func NewVZHIDEventFilter() VZHIDEventFilter {
	class := getVZHIDEventFilterClass()
	rv := objc.Send[VZHIDEventFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDEventFilter/getHIDReportsFromHIDEvent:
func (v VZHIDEventFilter) GetHIDReportsFromHIDEvent(hIDEvent objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("getHIDReportsFromHIDEvent:"), hIDEvent)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDEventFilter/getHIDReportsFromNSEvent:
func (v VZHIDEventFilter) GetHIDReportsFromNSEvent(nSEvent objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("getHIDReportsFromNSEvent:"), nSEvent)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDEventFilter/updateCoordinateTransform:isFlipped:
func (v VZHIDEventFilter) UpdateCoordinateTransformIsFlipped(transform corefoundation.CGRect, flipped bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("updateCoordinateTransform:isFlipped:"), transform, flipped)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDEventFilter/hasEventTranslators
func (_VZHIDEventFilterClass VZHIDEventFilterClass) HasEventTranslators() bool {
	rv := objc.Send[bool](objc.ID(_VZHIDEventFilterClass.class), objc.Sel("hasEventTranslators"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDEventFilter/isEnabled
func (_VZHIDEventFilterClass VZHIDEventFilterClass) IsEnabled() bool {
	rv := objc.Send[bool](objc.ID(_VZHIDEventFilterClass.class), objc.Sel("isEnabled"))
	return rv
}
