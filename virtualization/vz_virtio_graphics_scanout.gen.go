// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioGraphicsScanout] class.
var (
	_VZVirtioGraphicsScanoutClass     VZVirtioGraphicsScanoutClass
	_VZVirtioGraphicsScanoutClassOnce sync.Once
)

func getVZVirtioGraphicsScanoutClass() VZVirtioGraphicsScanoutClass {
	_VZVirtioGraphicsScanoutClassOnce.Do(func() {
		_VZVirtioGraphicsScanoutClass = VZVirtioGraphicsScanoutClass{class: objc.GetClass("VZVirtioGraphicsScanout")}
	})
	return _VZVirtioGraphicsScanoutClass
}

// GetVZVirtioGraphicsScanoutClass returns the class object for VZVirtioGraphicsScanout.
func GetVZVirtioGraphicsScanoutClass() VZVirtioGraphicsScanoutClass {
	return getVZVirtioGraphicsScanoutClass()
}

type VZVirtioGraphicsScanoutClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioGraphicsScanoutClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioGraphicsScanoutClass) Alloc() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A Virtio graphics scanout that corresponds to a Virtio graphics scanout
// configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanout
type VZVirtioGraphicsScanout struct {
	VZGraphicsDisplay
}

// VZVirtioGraphicsScanoutFromID constructs a [VZVirtioGraphicsScanout] from an objc.ID.
//
// A Virtio graphics scanout that corresponds to a Virtio graphics scanout
// configuration.
func VZVirtioGraphicsScanoutFromID(id objc.ID) VZVirtioGraphicsScanout {
	return VZVirtioGraphicsScanout{VZGraphicsDisplay: VZGraphicsDisplayFromID(id)}
}

// NOTE: VZVirtioGraphicsScanout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioGraphicsScanout] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanout
type IVZVirtioGraphicsScanout interface {
	IVZGraphicsDisplay
}

// Init initializes the instance.
func (v VZVirtioGraphicsScanout) Init() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioGraphicsScanout) Autorelease() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsScanout creates a new VZVirtioGraphicsScanout instance.
func NewVZVirtioGraphicsScanout() VZVirtioGraphicsScanout {
	class := getVZVirtioGraphicsScanoutClass()
	rv := objc.Send[VZVirtioGraphicsScanout](objc.ID(class.class), objc.Sel("new"))
	return rv
}
