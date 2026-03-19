// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGraphicsDisplay] class.
var (
	_VZGraphicsDisplayClass     VZGraphicsDisplayClass
	_VZGraphicsDisplayClassOnce sync.Once
)

func getVZGraphicsDisplayClass() VZGraphicsDisplayClass {
	_VZGraphicsDisplayClassOnce.Do(func() {
		_VZGraphicsDisplayClass = VZGraphicsDisplayClass{class: objc.GetClass("VZGraphicsDisplay")}
	})
	return _VZGraphicsDisplayClass
}

// GetVZGraphicsDisplayClass returns the class object for VZGraphicsDisplay.
func GetVZGraphicsDisplayClass() VZGraphicsDisplayClass {
	return getVZGraphicsDisplayClass()
}

type VZGraphicsDisplayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGraphicsDisplayClass) Alloc() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a graphics display in a VM.
//
// # Overview
// 
// Don’t instantiate a [VZGraphicsDisplay] directly. Graphics displays are
// first configured on a [VZGraphicsDeviceConfiguration] subclass. When you
// create a [VZVirtualMachine] from the configuration, the displays are
// available through the [VZGraphicsDisplay.Displays] property of the configuration’s
// [VZGraphicsDevice].
//
// # Getting the display size
//
//   - [VZGraphicsDisplay.SizeInPixels]: Returns the size of the display, in pixels.
//
// # Observing changes to the display configuration
//
//   - [VZGraphicsDisplay.AddObserver]: Adds an observer to notify about display configuration changes.
//   - [VZGraphicsDisplay.RemoveObserver]: Removes a display configuration change observer.
//
// # Changing the display configuration
//
//   - [VZGraphicsDisplay.ReconfigureWithSizeInPixelsError]: Resize this display with the new dimensions you provide.
//   - [VZGraphicsDisplay.ReconfigureWithConfigurationError]: Reconfigure this display with the new display configuration you provide.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay
type VZGraphicsDisplay struct {
	objectivec.Object
}

// VZGraphicsDisplayFromID constructs a [VZGraphicsDisplay] from an objc.ID.
//
// A class that represents a graphics display in a VM.
func VZGraphicsDisplayFromID(id objc.ID) VZGraphicsDisplay {
	return VZGraphicsDisplay{objectivec.Object{ID: id}}
}
// NOTE: VZGraphicsDisplay adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZGraphicsDisplay] class.
//
// # Getting the display size
//
//   - [IVZGraphicsDisplay.SizeInPixels]: Returns the size of the display, in pixels.
//
// # Observing changes to the display configuration
//
//   - [IVZGraphicsDisplay.AddObserver]: Adds an observer to notify about display configuration changes.
//   - [IVZGraphicsDisplay.RemoveObserver]: Removes a display configuration change observer.
//
// # Changing the display configuration
//
//   - [IVZGraphicsDisplay.ReconfigureWithSizeInPixelsError]: Resize this display with the new dimensions you provide.
//   - [IVZGraphicsDisplay.ReconfigureWithConfigurationError]: Reconfigure this display with the new display configuration you provide.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay
type IVZGraphicsDisplay interface {
	objectivec.IObject

	// Topic: Getting the display size

	// Returns the size of the display, in pixels.
	SizeInPixels() corefoundation.CGSize

	// Topic: Observing changes to the display configuration

	// Adds an observer to notify about display configuration changes.
	AddObserver(observer VZGraphicsDisplayObserver)
	// Removes a display configuration change observer.
	RemoveObserver(observer VZGraphicsDisplayObserver)

	// Topic: Changing the display configuration

	// Resize this display with the new dimensions you provide.
	ReconfigureWithSizeInPixelsError(sizeInPixels corefoundation.CGSize) (bool, error)
	// Reconfigure this display with the new display configuration you provide.
	ReconfigureWithConfigurationError(configuration IVZGraphicsDisplayConfiguration) (bool, error)

	// The list of graphics displays configured for this graphics device.
	Displays() IVZGraphicsDisplay
	SetDisplays(value IVZGraphicsDisplay)
}

// Init initializes the instance.
func (g VZGraphicsDisplay) Init() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGraphicsDisplay) Autorelease() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDisplay creates a new VZGraphicsDisplay instance.
func NewVZGraphicsDisplay() VZGraphicsDisplay {
	class := getVZGraphicsDisplayClass()
	rv := objc.Send[VZGraphicsDisplay](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds an observer to notify about display configuration changes.
//
// observer: The new observer the framework notifies about display configuration
// changes.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/addObserver(_:)
func (g VZGraphicsDisplay) AddObserver(observer VZGraphicsDisplayObserver) {
	objc.Send[objc.ID](g.ID, objc.Sel("addObserver:"), observer)
}
// Removes a display configuration change observer.
//
// observer: The observer to remove from notifications about display configuration
// changes.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/removeObserver(_:)
func (g VZGraphicsDisplay) RemoveObserver(observer VZGraphicsDisplayObserver) {
	objc.Send[objc.ID](g.ID, objc.Sel("removeObserver:"), observer)
}
// Resize this display with the new dimensions you provide.
//
// sizeInPixels: The new display width and height in pixels.
//
// # Discussion
// 
// If successful, the framework passes the new size to the guest but the guest
// may or may not respond to the new size. If the guest doesn’t use the new
// size, the Virtualization framework doesn’t return an error.
// 
// Resizing the display triggers a display state change that you can track by
// adopting the [VZGraphicsDisplayObserver] protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/reconfigure(sizeInPixels:)
func (g VZGraphicsDisplay) ReconfigureWithSizeInPixelsError(sizeInPixels corefoundation.CGSize) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](g.ID, objc.Sel("reconfigureWithSizeInPixels:error:"), sizeInPixels, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("reconfigureWithSizeInPixels:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Reconfigure this display with the new display configuration you provide.
//
// configuration: The new [VZGraphicsDisplayConfiguration] configuration.
//
// # Discussion
// 
// If successful, the framework passes the new configuration to the guest, but
// the guest may or may not respond to parts of the configuration. If the
// guest doesn’t use the new configuration, the Virtualization framework
// doesn’t return an error.
// 
// Reconfiguration of the display triggers a display state change that you can
// track by adopting the [VZGraphicsDisplayObserver] protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/reconfigure(configuration:)
func (g VZGraphicsDisplay) ReconfigureWithConfigurationError(configuration IVZGraphicsDisplayConfiguration) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](g.ID, objc.Sel("reconfigureWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("reconfigureWithConfiguration:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the size of the display, in pixels.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/sizeInPixels
func (g VZGraphicsDisplay) SizeInPixels() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](g.ID, objc.Sel("sizeInPixels"))
	return corefoundation.CGSize(rv)
}
// The list of graphics displays configured for this graphics device.
//
// See: https://developer.apple.com/documentation/virtualization/vzgraphicsdevice/displays
func (g VZGraphicsDisplay) Displays() IVZGraphicsDisplay {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("displays"))
	return VZGraphicsDisplayFromID(objc.ID(rv))
}
func (g VZGraphicsDisplay) SetDisplays(value IVZGraphicsDisplay) {
	objc.Send[struct{}](g.ID, objc.Sel("setDisplays:"), value)
}

