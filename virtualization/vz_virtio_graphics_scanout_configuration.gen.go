// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioGraphicsScanoutConfiguration] class.
var (
	_VZVirtioGraphicsScanoutConfigurationClass     VZVirtioGraphicsScanoutConfigurationClass
	_VZVirtioGraphicsScanoutConfigurationClassOnce sync.Once
)

func getVZVirtioGraphicsScanoutConfigurationClass() VZVirtioGraphicsScanoutConfigurationClass {
	_VZVirtioGraphicsScanoutConfigurationClassOnce.Do(func() {
		_VZVirtioGraphicsScanoutConfigurationClass = VZVirtioGraphicsScanoutConfigurationClass{class: objc.GetClass("VZVirtioGraphicsScanoutConfiguration")}
	})
	return _VZVirtioGraphicsScanoutConfigurationClass
}

// GetVZVirtioGraphicsScanoutConfigurationClass returns the class object for VZVirtioGraphicsScanoutConfiguration.
func GetVZVirtioGraphicsScanoutConfigurationClass() VZVirtioGraphicsScanoutConfigurationClass {
	return getVZVirtioGraphicsScanoutConfigurationClass()
}

type VZVirtioGraphicsScanoutConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioGraphicsScanoutConfigurationClass) Alloc() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The configuration for a Virtio graphics device that configures the
// dimensions of the graphics device for a Linux VM.
//
// # Overview
// 
// Use a [VZVirtioGraphicsScanoutConfiguration] to configure the width and
// height of a Virtio graphics device.
//
// # Creating the configuration object
//
//   - [VZVirtioGraphicsScanoutConfiguration.InitWithWidthInPixelsHeightInPixels]: Creates a Virtio graphics device with the specified dimensions.
//
// # Instance properties
//
//   - [VZVirtioGraphicsScanoutConfiguration.HeightInPixels]: An integer value that describes the height of the graphics device in pixels.
//   - [VZVirtioGraphicsScanoutConfiguration.SetHeightInPixels]
//   - [VZVirtioGraphicsScanoutConfiguration.WidthInPixels]: An integer value that describes the width of the graphics device in pixels.
//   - [VZVirtioGraphicsScanoutConfiguration.SetWidthInPixels]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanoutConfiguration
type VZVirtioGraphicsScanoutConfiguration struct {
	VZGraphicsDisplayConfiguration
}

// VZVirtioGraphicsScanoutConfigurationFromID constructs a [VZVirtioGraphicsScanoutConfiguration] from an objc.ID.
//
// The configuration for a Virtio graphics device that configures the
// dimensions of the graphics device for a Linux VM.
func VZVirtioGraphicsScanoutConfigurationFromID(id objc.ID) VZVirtioGraphicsScanoutConfiguration {
	return VZVirtioGraphicsScanoutConfiguration{VZGraphicsDisplayConfiguration: VZGraphicsDisplayConfigurationFromID(id)}
}
// NOTE: VZVirtioGraphicsScanoutConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioGraphicsScanoutConfiguration] class.
//
// # Creating the configuration object
//
//   - [IVZVirtioGraphicsScanoutConfiguration.InitWithWidthInPixelsHeightInPixels]: Creates a Virtio graphics device with the specified dimensions.
//
// # Instance properties
//
//   - [IVZVirtioGraphicsScanoutConfiguration.HeightInPixels]: An integer value that describes the height of the graphics device in pixels.
//   - [IVZVirtioGraphicsScanoutConfiguration.SetHeightInPixels]
//   - [IVZVirtioGraphicsScanoutConfiguration.WidthInPixels]: An integer value that describes the width of the graphics device in pixels.
//   - [IVZVirtioGraphicsScanoutConfiguration.SetWidthInPixels]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanoutConfiguration
type IVZVirtioGraphicsScanoutConfiguration interface {
	IVZGraphicsDisplayConfiguration

	// Topic: Creating the configuration object

	// Creates a Virtio graphics device with the specified dimensions.
	InitWithWidthInPixelsHeightInPixels(widthInPixels int, heightInPixels int) VZVirtioGraphicsScanoutConfiguration

	// Topic: Instance properties

	// An integer value that describes the height of the graphics device in pixels.
	HeightInPixels() int
	SetHeightInPixels(value int)
	// An integer value that describes the width of the graphics device in pixels.
	WidthInPixels() int
	SetWidthInPixels(value int)

	// The array of output devices.
	Scanouts() IVZVirtioGraphicsScanoutConfiguration
	SetScanouts(value IVZVirtioGraphicsScanoutConfiguration)
}

// Init initializes the instance.
func (v VZVirtioGraphicsScanoutConfiguration) Init() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioGraphicsScanoutConfiguration) Autorelease() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsScanoutConfiguration creates a new VZVirtioGraphicsScanoutConfiguration instance.
func NewVZVirtioGraphicsScanoutConfiguration() VZVirtioGraphicsScanoutConfiguration {
	class := getVZVirtioGraphicsScanoutConfigurationClass()
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a Virtio graphics device with the specified dimensions.
//
// widthInPixels: The graphics device width in pixels.
//
// heightInPixels: The graphics device height in pixels.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanoutConfiguration/init(widthInPixels:heightInPixels:)
func NewVirtioGraphicsScanoutConfigurationWithWidthInPixelsHeightInPixels(widthInPixels int, heightInPixels int) VZVirtioGraphicsScanoutConfiguration {
	instance := getVZVirtioGraphicsScanoutConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWidthInPixels:heightInPixels:"), widthInPixels, heightInPixels)
	return VZVirtioGraphicsScanoutConfigurationFromID(rv)
}

// Creates a Virtio graphics device with the specified dimensions.
//
// widthInPixels: The graphics device width in pixels.
//
// heightInPixels: The graphics device height in pixels.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanoutConfiguration/init(widthInPixels:heightInPixels:)
func (v VZVirtioGraphicsScanoutConfiguration) InitWithWidthInPixelsHeightInPixels(widthInPixels int, heightInPixels int) VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](v.ID, objc.Sel("initWithWidthInPixels:heightInPixels:"), widthInPixels, heightInPixels)
	return rv
}

// An integer value that describes the height of the graphics device in
// pixels.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanoutConfiguration/heightInPixels
func (v VZVirtioGraphicsScanoutConfiguration) HeightInPixels() int {
	rv := objc.Send[int](v.ID, objc.Sel("heightInPixels"))
	return rv
}
func (v VZVirtioGraphicsScanoutConfiguration) SetHeightInPixels(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setHeightInPixels:"), value)
}

// An integer value that describes the width of the graphics device in pixels.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanoutConfiguration/widthInPixels
func (v VZVirtioGraphicsScanoutConfiguration) WidthInPixels() int {
	rv := objc.Send[int](v.ID, objc.Sel("widthInPixels"))
	return rv
}
func (v VZVirtioGraphicsScanoutConfiguration) SetWidthInPixels(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setWidthInPixels:"), value)
}

// The array of output devices.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtiographicsdeviceconfiguration/scanouts
func (v VZVirtioGraphicsScanoutConfiguration) Scanouts() IVZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("scanouts"))
	return VZVirtioGraphicsScanoutConfigurationFromID(objc.ID(rv))
}
func (v VZVirtioGraphicsScanoutConfiguration) SetScanouts(value IVZVirtioGraphicsScanoutConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setScanouts:"), value)
}

