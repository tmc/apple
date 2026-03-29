// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VirtualizationUniversalHIDKitTranslators] class.
var (
	_VirtualizationUniversalHIDKitTranslatorsClass     VirtualizationUniversalHIDKitTranslatorsClass
	_VirtualizationUniversalHIDKitTranslatorsClassOnce sync.Once
)

func getVirtualizationUniversalHIDKitTranslatorsClass() VirtualizationUniversalHIDKitTranslatorsClass {
	_VirtualizationUniversalHIDKitTranslatorsClassOnce.Do(func() {
		_VirtualizationUniversalHIDKitTranslatorsClass = VirtualizationUniversalHIDKitTranslatorsClass{class: objc.GetClass("Virtualization.UniversalHIDKitTranslators")}
	})
	return _VirtualizationUniversalHIDKitTranslatorsClass
}

// GetVirtualizationUniversalHIDKitTranslatorsClass returns the class object for Virtualization.UniversalHIDKitTranslators.
func GetVirtualizationUniversalHIDKitTranslatorsClass() VirtualizationUniversalHIDKitTranslatorsClass {
	return getVirtualizationUniversalHIDKitTranslatorsClass()
}

type VirtualizationUniversalHIDKitTranslatorsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VirtualizationUniversalHIDKitTranslatorsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VirtualizationUniversalHIDKitTranslatorsClass) Alloc() VirtualizationUniversalHIDKitTranslators {
	rv := objc.Send[VirtualizationUniversalHIDKitTranslators](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/Virtualization.UniversalHIDKitTranslators
type VirtualizationUniversalHIDKitTranslators struct {
	objectivec.Object
}

// VirtualizationUniversalHIDKitTranslatorsFromID constructs a [VirtualizationUniversalHIDKitTranslators] from an objc.ID.
func VirtualizationUniversalHIDKitTranslatorsFromID(id objc.ID) VirtualizationUniversalHIDKitTranslators {
	return VirtualizationUniversalHIDKitTranslators{objectivec.Object{ID: id}}
}
// NOTE: VirtualizationUniversalHIDKitTranslators struct embeds objectivec.Object (parent type unavailable) but
// IVirtualizationUniversalHIDKitTranslators embeds the parent interface; skip compile-time assertion.

// An interface definition for the [VirtualizationUniversalHIDKitTranslators] class.
//
// See: https://developer.apple.com/documentation/Virtualization/Virtualization.UniversalHIDKitTranslators
type IVirtualizationUniversalHIDKitTranslators interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VirtualizationUniversalHIDKitTranslators) Init() VirtualizationUniversalHIDKitTranslators {
	rv := objc.Send[VirtualizationUniversalHIDKitTranslators](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VirtualizationUniversalHIDKitTranslators) Autorelease() VirtualizationUniversalHIDKitTranslators {
	rv := objc.Send[VirtualizationUniversalHIDKitTranslators](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVirtualizationUniversalHIDKitTranslators creates a new VirtualizationUniversalHIDKitTranslators instance.
func NewVirtualizationUniversalHIDKitTranslators() VirtualizationUniversalHIDKitTranslators {
	class := getVirtualizationUniversalHIDKitTranslatorsClass()
	rv := objc.Send[VirtualizationUniversalHIDKitTranslators](objc.ID(class.class), objc.Sel("new"))
	return rv
}

