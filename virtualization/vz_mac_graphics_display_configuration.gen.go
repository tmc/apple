// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
)

// The class instance for the [VZMacGraphicsDisplayConfiguration] class.
var (
	_VZMacGraphicsDisplayConfigurationClass     VZMacGraphicsDisplayConfigurationClass
	_VZMacGraphicsDisplayConfigurationClassOnce sync.Once
)

func getVZMacGraphicsDisplayConfigurationClass() VZMacGraphicsDisplayConfigurationClass {
	_VZMacGraphicsDisplayConfigurationClassOnce.Do(func() {
		_VZMacGraphicsDisplayConfigurationClass = VZMacGraphicsDisplayConfigurationClass{class: objc.GetClass("VZMacGraphicsDisplayConfiguration")}
	})
	return _VZMacGraphicsDisplayConfigurationClass
}

// GetVZMacGraphicsDisplayConfigurationClass returns the class object for VZMacGraphicsDisplayConfiguration.
func GetVZMacGraphicsDisplayConfigurationClass() VZMacGraphicsDisplayConfigurationClass {
	return getVZMacGraphicsDisplayConfigurationClass()
}

type VZMacGraphicsDisplayConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacGraphicsDisplayConfigurationClass) Alloc() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The configuration for a Mac graphics device.
//
// # Overview
// 
// Use this device to attach a display that’s shown in a
// [VZVirtualMachineView].
//
// # Creating the display configuration
//
//   - [VZMacGraphicsDisplayConfiguration.InitForScreenSizeInPoints]: Create a display configuration suitable for showing on the specified screen.
//   - [VZMacGraphicsDisplayConfiguration.InitWithWidthInPixelsHeightInPixelsPixelsPerInch]: Create a display configuration with the specified pixel dimensions and pixel density.
//
// # Configuring the display properties
//
//   - [VZMacGraphicsDisplayConfiguration.HeightInPixels]: The height of the display, in pixels.
//   - [VZMacGraphicsDisplayConfiguration.SetHeightInPixels]
//   - [VZMacGraphicsDisplayConfiguration.WidthInPixels]: The width of the display, in pixels.
//   - [VZMacGraphicsDisplayConfiguration.SetWidthInPixels]
//   - [VZMacGraphicsDisplayConfiguration.PixelsPerInch]: The pixel density in pixels per inch.
//   - [VZMacGraphicsDisplayConfiguration.SetPixelsPerInch]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration
type VZMacGraphicsDisplayConfiguration struct {
	VZGraphicsDisplayConfiguration
}

// VZMacGraphicsDisplayConfigurationFromID constructs a [VZMacGraphicsDisplayConfiguration] from an objc.ID.
//
// The configuration for a Mac graphics device.
func VZMacGraphicsDisplayConfigurationFromID(id objc.ID) VZMacGraphicsDisplayConfiguration {
	return VZMacGraphicsDisplayConfiguration{VZGraphicsDisplayConfiguration: VZGraphicsDisplayConfigurationFromID(id)}
}
// NOTE: VZMacGraphicsDisplayConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZMacGraphicsDisplayConfiguration] class.
//
// # Creating the display configuration
//
//   - [IVZMacGraphicsDisplayConfiguration.InitForScreenSizeInPoints]: Create a display configuration suitable for showing on the specified screen.
//   - [IVZMacGraphicsDisplayConfiguration.InitWithWidthInPixelsHeightInPixelsPixelsPerInch]: Create a display configuration with the specified pixel dimensions and pixel density.
//
// # Configuring the display properties
//
//   - [IVZMacGraphicsDisplayConfiguration.HeightInPixels]: The height of the display, in pixels.
//   - [IVZMacGraphicsDisplayConfiguration.SetHeightInPixels]
//   - [IVZMacGraphicsDisplayConfiguration.WidthInPixels]: The width of the display, in pixels.
//   - [IVZMacGraphicsDisplayConfiguration.SetWidthInPixels]
//   - [IVZMacGraphicsDisplayConfiguration.PixelsPerInch]: The pixel density in pixels per inch.
//   - [IVZMacGraphicsDisplayConfiguration.SetPixelsPerInch]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration
type IVZMacGraphicsDisplayConfiguration interface {
	IVZGraphicsDisplayConfiguration

	// Topic: Creating the display configuration

	// Create a display configuration suitable for showing on the specified screen.
	InitForScreenSizeInPoints(screen appkit.NSScreen, sizeInPoints corefoundation.CGSize) VZMacGraphicsDisplayConfiguration
	// Create a display configuration with the specified pixel dimensions and pixel density.
	InitWithWidthInPixelsHeightInPixelsPixelsPerInch(widthInPixels int, heightInPixels int, pixelsPerInch int) VZMacGraphicsDisplayConfiguration

	// Topic: Configuring the display properties

	// The height of the display, in pixels.
	HeightInPixels() int
	SetHeightInPixels(value int)
	// The width of the display, in pixels.
	WidthInPixels() int
	SetWidthInPixels(value int)
	// The pixel density in pixels per inch.
	PixelsPerInch() int
	SetPixelsPerInch(value int)
}





