// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSpiceAgentCoreCapabilities] class.
var (
	_VZSpiceAgentCoreCapabilitiesClass     VZSpiceAgentCoreCapabilitiesClass
	_VZSpiceAgentCoreCapabilitiesClassOnce sync.Once
)

func getVZSpiceAgentCoreCapabilitiesClass() VZSpiceAgentCoreCapabilitiesClass {
	_VZSpiceAgentCoreCapabilitiesClassOnce.Do(func() {
		_VZSpiceAgentCoreCapabilitiesClass = VZSpiceAgentCoreCapabilitiesClass{class: objc.GetClass("_VZSpiceAgentCoreCapabilities")}
	})
	return _VZSpiceAgentCoreCapabilitiesClass
}

// GetVZSpiceAgentCoreCapabilitiesClass returns the class object for _VZSpiceAgentCoreCapabilities.
func GetVZSpiceAgentCoreCapabilitiesClass() VZSpiceAgentCoreCapabilitiesClass {
	return getVZSpiceAgentCoreCapabilitiesClass()
}

type VZSpiceAgentCoreCapabilitiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSpiceAgentCoreCapabilitiesClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSpiceAgentCoreCapabilitiesClass) Alloc() VZSpiceAgentCoreCapabilities {
	rv := objc.Send[VZSpiceAgentCoreCapabilities](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZSpiceAgentCoreCapabilities.Clipboard]
//   - [VZSpiceAgentCoreCapabilities.SetClipboard]
// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCoreCapabilities
type VZSpiceAgentCoreCapabilities struct {
	objectivec.Object
}

// VZSpiceAgentCoreCapabilitiesFromID constructs a [VZSpiceAgentCoreCapabilities] from an objc.ID.
func VZSpiceAgentCoreCapabilitiesFromID(id objc.ID) VZSpiceAgentCoreCapabilities {
	return VZSpiceAgentCoreCapabilities{objectivec.Object{ID: id}}
}
// Ensure VZSpiceAgentCoreCapabilities implements IVZSpiceAgentCoreCapabilities.
var _ IVZSpiceAgentCoreCapabilities = VZSpiceAgentCoreCapabilities{}

// An interface definition for the [VZSpiceAgentCoreCapabilities] class.
//
// # Methods
//
//   - [IVZSpiceAgentCoreCapabilities.Clipboard]
//   - [IVZSpiceAgentCoreCapabilities.SetClipboard]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCoreCapabilities
type IVZSpiceAgentCoreCapabilities interface {
	objectivec.IObject

	// Topic: Methods

	Clipboard() bool
	SetClipboard(value bool)
}

// Init initializes the instance.
func (v VZSpiceAgentCoreCapabilities) Init() VZSpiceAgentCoreCapabilities {
	rv := objc.Send[VZSpiceAgentCoreCapabilities](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSpiceAgentCoreCapabilities) Autorelease() VZSpiceAgentCoreCapabilities {
	rv := objc.Send[VZSpiceAgentCoreCapabilities](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSpiceAgentCoreCapabilities creates a new VZSpiceAgentCoreCapabilities instance.
func NewVZSpiceAgentCoreCapabilities() VZSpiceAgentCoreCapabilities {
	class := getVZSpiceAgentCoreCapabilitiesClass()
	rv := objc.Send[VZSpiceAgentCoreCapabilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCoreCapabilities/clipboard
func (v VZSpiceAgentCoreCapabilities) Clipboard() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("clipboard"))
	return rv
}
func (v VZSpiceAgentCoreCapabilities) SetClipboard(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setClipboard:"), value)
}

