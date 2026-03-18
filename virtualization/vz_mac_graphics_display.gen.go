// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacGraphicsDisplay] class.
var (
	_VZMacGraphicsDisplayClass     VZMacGraphicsDisplayClass
	_VZMacGraphicsDisplayClassOnce sync.Once
)

func getVZMacGraphicsDisplayClass() VZMacGraphicsDisplayClass {
	_VZMacGraphicsDisplayClassOnce.Do(func() {
		_VZMacGraphicsDisplayClass = VZMacGraphicsDisplayClass{class: objc.GetClass("VZMacGraphicsDisplay")}
	})
	return _VZMacGraphicsDisplayClass
}

// GetVZMacGraphicsDisplayClass returns the class object for VZMacGraphicsDisplay.
func GetVZMacGraphicsDisplayClass() VZMacGraphicsDisplayClass {
	return getVZMacGraphicsDisplayClass()
}

type VZMacGraphicsDisplayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacGraphicsDisplayClass) Alloc() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the graphics display on a Mac.
//
// # Getting the display’s pixel density
//
//   - [VZMacGraphicsDisplay.PixelsPerInch]: Returns the pixel density of the display in pixels per inch.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay
type VZMacGraphicsDisplay struct {
	VZGraphicsDisplay
}

// VZMacGraphicsDisplayFromID constructs a [VZMacGraphicsDisplay] from an objc.ID.
//
// An object that represents the graphics display on a Mac.
func VZMacGraphicsDisplayFromID(id objc.ID) VZMacGraphicsDisplay {
	return VZMacGraphicsDisplay{VZGraphicsDisplay: VZGraphicsDisplayFromID(id)}
}
// NOTE: VZMacGraphicsDisplay adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacGraphicsDisplay] class.
//
// # Getting the display’s pixel density
//
//   - [IVZMacGraphicsDisplay.PixelsPerInch]: Returns the pixel density of the display in pixels per inch.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay
type IVZMacGraphicsDisplay interface {
	IVZGraphicsDisplay

	// Topic: Getting the display’s pixel density

	// Returns the pixel density of the display in pixels per inch.
	PixelsPerInch() int
}

// Init initializes the instance.
func (m VZMacGraphicsDisplay) Init() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacGraphicsDisplay) Autorelease() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDisplay creates a new VZMacGraphicsDisplay instance.
func NewVZMacGraphicsDisplay() VZMacGraphicsDisplay {
	class := getVZMacGraphicsDisplayClass()
	rv := objc.Send[VZMacGraphicsDisplay](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the pixel density of the display in pixels per inch.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay/pixelsPerInch
func (m VZMacGraphicsDisplay) PixelsPerInch() int {
	rv := objc.Send[int](m.ID, objc.Sel("pixelsPerInch"))
	return rv
}