// Init initializes the instance.
func (m VZMacGraphicsDisplayConfiguration) Init() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacGraphicsDisplayConfiguration) Autorelease() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDisplayConfiguration creates a new VZMacGraphicsDisplayConfiguration instance.
func NewVZMacGraphicsDisplayConfiguration() VZMacGraphicsDisplayConfiguration {
	class := getVZMacGraphicsDisplayConfigurationClass()
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Create a display configuration suitable for showing on the specified
// screen.
//
// screen: The screen on which you intend to present the [VZVirtualMachineView] for
// the display.
//
// sizeInPoints: The intended logical size of the display.
//
// # Discussion
// 
// The framework initializes the pixel dimensions and pixel density based on
// the specified screen and size. An instance of macOS running in the VM may
// not necessarily provide a display mode with a backing scale factor matching
// the specified screen.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/init(for:sizeInPoints:)
func NewMacGraphicsDisplayConfigurationForScreenSizeInPoints(screen appkit.NSScreen, sizeInPoints corefoundation.CGSize) VZMacGraphicsDisplayConfiguration {
	instance := getVZMacGraphicsDisplayConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForScreen:sizeInPoints:"), screen, sizeInPoints)
	return VZMacGraphicsDisplayConfigurationFromID(rv)
}


// Create a display configuration with the specified pixel dimensions and
// pixel density.
//
// widthInPixels: The width of the display, in pixels.
//
// heightInPixels: The height of the display, in pixels.
//
// pixelsPerInch: The pixel density as a number of pixels per inch.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/init(widthInPixels:heightInPixels:pixelsPerInch:)
func NewMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch(widthInPixels int, heightInPixels int, pixelsPerInch int) VZMacGraphicsDisplayConfiguration {
	instance := getVZMacGraphicsDisplayConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWidthInPixels:heightInPixels:pixelsPerInch:"), widthInPixels, heightInPixels, pixelsPerInch)
	return VZMacGraphicsDisplayConfigurationFromID(rv)
}







// Create a display configuration suitable for showing on the specified
// screen.
//
// screen: The screen on which you intend to present the [VZVirtualMachineView] for
// the display.
//
// sizeInPoints: The intended logical size of the display.
//
// # Discussion
// 
// The framework initializes the pixel dimensions and pixel density based on
// the specified screen and size. An instance of macOS running in the VM may
// not necessarily provide a display mode with a backing scale factor matching
// the specified screen.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/init(for:sizeInPoints:)
func (m VZMacGraphicsDisplayConfiguration) InitForScreenSizeInPoints(screen appkit.NSScreen, sizeInPoints corefoundation.CGSize) VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](m.ID, objc.Sel("initForScreen:sizeInPoints:"), screen, sizeInPoints)
	return rv
}

// Create a display configuration with the specified pixel dimensions and
// pixel density.
//
// widthInPixels: The width of the display, in pixels.
//
// heightInPixels: The height of the display, in pixels.
//
// pixelsPerInch: The pixel density as a number of pixels per inch.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/init(widthInPixels:heightInPixels:pixelsPerInch:)
func (m VZMacGraphicsDisplayConfiguration) InitWithWidthInPixelsHeightInPixelsPixelsPerInch(widthInPixels int, heightInPixels int, pixelsPerInch int) VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](m.ID, objc.Sel("initWithWidthInPixels:heightInPixels:pixelsPerInch:"), widthInPixels, heightInPixels, pixelsPerInch)
	return rv
}












// The height of the display, in pixels.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/heightInPixels
func (m VZMacGraphicsDisplayConfiguration) HeightInPixels() int {
	rv := objc.Send[int](m.ID, objc.Sel("heightInPixels"))
	return rv
}
func (m VZMacGraphicsDisplayConfiguration) SetHeightInPixels(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setHeightInPixels:"), value)
}



// The width of the display, in pixels.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/widthInPixels
func (m VZMacGraphicsDisplayConfiguration) WidthInPixels() int {
	rv := objc.Send[int](m.ID, objc.Sel("widthInPixels"))
	return rv
}
func (m VZMacGraphicsDisplayConfiguration) SetWidthInPixels(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setWidthInPixels:"), value)
}



// The pixel density in pixels per inch.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/pixelsPerInch
func (m VZMacGraphicsDisplayConfiguration) PixelsPerInch() int {
	rv := objc.Send[int](m.ID, objc.Sel("pixelsPerInch"))
	return rv
}
func (m VZMacGraphicsDisplayConfiguration) SetPixelsPerInch(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setPixelsPerInch:"), value)
}
























