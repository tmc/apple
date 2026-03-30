// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VirtualizationUniversalHIDFilters] class.
var (
	_VirtualizationUniversalHIDFiltersClass     VirtualizationUniversalHIDFiltersClass
	_VirtualizationUniversalHIDFiltersClassOnce sync.Once
)

func getVirtualizationUniversalHIDFiltersClass() VirtualizationUniversalHIDFiltersClass {
	_VirtualizationUniversalHIDFiltersClassOnce.Do(func() {
		_VirtualizationUniversalHIDFiltersClass = VirtualizationUniversalHIDFiltersClass{class: objc.GetClass("Virtualization.UniversalHIDFilters")}
	})
	return _VirtualizationUniversalHIDFiltersClass
}

// GetVirtualizationUniversalHIDFiltersClass returns the class object for Virtualization.UniversalHIDFilters.
func GetVirtualizationUniversalHIDFiltersClass() VirtualizationUniversalHIDFiltersClass {
	return getVirtualizationUniversalHIDFiltersClass()
}

type VirtualizationUniversalHIDFiltersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VirtualizationUniversalHIDFiltersClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VirtualizationUniversalHIDFiltersClass) Alloc() VirtualizationUniversalHIDFilters {
	rv := objc.Send[VirtualizationUniversalHIDFilters](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/Virtualization.UniversalHIDFilters
type VirtualizationUniversalHIDFilters struct {
	objectivec.Object
}

// VirtualizationUniversalHIDFiltersFromID constructs a [VirtualizationUniversalHIDFilters] from an objc.ID.
func VirtualizationUniversalHIDFiltersFromID(id objc.ID) VirtualizationUniversalHIDFilters {
	return VirtualizationUniversalHIDFilters{objectivec.Object{ID: id}}
}

// NOTE: VirtualizationUniversalHIDFilters struct embeds objectivec.Object (parent type unavailable) but
// IVirtualizationUniversalHIDFilters embeds the parent interface; skip compile-time assertion.

// An interface definition for the [VirtualizationUniversalHIDFilters] class.
//
// See: https://developer.apple.com/documentation/Virtualization/Virtualization.UniversalHIDFilters
type IVirtualizationUniversalHIDFilters interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VirtualizationUniversalHIDFilters) Init() VirtualizationUniversalHIDFilters {
	rv := objc.Send[VirtualizationUniversalHIDFilters](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VirtualizationUniversalHIDFilters) Autorelease() VirtualizationUniversalHIDFilters {
	rv := objc.Send[VirtualizationUniversalHIDFilters](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVirtualizationUniversalHIDFilters creates a new VirtualizationUniversalHIDFilters instance.
func NewVirtualizationUniversalHIDFilters() VirtualizationUniversalHIDFilters {
	class := getVirtualizationUniversalHIDFiltersClass()
	rv := objc.Send[VirtualizationUniversalHIDFilters](objc.ID(class.class), objc.Sel("new"))
	return rv
}
